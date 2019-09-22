package parser

import "fmt"

type DebugCommand struct {
	Values []interface{}
}

func (hc *DebugCommand) Run(l *lex) interface{} {
	for _, y := range hc.Values {
		fmt.Printf("%+v ", y)
	}
	fmt.Printf("\n")
	return nil
}
