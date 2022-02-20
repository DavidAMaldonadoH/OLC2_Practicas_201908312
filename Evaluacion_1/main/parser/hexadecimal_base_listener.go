// Code generated from Hexadecimal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Hexadecimal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHexadecimalListener is a complete listener for a parse tree produced by HexadecimalParser.
type BaseHexadecimalListener struct{}

var _ HexadecimalListener = &BaseHexadecimalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHexadecimalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHexadecimalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHexadecimalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHexadecimalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseHexadecimalListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseHexadecimalListener) ExitStart(ctx *StartContext) {}

// EnterTotal is called when production total is entered.
func (s *BaseHexadecimalListener) EnterTotal(ctx *TotalContext) {}

// ExitTotal is called when production total is exited.
func (s *BaseHexadecimalListener) ExitTotal(ctx *TotalContext) {}

// EnterLista is called when production lista is entered.
func (s *BaseHexadecimalListener) EnterLista(ctx *ListaContext) {}

// ExitLista is called when production lista is exited.
func (s *BaseHexadecimalListener) ExitLista(ctx *ListaContext) {}

// EnterDigit is called when production digit is entered.
func (s *BaseHexadecimalListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production digit is exited.
func (s *BaseHexadecimalListener) ExitDigit(ctx *DigitContext) {}
