// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)
}
