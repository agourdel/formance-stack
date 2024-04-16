package ledgerstore

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	ledger "github.com/formancehq/ledger/internal"
	"github.com/formancehq/stack/libs/go-libs/bun/bunpaginate"
	lquery "github.com/formancehq/stack/libs/go-libs/query"
	"github.com/uptrace/bun"
)

func convertOperatorToSQL(operator string) string {
		switch operator {
		case "$match":
			return "="
		case "$lt":
			return "<"
		case "$gt":
			return ">"
		case "$lte":
			return "<="
		case "$gte":
			return ">="
		}
		panic("unreachable")
}

type filter struct {
	condition 	string 
	args 		[]any
}

func (f *filter) Clean(){
	
	startRegex := regexp.MustCompile(`^\s*(and|or)`) // Start by space(s) then 'and|or'
	endRegex := regexp.MustCompile(`(and|or)\s*$`) // End by 'and|or' then space(s)
	voidRegex := regexp.MustCompile(`\(\s*\)`) // Parenthesis with only empty space(s) inside
	orAndRegex := regexp.MustCompile(`\(\s*(or|and)?\s*\)`) // Parenthesis with only empty space(s) and 'and|or'
	emptyRegex := regexp.MustCompile(`^\s*$`) // string with only space(s)
	
	f.condition = strings.ReplaceAll(f.condition, "()", "")
	f.condition = orAndRegex.ReplaceAllString(f.condition, "")
	f.condition = startRegex.ReplaceAllString(f.condition, "")
	f.condition = endRegex.ReplaceAllString(f.condition, "")
	f.condition = voidRegex.ReplaceAllString(f.condition, "")
	f.condition = emptyRegex.ReplaceAllString(f.condition, "")
}

func getAddressFilter(qb lquery.Builder, q GetVolumesWithBalancesQuery) (filter, error){
	var (
		subQuery string
		args     []any
		err      error
	)

	if q.Options.QueryBuilder != nil {
		
		subQuery, args, err = q.Options.QueryBuilder.Build(lquery.ContextFn(func(key, operator string, value any) (string, []any, error) {
			switch {
			case key == "account":
				// TODO: Should allow comparison operator only if segments not used
				if operator != "$match" {
					return "", nil, newErrInvalidQuery("'address' column can only be used with $match")
				}

				switch address := value.(type) {
				case string:
					return filterAccountAddress(address, "account_address"), nil, nil
				default:
					return "", nil, newErrInvalidQuery("unexpected type %T for column 'address'", address)
				}
			default:
				return "", nil, nil
			}
		}))

	}

	return filter{condition:subQuery, args:args}, err

}

func getMetadataFilter(qb lquery.Builder, q GetVolumesWithBalancesQuery) (filter, error){
	var (
		subQuery string
		args     []any
		err      error
	)

	if q.Options.QueryBuilder != nil {
		subQuery, args, err = q.Options.QueryBuilder.Build(lquery.ContextFn(func(key, operator string, value any) (string, []any, error) {
			switch {
			case metadataRegex.Match([]byte(key)):
				if operator != "$match" {
					return "", nil, newErrInvalidQuery("'metadata' column can only be used with $match")
				}
				match := metadataRegex.FindAllStringSubmatch(key, 3)
				key := "metadata"

				return key + " @> ?", []any{map[string]any{
					match[0][1]: value,
				}}, nil
			default:
				return "", nil, nil
			}
		}))
		
	}

	return filter{condition:subQuery, args: args}, err

}

func getBalancesAndAssetFilter(qb lquery.Builder, q GetVolumesWithBalancesQuery)(filter, error){
	
	balanceRegex := regexp.MustCompile("balance\\[(.*)\\]")
	
	var (
		subQuery string
		args     []any
		err      error
	)

	if q.Options.QueryBuilder != nil {
		
		subQuery, args, err = q.Options.QueryBuilder.Build(lquery.ContextFn(func(key, operator string, value any) (string, []any, error) {
			switch {
			case balanceRegex.Match([]byte(key)):
				match := balanceRegex.FindAllStringSubmatch(key, 2)
				
				var condition string = fmt.Sprintf(`sum(case when not is_source then amount else -amount end) %s ? AND asset = ?`, convertOperatorToSQL(operator))
				var args = []any{value,match[0][1]}
				return condition, args, nil
			default:
				return "", nil, nil
			}
		}))

	}

	return filter{condition:subQuery, args:args}, err

}

func (store *Store) buildVolumesWithBalancesQuery(query *bun.SelectQuery, q GetVolumesWithBalancesQuery, filterAddress filter, filterMetadata filter, filterBalancesAndAsset filter) *bun.SelectQuery {

	filtersForVolumes := q.Options.Options
	dateFilterColumn := "effective_date"

	if filtersForVolumes.UseInsertionDate {
		dateFilterColumn = "insertion_date"
	}

	query = query.
		Column("asset").
		ColumnExpr("sum(case when not is_source then amount else 0 end) as input").
		ColumnExpr("sum(case when is_source then amount else 0 end) as output").
		ColumnExpr("sum(case when not is_source then amount else -amount end) as balance").
		Table("moves"). 
		Where("ledger = ?", store.name).
		Apply(filterPIT(filtersForVolumes.PIT, dateFilterColumn)).
		Apply(filterOOT(filtersForVolumes.OOT, dateFilterColumn)).
		GroupExpr("account, asset")

	
	if filtersForVolumes.GroupLvl > 0 {
		query = query.ColumnExpr(fmt.Sprintf(`(array_to_string((string_to_array(account_address, ':'))[1:LEAST(array_length(string_to_array(account_address, ':'),1),%d)],':')) as account`, filtersForVolumes.GroupLvl))
	} else {
		query = query.ColumnExpr("account_address as account")
	}

	if filterAddress.condition != "" {
		query = query.Where(filterAddress.condition, filterAddress.args...)
	}

	if filterMetadata.condition != "" {
		query = query.
		ColumnExpr("acc.metadata as metadata").
		Join("left join account as acc").JoinOn("acc.seq = moves.accounts_seq"). 
		Where(filterMetadata.condition, filterMetadata.args...)		
	}
		
	if filterBalancesAndAsset.condition != ""{
		assets := make([]string, len(filterBalancesAndAsset.args)) 
		for i,v := range filterBalancesAndAsset.args {
			if (i%2>0){
				assets = append(assets, v.(string))
			}
		}

		query = query.
		Where("asset IN (?)", bun.In(assets)).
		Having(filterBalancesAndAsset.condition, filterBalancesAndAsset.args...)
	}
	
	return query
}

func (store *Store) GetVolumesWithBalances(ctx context.Context, q GetVolumesWithBalancesQuery) (*bunpaginate.Cursor[ledger.VolumesWithBalanceByAssetByAccount], error) {
	var err error
	filterAddress, filterMetadata, filterBalancesAndAsset := filter{condition:"", args:nil}, filter{condition:"", args:nil}, filter{condition:"", args:nil}
	 
	
	if q.Options.QueryBuilder != nil {
		filterAddress, err = getAddressFilter(q.Options.QueryBuilder, q)
		filterAddress.Clean()
		if err != nil { return nil, err}
		filterMetadata, err = getMetadataFilter(q.Options.QueryBuilder, q)
		filterMetadata.Clean()
		if err != nil { return nil, err}
		filterBalancesAndAsset, err = getBalancesAndAssetFilter(q.Options.QueryBuilder, q)
		filterBalancesAndAsset.Clean()
	}

	return paginateWithOffsetWithoutModel[PaginatedQueryOptions[FiltersForVolumes], ledger.VolumesWithBalanceByAssetByAccount](
		store, ctx, (*bunpaginate.OffsetPaginatedQuery[PaginatedQueryOptions[FiltersForVolumes]])(&q),
		func(query *bun.SelectQuery) *bun.SelectQuery {
			return store.buildVolumesWithBalancesQuery(query, q, filterAddress, filterMetadata, filterBalancesAndAsset)
		},
	)
}

type GetVolumesWithBalancesQuery bunpaginate.OffsetPaginatedQuery[PaginatedQueryOptions[FiltersForVolumes]]

func NewGetVolumesWithBalancesQuery(opts PaginatedQueryOptions[FiltersForVolumes]) GetVolumesWithBalancesQuery {
	return GetVolumesWithBalancesQuery{
		PageSize: opts.PageSize,
		Order:    bunpaginate.OrderAsc,
		Options:  opts,
	}
}
