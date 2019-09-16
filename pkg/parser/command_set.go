package parser

import (
	"fmt"
	"strings"
)

type SetCommand struct {
	Name  string
	Value interface{}
}

func (hc *SetCommand) Run() interface{} {
	if strings.Contains(hc.Name, ".") {
		panic(fmt.Errorf("nested variables are not supported yet"))

	}
	globalvars[hc.Name] = hc.Value
	return "we got a SET Command here"

}
