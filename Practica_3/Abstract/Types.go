package abstract

type Type int64

const (
	Int Type = iota
	Double
)

func (t Type) String() string {
	switch t {
	case Int:
		return "integer"
	default:
		return "double"
	}
}