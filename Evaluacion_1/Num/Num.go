package num

type Num struct {
	Val1 int
	Val2 float64
	Aux  float64
}

func NewNum() Num {
	return Num{Val1: 0, Val2: 0.0, Aux: 0.0}
}

func (n Num) GetValue1() int {
	return n.Val1
}

func (n *Num) SetValue1(val int) {
	n.Val1 = val
}

func (n *Num) GetValue2() float64 {
	return n.Val2
}

func (n *Num) SetValue2(val float64) {
	n.Val2 = val
}

func (n *Num) GetAux() float64 {
	return n.Aux
}

func (n *Num) SetAux(val float64) {
	n.Aux = val
}