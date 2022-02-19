// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterCoordinates is called when entering the coordinates production.
	EnterCoordinates(c *CoordinatesContext)

	// EnterCoordinate is called when entering the coordinate production.
	EnterCoordinate(c *CoordinateContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitCoordinates is called when exiting the coordinates production.
	ExitCoordinates(c *CoordinatesContext)

	// ExitCoordinate is called when exiting the coordinate production.
	ExitCoordinate(c *CoordinateContext)
}
