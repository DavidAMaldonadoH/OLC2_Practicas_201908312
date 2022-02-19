// Code generated from c:\Users\David Maldonado\OneDrive\Documentos\USAC\2022_1S\0781 - LABORATORIO de ORGANIZACION DE LENGUAJES Y COMPILADORES 2\Practicas\Practica_4\main\Gramatica.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import (
	arrayList "github.com/colegno/arraylist"
	m "practica4/Coordinate"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 28, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 14,
	10, 3, 12, 3, 14, 3, 17, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 2, 2, 25, 2, 8, 3, 2, 2, 2, 4,
	15, 3, 2, 2, 2, 6, 20, 3, 2, 2, 2, 8, 9, 5, 4, 3, 2, 9, 10, 7, 2, 2, 3,
	10, 11, 8, 2, 1, 2, 11, 3, 3, 2, 2, 2, 12, 14, 5, 6, 4, 2, 13, 12, 3, 2,
	2, 2, 14, 17, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 18,
	3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 18, 19, 8, 3, 1, 2, 19, 5, 3, 2, 2, 2,
	20, 21, 7, 3, 2, 2, 21, 22, 7, 6, 2, 2, 22, 23, 7, 5, 2, 2, 23, 24, 7,
	6, 2, 2, 24, 25, 7, 4, 2, 2, 25, 26, 8, 4, 1, 2, 26, 7, 3, 2, 2, 2, 3,
	15,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "','",
}
var symbolicNames = []string{
	"", "LPAR", "RPAR", "COMMA", "NUM", "WS",
}

var ruleNames = []string{
	"start", "coordinates", "coordinate",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GramaticaParser struct {
	*antlr.BaseParser
}

func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	this := new(GramaticaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gramatica.g4"

	return this
}

// GramaticaParser tokens.
const (
	GramaticaParserEOF   = antlr.TokenEOF
	GramaticaParserLPAR  = 1
	GramaticaParserRPAR  = 2
	GramaticaParserCOMMA = 3
	GramaticaParserNUM   = 4
	GramaticaParserWS    = 5
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start       = 0
	GramaticaParserRULE_coordinates = 1
	GramaticaParserRULE_coordinate  = 2
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_coordinates returns the _coordinates rule contexts.
	Get_coordinates() ICoordinatesContext

	// Set_coordinates sets the _coordinates rule contexts.
	Set_coordinates(ICoordinatesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	lista        *arrayList.List
	_coordinates ICoordinatesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_coordinates() ICoordinatesContext { return s._coordinates }

func (s *StartContext) Set_coordinates(v ICoordinatesContext) { s._coordinates = v }

func (s *StartContext) GetLista() *arrayList.List { return s.lista }

func (s *StartContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *StartContext) Coordinates() ICoordinatesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinatesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordinatesContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramaticaParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)

		var _x = p.Coordinates()

		localctx.(*StartContext)._coordinates = _x
	}
	{
		p.SetState(7)
		p.Match(GramaticaParserEOF)
	}

	localctx.(*StartContext).lista = localctx.(*StartContext).Get_coordinates().GetL()

	return localctx
}

// ICoordinatesContext is an interface to support dynamic dispatch.
type ICoordinatesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_coordinate returns the _coordinate rule contexts.
	Get_coordinate() ICoordinateContext

	// Set_coordinate sets the _coordinate rule contexts.
	Set_coordinate(ICoordinateContext)

	// GetE returns the e rule context list.
	GetE() []ICoordinateContext

	// SetE sets the e rule context list.
	SetE([]ICoordinateContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsCoordinatesContext differentiates from other interfaces.
	IsCoordinatesContext()
}

type CoordinatesContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	_coordinate ICoordinateContext
	e           []ICoordinateContext
}

func NewEmptyCoordinatesContext() *CoordinatesContext {
	var p = new(CoordinatesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_coordinates
	return p
}

func (*CoordinatesContext) IsCoordinatesContext() {}

func NewCoordinatesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordinatesContext {
	var p = new(CoordinatesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_coordinates

	return p
}

func (s *CoordinatesContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordinatesContext) Get_coordinate() ICoordinateContext { return s._coordinate }

func (s *CoordinatesContext) Set_coordinate(v ICoordinateContext) { s._coordinate = v }

func (s *CoordinatesContext) GetE() []ICoordinateContext { return s.e }

func (s *CoordinatesContext) SetE(v []ICoordinateContext) { s.e = v }

func (s *CoordinatesContext) GetL() *arrayList.List { return s.l }

func (s *CoordinatesContext) SetL(v *arrayList.List) { s.l = v }

func (s *CoordinatesContext) AllCoordinate() []ICoordinateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICoordinateContext)(nil)).Elem())
	var tst = make([]ICoordinateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICoordinateContext)
		}
	}

	return tst
}

func (s *CoordinatesContext) Coordinate(i int) ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *CoordinatesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinatesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Coordinates() (localctx ICoordinatesContext) {
	localctx = NewCoordinatesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GramaticaParserRULE_coordinates)

	localctx.(*CoordinatesContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GramaticaParserLPAR {
		{
			p.SetState(10)

			var _x = p.Coordinate()

			localctx.(*CoordinatesContext)._coordinate = _x
		}
		localctx.(*CoordinatesContext).e = append(localctx.(*CoordinatesContext).e, localctx.(*CoordinatesContext)._coordinate)

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listCoord := localctx.(*CoordinatesContext).GetE()
	for _, e := range listCoord {
		localctx.(*CoordinatesContext).l.Add(e.GetC())
	}

	return localctx
}

// ICoordinateContext is an interface to support dynamic dispatch.
type ICoordinateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left token.
	GetLeft() antlr.Token

	// GetRight returns the right token.
	GetRight() antlr.Token

	// SetLeft sets the left token.
	SetLeft(antlr.Token)

	// SetRight sets the right token.
	SetRight(antlr.Token)

	// GetC returns the c attribute.
	GetC() m.Coordinate

	// SetC sets the c attribute.
	SetC(m.Coordinate)

	// IsCoordinateContext differentiates from other interfaces.
	IsCoordinateContext()
}

type CoordinateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	c      m.Coordinate
	left   antlr.Token
	right  antlr.Token
}

func NewEmptyCoordinateContext() *CoordinateContext {
	var p = new(CoordinateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_coordinate
	return p
}

func (*CoordinateContext) IsCoordinateContext() {}

func NewCoordinateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordinateContext {
	var p = new(CoordinateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_coordinate

	return p
}

func (s *CoordinateContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordinateContext) GetLeft() antlr.Token { return s.left }

func (s *CoordinateContext) GetRight() antlr.Token { return s.right }

func (s *CoordinateContext) SetLeft(v antlr.Token) { s.left = v }

func (s *CoordinateContext) SetRight(v antlr.Token) { s.right = v }

func (s *CoordinateContext) GetC() m.Coordinate { return s.c }

func (s *CoordinateContext) SetC(v m.Coordinate) { s.c = v }

func (s *CoordinateContext) LPAR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLPAR, 0)
}

func (s *CoordinateContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCOMMA, 0)
}

func (s *CoordinateContext) RPAR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserRPAR, 0)
}

func (s *CoordinateContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserNUM)
}

func (s *CoordinateContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, i)
}

func (s *CoordinateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Coordinate() (localctx ICoordinateContext) {
	localctx = NewCoordinateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GramaticaParserRULE_coordinate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Match(GramaticaParserLPAR)
	}
	{
		p.SetState(19)

		var _m = p.Match(GramaticaParserNUM)

		localctx.(*CoordinateContext).left = _m
	}
	{
		p.SetState(20)
		p.Match(GramaticaParserCOMMA)
	}
	{
		p.SetState(21)

		var _m = p.Match(GramaticaParserNUM)

		localctx.(*CoordinateContext).right = _m
	}
	{
		p.SetState(22)
		p.Match(GramaticaParserRPAR)
	}

	localctx.(*CoordinateContext).c = m.NewCoordinate((func() string {
		if localctx.(*CoordinateContext).GetLeft() == nil {
			return ""
		} else {
			return localctx.(*CoordinateContext).GetLeft().GetText()
		}
	}()), (func() string {
		if localctx.(*CoordinateContext).GetRight() == nil {
			return ""
		} else {
			return localctx.(*CoordinateContext).GetRight().GetText()
		}
	}()))

	return localctx
}
