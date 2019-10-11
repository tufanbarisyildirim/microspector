package parser

import "errors"

var (
	ErrStopExecution = errors.New("unable to find the key")
)

type EndCommand struct {
	Expr Expression
	When Expression
}

func (ec *EndCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	if ec.When == nil {
		return ErrStopExecution
	}

	if IsTrue(ec.When.Evaluate(l)) {
		return ErrStopExecution
	}

	return nil
}

func (ec *EndCommand) SetWhen(expr Expression) {
	ec.When = expr
}
