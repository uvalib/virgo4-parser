// Code generated from VirgoQuery.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 136, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 34, 10, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 7, 3, 40, 10, 3, 12, 3, 14, 3, 43, 11, 3, 3, 4, 3, 4, 
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 
	3, 4, 3, 4, 3, 4, 5, 4, 62, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 79, 10, 8, 
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 
	10, 3, 10, 5, 10, 93, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 
	7, 10, 101, 10, 10, 12, 10, 14, 10, 104, 11, 10, 3, 11, 3, 11, 3, 11, 3, 
	11, 3, 11, 7, 11, 111, 10, 11, 12, 11, 14, 11, 114, 11, 11, 3, 12, 3, 12, 
	3, 12, 3, 12, 3, 12, 5, 12, 121, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 
	12, 3, 12, 3, 12, 3, 12, 7, 12, 131, 10, 12, 12, 12, 14, 12, 134, 11, 12, 
	3, 12, 2, 6, 4, 18, 20, 22, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 
	2, 3, 3, 2, 7, 11, 2, 143, 2, 24, 3, 2, 2, 2, 4, 33, 3, 2, 2, 2, 6, 61, 
	3, 2, 2, 2, 8, 63, 3, 2, 2, 2, 10, 65, 3, 2, 2, 2, 12, 67, 3, 2, 2, 2, 
	14, 78, 3, 2, 2, 2, 16, 80, 3, 2, 2, 2, 18, 92, 3, 2, 2, 2, 20, 105, 3, 
	2, 2, 2, 22, 120, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26, 7, 2, 2, 3, 26, 
	3, 3, 2, 2, 2, 27, 28, 8, 3, 1, 2, 28, 29, 7, 3, 2, 2, 29, 30, 5, 4, 3, 
	2, 30, 31, 7, 4, 2, 2, 31, 34, 3, 2, 2, 2, 32, 34, 5, 6, 4, 2, 33, 27, 
	3, 2, 2, 2, 33, 32, 3, 2, 2, 2, 34, 41, 3, 2, 2, 2, 35, 36, 12, 5, 2, 2, 
	36, 37, 5, 12, 7, 2, 37, 38, 5, 4, 3, 6, 38, 40, 3, 2, 2, 2, 39, 35, 3, 
	2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 
	5, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 45, 5, 8, 5, 2, 45, 46, 7, 6, 2, 
	2, 46, 47, 7, 13, 2, 2, 47, 48, 5, 18, 10, 2, 48, 49, 7, 19, 2, 2, 49, 
	62, 3, 2, 2, 2, 50, 51, 5, 8, 5, 2, 51, 52, 7, 6, 2, 2, 52, 53, 7, 13, 
	2, 2, 53, 54, 7, 19, 2, 2, 54, 62, 3, 2, 2, 2, 55, 56, 5, 10, 6, 2, 56, 
	57, 7, 6, 2, 2, 57, 58, 7, 13, 2, 2, 58, 59, 5, 14, 8, 2, 59, 60, 7, 19, 
	2, 2, 60, 62, 3, 2, 2, 2, 61, 44, 3, 2, 2, 2, 61, 50, 3, 2, 2, 2, 61, 55, 
	3, 2, 2, 2, 62, 7, 3, 2, 2, 2, 63, 64, 9, 2, 2, 2, 64, 9, 3, 2, 2, 2, 65, 
	66, 7, 12, 2, 2, 66, 11, 3, 2, 2, 2, 67, 68, 7, 5, 2, 2, 68, 13, 3, 2, 
	2, 2, 69, 70, 5, 16, 9, 2, 70, 71, 7, 22, 2, 2, 71, 72, 5, 16, 9, 2, 72, 
	79, 3, 2, 2, 2, 73, 74, 7, 24, 2, 2, 74, 79, 5, 16, 9, 2, 75, 76, 7, 23, 
	2, 2, 76, 79, 5, 16, 9, 2, 77, 79, 5, 16, 9, 2, 78, 69, 3, 2, 2, 2, 78, 
	73, 3, 2, 2, 2, 78, 75, 3, 2, 2, 2, 78, 77, 3, 2, 2, 2, 79, 15, 3, 2, 2, 
	2, 80, 81, 7, 25, 2, 2, 81, 17, 3, 2, 2, 2, 82, 83, 8, 10, 1, 2, 83, 84, 
	7, 16, 2, 2, 84, 85, 5, 22, 12, 2, 85, 86, 7, 16, 2, 2, 86, 93, 3, 2, 2, 
	2, 87, 88, 7, 3, 2, 2, 88, 89, 5, 18, 10, 2, 89, 90, 7, 4, 2, 2, 90, 93, 
	3, 2, 2, 2, 91, 93, 5, 20, 11, 2, 92, 82, 3, 2, 2, 2, 92, 87, 3, 2, 2, 
	2, 92, 91, 3, 2, 2, 2, 93, 102, 3, 2, 2, 2, 94, 95, 12, 5, 2, 2, 95, 96, 
	5, 12, 7, 2, 96, 97, 5, 18, 10, 6, 97, 101, 3, 2, 2, 2, 98, 99, 12, 4, 
	2, 2, 99, 101, 5, 18, 10, 5, 100, 94, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 
	101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 
	19, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 106, 8, 11, 1, 2, 106, 107, 
	7, 20, 2, 2, 107, 112, 3, 2, 2, 2, 108, 109, 12, 4, 2, 2, 109, 111, 7, 
	20, 2, 2, 110, 108, 3, 2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 110, 3, 2, 2, 
	2, 112, 113, 3, 2, 2, 2, 113, 21, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 
	116, 8, 12, 1, 2, 116, 121, 7, 20, 2, 2, 117, 121, 7, 3, 2, 2, 118, 121, 
	7, 4, 2, 2, 119, 121, 5, 12, 7, 2, 120, 115, 3, 2, 2, 2, 120, 117, 3, 2, 
	2, 2, 120, 118, 3, 2, 2, 2, 120, 119, 3, 2, 2, 2, 121, 132, 3, 2, 2, 2, 
	122, 123, 12, 10, 2, 2, 123, 131, 7, 20, 2, 2, 124, 125, 12, 9, 2, 2, 125, 
	131, 7, 3, 2, 2, 126, 127, 12, 8, 2, 2, 127, 131, 7, 4, 2, 2, 128, 129, 
	12, 7, 2, 2, 129, 131, 5, 12, 7, 2, 130, 122, 3, 2, 2, 2, 130, 124, 3, 
	2, 2, 2, 130, 126, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 134, 3, 2, 2, 
	2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 23, 3, 2, 2, 2, 134, 
	132, 3, 2, 2, 2, 13, 33, 41, 61, 78, 92, 100, 102, 112, 120, 130, 132,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'title'", "'author'", "'subject'", "'keyword'", "'identifier'", 
	"'date'", "", "", "", "'\"'", "'['", "']'",
}
var symbolicNames = []string{
	"", "LPAREN", "RPAREN", "BOOLEAN", "COLON", "TITLE", "AUTHOR", "SUBJECT", 
	"KEYWORD", "IDENTIFIER", "DATE", "LBRACE", "WS1", "ERROR_CHARACTER", "QUOTE", 
	"LBRACKET", "RBRACKET", "RBRACE", "SEARCH_WORD", "WS2", "TO", "AFTER", 
	"BEFORE", "DATE_STRING", "WS3",
}

var ruleNames = []string{
	"query", "query_parts", "field_query", "field_type", "range_field_type", 
	"boolean_op", "range_search_string", "date_string", "search_string", "search_part", 
	"quoted_search_part",
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
	VirgoQueryEOF = antlr.TokenEOF
	VirgoQueryLPAREN = 1
	VirgoQueryRPAREN = 2
	VirgoQueryBOOLEAN = 3
	VirgoQueryCOLON = 4
	VirgoQueryTITLE = 5
	VirgoQueryAUTHOR = 6
	VirgoQuerySUBJECT = 7
	VirgoQueryKEYWORD = 8
	VirgoQueryIDENTIFIER = 9
	VirgoQueryDATE = 10
	VirgoQueryLBRACE = 11
	VirgoQueryWS1 = 12
	VirgoQueryERROR_CHARACTER = 13
	VirgoQueryQUOTE = 14
	VirgoQueryLBRACKET = 15
	VirgoQueryRBRACKET = 16
	VirgoQueryRBRACE = 17
	VirgoQuerySEARCH_WORD = 18
	VirgoQueryWS2 = 19
	VirgoQueryTO = 20
	VirgoQueryAFTER = 21
	VirgoQueryBEFORE = 22
	VirgoQueryDATE_STRING = 23
	VirgoQueryWS3 = 24
)

// VirgoQuery rules.
const (
	VirgoQueryRULE_query = 0
	VirgoQueryRULE_query_parts = 1
	VirgoQueryRULE_field_query = 2
	VirgoQueryRULE_field_type = 3
	VirgoQueryRULE_range_field_type = 4
	VirgoQueryRULE_boolean_op = 5
	VirgoQueryRULE_range_search_string = 6
	VirgoQueryRULE_date_string = 7
	VirgoQueryRULE_search_string = 8
	VirgoQueryRULE_search_part = 9
	VirgoQueryRULE_quoted_search_part = 10
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
		p.SetState(22)
		p.query_parts(0)
	}
	{
		p.SetState(23)
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VirgoQueryLPAREN:
		{
			p.SetState(26)
			p.Match(VirgoQueryLPAREN)
		}
		{
			p.SetState(27)
			p.query_parts(0)
		}
		{
			p.SetState(28)
			p.Match(VirgoQueryRPAREN)
		}


	case VirgoQueryTITLE, VirgoQueryAUTHOR, VirgoQuerySUBJECT, VirgoQueryKEYWORD, VirgoQueryIDENTIFIER, VirgoQueryDATE:
		{
			p.SetState(30)
			p.Field_query()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(39)
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
			p.SetState(33)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(34)
				p.Boolean_op()
			}
			{
				p.SetState(35)
				p.query_parts(4)
			}


		}
		p.SetState(41)
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

func (s *Field_queryContext) Range_field_type() IRange_field_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_field_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRange_field_typeContext)
}

func (s *Field_queryContext) Range_search_string() IRange_search_stringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_search_stringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRange_search_stringContext)
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

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Field_type()
		}
		{
			p.SetState(43)
			p.Match(VirgoQueryCOLON)
		}
		{
			p.SetState(44)
			p.Match(VirgoQueryLBRACE)
		}
		{
			p.SetState(45)
			p.search_string(0)
		}
		{
			p.SetState(46)
			p.Match(VirgoQueryRBRACE)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Field_type()
		}
		{
			p.SetState(49)
			p.Match(VirgoQueryCOLON)
		}
		{
			p.SetState(50)
			p.Match(VirgoQueryLBRACE)
		}
		{
			p.SetState(51)
			p.Match(VirgoQueryRBRACE)
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)
			p.Range_field_type()
		}
		{
			p.SetState(54)
			p.Match(VirgoQueryCOLON)
		}
		{
			p.SetState(55)
			p.Match(VirgoQueryLBRACE)
		}
		{
			p.SetState(56)
			p.Range_search_string()
		}
		{
			p.SetState(57)
			p.Match(VirgoQueryRBRACE)
		}

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
	{
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << VirgoQueryTITLE) | (1 << VirgoQueryAUTHOR) | (1 << VirgoQuerySUBJECT) | (1 << VirgoQueryKEYWORD) | (1 << VirgoQueryIDENTIFIER))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IRange_field_typeContext is an interface to support dynamic dispatch.
type IRange_field_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRange_field_typeContext differentiates from other interfaces.
	IsRange_field_typeContext()
}

type Range_field_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_field_typeContext() *Range_field_typeContext {
	var p = new(Range_field_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_range_field_type
	return p
}

func (*Range_field_typeContext) IsRange_field_typeContext() {}

func NewRange_field_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_field_typeContext {
	var p = new(Range_field_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_range_field_type

	return p
}

func (s *Range_field_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_field_typeContext) DATE() antlr.TerminalNode {
	return s.GetToken(VirgoQueryDATE, 0)
}

func (s *Range_field_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_field_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Range_field_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterRange_field_type(s)
	}
}

func (s *Range_field_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitRange_field_type(s)
	}
}

func (s *Range_field_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitRange_field_type(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *VirgoQuery) Range_field_type() (localctx IRange_field_typeContext) {
	localctx = NewRange_field_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VirgoQueryRULE_range_field_type)

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
		p.SetState(63)
		p.Match(VirgoQueryDATE)
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
	p.EnterRule(localctx, 10, VirgoQueryRULE_boolean_op)

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
		p.SetState(65)
		p.Match(VirgoQueryBOOLEAN)
	}



	return localctx
}


// IRange_search_stringContext is an interface to support dynamic dispatch.
type IRange_search_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRange_search_stringContext differentiates from other interfaces.
	IsRange_search_stringContext()
}

type Range_search_stringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_search_stringContext() *Range_search_stringContext {
	var p = new(Range_search_stringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_range_search_string
	return p
}

func (*Range_search_stringContext) IsRange_search_stringContext() {}

func NewRange_search_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_search_stringContext {
	var p = new(Range_search_stringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_range_search_string

	return p
}

func (s *Range_search_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_search_stringContext) AllDate_string() []IDate_stringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDate_stringContext)(nil)).Elem())
	var tst = make([]IDate_stringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDate_stringContext)
		}
	}

	return tst
}

func (s *Range_search_stringContext) Date_string(i int) IDate_stringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_stringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDate_stringContext)
}

func (s *Range_search_stringContext) TO() antlr.TerminalNode {
	return s.GetToken(VirgoQueryTO, 0)
}

func (s *Range_search_stringContext) BEFORE() antlr.TerminalNode {
	return s.GetToken(VirgoQueryBEFORE, 0)
}

func (s *Range_search_stringContext) AFTER() antlr.TerminalNode {
	return s.GetToken(VirgoQueryAFTER, 0)
}

func (s *Range_search_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_search_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Range_search_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterRange_search_string(s)
	}
}

func (s *Range_search_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitRange_search_string(s)
	}
}

func (s *Range_search_stringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitRange_search_string(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *VirgoQuery) Range_search_string() (localctx IRange_search_stringContext) {
	localctx = NewRange_search_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VirgoQueryRULE_range_search_string)

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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Date_string()
		}
		{
			p.SetState(68)
			p.Match(VirgoQueryTO)
		}
		{
			p.SetState(69)
			p.Date_string()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(VirgoQueryBEFORE)
		}
		{
			p.SetState(72)
			p.Date_string()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Match(VirgoQueryAFTER)
		}
		{
			p.SetState(74)
			p.Date_string()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Date_string()
		}

	}


	return localctx
}


// IDate_stringContext is an interface to support dynamic dispatch.
type IDate_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_stringContext differentiates from other interfaces.
	IsDate_stringContext()
}

type Date_stringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_stringContext() *Date_stringContext {
	var p = new(Date_stringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_date_string
	return p
}

func (*Date_stringContext) IsDate_stringContext() {}

func NewDate_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_stringContext {
	var p = new(Date_stringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_date_string

	return p
}

func (s *Date_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_stringContext) DATE_STRING() antlr.TerminalNode {
	return s.GetToken(VirgoQueryDATE_STRING, 0)
}

func (s *Date_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Date_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterDate_string(s)
	}
}

func (s *Date_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitDate_string(s)
	}
}

func (s *Date_stringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitDate_string(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *VirgoQuery) Date_string() (localctx IDate_stringContext) {
	localctx = NewDate_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VirgoQueryRULE_date_string)

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
		p.SetState(78)
		p.Match(VirgoQueryDATE_STRING)
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

func (s *Search_stringContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(VirgoQueryQUOTE)
}

func (s *Search_stringContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(VirgoQueryQUOTE, i)
}

func (s *Search_stringContext) Quoted_search_part() IQuoted_search_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuoted_search_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuoted_search_partContext)
}

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
	_startState := 16
	p.EnterRecursionRule(localctx, 16, VirgoQueryRULE_search_string, _p)

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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VirgoQueryQUOTE:
		{
			p.SetState(81)
			p.Match(VirgoQueryQUOTE)
		}
		{
			p.SetState(82)
			p.quoted_search_part(0)
		}
		{
			p.SetState(83)
			p.Match(VirgoQueryQUOTE)
		}


	case VirgoQueryLPAREN:
		{
			p.SetState(85)
			p.Match(VirgoQueryLPAREN)
		}
		{
			p.SetState(86)
			p.search_string(0)
		}
		{
			p.SetState(87)
			p.Match(VirgoQueryRPAREN)
		}


	case VirgoQuerySEARCH_WORD:
		{
			p.SetState(89)
			p.search_part(0)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSearch_stringContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_search_string)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(93)
					p.Boolean_op()
				}
				{
					p.SetState(94)
					p.search_string(4)
				}


			case 2:
				localctx = NewSearch_stringContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_search_string)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(97)
					p.search_string(3)
				}

			}

		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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

func (s *Search_partContext) SEARCH_WORD() antlr.TerminalNode {
	return s.GetToken(VirgoQuerySEARCH_WORD, 0)
}

func (s *Search_partContext) Search_part() ISearch_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISearch_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISearch_partContext)
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
	_startState := 18
	p.EnterRecursionRule(localctx, 18, VirgoQueryRULE_search_part, _p)

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
		p.SetState(104)
		p.Match(VirgoQuerySEARCH_WORD)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSearch_partContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_search_part)
			p.SetState(106)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(107)
				p.Match(VirgoQuerySEARCH_WORD)
			}


		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}



	return localctx
}


// IQuoted_search_partContext is an interface to support dynamic dispatch.
type IQuoted_search_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuoted_search_partContext differentiates from other interfaces.
	IsQuoted_search_partContext()
}

type Quoted_search_partContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuoted_search_partContext() *Quoted_search_partContext {
	var p = new(Quoted_search_partContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VirgoQueryRULE_quoted_search_part
	return p
}

func (*Quoted_search_partContext) IsQuoted_search_partContext() {}

func NewQuoted_search_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Quoted_search_partContext {
	var p = new(Quoted_search_partContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_quoted_search_part

	return p
}

func (s *Quoted_search_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Quoted_search_partContext) SEARCH_WORD() antlr.TerminalNode {
	return s.GetToken(VirgoQuerySEARCH_WORD, 0)
}

func (s *Quoted_search_partContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryLPAREN, 0)
}

func (s *Quoted_search_partContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryRPAREN, 0)
}

func (s *Quoted_search_partContext) Boolean_op() IBoolean_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_opContext)
}

func (s *Quoted_search_partContext) Quoted_search_part() IQuoted_search_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuoted_search_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuoted_search_partContext)
}

func (s *Quoted_search_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Quoted_search_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Quoted_search_partContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.EnterQuoted_search_part(s)
	}
}

func (s *Quoted_search_partContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VirgoQueryListener); ok {
		listenerT.ExitQuoted_search_part(s)
	}
}

func (s *Quoted_search_partContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VirgoQueryVisitor:
		return t.VisitQuoted_search_part(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *VirgoQuery) Quoted_search_part() (localctx IQuoted_search_partContext) {
	return p.quoted_search_part(0)
}

func (p *VirgoQuery) quoted_search_part(_p int) (localctx IQuoted_search_partContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQuoted_search_partContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQuoted_search_partContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, VirgoQueryRULE_quoted_search_part, _p)

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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case VirgoQuerySEARCH_WORD:
		{
			p.SetState(114)
			p.Match(VirgoQuerySEARCH_WORD)
		}


	case VirgoQueryLPAREN:
		{
			p.SetState(115)
			p.Match(VirgoQueryLPAREN)
		}


	case VirgoQueryRPAREN:
		{
			p.SetState(116)
			p.Match(VirgoQueryRPAREN)
		}


	case VirgoQueryBOOLEAN:
		{
			p.SetState(117)
			p.Boolean_op()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewQuoted_search_partContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_quoted_search_part)
				p.SetState(120)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(121)
					p.Match(VirgoQuerySEARCH_WORD)
				}


			case 2:
				localctx = NewQuoted_search_partContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_quoted_search_part)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(123)
					p.Match(VirgoQueryLPAREN)
				}


			case 3:
				localctx = NewQuoted_search_partContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_quoted_search_part)
				p.SetState(124)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(125)
					p.Match(VirgoQueryRPAREN)
				}


			case 4:
				localctx = NewQuoted_search_partContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_quoted_search_part)
				p.SetState(126)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(127)
					p.Boolean_op()
				}

			}

		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}



	return localctx
}


func (p *VirgoQuery) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
			var t *Query_partsContext = nil
			if localctx != nil { t = localctx.(*Query_partsContext) }
			return p.Query_parts_Sempred(t, predIndex)

	case 8:
			var t *Search_stringContext = nil
			if localctx != nil { t = localctx.(*Search_stringContext) }
			return p.Search_string_Sempred(t, predIndex)

	case 9:
			var t *Search_partContext = nil
			if localctx != nil { t = localctx.(*Search_partContext) }
			return p.Search_part_Sempred(t, predIndex)

	case 10:
			var t *Quoted_search_partContext = nil
			if localctx != nil { t = localctx.(*Quoted_search_partContext) }
			return p.Quoted_search_part_Sempred(t, predIndex)


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

func (p *VirgoQuery) Quoted_search_part_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
			return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
			return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

