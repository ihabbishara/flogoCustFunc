package mod10esr

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
)

func TestFnmod10esr_Eval(t *testing.T) {
	f := &mod10esr{}

	v, err := function.Eval(f, "0946827135", "988138","0000096950","400000001")
	if err != nil {
		fmt.Println("error occured:   ", err)
	}
	fmt.Println(v)
}