// Code generated from c:\Users\David Maldonado\OneDrive\Documentos\USAC\2022_1S\0781 - LABORATORIO de ORGANIZACION DE LENGUAJES Y COMPILADORES 2\Practicas\Evaluacion_1\main\Hexadecimal.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 80, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18,
	9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 6, 19, 75, 10, 19, 13, 19, 14, 19, 76,
	3, 19, 3, 19, 2, 2, 20, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35,
	19, 37, 20, 3, 2, 3, 5, 2, 11, 12, 15, 15, 34, 34, 2, 80, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35,
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3, 39, 3, 2, 2, 2, 5, 41, 3, 2, 2, 2, 7,
	43, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 47, 3, 2, 2, 2, 13, 49, 3, 2, 2,
	2, 15, 51, 3, 2, 2, 2, 17, 53, 3, 2, 2, 2, 19, 55, 3, 2, 2, 2, 21, 57,
	3, 2, 2, 2, 23, 59, 3, 2, 2, 2, 25, 61, 3, 2, 2, 2, 27, 63, 3, 2, 2, 2,
	29, 65, 3, 2, 2, 2, 31, 67, 3, 2, 2, 2, 33, 69, 3, 2, 2, 2, 35, 71, 3,
	2, 2, 2, 37, 74, 3, 2, 2, 2, 39, 40, 7, 50, 2, 2, 40, 4, 3, 2, 2, 2, 41,
	42, 7, 51, 2, 2, 42, 6, 3, 2, 2, 2, 43, 44, 7, 52, 2, 2, 44, 8, 3, 2, 2,
	2, 45, 46, 7, 53, 2, 2, 46, 10, 3, 2, 2, 2, 47, 48, 7, 54, 2, 2, 48, 12,
	3, 2, 2, 2, 49, 50, 7, 55, 2, 2, 50, 14, 3, 2, 2, 2, 51, 52, 7, 56, 2,
	2, 52, 16, 3, 2, 2, 2, 53, 54, 7, 57, 2, 2, 54, 18, 3, 2, 2, 2, 55, 56,
	7, 58, 2, 2, 56, 20, 3, 2, 2, 2, 57, 58, 7, 59, 2, 2, 58, 22, 3, 2, 2,
	2, 59, 60, 7, 67, 2, 2, 60, 24, 3, 2, 2, 2, 61, 62, 7, 68, 2, 2, 62, 26,
	3, 2, 2, 2, 63, 64, 7, 69, 2, 2, 64, 28, 3, 2, 2, 2, 65, 66, 7, 70, 2,
	2, 66, 30, 3, 2, 2, 2, 67, 68, 7, 71, 2, 2, 68, 32, 3, 2, 2, 2, 69, 70,
	7, 72, 2, 2, 70, 34, 3, 2, 2, 2, 71, 72, 7, 48, 2, 2, 72, 36, 3, 2, 2,
	2, 73, 75, 9, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 74,
	3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 8, 19, 2, 2,
	79, 38, 3, 2, 2, 2, 4, 2, 76, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'0'", "'1'", "'2'", "'3'", "'4'", "'5'", "'6'", "'7'", "'8'", "'9'",
	"'A'", "'B'", "'C'", "'D'", "'E'", "'F'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "CERO", "UNO", "DOS", "TRES", "CUATRO", "CINCO", "SEIS", "SIETE", "OCHO",
	"NUEVE", "DIEZ", "ONCE", "DOCE", "TRECE", "CATORCE", "QUINCE", "PUNTO",
	"WS",
}

var lexerRuleNames = []string{
	"CERO", "UNO", "DOS", "TRES", "CUATRO", "CINCO", "SEIS", "SIETE", "OCHO",
	"NUEVE", "DIEZ", "ONCE", "DOCE", "TRECE", "CATORCE", "QUINCE", "PUNTO",
	"WS",
}

type HexadecimalLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewHexadecimalLexer(input antlr.CharStream) *HexadecimalLexer {

	l := new(HexadecimalLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Hexadecimal.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// HexadecimalLexer tokens.
const (
	HexadecimalLexerCERO    = 1
	HexadecimalLexerUNO     = 2
	HexadecimalLexerDOS     = 3
	HexadecimalLexerTRES    = 4
	HexadecimalLexerCUATRO  = 5
	HexadecimalLexerCINCO   = 6
	HexadecimalLexerSEIS    = 7
	HexadecimalLexerSIETE   = 8
	HexadecimalLexerOCHO    = 9
	HexadecimalLexerNUEVE   = 10
	HexadecimalLexerDIEZ    = 11
	HexadecimalLexerONCE    = 12
	HexadecimalLexerDOCE    = 13
	HexadecimalLexerTRECE   = 14
	HexadecimalLexerCATORCE = 15
	HexadecimalLexerQUINCE  = 16
	HexadecimalLexerPUNTO   = 17
	HexadecimalLexerWS      = 18
)
