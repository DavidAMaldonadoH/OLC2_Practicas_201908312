// Code generated from Hexadecimal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Hexadecimal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HexadecimalListener is a complete listener for a parse tree produced by HexadecimalParser.
type HexadecimalListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterTotal is called when entering the total production.
	EnterTotal(c *TotalContext)

	// EnterLista is called when entering the lista production.
	EnterLista(c *ListaContext)

	// EnterDigit is called when entering the digit production.
	EnterDigit(c *DigitContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitTotal is called when exiting the total production.
	ExitTotal(c *TotalContext)

	// ExitLista is called when exiting the lista production.
	ExitLista(c *ListaContext)

	// ExitDigit is called when exiting the digit production.
	ExitDigit(c *DigitContext)
}
