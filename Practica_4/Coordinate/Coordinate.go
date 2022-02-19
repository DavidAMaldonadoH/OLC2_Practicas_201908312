package coordinate

import (
	"fmt"
	"math"
	"strconv"
)

type Coordinate struct {
	X float64
	Y float64
}

func (c Coordinate) GetX() float64 {
	return c.X
}

func (c Coordinate) GetY() float64 {
	return c.X
}

func NewCoordinate(x string, y string) Coordinate {
	p1, _ := strconv.ParseFloat(x, 64)
	p2, _ := strconv.ParseFloat(y, 64)
	return Coordinate{p1, p2}
}

func (c Coordinate) ShowInfo() {
	fmt.Print("(")
	fmt.Print(c.X)
	fmt.Print(", ")
	fmt.Print(c.Y)
	fmt.Print(")\n")
}

func CalcDistance(x1 float64, x2 float64, y1 float64, y2 float64) float64  {
	val1 := math.Pow(x1-x2, 2)
	val2 := math.Pow(y1-y2, 2)
	result := math.Sqrt(val1+val2)
	return result
}
