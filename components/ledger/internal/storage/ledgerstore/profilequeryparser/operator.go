package profilequeryparser

import(
	"errors"
	"fmt"
)

type Operator string

const (
	And Operator = "AND"
	Or  Operator = "OR"

	Not Operator = "NOT"

	Match Operator = "MATCH"
	Lte   Operator = "LTE"
	Lt    Operator = "LT"
	Gte   Operator = "GTE"
	Gt    Operator = "GT"
)


func (o Operator) New(s string) (Operator, error) {
	switch s {
	case "$and":
		return And, nil
	case "$or":
		return Or, nil
	case "$not":
		return Not, nil
	case "$match":
		return Match, nil
	case "$lte":
		return Lte, nil
	case "$lt":
		return Lt, nil
	case "$gte":
		return Gte, nil
	case "$gt":
		return Gt, nil
	default:
		return "", errors.New(fmt.Sprintf("Undefined Operator %s in query filters", s))
	}
}

func (o Operator) IsLogical() bool {
	switch o {
		case And,Or: return true 
	}
	return false
}

func (o Operator) IsPreCondition() bool {
	switch o {
	case Not: return true
	}
	return false
}

func (o Operator) IsCondition() bool {
	switch o {
	case Match,Lte,Lt,Gte,Gt:return true
	}
	return false
}


type NextOperator map[Operator]RulesNext



type RulesNext struct {
	list []Operator 
	max int 	
}

var a = RulesNext{
	list : make([]Operator, 0),
	max: 0,
}

var nextOperators NextOperator = map[Operator]RulesNext{
	And: RulesNext{ list: []Operator{And, Or, Not, Match, Lte, Lt, Gte, Gt},max:-1},
	Or: RulesNext{ list: []Operator{And, Or, Not, Match, Lte, Lt, Gte, Gt},max:-1},
	Not: RulesNext{list:[]Operator{Match, Lte, Lt, Gte, Gt},max:1},
	Match: RulesNext{list:[]Operator{},max:1},
	Lte: RulesNext{list:[]Operator{}, max:1},
	Lt: RulesNext{list:[]Operator{}, max:1},
	Gte: RulesNext{list:[]Operator{}, max:1},
	Gt: RulesNext{list:[]Operator{}, max:1},
}