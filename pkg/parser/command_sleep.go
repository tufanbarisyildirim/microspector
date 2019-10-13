package parser

import (
	"time"
)

//A command to block the current thread
type SleepCommand struct {
	Millisecond int64
	When        Expression
	Async       bool
}

func (sc *SleepCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	if sc.When == nil || IsTrue(sc.When.Evaluate(l)) {
		time.Sleep(time.Duration(sc.Millisecond) * time.Millisecond)
	}
	return nil

}

func (sc *SleepCommand) SetAsync(async bool) {
	sc.Async = async
}

func (sc *SleepCommand) SetWhen(expr Expression) {
	sc.When = expr
}

func (sc *SleepCommand) IsAsync() bool {
	return sc.Async
}
