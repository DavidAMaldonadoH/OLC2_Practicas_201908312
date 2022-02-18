package abstract

type Any interface {
	String() string
}

type Entero struct {
	Value int
	Any
}

type Decimal struct {
	Value float64
	Any
}

type Retorno struct {
	Value string
	Tipo  Type
}
