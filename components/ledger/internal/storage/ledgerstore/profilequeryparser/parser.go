package profilequeryparser

import (
	"encoding/json"
	"errors"
)


type Node interface{
	Filters|Operator|Criterion
}

type Filters[T any] map[string]T

func parseRawData[T any](rawFilterQuery []byte) (Filters[T], error){
	if(len(rawFilterQuery) == 0) {return nil, errors.New("Empty RawFilterQuery")}

	 filters := make(Filters[T])
	 if err := json.Unmarshal(rawFilterQuery, &filters); err!=nil{
		return nil, err
	 }

	 return filters,nil
} 