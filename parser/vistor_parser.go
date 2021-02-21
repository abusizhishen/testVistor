// Code generated from Vistor.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Vistor

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 80, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 25, 10,
	3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 33, 10, 5, 12, 5, 14, 5, 36,
	11, 5, 3, 5, 3, 5, 7, 5, 40, 10, 5, 12, 5, 14, 5, 43, 11, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 51, 10, 6, 12, 6, 14, 6, 54, 11, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 61, 10, 7, 12, 7, 14, 7, 64, 11, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 5, 8, 70, 10, 8, 3, 9, 7, 9, 73, 10, 9, 12, 9, 14, 9,
	76, 11, 9, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 4,
	3, 2, 3, 6, 3, 2, 8, 9, 2, 78, 2, 18, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6,
	26, 3, 2, 2, 2, 8, 28, 3, 2, 2, 2, 10, 46, 3, 2, 2, 2, 12, 57, 3, 2, 2,
	2, 14, 69, 3, 2, 2, 2, 16, 74, 3, 2, 2, 2, 18, 19, 7, 12, 2, 2, 19, 3,
	3, 2, 2, 2, 20, 21, 5, 2, 2, 2, 21, 22, 9, 2, 2, 2, 22, 24, 5, 2, 2, 2,
	23, 25, 7, 7, 2, 2, 24, 23, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 5, 3, 2,
	2, 2, 26, 27, 9, 3, 2, 2, 27, 7, 3, 2, 2, 2, 28, 29, 7, 14, 2, 2, 29, 30,
	5, 6, 4, 2, 30, 34, 7, 10, 2, 2, 31, 33, 5, 14, 8, 2, 32, 31, 3, 2, 2,
	2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 37,
	3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 41, 7, 11, 2, 2, 38, 40, 5, 10, 6,
	2, 39, 38, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42,
	3, 2, 2, 2, 42, 44, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 45, 5, 12, 7, 2,
	45, 9, 3, 2, 2, 2, 46, 47, 7, 15, 2, 2, 47, 48, 5, 6, 4, 2, 48, 52, 7,
	10, 2, 2, 49, 51, 5, 14, 8, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2,
	52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 52, 3,
	2, 2, 2, 55, 56, 7, 11, 2, 2, 56, 11, 3, 2, 2, 2, 57, 58, 7, 16, 2, 2,
	58, 62, 7, 10, 2, 2, 59, 61, 5, 14, 8, 2, 60, 59, 3, 2, 2, 2, 61, 64, 3,
	2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64,
	62, 3, 2, 2, 2, 65, 66, 7, 11, 2, 2, 66, 13, 3, 2, 2, 2, 67, 70, 5, 8,
	5, 2, 68, 70, 5, 4, 3, 2, 69, 67, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 15,
	3, 2, 2, 2, 71, 73, 5, 14, 8, 2, 72, 71, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2,
	74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 77, 3, 2, 2, 2, 76, 74, 3,
	2, 2, 2, 77, 78, 7, 2, 2, 3, 78, 17, 3, 2, 2, 2, 9, 24, 34, 41, 52, 62,
	69, 74,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "';'", "'true'", "'false'", "'{'", "'}'",
	"", "", "'if'", "'elsif'", "'else'", "'end'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "NUM", "WS", "IF", "ELSIF", "ELSE",
	"END",
}

var ruleNames = []string{
	"num", "calculate", "condition", "ifStatement", "elseIfStatement", "elseStatement",
	"statement", "start",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type VistorParser struct {
	*antlr.BaseParser
}

func NewVistorParser(input antlr.TokenStream) *VistorParser {
	this := new(VistorParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Vistor.g4"

	return this
}

// VistorParser tokens.
const (
	VistorParserEOF   = antlr.TokenEOF
	VistorParserT__0  = 1
	VistorParserT__1  = 2
	VistorParserT__2  = 3
	VistorParserT__3  = 4
	VistorParserT__4  = 5
	VistorParserT__5  = 6
	VistorParserT__6  = 7
	VistorParserT__7  = 8
	VistorParserT__8  = 9
	VistorParserNUM   = 10
	VistorParserWS    = 11
	VistorParserIF    = 12
	VistorParserELSIF = 13
	VistorParserELSE  = 14
	VistorParserEND   = 15
)

// VistorParser rules.
const (
	VistorParserRULE_num             = 0
	VistorParserRULE_calculate       = 1
	VistorParserRULE_condition       = 2
	VistorParserRULE_ifStatement     = 3
	VistorParserRULE_elseIfStatement = 4
	VistorParserRULE_elseStatement   = 5
	VistorParserRULE_statement       = 6
	VistorParserRULE_start           = 7
)

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VistorParserRULE_num
	return p
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VistorParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) NUM() antlr.TerminalNode {
	return s.GetToken(VistorParserNUM, 0)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.ExitNum(s)
	}
}

func (s *NumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VistorVisitor:
		return t.VisitNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VistorParser) Num() (localctx INumContext) {
	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VistorParserRULE_num)

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
		p.SetState(16)
		p.Match(VistorParserNUM)
	}

	return localctx
}

// ICalculateContext is an interface to support dynamic dispatch.
type ICalculateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsCalculateContext differentiates from other interfaces.
	IsCalculateContext()
}

type CalculateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyCalculateContext() *CalculateContext {
	var p = new(CalculateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VistorParserRULE_calculate
	return p
}

func (*CalculateContext) IsCalculateContext() {}

func NewCalculateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalculateContext {
	var p = new(CalculateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VistorParserRULE_calculate

	return p
}

func (s *CalculateContext) GetParser() antlr.Parser { return s.parser }

func (s *CalculateContext) GetOp() antlr.Token { return s.op }

func (s *CalculateContext) SetOp(v antlr.Token) { s.op = v }

func (s *CalculateContext) AllNum() []INumContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumContext)(nil)).Elem())
	var tst = make([]INumContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumContext)
		}
	}

	return tst
}

func (s *CalculateContext) Num(i int) INumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *CalculateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalculateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.EnterCalculate(s)
	}
}

func (s *CalculateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.ExitCalculate(s)
	}
}

func (s *CalculateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VistorVisitor:
		return t.VisitCalculate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VistorParser) Calculate() (localctx ICalculateContext) {
	localctx = NewCalculateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VistorParserRULE_calculate)
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
	{
		p.SetState(18)
		p.Num()
	}
	{
		p.SetState(19)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CalculateContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<VistorParserT__0)|(1<<VistorParserT__1)|(1<<VistorParserT__2)|(1<<VistorParserT__3))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CalculateContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(20)
		p.Num()
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == VistorParserT__4 {
		{
			p.SetState(21)
			p.Match(VistorParserT__4)
		}

	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VistorParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VistorParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }
func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VistorVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VistorParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VistorParserRULE_condition)
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
	{
		p.SetState(24)
		_la = p.GetTokenStream().LA(1)

		if !(_la == VistorParserT__5 || _la == VistorParserT__6) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VistorParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VistorParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(VistorParserIF, 0)
}

func (s *IfStatementContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfStatementContext) ElseStatement() IElseStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStatementContext)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) AllElseIfStatement() []IElseIfStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfStatementContext)(nil)).Elem())
	var tst = make([]IElseIfStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) ElseIfStatement(i int) IElseIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VistorVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VistorParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VistorParserRULE_ifStatement)
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
	{
		p.SetState(26)
		p.Match(VistorParserIF)
	}
	{
		p.SetState(27)
		p.Condition()
	}
	{
		p.SetState(28)
		p.Match(VistorParserT__7)
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VistorParserNUM || _la == VistorParserIF {
		{
			p.SetState(29)
			p.Statement()
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
		p.Match(VistorParserT__8)
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VistorParserELSIF {
		{
			p.SetState(36)
			p.ElseIfStatement()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
		p.ElseStatement()
	}

	return localctx
}

// IElseIfStatementContext is an interface to support dynamic dispatch.
type IElseIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfStatementContext differentiates from other interfaces.
	IsElseIfStatementContext()
}

type ElseIfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStatementContext() *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VistorParserRULE_elseIfStatement
	return p
}

func (*ElseIfStatementContext) IsElseIfStatementContext() {}

func NewElseIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VistorParserRULE_elseIfStatement

	return p
}

func (s *ElseIfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStatementContext) ELSIF() antlr.TerminalNode {
	return s.GetToken(VistorParserELSIF, 0)
}

func (s *ElseIfStatementContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ElseIfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ElseIfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ElseIfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.EnterElseIfStatement(s)
	}
}

func (s *ElseIfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.ExitElseIfStatement(s)
	}
}

func (s *ElseIfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VistorVisitor:
		return t.VisitElseIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VistorParser) ElseIfStatement() (localctx IElseIfStatementContext) {
	localctx = NewElseIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VistorParserRULE_elseIfStatement)
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
	{
		p.SetState(44)
		p.Match(VistorParserELSIF)
	}
	{
		p.SetState(45)
		p.Condition()
	}
	{
		p.SetState(46)
		p.Match(VistorParserT__7)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VistorParserNUM || _la == VistorParserIF {
		{
			p.SetState(47)
			p.Statement()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(VistorParserT__8)
	}

	return localctx
}

// IElseStatementContext is an interface to support dynamic dispatch.
type IElseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatementContext differentiates from other interfaces.
	IsElseStatementContext()
}

type ElseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatementContext() *ElseStatementContext {
	var p = new(ElseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VistorParserRULE_elseStatement
	return p
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VistorParserRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(VistorParserELSE, 0)
}

func (s *ElseStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ElseStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.EnterElseStatement(s)
	}
}

func (s *ElseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.ExitElseStatement(s)
	}
}

func (s *ElseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VistorVisitor:
		return t.VisitElseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VistorParser) ElseStatement() (localctx IElseStatementContext) {
	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VistorParserRULE_elseStatement)
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
	{
		p.SetState(55)
		p.Match(VistorParserELSE)
	}
	{
		p.SetState(56)
		p.Match(VistorParserT__7)
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VistorParserNUM || _la == VistorParserIF {
		{
			p.SetState(57)
			p.Statement()
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.Match(VistorParserT__8)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VistorParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VistorParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) Calculate() ICalculateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VistorVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VistorParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VistorParserRULE_statement)

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
	case VistorParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.IfStatement()
		}

	case VistorParserNUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Calculate()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VistorParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VistorParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(VistorParserEOF, 0)
}

func (s *StartContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StartContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VistorListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VistorVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VistorParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VistorParserRULE_start)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == VistorParserNUM || _la == VistorParserIF {
		{
			p.SetState(69)
			p.Statement()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Match(VistorParserEOF)
	}

	return localctx
}
