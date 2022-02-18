package abstract

type Expression interface {
	Execute() Retorno
}

func GetDominantType(type1 Type, type2 Type, tablaOp [][]Type) Type {
	tipo := tablaOp[type1][type2]
	return tipo
}