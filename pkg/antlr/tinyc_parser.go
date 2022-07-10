// Code generated from tinyc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // tinyc

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type tinycParser struct {
	*antlr.BaseParser
}

var tinycParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tinycParserInit() {
	staticData := &tinycParserStaticData
	staticData.literalNames = []string{
		"", "'if'", "'else'", "'while'", "'do'", "';'", "'{'", "'}'", "'('",
		"')'", "'='", "'<'", "", "", "", "'+'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "STRING", "INT", "WS",
		"PLUS", "MINUS",
	}
	staticData.ruleNames = []string{
		"program", "statement", "paren_expr", "expr", "test", "sum_", "term",
		"id_", "integer",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 99, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 4, 0, 20, 8, 0,
		11, 0, 12, 0, 21, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 5, 1, 46, 8, 1, 10, 1, 12, 1, 49, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 56, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 67, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 74, 8, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 85, 8, 5, 10, 5, 12,
		5, 88, 9, 5, 1, 6, 1, 6, 1, 6, 3, 6, 93, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 0, 1, 10, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 0, 103, 0, 19, 1,
		0, 0, 0, 2, 55, 1, 0, 0, 0, 4, 57, 1, 0, 0, 0, 6, 66, 1, 0, 0, 0, 8, 73,
		1, 0, 0, 0, 10, 75, 1, 0, 0, 0, 12, 92, 1, 0, 0, 0, 14, 94, 1, 0, 0, 0,
		16, 96, 1, 0, 0, 0, 18, 20, 3, 2, 1, 0, 19, 18, 1, 0, 0, 0, 20, 21, 1,
		0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 1, 1, 0, 0, 0, 23,
		24, 5, 1, 0, 0, 24, 25, 3, 4, 2, 0, 25, 26, 3, 2, 1, 0, 26, 56, 1, 0, 0,
		0, 27, 28, 5, 1, 0, 0, 28, 29, 3, 4, 2, 0, 29, 30, 3, 2, 1, 0, 30, 31,
		5, 2, 0, 0, 31, 32, 3, 2, 1, 0, 32, 56, 1, 0, 0, 0, 33, 34, 5, 3, 0, 0,
		34, 35, 3, 4, 2, 0, 35, 36, 3, 2, 1, 0, 36, 56, 1, 0, 0, 0, 37, 38, 5,
		4, 0, 0, 38, 39, 3, 2, 1, 0, 39, 40, 5, 3, 0, 0, 40, 41, 3, 4, 2, 0, 41,
		42, 5, 5, 0, 0, 42, 56, 1, 0, 0, 0, 43, 47, 5, 6, 0, 0, 44, 46, 3, 2, 1,
		0, 45, 44, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48,
		1, 0, 0, 0, 48, 50, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 56, 5, 7, 0, 0,
		51, 52, 3, 6, 3, 0, 52, 53, 5, 5, 0, 0, 53, 56, 1, 0, 0, 0, 54, 56, 5,
		5, 0, 0, 55, 23, 1, 0, 0, 0, 55, 27, 1, 0, 0, 0, 55, 33, 1, 0, 0, 0, 55,
		37, 1, 0, 0, 0, 55, 43, 1, 0, 0, 0, 55, 51, 1, 0, 0, 0, 55, 54, 1, 0, 0,
		0, 56, 3, 1, 0, 0, 0, 57, 58, 5, 8, 0, 0, 58, 59, 3, 6, 3, 0, 59, 60, 5,
		9, 0, 0, 60, 5, 1, 0, 0, 0, 61, 67, 3, 8, 4, 0, 62, 63, 3, 14, 7, 0, 63,
		64, 5, 10, 0, 0, 64, 65, 3, 6, 3, 0, 65, 67, 1, 0, 0, 0, 66, 61, 1, 0,
		0, 0, 66, 62, 1, 0, 0, 0, 67, 7, 1, 0, 0, 0, 68, 74, 3, 10, 5, 0, 69, 70,
		3, 10, 5, 0, 70, 71, 5, 11, 0, 0, 71, 72, 3, 10, 5, 0, 72, 74, 1, 0, 0,
		0, 73, 68, 1, 0, 0, 0, 73, 69, 1, 0, 0, 0, 74, 9, 1, 0, 0, 0, 75, 76, 6,
		5, -1, 0, 76, 77, 3, 12, 6, 0, 77, 86, 1, 0, 0, 0, 78, 79, 10, 2, 0, 0,
		79, 80, 5, 15, 0, 0, 80, 85, 3, 12, 6, 0, 81, 82, 10, 1, 0, 0, 82, 83,
		5, 16, 0, 0, 83, 85, 3, 12, 6, 0, 84, 78, 1, 0, 0, 0, 84, 81, 1, 0, 0,
		0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 11,
		1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 93, 3, 14, 7, 0, 90, 93, 3, 16, 8,
		0, 91, 93, 3, 4, 2, 0, 92, 89, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 91,
		1, 0, 0, 0, 93, 13, 1, 0, 0, 0, 94, 95, 5, 12, 0, 0, 95, 15, 1, 0, 0, 0,
		96, 97, 5, 13, 0, 0, 97, 17, 1, 0, 0, 0, 8, 21, 47, 55, 66, 73, 84, 86,
		92,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// tinycParserInit initializes any static state used to implement tinycParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewtinycParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TinycParserInit() {
	staticData := &tinycParserStaticData
	staticData.once.Do(tinycParserInit)
}

// NewtinycParser produces a new parser instance for the optional input antlr.TokenStream.
func NewtinycParser(input antlr.TokenStream) *tinycParser {
	TinycParserInit()
	this := new(tinycParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tinycParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "tinyc.g4"

	return this
}

// tinycParser tokens.
const (
	tinycParserEOF    = antlr.TokenEOF
	tinycParserT__0   = 1
	tinycParserT__1   = 2
	tinycParserT__2   = 3
	tinycParserT__3   = 4
	tinycParserT__4   = 5
	tinycParserT__5   = 6
	tinycParserT__6   = 7
	tinycParserT__7   = 8
	tinycParserT__8   = 9
	tinycParserT__9   = 10
	tinycParserT__10  = 11
	tinycParserSTRING = 12
	tinycParserINT    = 13
	tinycParserWS     = 14
	tinycParserPLUS   = 15
	tinycParserMINUS  = 16
)

// tinycParser rules.
const (
	tinycParserRULE_program    = 0
	tinycParserRULE_statement  = 1
	tinycParserRULE_paren_expr = 2
	tinycParserRULE_expr       = 3
	tinycParserRULE_test       = 4
	tinycParserRULE_sum_       = 5
	tinycParserRULE_term       = 6
	tinycParserRULE_id_        = 7
	tinycParserRULE_integer    = 8
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinycParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tinycParserRULE_program)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tinycParserT__0)|(1<<tinycParserT__2)|(1<<tinycParserT__3)|(1<<tinycParserT__4)|(1<<tinycParserT__5)|(1<<tinycParserT__7)|(1<<tinycParserSTRING)|(1<<tinycParserINT))) != 0) {
		{
			p.SetState(18)
			p.Statement()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = tinycParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Paren_expr() IParen_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParen_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParen_exprContext)
}

func (s *StatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tinycParserRULE_statement)
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

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Match(tinycParserT__0)
		}
		{
			p.SetState(24)
			p.Paren_expr()
		}
		{
			p.SetState(25)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.Match(tinycParserT__0)
		}
		{
			p.SetState(28)
			p.Paren_expr()
		}
		{
			p.SetState(29)
			p.Statement()
		}
		{
			p.SetState(30)
			p.Match(tinycParserT__1)
		}
		{
			p.SetState(31)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.Match(tinycParserT__2)
		}
		{
			p.SetState(34)
			p.Paren_expr()
		}
		{
			p.SetState(35)
			p.Statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(37)
			p.Match(tinycParserT__3)
		}
		{
			p.SetState(38)
			p.Statement()
		}
		{
			p.SetState(39)
			p.Match(tinycParserT__2)
		}
		{
			p.SetState(40)
			p.Paren_expr()
		}
		{
			p.SetState(41)
			p.Match(tinycParserT__4)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(43)
			p.Match(tinycParserT__5)
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tinycParserT__0)|(1<<tinycParserT__2)|(1<<tinycParserT__3)|(1<<tinycParserT__4)|(1<<tinycParserT__5)|(1<<tinycParserT__7)|(1<<tinycParserSTRING)|(1<<tinycParserINT))) != 0 {
			{
				p.SetState(44)
				p.Statement()
			}

			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(50)
			p.Match(tinycParserT__6)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(51)
			p.Expr()
		}
		{
			p.SetState(52)
			p.Match(tinycParserT__4)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(54)
			p.Match(tinycParserT__4)
		}

	}

	return localctx
}

// IParen_exprContext is an interface to support dynamic dispatch.
type IParen_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParen_exprContext differentiates from other interfaces.
	IsParen_exprContext()
}

type Paren_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParen_exprContext() *Paren_exprContext {
	var p = new(Paren_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinycParserRULE_paren_expr
	return p
}

func (*Paren_exprContext) IsParen_exprContext() {}

func NewParen_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Paren_exprContext {
	var p = new(Paren_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_paren_expr

	return p
}

func (s *Paren_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Paren_exprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Paren_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Paren_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Paren_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterParen_expr(s)
	}
}

func (s *Paren_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitParen_expr(s)
	}
}

func (s *Paren_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitParen_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Paren_expr() (localctx IParen_exprContext) {
	this := p
	_ = this

	localctx = NewParen_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tinycParserRULE_paren_expr)

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
		p.SetState(57)
		p.Match(tinycParserT__7)
	}
	{
		p.SetState(58)
		p.Expr()
	}
	{
		p.SetState(59)
		p.Match(tinycParserT__8)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinycParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Test() ITestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *ExprContext) Id_() IId_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *ExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tinycParserRULE_expr)

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

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Test()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Id_()
		}
		{
			p.SetState(63)
			p.Match(tinycParserT__9)
		}
		{
			p.SetState(64)
			p.Expr()
		}

	}

	return localctx
}

// ITestContext is an interface to support dynamic dispatch.
type ITestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestContext differentiates from other interfaces.
	IsTestContext()
}

type TestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestContext() *TestContext {
	var p = new(TestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinycParserRULE_test
	return p
}

func (*TestContext) IsTestContext() {}

func NewTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestContext {
	var p = new(TestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_test

	return p
}

func (s *TestContext) GetParser() antlr.Parser { return s.parser }

func (s *TestContext) AllSum_() []ISum_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISum_Context); ok {
			len++
		}
	}

	tst := make([]ISum_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISum_Context); ok {
			tst[i] = t.(ISum_Context)
			i++
		}
	}

	return tst
}

func (s *TestContext) Sum_(i int) ISum_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISum_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISum_Context)
}

func (s *TestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterTest(s)
	}
}

func (s *TestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitTest(s)
	}
}

func (s *TestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitTest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Test() (localctx ITestContext) {
	this := p
	_ = this

	localctx = NewTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tinycParserRULE_test)

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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.sum_(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.sum_(0)
		}
		{
			p.SetState(70)
			p.Match(tinycParserT__10)
		}
		{
			p.SetState(71)
			p.sum_(0)
		}

	}

	return localctx
}

// ISum_Context is an interface to support dynamic dispatch.
type ISum_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSum_Context differentiates from other interfaces.
	IsSum_Context()
}

type Sum_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySum_Context() *Sum_Context {
	var p = new(Sum_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinycParserRULE_sum_
	return p
}

func (*Sum_Context) IsSum_Context() {}

func NewSum_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sum_Context {
	var p = new(Sum_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_sum_

	return p
}

func (s *Sum_Context) GetParser() antlr.Parser { return s.parser }

func (s *Sum_Context) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Sum_Context) Sum_() ISum_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISum_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISum_Context)
}

func (s *Sum_Context) PLUS() antlr.TerminalNode {
	return s.GetToken(tinycParserPLUS, 0)
}

func (s *Sum_Context) MINUS() antlr.TerminalNode {
	return s.GetToken(tinycParserMINUS, 0)
}

func (s *Sum_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sum_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sum_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterSum_(s)
	}
}

func (s *Sum_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitSum_(s)
	}
}

func (s *Sum_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitSum_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Sum_() (localctx ISum_Context) {
	return p.sum_(0)
}

func (p *tinycParser) sum_(_p int) (localctx ISum_Context) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSum_Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISum_Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, tinycParserRULE_sum_, _p)

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
		p.SetState(76)
		p.Term()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSum_Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_sum_)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(79)
					p.Match(tinycParserPLUS)
				}
				{
					p.SetState(80)
					p.Term()
				}

			case 2:
				localctx = NewSum_Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tinycParserRULE_sum_)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(82)
					p.Match(tinycParserMINUS)
				}
				{
					p.SetState(83)
					p.Term()
				}

			}

		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinycParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Id_() IId_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *TermContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *TermContext) Paren_expr() IParen_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParen_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParen_exprContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tinycParserRULE_term)

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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tinycParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Id_()
		}

	case tinycParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Integer()
		}

	case tinycParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Paren_expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IId_Context is an interface to support dynamic dispatch.
type IId_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsId_Context differentiates from other interfaces.
	IsId_Context()
}

type Id_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_Context() *Id_Context {
	var p = new(Id_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinycParserRULE_id_
	return p
}

func (*Id_Context) IsId_Context() {}

func NewId_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_Context {
	var p = new(Id_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_id_

	return p
}

func (s *Id_Context) GetParser() antlr.Parser { return s.parser }

func (s *Id_Context) STRING() antlr.TerminalNode {
	return s.GetToken(tinycParserSTRING, 0)
}

func (s *Id_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterId_(s)
	}
}

func (s *Id_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitId_(s)
	}
}

func (s *Id_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitId_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Id_() (localctx IId_Context) {
	this := p
	_ = this

	localctx = NewId_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tinycParserRULE_id_)

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
		p.SetState(94)
		p.Match(tinycParserSTRING)
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinycParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinycParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(tinycParserINT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinycListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tinycVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tinycParser) Integer() (localctx IIntegerContext) {
	this := p
	_ = this

	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tinycParserRULE_integer)

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
		p.SetState(96)
		p.Match(tinycParserINT)
	}

	return localctx
}

func (p *tinycParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *Sum_Context = nil
		if localctx != nil {
			t = localctx.(*Sum_Context)
		}
		return p.Sum__Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *tinycParser) Sum__Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
