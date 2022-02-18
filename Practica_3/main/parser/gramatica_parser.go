// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import (
	abs "practica3/Abstract"
	exp "practica3/Expressions"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 33, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 16, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 7, 3, 28, 10, 3, 12, 3, 14, 3, 31, 11, 3, 3, 3, 2, 3, 4, 4, 2,
	4, 2, 4, 3, 2, 5, 6, 3, 2, 3, 4, 2, 33, 2, 6, 3, 2, 2, 2, 4, 15, 3, 2,
	2, 2, 6, 7, 5, 4, 3, 2, 7, 8, 7, 2, 2, 3, 8, 9, 8, 2, 1, 2, 9, 3, 3, 2,
	2, 2, 10, 11, 8, 3, 1, 2, 11, 12, 7, 7, 2, 2, 12, 16, 8, 3, 1, 2, 13, 14,
	7, 8, 2, 2, 14, 16, 8, 3, 1, 2, 15, 10, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2,
	16, 29, 3, 2, 2, 2, 17, 18, 12, 6, 2, 2, 18, 19, 9, 2, 2, 2, 19, 20, 5,
	4, 3, 7, 20, 21, 8, 3, 1, 2, 21, 28, 3, 2, 2, 2, 22, 23, 12, 5, 2, 2, 23,
	24, 9, 3, 2, 2, 24, 25, 5, 4, 3, 6, 25, 26, 8, 3, 1, 2, 26, 28, 3, 2, 2,
	2, 27, 17, 3, 2, 2, 2, 27, 22, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27,
	3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 5, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2,
	5, 15, 27, 29,
}
var literalNames = []string{
	"", "'+'", "'-'", "'*'", "'/'",
}
var symbolicNames = []string{
	"", "PLUS", "MINUS", "TIMES", "DIVISION", "ENTERO", "DECIMAL", "WS",
}

var ruleNames = []string{
	"start", "exp",
}

type GramaticaParser struct {
	*antlr.BaseParser
}

// NewGramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GramaticaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	this := new(GramaticaParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
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
	GramaticaParserEOF      = antlr.TokenEOF
	GramaticaParserPLUS     = 1
	GramaticaParserMINUS    = 2
	GramaticaParserTIMES    = 3
	GramaticaParserDIVISION = 4
	GramaticaParserENTERO   = 5
	GramaticaParserDECIMAL  = 6
	GramaticaParserWS       = 7
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start = 0
	GramaticaParserRULE_exp   = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IExpContext

	// SetInit sets the init rule contexts.
	SetInit(IExpContext)

	// GetE returns the e attribute.
	GetE() abs.Expression

	// SetE sets the e attribute.
	SetE(abs.Expression)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	e      abs.Expression
	init   IExpContext
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

func (s *StartContext) GetInit() IExpContext { return s.init }

func (s *StartContext) SetInit(v IExpContext) { s.init = v }

func (s *StartContext) GetE() abs.Expression { return s.e }

func (s *StartContext) SetE(v abs.Expression) { s.e = v }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *StartContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *GramaticaParser) Start() (localctx IStartContext) {
	this := p
	_ = this

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
		p.SetState(4)

		var _x = p.exp(0)

		localctx.(*StartContext).init = _x
	}
	{
		p.SetState(5)
		p.Match(GramaticaParserEOF)
	}

	localctx.(*StartContext).e = localctx.(*StartContext).GetInit().GetE()
	result := localctx.(*StartContext).e.Execute()
	fmt.Println(result.Value)

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_DECIMAL returns the _DECIMAL token.
	Get_DECIMAL() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_DECIMAL sets the _DECIMAL token.
	Set_DECIMAL(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpContext

	// GetRight returns the right rule contexts.
	GetRight() IExpContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpContext)

	// GetE returns the e attribute.
	GetE() abs.Expression

	// SetE sets the e attribute.
	SetE(abs.Expression)

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	e        abs.Expression
	left     IExpContext
	_ENTERO  antlr.Token
	_DECIMAL antlr.Token
	op       antlr.Token
	right    IExpContext
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *ExpContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *ExpContext) GetOp() antlr.Token { return s.op }

func (s *ExpContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *ExpContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *ExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpContext) GetLeft() IExpContext { return s.left }

func (s *ExpContext) GetRight() IExpContext { return s.right }

func (s *ExpContext) SetLeft(v IExpContext) { s.left = v }

func (s *ExpContext) SetRight(v IExpContext) { s.right = v }

func (s *ExpContext) GetE() abs.Expression { return s.e }

func (s *ExpContext) SetE(v abs.Expression) { s.e = v }

func (s *ExpContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GramaticaParserENTERO, 0)
}

func (s *ExpContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDECIMAL, 0)
}

func (s *ExpContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) TIMES() antlr.TerminalNode {
	return s.GetToken(GramaticaParserTIMES, 0)
}

func (s *ExpContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDIVISION, 0)
}

func (s *ExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPLUS, 0)
}

func (s *ExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMINUS, 0)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *GramaticaParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *GramaticaParser) exp(_p int) (localctx IExpContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, GramaticaParserRULE_exp, _p)
	var _la int

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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserENTERO:
		{
			p.SetState(9)

			var _m = p.Match(GramaticaParserENTERO)

			localctx.(*ExpContext)._ENTERO = _m
		}

		localctx.(*ExpContext).e = exp.NewLiteral(1, 1, abs.Int, (func() string {
			if localctx.(*ExpContext).Get_ENTERO() == nil {
				return ""
			} else {
				return localctx.(*ExpContext).Get_ENTERO().GetText()
			}
		}()))

	case GramaticaParserDECIMAL:
		{
			p.SetState(11)

			var _m = p.Match(GramaticaParserDECIMAL)

			localctx.(*ExpContext)._DECIMAL = _m
		}

		localctx.(*ExpContext).e = exp.NewLiteral(1, 1, abs.Double, (func() string {
			if localctx.(*ExpContext).Get_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*ExpContext).Get_DECIMAL().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(25)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp)
				p.SetState(15)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(16)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserTIMES || _la == GramaticaParserDIVISION) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(17)

					var _x = p.exp(5)

					localctx.(*ExpContext).right = _x
				}

				if (func() string {
					if localctx.(*ExpContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpContext).GetOp().GetText()
					}
				}()) == "*" {
					localctx.(*ExpContext).e = exp.NewArithmetic(1, 1, localctx.(*ExpContext).GetLeft().GetE(), localctx.(*ExpContext).GetRight().GetE(), exp.Multiplicacion)
				} else {
					localctx.(*ExpContext).e = exp.NewArithmetic(1, 1, localctx.(*ExpContext).GetLeft().GetE(), localctx.(*ExpContext).GetRight().GetE(), exp.Division)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(21)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserPLUS || _la == GramaticaParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(22)

					var _x = p.exp(4)

					localctx.(*ExpContext).right = _x
				}

				if (func() string {
					if localctx.(*ExpContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpContext).GetOp().GetText()
					}
				}()) == "+" {
					localctx.(*ExpContext).e = exp.NewArithmetic(1, 1, localctx.(*ExpContext).GetLeft().GetE(), localctx.(*ExpContext).GetRight().GetE(), exp.Suma)
				} else {
					localctx.(*ExpContext).e = exp.NewArithmetic(1, 1, localctx.(*ExpContext).GetLeft().GetE(), localctx.(*ExpContext).GetRight().GetE(), exp.Resta)
				}

			}

		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
