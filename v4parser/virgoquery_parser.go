// Generated from VirgoQuery.g4 by ANTLR 4.7.

package v4parser // VirgoQuery
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 81, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 26,
	10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 32, 10, 3, 12, 3, 14, 3, 35, 11, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 53, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 7, 7, 61, 10, 7, 12, 7, 14, 7, 64, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 5, 8, 72, 10, 8, 3, 8, 3, 8, 7, 8, 76, 10, 8, 12, 8, 14, 8,
	79, 11, 8, 3, 8, 2, 5, 4, 12, 14, 9, 2, 4, 6, 8, 10, 12, 14, 2, 3, 3, 2,
	7, 11, 2, 80, 2, 16, 3, 2, 2, 2, 4, 25, 3, 2, 2, 2, 6, 36, 3, 2, 2, 2,
	8, 42, 3, 2, 2, 2, 10, 44, 3, 2, 2, 2, 12, 52, 3, 2, 2, 2, 14, 71, 3, 2,
	2, 2, 16, 17, 5, 4, 3, 2, 17, 18, 7, 2, 2, 3, 18, 3, 3, 2, 2, 2, 19, 20,
	8, 3, 1, 2, 20, 21, 7, 3, 2, 2, 21, 22, 5, 4, 3, 2, 22, 23, 7, 4, 2, 2,
	23, 26, 3, 2, 2, 2, 24, 26, 5, 6, 4, 2, 25, 19, 3, 2, 2, 2, 25, 24, 3,
	2, 2, 2, 26, 33, 3, 2, 2, 2, 27, 28, 12, 5, 2, 2, 28, 29, 5, 10, 6, 2,
	29, 30, 5, 4, 3, 6, 30, 32, 3, 2, 2, 2, 31, 27, 3, 2, 2, 2, 32, 35, 3,
	2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35,
	33, 3, 2, 2, 2, 36, 37, 5, 8, 5, 2, 37, 38, 7, 6, 2, 2, 38, 39, 7, 12,
	2, 2, 39, 40, 5, 12, 7, 2, 40, 41, 7, 17, 2, 2, 41, 7, 3, 2, 2, 2, 42,
	43, 9, 2, 2, 2, 43, 9, 3, 2, 2, 2, 44, 45, 7, 5, 2, 2, 45, 11, 3, 2, 2,
	2, 46, 47, 8, 7, 1, 2, 47, 48, 7, 3, 2, 2, 48, 49, 5, 12, 7, 2, 49, 50,
	7, 4, 2, 2, 50, 53, 3, 2, 2, 2, 51, 53, 5, 14, 8, 2, 52, 46, 3, 2, 2, 2,
	52, 51, 3, 2, 2, 2, 53, 62, 3, 2, 2, 2, 54, 55, 12, 5, 2, 2, 55, 56, 5,
	10, 6, 2, 56, 57, 5, 12, 7, 6, 57, 61, 3, 2, 2, 2, 58, 59, 12, 4, 2, 2,
	59, 61, 5, 14, 8, 2, 60, 54, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 64, 3,
	2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 13, 3, 2, 2, 2, 64,
	62, 3, 2, 2, 2, 65, 66, 8, 8, 1, 2, 66, 67, 7, 14, 2, 2, 67, 68, 5, 14,
	8, 2, 68, 69, 7, 14, 2, 2, 69, 72, 3, 2, 2, 2, 70, 72, 7, 18, 2, 2, 71,
	65, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 77, 3, 2, 2, 2, 73, 74, 12, 4,
	2, 2, 74, 76, 7, 18, 2, 2, 75, 73, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77,
	75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 15, 3, 2, 2, 2, 79, 77, 3, 2, 2,
	2, 9, 25, 33, 52, 60, 62, 71, 77,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "':'", "'title'", "'author'", "'subject'", "'keyword'",
	"'identifier'", "'{'", "", "'\"'", "'['", "']'", "'}'",
}
var symbolicNames = []string{
	"", "LPAREN", "RPAREN", "BOOLEAN", "COLON", "TITLE", "AUTHOR", "SUBJECT",
	"KEYWORD", "IDENTIFIER", "LBRACE", "WS1", "QUOTE", "LBRACKET", "RBRACKET",
	"RBRACE", "SEARCH_WORD", "WS2",
}

var ruleNames = []string{
	"query", "query_parts", "field_query", "field_type", "boolean_op", "search_string",
	"search_part",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type VirgoQuery struct {
	*antlr.BaseParser
}

func NewVirgoQuery(input antlr.TokenStream) *VirgoQuery {
	this := new(VirgoQuery)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "VirgoQuery.g4"

	return this
}

// VirgoQuery tokens.
const (
	VirgoQueryEOF         = antlr.TokenEOF
	VirgoQueryLPAREN      = 1
	VirgoQueryRPAREN      = 2
	VirgoQueryBOOLEAN     = 3
	VirgoQueryCOLON       = 4
	VirgoQueryTITLE       = 5
	VirgoQueryAUTHOR      = 6
	VirgoQuerySUBJECT     = 7
	VirgoQueryKEYWORD     = 8
	VirgoQueryIDENTIFIER  = 9
	VirgoQueryLBRACE      = 10
	VirgoQueryWS1         = 11
	VirgoQueryQUOTE       = 12
	VirgoQueryLBRACKET    = 13
	VirgoQueryRBRACKET    = 14
	VirgoQueryRBRACE      = 15
	VirgoQuerySEARCH_WORD = 16
	VirgoQueryWS2         = 17
)

// VirgoQuery rules.
const (
	VirgoQueryRULE_query         = 0
	VirgoQueryRULE_query_parts   = 1
	VirgoQueryRULE_field_query   = 2
	VirgoQueryRULE_field_type    = 3
	VirgoQueryRULE_boolean_op    = 4
	VirgoQueryRULE_search_string = 5
	VirgoQueryRULE_search_part   = 6
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Query_parts() IQuery_partsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuery_partsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuery_partsContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(VirgoQueryEOF, 0)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VirgoQuery) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VirgoQueryRULE_query)

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
		p.SetState(14)
		p.query_parts(0)
	}
	{
		p.SetState(15)
		p.Match(VirgoQueryEOF)
	}

	return localctx
}

// IQuery_partsContext is an interface to support dynamic dispatch.
type IQuery_partsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuery_partsContext differentiates from other interfaces.
	IsQuery_partsContext()
}

type Query_partsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuery_partsContext() *Query_partsContext {
	var p = new(Query_partsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_query_parts
	return p
}

func (*Query_partsContext) IsQuery_partsContext() {}

func NewQuery_partsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Query_partsContext {
	var p = new(Query_partsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_query_parts

	return p
}

func (s *Query_partsContext) GetParser() antlr.Parser { return s.parser }

func (s *Query_partsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryLPAREN, 0)
}

func (s *Query_partsContext) AllQuery_parts() []IQuery_partsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuery_partsContext)(nil)).Elem())
	var tst = make([]IQuery_partsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuery_partsContext)
		}
	}

	return tst
}

func (s *Query_partsContext) Query_parts(i int) IQuery_partsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuery_partsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuery_partsContext)
}

func (s *Query_partsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryRPAREN, 0)
}

func (s *Query_partsContext) Field_query() IField_queryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_queryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_queryContext)
}

func (s *Query_partsContext) Boolean_op() IBoolean_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_opContext)
}

func (s *Query_partsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Query_partsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Query_partsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterQuery_parts(s)
	}
}

func (s *Query_partsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitQuery_parts(s)
	}
}

func (s *Query_partsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitQuery_parts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VirgoQuery) Query_parts() (localctx IQuery_partsContext) {
	return p.query_parts(0)
}

func (p *VirgoQuery) query_parts(_p int) (localctx IQuery_partsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQuery_partsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQuery_partsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, VirgoQueryRULE_query_parts, _p)

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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VirgoQueryLPAREN:
		{
			p.SetState(18)
			p.Match(VirgoQueryLPAREN)
		}
		{
			p.SetState(19)
			p.query_parts(0)
		}
		{
			p.SetState(20)
			p.Match(VirgoQueryRPAREN)
		}

	case VirgoQueryTITLE, VirgoQueryAUTHOR, VirgoQuerySUBJECT, VirgoQueryKEYWORD, VirgoQueryIDENTIFIER:
		{
			p.SetState(22)
			p.Field_query()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewQuery_partsContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_query_parts)
			p.SetState(25)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(26)
				p.Boolean_op()
			}
			{
				p.SetState(27)
				p.query_parts(4)
			}

		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IField_queryContext is an interface to support dynamic dispatch.
type IField_queryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_queryContext differentiates from other interfaces.
	IsField_queryContext()
}

type Field_queryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_queryContext() *Field_queryContext {
	var p = new(Field_queryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_field_query
	return p
}

func (*Field_queryContext) IsField_queryContext() {}

func NewField_queryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_queryContext {
	var p = new(Field_queryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_field_query

	return p
}

func (s *Field_queryContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_queryContext) Field_type() IField_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Field_queryContext) COLON() antlr.TerminalNode {
	return s.GetToken(VirgoQueryCOLON, 0)
}

func (s *Field_queryContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(VirgoQueryLBRACE, 0)
}

func (s *Field_queryContext) Search_string() ISearch_stringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISearch_stringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISearch_stringContext)
}

func (s *Field_queryContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VirgoQueryRBRACE, 0)
}

func (s *Field_queryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_queryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_queryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterField_query(s)
	}
}

func (s *Field_queryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitField_query(s)
	}
}

func (s *Field_queryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitField_query(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VirgoQuery) Field_query() (localctx IField_queryContext) {
	localctx = NewField_queryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VirgoQueryRULE_field_query)

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
		p.SetState(34)
		p.Field_type()
	}
	{
		p.SetState(35)
		p.Match(VirgoQueryCOLON)
	}
	{
		p.SetState(36)
		p.Match(VirgoQueryLBRACE)
	}
	{
		p.SetState(37)
		p.search_string(0)
	}
	{
		p.SetState(38)
		p.Match(VirgoQueryRBRACE)
	}

	return localctx
}

// IField_typeContext is an interface to support dynamic dispatch.
type IField_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_typeContext differentiates from other interfaces.
	IsField_typeContext()
}

type Field_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_typeContext() *Field_typeContext {
	var p = new(Field_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_field_type
	return p
}

func (*Field_typeContext) IsField_typeContext() {}

func NewField_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_typeContext {
	var p = new(Field_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_field_type

	return p
}

func (s *Field_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_typeContext) TITLE() antlr.TerminalNode {
	return s.GetToken(VirgoQueryTITLE, 0)
}

func (s *Field_typeContext) AUTHOR() antlr.TerminalNode {
	return s.GetToken(VirgoQueryAUTHOR, 0)
}

func (s *Field_typeContext) SUBJECT() antlr.TerminalNode {
	return s.GetToken(VirgoQuerySUBJECT, 0)
}

func (s *Field_typeContext) KEYWORD() antlr.TerminalNode {
	return s.GetToken(VirgoQueryKEYWORD, 0)
}

func (s *Field_typeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VirgoQueryIDENTIFIER, 0)
}

func (s *Field_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterField_type(s)
	}
}

func (s *Field_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitField_type(s)
	}
}

func (s *Field_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitField_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VirgoQuery) Field_type() (localctx IField_typeContext) {
	localctx = NewField_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VirgoQueryRULE_field_type)
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
	p.SetState(40)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<VirgoQueryTITLE)|(1<<VirgoQueryAUTHOR)|(1<<VirgoQuerySUBJECT)|(1<<VirgoQueryKEYWORD)|(1<<VirgoQueryIDENTIFIER))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IBoolean_opContext is an interface to support dynamic dispatch.
type IBoolean_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolean_opContext differentiates from other interfaces.
	IsBoolean_opContext()
}

type Boolean_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_opContext() *Boolean_opContext {
	var p = new(Boolean_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_boolean_op
	return p
}

func (*Boolean_opContext) IsBoolean_opContext() {}

func NewBoolean_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_opContext {
	var p = new(Boolean_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_boolean_op

	return p
}

func (s *Boolean_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_opContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryBOOLEAN, 0)
}

func (s *Boolean_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterBoolean_op(s)
	}
}

func (s *Boolean_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitBoolean_op(s)
	}
}

func (s *Boolean_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitBoolean_op(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VirgoQuery) Boolean_op() (localctx IBoolean_opContext) {
	localctx = NewBoolean_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VirgoQueryRULE_boolean_op)

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
		p.SetState(42)
		p.Match(VirgoQueryBOOLEAN)
	}

	return localctx
}

// ISearch_stringContext is an interface to support dynamic dispatch.
type ISearch_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSearch_stringContext differentiates from other interfaces.
	IsSearch_stringContext()
}

type Search_stringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearch_stringContext() *Search_stringContext {
	var p = new(Search_stringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_search_string
	return p
}

func (*Search_stringContext) IsSearch_stringContext() {}

func NewSearch_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Search_stringContext {
	var p = new(Search_stringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_search_string

	return p
}

func (s *Search_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Search_stringContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryLPAREN, 0)
}

func (s *Search_stringContext) AllSearch_string() []ISearch_stringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISearch_stringContext)(nil)).Elem())
	var tst = make([]ISearch_stringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISearch_stringContext)
		}
	}

	return tst
}

func (s *Search_stringContext) Search_string(i int) ISearch_stringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISearch_stringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISearch_stringContext)
}

func (s *Search_stringContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryRPAREN, 0)
}

func (s *Search_stringContext) Search_part() ISearch_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISearch_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISearch_partContext)
}

func (s *Search_stringContext) Boolean_op() IBoolean_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_opContext)
}

func (s *Search_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Search_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Search_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterSearch_string(s)
	}
}

func (s *Search_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitSearch_string(s)
	}
}

func (s *Search_stringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitSearch_string(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VirgoQuery) Search_string() (localctx ISearch_stringContext) {
	return p.search_string(0)
}

func (p *VirgoQuery) search_string(_p int) (localctx ISearch_stringContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSearch_stringContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISearch_stringContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, VirgoQueryRULE_search_string, _p)

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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VirgoQueryLPAREN:
		{
			p.SetState(45)
			p.Match(VirgoQueryLPAREN)
		}
		{
			p.SetState(46)
			p.search_string(0)
		}
		{
			p.SetState(47)
			p.Match(VirgoQueryRPAREN)
		}

	case VirgoQueryQUOTE, VirgoQuerySEARCH_WORD:
		{
			p.SetState(49)
			p.search_part(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSearch_stringContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_search_string)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(53)
					p.Boolean_op()
				}
				{
					p.SetState(54)
					p.search_string(4)
				}

			case 2:
				localctx = NewSearch_stringContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_search_string)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(57)
					p.search_part(0)
				}

			}

		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// ISearch_partContext is an interface to support dynamic dispatch.
type ISearch_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSearch_partContext differentiates from other interfaces.
	IsSearch_partContext()
}

type Search_partContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearch_partContext() *Search_partContext {
	var p = new(Search_partContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_search_part
	return p
}

func (*Search_partContext) IsSearch_partContext() {}

func NewSearch_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Search_partContext {
	var p = new(Search_partContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_search_part

	return p
}

func (s *Search_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Search_partContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(VirgoQueryQUOTE)
}

func (s *Search_partContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(VirgoQueryQUOTE, i)
}

func (s *Search_partContext) Search_part() ISearch_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISearch_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISearch_partContext)
}

func (s *Search_partContext) SEARCH_WORD() antlr.TerminalNode {
	return s.GetToken(VirgoQuerySEARCH_WORD, 0)
}

func (s *Search_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Search_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Search_partContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterSearch_part(s)
	}
}

func (s *Search_partContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitSearch_part(s)
	}
}

func (s *Search_partContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitSearch_part(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VirgoQuery) Search_part() (localctx ISearch_partContext) {
	return p.search_part(0)
}

func (p *VirgoQuery) search_part(_p int) (localctx ISearch_partContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSearch_partContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISearch_partContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, VirgoQueryRULE_search_part, _p)

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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VirgoQueryQUOTE:
		{
			p.SetState(64)
			p.Match(VirgoQueryQUOTE)
		}
		{
			p.SetState(65)
			p.search_part(0)
		}
		{
			p.SetState(66)
			p.Match(VirgoQueryQUOTE)
		}

	case VirgoQuerySEARCH_WORD:
		{
			p.SetState(68)
			p.Match(VirgoQuerySEARCH_WORD)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSearch_partContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_search_part)
			p.SetState(71)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(72)
				p.Match(VirgoQuerySEARCH_WORD)
			}

		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

func (p *VirgoQuery) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *Query_partsContext = nil
		if localctx != nil {
			t = localctx.(*Query_partsContext)
		}
		return p.Query_parts_Sempred(t, predIndex)

	case 5:
		var t *Search_stringContext = nil
		if localctx != nil {
			t = localctx.(*Search_stringContext)
		}
		return p.Search_string_Sempred(t, predIndex)

	case 6:
		var t *Search_partContext = nil
		if localctx != nil {
			t = localctx.(*Search_partContext)
		}
		return p.Search_part_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VirgoQuery) Query_parts_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *VirgoQuery) Search_string_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *VirgoQuery) Search_part_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
