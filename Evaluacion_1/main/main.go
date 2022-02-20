package main

import (
	"Evaluacion1/main/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Listener struct {
	*parser.BaseHexadecimalListener
}

func analizar(input string) {
	fmt.Println(input)

	// create input stream
	stream := antlr.NewInputStream(input)
	// create lexer
	lexer := parser.NewHexadecimalLexer(stream)
	// create tokenStream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// create parser
	parser := parser.NewHexadecimalParser(tokenStream)
	// get tree
	tree := parser.Start()

	// activate listener
	listen := Listener{}
	antlr.ParseTreeWalkerDefault.Walk(&listen, tree)
}

func main() {
	input := "2G.10"
	analizar(input)
}