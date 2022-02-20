// Code generated from c:\Users\David Maldonado\OneDrive\Documentos\USAC\2022_1S\0781 - LABORATORIO de ORGANIZACION DE LENGUAJES Y COMPILADORES 2\Practicas\Evaluacion_1\main\Hexadecimal.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Hexadecimal

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import (
	n "Evaluacion1/num"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 72, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 23, 10, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 33, 10, 4, 12, 4, 14, 4, 36,
	11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 70, 10, 5,
	3, 5, 2, 3, 6, 6, 2, 4, 6, 8, 2, 2, 2, 84, 2, 10, 3, 2, 2, 2, 4, 22, 3,
	2, 2, 2, 6, 24, 3, 2, 2, 2, 8, 69, 3, 2, 2, 2, 10, 11, 5, 4, 3, 2, 11,
	12, 7, 2, 2, 3, 12, 13, 8, 2, 1, 2, 13, 3, 3, 2, 2, 2, 14, 15, 5, 6, 4,
	2, 15, 16, 7, 19, 2, 2, 16, 17, 5, 6, 4, 2, 17, 18, 8, 3, 1, 2, 18, 23,
	3, 2, 2, 2, 19, 20, 5, 6, 4, 2, 20, 21, 8, 3, 1, 2, 21, 23, 3, 2, 2, 2,
	22, 14, 3, 2, 2, 2, 22, 19, 3, 2, 2, 2, 23, 5, 3, 2, 2, 2, 24, 25, 8, 4,
	1, 2, 25, 26, 5, 8, 5, 2, 26, 27, 8, 4, 1, 2, 27, 34, 3, 2, 2, 2, 28, 29,
	12, 4, 2, 2, 29, 30, 5, 8, 5, 2, 30, 31, 8, 4, 1, 2, 31, 33, 3, 2, 2, 2,
	32, 28, 3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3,
	2, 2, 2, 35, 7, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 38, 7, 3, 2, 2, 38,
	70, 8, 5, 1, 2, 39, 40, 7, 4, 2, 2, 40, 70, 8, 5, 1, 2, 41, 42, 7, 5, 2,
	2, 42, 70, 8, 5, 1, 2, 43, 44, 7, 6, 2, 2, 44, 70, 8, 5, 1, 2, 45, 46,
	7, 7, 2, 2, 46, 70, 8, 5, 1, 2, 47, 48, 7, 8, 2, 2, 48, 70, 8, 5, 1, 2,
	49, 50, 7, 9, 2, 2, 50, 70, 8, 5, 1, 2, 51, 52, 7, 10, 2, 2, 52, 70, 8,
	5, 1, 2, 53, 54, 7, 11, 2, 2, 54, 70, 8, 5, 1, 2, 55, 56, 7, 12, 2, 2,
	56, 70, 8, 5, 1, 2, 57, 58, 7, 13, 2, 2, 58, 70, 8, 5, 1, 2, 59, 60, 7,
	14, 2, 2, 60, 70, 8, 5, 1, 2, 61, 62, 7, 15, 2, 2, 62, 70, 8, 5, 1, 2,
	63, 64, 7, 16, 2, 2, 64, 70, 8, 5, 1, 2, 65, 66, 7, 17, 2, 2, 66, 70, 8,
	5, 1, 2, 67, 68, 7, 18, 2, 2, 68, 70, 8, 5, 1, 2, 69, 37, 3, 2, 2, 2, 69,
	39, 3, 2, 2, 2, 69, 41, 3, 2, 2, 2, 69, 43, 3, 2, 2, 2, 69, 45, 3, 2, 2,
	2, 69, 47, 3, 2, 2, 2, 69, 49, 3, 2, 2, 2, 69, 51, 3, 2, 2, 2, 69, 53,
	3, 2, 2, 2, 69, 55, 3, 2, 2, 2, 69, 57, 3, 2, 2, 2, 69, 59, 3, 2, 2, 2,
	69, 61, 3, 2, 2, 2, 69, 63, 3, 2, 2, 2, 69, 65, 3, 2, 2, 2, 69, 67, 3,
	2, 2, 2, 70, 9, 3, 2, 2, 2, 5, 22, 34, 69,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'0'", "'1'", "'2'", "'3'", "'4'", "'5'", "'6'", "'7'", "'8'", "'9'",
	"'A'", "'B'", "'C'", "'D'", "'E'", "'F'", "'.'",
}
var symbolicNames = []string{
	"", "CERO", "UNO", "DOS", "TRES", "CUATRO", "CINCO", "SEIS", "SIETE", "OCHO",
	"NUEVE", "DIEZ", "ONCE", "DOCE", "TRECE", "CATORCE", "QUINCE", "PUNTO",
	"WS",
}

var ruleNames = []string{
	"start", "total", "lista", "digit",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type HexadecimalParser struct {
	*antlr.BaseParser
}

func NewHexadecimalParser(input antlr.TokenStream) *HexadecimalParser {
	this := new(HexadecimalParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Hexadecimal.g4"

	return this
}

// HexadecimalParser tokens.
const (
	HexadecimalParserEOF     = antlr.TokenEOF
	HexadecimalParserCERO    = 1
	HexadecimalParserUNO     = 2
	HexadecimalParserDOS     = 3
	HexadecimalParserTRES    = 4
	HexadecimalParserCUATRO  = 5
	HexadecimalParserCINCO   = 6
	HexadecimalParserSEIS    = 7
	HexadecimalParserSIETE   = 8
	HexadecimalParserOCHO    = 9
	HexadecimalParserNUEVE   = 10
	HexadecimalParserDIEZ    = 11
	HexadecimalParserONCE    = 12
	HexadecimalParserDOCE    = 13
	HexadecimalParserTRECE   = 14
	HexadecimalParserCATORCE = 15
	HexadecimalParserQUINCE  = 16
	HexadecimalParserPUNTO   = 17
	HexadecimalParserWS      = 18
)

// HexadecimalParser rules.
const (
	HexadecimalParserRULE_start = 0
	HexadecimalParserRULE_total = 1
	HexadecimalParserRULE_lista = 2
	HexadecimalParserRULE_digit = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_total returns the _total rule contexts.
	Get_total() ITotalContext

	// Set_total sets the _total rule contexts.
	Set_total(ITotalContext)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_total ITotalContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HexadecimalParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HexadecimalParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_total() ITotalContext { return s._total }

func (s *StartContext) Set_total(v ITotalContext) { s._total = v }

func (s *StartContext) Total() ITotalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITotalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITotalContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *HexadecimalParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HexadecimalParserRULE_start)

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
		p.SetState(8)

		var _x = p.Total()

		localctx.(*StartContext)._total = _x
	}
	{
		p.SetState(9)
		p.Match(HexadecimalParserEOF)
	}

	fmt.Println(localctx.(*StartContext).Get_total().GetRes())

	return localctx
}

// ITotalContext is an interface to support dynamic dispatch.
type ITotalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL1 returns the l1 rule contexts.
	GetL1() IListaContext

	// GetL2 returns the l2 rule contexts.
	GetL2() IListaContext

	// SetL1 sets the l1 rule contexts.
	SetL1(IListaContext)

	// SetL2 sets the l2 rule contexts.
	SetL2(IListaContext)

	// GetRes returns the res attribute.
	GetRes() float64

	// SetRes sets the res attribute.
	SetRes(float64)

	// IsTotalContext differentiates from other interfaces.
	IsTotalContext()
}

type TotalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	res    float64
	l1     IListaContext
	l2     IListaContext
}

func NewEmptyTotalContext() *TotalContext {
	var p = new(TotalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HexadecimalParserRULE_total
	return p
}

func (*TotalContext) IsTotalContext() {}

func NewTotalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TotalContext {
	var p = new(TotalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HexadecimalParserRULE_total

	return p
}

func (s *TotalContext) GetParser() antlr.Parser { return s.parser }

func (s *TotalContext) GetL1() IListaContext { return s.l1 }

func (s *TotalContext) GetL2() IListaContext { return s.l2 }

func (s *TotalContext) SetL1(v IListaContext) { s.l1 = v }

func (s *TotalContext) SetL2(v IListaContext) { s.l2 = v }

func (s *TotalContext) GetRes() float64 { return s.res }

func (s *TotalContext) SetRes(v float64) { s.res = v }

func (s *TotalContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserPUNTO, 0)
}

func (s *TotalContext) AllLista() []IListaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IListaContext)(nil)).Elem())
	var tst = make([]IListaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IListaContext)
		}
	}

	return tst
}

func (s *TotalContext) Lista(i int) IListaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IListaContext)
}

func (s *TotalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TotalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *HexadecimalParser) Total() (localctx ITotalContext) {
	localctx = NewTotalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HexadecimalParserRULE_total)

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

	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)

			var _x = p.lista(0)

			localctx.(*TotalContext).l1 = _x
		}
		{
			p.SetState(13)
			p.Match(HexadecimalParserPUNTO)
		}
		{
			p.SetState(14)

			var _x = p.lista(0)

			localctx.(*TotalContext).l2 = _x
		}

		localctx.(*TotalContext).res = float64(localctx.(*TotalContext).GetL1().GetNumero().GetValue1()) + localctx.(*TotalContext).GetL2().GetNumero().GetValue2()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)

			var _x = p.lista(0)

			localctx.(*TotalContext).l1 = _x
		}

		localctx.(*TotalContext).res = float64(localctx.(*TotalContext).GetL1().GetNumero().GetValue1())

	}

	return localctx
}

// IListaContext is an interface to support dynamic dispatch.
type IListaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetX returns the x rule contexts.
	GetX() IListaContext

	// GetY returns the y rule contexts.
	GetY() IDigitContext

	// SetX sets the x rule contexts.
	SetX(IListaContext)

	// SetY sets the y rule contexts.
	SetY(IDigitContext)

	// GetNumero returns the numero attribute.
	GetNumero() *n.Num

	// SetNumero sets the numero attribute.
	SetNumero(*n.Num)

	// IsListaContext differentiates from other interfaces.
	IsListaContext()
}

type ListaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	numero *n.Num
	x      IListaContext
	y      IDigitContext
}

func NewEmptyListaContext() *ListaContext {
	var p = new(ListaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HexadecimalParserRULE_lista
	return p
}

func (*ListaContext) IsListaContext() {}

func NewListaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaContext {
	var p = new(ListaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HexadecimalParserRULE_lista

	return p
}

func (s *ListaContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaContext) GetX() IListaContext { return s.x }

func (s *ListaContext) GetY() IDigitContext { return s.y }

func (s *ListaContext) SetX(v IListaContext) { s.x = v }

func (s *ListaContext) SetY(v IDigitContext) { s.y = v }

func (s *ListaContext) GetNumero() *n.Num { return s.numero }

func (s *ListaContext) SetNumero(v *n.Num) { s.numero = v }

func (s *ListaContext) Digit() IDigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDigitContext)
}

func (s *ListaContext) Lista() IListaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaContext)
}

func (s *ListaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *HexadecimalParser) Lista() (localctx IListaContext) {
	return p.lista(0)
}

func (p *HexadecimalParser) lista(_p int) (localctx IListaContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, HexadecimalParserRULE_lista, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)

		var _x = p.Digit()

		localctx.(*ListaContext).y = _x
	}

	valor, _ := strconv.Atoi(localctx.(*ListaContext).GetY().GetS())
	localctx.(*ListaContext).numero = &n.Num{}
	localctx.(*ListaContext).numero.SetValue1(valor)
	localctx.(*ListaContext).numero.SetAux(0.0625)
	localctx.(*ListaContext).numero.SetValue2(float64(valor) * 0.0625)

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaContext(p, _parentctx, _parentState)
			localctx.(*ListaContext).x = _prevctx
			p.PushNewRecursionContext(localctx, _startState, HexadecimalParserRULE_lista)
			p.SetState(26)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(27)

				var _x = p.Digit()

				localctx.(*ListaContext).y = _x
			}

			valor, _ := strconv.Atoi(localctx.(*ListaContext).GetY().GetS())
			valor1 := localctx.(*ListaContext).GetX().GetNumero().GetValue1()*16 + valor
			localctx.(*ListaContext).numero = localctx.(*ListaContext).GetX().GetNumero()
			localctx.(*ListaContext).numero.SetAux(localctx.(*ListaContext).GetX().GetNumero().GetAux() * 0.0625)
			valor2 := localctx.(*ListaContext).GetX().GetNumero().GetValue2() + float64(valor)*localctx.(*ListaContext).GetX().GetNumero().GetAux()
			localctx.(*ListaContext).numero.SetValue1(valor1)
			localctx.(*ListaContext).numero.SetValue2(valor2)

		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IDigitContext is an interface to support dynamic dispatch.
type IDigitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetS returns the s attribute.
	GetS() string

	// SetS sets the s attribute.
	SetS(string)

	// IsDigitContext differentiates from other interfaces.
	IsDigitContext()
}

type DigitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	s      string
}

func NewEmptyDigitContext() *DigitContext {
	var p = new(DigitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HexadecimalParserRULE_digit
	return p
}

func (*DigitContext) IsDigitContext() {}

func NewDigitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DigitContext {
	var p = new(DigitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HexadecimalParserRULE_digit

	return p
}

func (s *DigitContext) GetParser() antlr.Parser { return s.parser }

func (s *DigitContext) GetS() string { return s.s }

func (s *DigitContext) SetS(v string) { s.s = v }

func (s *DigitContext) CERO() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserCERO, 0)
}

func (s *DigitContext) UNO() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserUNO, 0)
}

func (s *DigitContext) DOS() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserDOS, 0)
}

func (s *DigitContext) TRES() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserTRES, 0)
}

func (s *DigitContext) CUATRO() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserCUATRO, 0)
}

func (s *DigitContext) CINCO() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserCINCO, 0)
}

func (s *DigitContext) SEIS() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserSEIS, 0)
}

func (s *DigitContext) SIETE() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserSIETE, 0)
}

func (s *DigitContext) OCHO() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserOCHO, 0)
}

func (s *DigitContext) NUEVE() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserNUEVE, 0)
}

func (s *DigitContext) DIEZ() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserDIEZ, 0)
}

func (s *DigitContext) ONCE() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserONCE, 0)
}

func (s *DigitContext) DOCE() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserDOCE, 0)
}

func (s *DigitContext) TRECE() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserTRECE, 0)
}

func (s *DigitContext) CATORCE() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserCATORCE, 0)
}

func (s *DigitContext) QUINCE() antlr.TerminalNode {
	return s.GetToken(HexadecimalParserQUINCE, 0)
}

func (s *DigitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *HexadecimalParser) Digit() (localctx IDigitContext) {
	localctx = NewDigitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HexadecimalParserRULE_digit)

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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HexadecimalParserCERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.Match(HexadecimalParserCERO)
		}
		localctx.(*DigitContext).s = "0"

	case HexadecimalParserUNO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(HexadecimalParserUNO)
		}
		localctx.(*DigitContext).s = "1"

	case HexadecimalParserDOS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.Match(HexadecimalParserDOS)
		}
		localctx.(*DigitContext).s = "2"

	case HexadecimalParserTRES:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Match(HexadecimalParserTRES)
		}
		localctx.(*DigitContext).s = "3"

	case HexadecimalParserCUATRO:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(43)
			p.Match(HexadecimalParserCUATRO)
		}
		localctx.(*DigitContext).s = "4"

	case HexadecimalParserCINCO:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(45)
			p.Match(HexadecimalParserCINCO)
		}
		localctx.(*DigitContext).s = "5"

	case HexadecimalParserSEIS:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(47)
			p.Match(HexadecimalParserSEIS)
		}
		localctx.(*DigitContext).s = "6"

	case HexadecimalParserSIETE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(49)
			p.Match(HexadecimalParserSIETE)
		}
		localctx.(*DigitContext).s = "7"

	case HexadecimalParserOCHO:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(51)
			p.Match(HexadecimalParserOCHO)
		}
		localctx.(*DigitContext).s = "8"

	case HexadecimalParserNUEVE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(53)
			p.Match(HexadecimalParserNUEVE)
		}
		localctx.(*DigitContext).s = "9"

	case HexadecimalParserDIEZ:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(55)
			p.Match(HexadecimalParserDIEZ)
		}
		localctx.(*DigitContext).s = "10"

	case HexadecimalParserONCE:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(57)
			p.Match(HexadecimalParserONCE)
		}
		localctx.(*DigitContext).s = "11"

	case HexadecimalParserDOCE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(59)
			p.Match(HexadecimalParserDOCE)
		}
		localctx.(*DigitContext).s = "12"

	case HexadecimalParserTRECE:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(61)
			p.Match(HexadecimalParserTRECE)
		}
		localctx.(*DigitContext).s = "13"

	case HexadecimalParserCATORCE:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(63)
			p.Match(HexadecimalParserCATORCE)
		}
		localctx.(*DigitContext).s = "14"

	case HexadecimalParserQUINCE:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(65)
			p.Match(HexadecimalParserQUINCE)
		}
		localctx.(*DigitContext).s = "15"

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *HexadecimalParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ListaContext = nil
		if localctx != nil {
			t = localctx.(*ListaContext)
		}
		return p.Lista_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *HexadecimalParser) Lista_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
