package main

import (
	"bufio"
	"log"
	"os"
	"practica3/main/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)


type TreeSyntax struct {
	*parser.BaseGramaticaListener
}

func NewTreeSyntax() *TreeSyntax {
	return new(TreeSyntax)
}

func interpretar(in string)  {
	is := antlr.NewInputStream(in)

	lexer := parser.NewGramaticaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGramaticaParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(NewTreeSyntax(), p.Start())
}

func main() {
	file, err := os.Open("entrada.txt")

	if err != nil {
		log.Fatal(err)
	}
	// se invoca hasta el final
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	text := ""

	for fileScanner.Scan() {
		text += fileScanner.Text()
	}

	interpretar(text)
}