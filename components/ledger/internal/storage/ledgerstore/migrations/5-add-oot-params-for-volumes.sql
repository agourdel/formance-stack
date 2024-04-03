drop function if exists  get_all_account_effective_volumes;
drop function if exists get_all_account_volumes;
drop function if exists get_account_aggregated_effective_volumes;
drop function if exists get_account_aggregated_volumes;


create or replace function get_naive_all_account_effective_volumes(_ledger varchar, _account varchar, _after timestamp default null, _before timestamp default null)
    returns setof volumes_with_asset
    language sql
    stable
as
$$
with results as (
  select 
    asset,
    sum(case when not is_source then amount else 0 end) as inputs,
    sum(case when is_source then amount else 0 end) as outputs
  from moves 
  where ledger = _ledger 
  and account_address = _account
  and (_after is null or effective_date >= _after)
  and (_before is null or effective_date <= _before) 
  group by asset 
)
select results.asset, row(results.inputs, results.outputs)::volumes as volumes 
from results
$$;

create or replace function get_all_account_effective_volumes(_ledger varchar, _account varchar, _before timestamp default null)
    returns setof volumes_with_asset
    language sql
    stable
as
$$
with all_assets as (select v.v as asset
                    from get_all_assets(_ledger) v),
     moves as (select m.*
               from all_assets assets
                        join lateral (
                   select *
                   from moves s
                   where (_before is null or s.effective_date <= _before)
                     and s.account_address = _account
                     and s.asset = assets.asset
                     and s.ledger = _ledger
                   order by effective_date desc, seq desc
                   limit 1
                   ) m on true)
select moves.asset, moves.post_commit_effective_volumes
from moves
$$;

create or replace function get_naive_all_account_volumes(_ledger varchar, _account varchar, _after timestamp default null, _before timestamp default null)
    returns setof volumes_with_asset
    language sql
    stable
as
$$
with results as (
  select 
    asset,
    sum(case when not is_source then amount else 0 end) as inputs,
    sum(case when is_source then amount else 0 end) as outputs
  from moves 
  where ledger = _ledger 
  and account_address = _account
  and (_after is null or insertion_date >= _after)
  and (_before is null or insertion_date <= _before) 
  group by asset 
)
select results.asset, row(results.inputs, results.outputs)::volumes as volumes 
from results
$$;

create or replace function get_all_account_volumes(_ledger varchar, _account varchar, _before timestamp default null)
    returns setof volumes_with_asset
    language sql
    stable
as
$$
with all_assets as (select v.v as asset
                    from get_all_assets(_ledger) v),
     moves as (select m.*
               from all_assets assets
                        join lateral (
                   select *
                   from moves s
                   where (_before is null or s.insertion_date <= _before)
                     and s.account_address = _account
                     and s.asset = assets.asset
                     and s.ledger = _ledger
                   order by seq desc
                   limit 1
                   ) m on true)
select moves.asset, moves.post_commit_volumes
from moves
$$;


create or replace function get_naive_account_aggregated_effective_volumes(_ledger varchar, _account_address varchar, _after timestamp default null,
                                                         _before timestamp default null)
    returns jsonb
    language sql
    stable
as
$$
select aggregate_objects(volumes_to_jsonb(volumes_with_asset))
from get_naive_all_account_effective_volumes(_ledger, _account_address, _after, _before) volumes_with_asset
$$;

create or replace function get_account_aggregated_effective_volumes(_ledger varchar, _account_address varchar,
                                                         _before timestamp default null)
    returns jsonb
    language sql
    stable
as
$$
select aggregate_objects(volumes_to_jsonb(volumes_with_asset))
from get_all_account_effective_volumes(_ledger, _account_address, _before) volumes_with_asset
$$;

create or replace function get_naive_account_aggregated_volumes(_ledger varchar, _account_address varchar, _after timestamp default null,
                                               _before timestamp default null)
    returns jsonb
    language sql
    stable
    parallel safe
as
$$
select aggregate_objects(volumes_to_jsonb(volumes_with_asset))
from get_naive_all_account_volumes(_ledger, _account_address, _after, _before) volumes_with_asset
$$;

create or replace function get_account_aggregated_volumes(_ledger varchar, _account_address varchar,
                                               _before timestamp default null)
    returns jsonb
    language sql
    stable
    parallel safe
as
$$
select aggregate_objects(volumes_to_jsonb(volumes_with_asset))
from get_all_account_volumes(_ledger, _account_address, _before) volumes_with_asset
$$;

