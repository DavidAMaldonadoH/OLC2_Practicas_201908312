package main

import (
	"fmt"
	coordinate "practica4/Coordinate"
	"practica4/main/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)


type TreeSyntax struct {
	*parser.BaseGramaticaListener
}

func NewTreeSyntax() *TreeSyntax {
	return new(TreeSyntax)
}

func (t *TreeSyntax) ExitStart(ctx *parser.StartContext) {
	result := ctx.GetLista()
	var val float64
	fmt.Println(result.ToArray())
	for i := range result.ToArray() {
		if i != (result.Len()-1) {
			c1 := result.GetValue(i).(coordinate.Coordinate)
			c2 := result.GetValue(i+1).(coordinate.Coordinate)
			val += coordinate.CalcDistance(c1.X, c2.X, c1.Y, c2.Y)
		}
	}
	fmt.Printf("La distancia es: %.2f", val)
}

func main() {
	is, _ := antlr.NewFileStream("entrada.txt")
	// Create the Lexer
	lexer := parser.NewGramaticaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGramaticaParser(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeSyntax(), tree)
}