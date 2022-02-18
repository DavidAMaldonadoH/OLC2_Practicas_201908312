package expressions

import (
	"fmt"
	abs "practica3/Abstract"
	tabs "practica3/Tablas"
	"strconv"
)

type Arithmetic struct {
	Line int
	Column int
	Left   abs.Expression
	Right  abs.Expression
	A_type   ArithmeticType
	abs.Expression
}

func NewArithmetic(line int, column int, left abs.Expression, right abs.Expression, a_type ArithmeticType) Arithmetic {
	return Arithmetic{Line: line, Column: column, Left: left, Right: right, A_type: a_type}
}

func (ae Arithmetic) GetLeft() abs.Expression {
	return ae.Left
}

func (ae Arithmetic) GetRight() abs.Expression {
	return ae.Right
}

func (ae Arithmetic) Execute() abs.Retorno {
	leftOperand := ae.Left.Execute()
	rightOperand := ae.Right.Execute()
	result := abs.Retorno{}
	dominantType := abs.GetDominantType(leftOperand.Tipo, rightOperand.Tipo, tabs.Tipos)
	if ae.A_type == Suma {
		if dominantType == abs.Int {
			value1, err1 := strconv.Atoi(leftOperand.Value)
			value2, err2 := strconv.Atoi(rightOperand.Value)
			result.Value = strconv.Itoa(value1 + value2)
			result.Tipo = abs.Int
			if (err1 != nil) || (err2 != nil) {
					fmt.Println("Error en el casteo")
			}
		} else {
			value1, err1 := strconv.ParseFloat(leftOperand.Value, 64)
			value2, err2 := strconv.ParseFloat(rightOperand.Value, 64)
			result.Value = fmt.Sprintf("%v", value1 + value2)
			result.Tipo = abs.Double
			if (err1 != nil) || (err2 != nil) {
				fmt.Println("Error en el casteo")
			}
		}
	} else if ae.A_type == Resta {
		if dominantType == abs.Int {
			value1, err1 := strconv.Atoi(leftOperand.Value)
			value2, err2 := strconv.Atoi(rightOperand.Value)
			result.Value = strconv.Itoa(value1 - value2)
			result.Tipo = abs.Int
			if (err1 != nil) || (err2 != nil) {
					fmt.Println("Error en el casteo")
			}
		} else {
			value1, err1 := strconv.ParseFloat(leftOperand.Value, 64)
			value2, err2 := strconv.ParseFloat(rightOperand.Value, 64)
			result.Value = fmt.Sprintf("%v", value1 - value2)
			result.Tipo = abs.Double
			if (err1 != nil) || (err2 != nil) {
				fmt.Println("Error en el casteo")
			}
		}
	} else if ae.A_type == Multiplicacion {
		if dominantType == abs.Int {
			value1, err1 := strconv.Atoi(leftOperand.Value)
			value2, err2 := strconv.Atoi(rightOperand.Value)
			result.Value = strconv.Itoa(value1 * value2)
			result.Tipo = abs.Int
			if (err1 != nil) || (err2 != nil) {
					fmt.Println("Error en el casteo")
			}
		} else {
			value1, err1 := strconv.ParseFloat(leftOperand.Value, 64)
			value2, err2 := strconv.ParseFloat(rightOperand.Value, 64)
			result.Value = fmt.Sprintf("%v", value1 * value2)
			result.Tipo = abs.Double
			if (err1 != nil) || (err2 != nil) {
				fmt.Println("Error en el casteo")
			}
		}
	} else {
		value1, err1 := strconv.ParseFloat(leftOperand.Value, 64)
		value2, err2 := strconv.ParseFloat(rightOperand.Value, 64)
		result.Value = fmt.Sprintf("%v", value1 / value2)
		result.Tipo = abs.Double
		if (err1 != nil) || (err2 != nil) {
			fmt.Println("Error en el casteo")
		}	
	}
	return result
}

type ArithmeticType int64

const (
	Suma ArithmeticType = iota
	Resta
	Multiplicacion 
	Division
)

func (at ArithmeticType) String() string {
	switch at {
	case Suma:
		return "suma"
	case Resta:
		return "resta"
	case Multiplicacion:
		return "multiplicacion"
	default:
		return "division"
	}
}
