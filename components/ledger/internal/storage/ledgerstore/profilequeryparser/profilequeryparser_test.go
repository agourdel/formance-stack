package profilequeryparser

import (
	"fmt"
	"testing"
)


func test(c ICriterion){
	println(fmt.Sprintln("YEs"))		
}

func TestMain(t *testing.T){

	c := Criterion{}
	test(&c)

	


}