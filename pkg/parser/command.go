package parser

type Command interface {
	Run(l *Lexer) interface{}
}
