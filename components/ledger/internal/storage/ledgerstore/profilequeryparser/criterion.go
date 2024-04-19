package profilequeryparser

import (
	"errors"
)



type ICriterion interface {
	Name() string
	Value() any
	Operator() Operator
	IsNot() bool

	Set(key string, operator Operator, value any) ICriterion
	SetNot(key string, operator Operator, value any) ICriterion
	Check(key string) bool

	Copy() ICriterion
	InverseWithNot() ICriterion
	Inverse() ICriterion
	Merge(ICriterion) (ICriterion,error)




}

type Criterion struct {
	name     string
	not      bool
	value    any
	operator Operator

	_Set   func(*Criterion, string, Operator, any) ICriterion
	_Check func(string) bool


}

func (c Criterion) Name() string {
	return c.name
}

func (c Criterion) Value() any {
	return c.value
}

func (c Criterion) Operator() Operator {
	return c.operator
}

func (c Criterion) IsNot() bool {
	return c.not
}

func (c Criterion) Set(key string, operator Operator, value any) ICriterion {
	c1 := c.Copy().(Criterion)
	c1.not = false
	return c1._Set(&c1, key, operator, value)
}

func (c Criterion) SetNot(key string, operator Operator, value any) ICriterion {
	c1 := c.Copy().(Criterion)
	c1.not = true
	return c._Set(&c1, key, operator, value)
}

func (c Criterion) Check(key string) bool {
	return c._Check(key)
}

func (c *Criterion) WithCustomSet(setExpress func(*Criterion, string, Operator, any) ICriterion) *Criterion {
	c._Set = setExpress
	return c
}


func (c Criterion) Copy() ICriterion {
	new := Criterion{
		name : c.name,
		not : c.not,
		value: c.value,
		operator: c.operator,
		_Set : c._Set,
		_Check: c._Check, 
	}

	return new
}

func (c Criterion) InverseWithNot() ICriterion{
	new := Criterion{
		name : c.name,
		not : !c.not,
		value: c.value,
		operator: c.operator,
		_Set : c._Set,
		_Check: c._Check, 
	}
	

	return new
}

func (c Criterion) Inverse() ICriterion{

	new := Criterion{
		name : c.name,
		not : !c.not,
		value: c.value,
		operator: c.operator,
		_Set : c._Set,
		_Check: c._Check, 
	}
	

	return new

}

func (c Criterion) IsSame(c1 ICriterion) bool {
	return c.name == c1.Name() && c.operator == c1.Operator() && c.not == c1.IsNot()
}

func (c Criterion) IsInverse(c1 ICriterion) bool {
	//process inverse
	return true
}

func (c Criterion) Merge(c1 ICriterion) (ICriterion, error) {
	if(!c.IsSame(c1)) {return Criterion{}, errors.New("Criterions are not the same, they can't be merged") }
	//Process merge
	return Criterion{}, nil
}




func NewCriterion(name string, checkExpress func(string) bool) *Criterion {

	var _set func(*Criterion, string, Operator, any) ICriterion = func(c *Criterion, key string, operator Operator, value any) ICriterion{
		c.operator = operator
		c.value = value
		return c
	}

	var c Criterion = Criterion{
		name: name,

		_Check: checkExpress,
		_Set:   _set,
	}

	return &c
}


