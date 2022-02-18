package expressions

import a "practica3/Abstract"

type Literal struct {
	Line int
	Column int
	Tipo a.Type
	Value string
	a.Expression
}

func NewLiteral(line int, column int, tipo a.Type, value string) Literal {
	return Literal{Line: line,Column: column, Tipo: tipo, Value: value}
}

func (e *Literal) GetType() a.Type {
	return e.Tipo
}

func (e *Literal) GetValue() string {
	return e.Value
}

func (e Literal) Execute() a.Retorno {
	return a.Retorno{Value: e.Value, Tipo: e.Tipo}
}