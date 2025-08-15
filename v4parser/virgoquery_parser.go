// Code generated from VirgoQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package v4parser // VirgoQuery
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VirgoQuery struct {
	*antlr.BaseParser
}

var VirgoQueryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func virgoqueryParserInit() {
	staticData := &VirgoQueryParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'title'", "'journal_title'", "'author'", "'subject'",
		"'keyword'", "'fulltext'", "'series'", "'published'", "'identifier'",
		"'filter'", "'date'", "", "", "", "", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "LPAREN", "RPAREN", "BOOLEAN", "COLON", "TITLE", "JOURNAL_TITLE",
		"AUTHOR", "SUBJECT", "KEYWORD", "FULLTEXT", "SERIES", "PUBLISHED", "IDENTIFIER",
		"FILTER", "DATE", "LBRACE", "WS1", "ERROR_CHARACTER", "QUOTE", "LBRACKET",
		"RBRACKET", "RBRACE", "SEARCH_WORD", "WS2", "TO", "AFTER", "BEFORE",
		"DATE_STRING", "WS3", "QUOTE_STR",
	}
	staticData.RuleNames = []string{
		"query", "query_parts", "field_query", "field_type", "range_field_type",
		"boolean_op", "range_search_string", "date_string", "search_string",
		"search_part", "quoted_search_part",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 133, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 32, 8,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 38, 8, 1, 10, 1, 12, 1, 41, 9, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 77, 8, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		90, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 98, 8, 8, 10, 8, 12,
		8, 101, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 108, 8, 9, 10, 9, 12,
		9, 111, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 118, 8, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 128, 8, 10,
		10, 10, 12, 10, 131, 9, 10, 1, 10, 0, 4, 2, 16, 18, 20, 11, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 0, 1, 1, 0, 5, 14, 140, 0, 22, 1, 0, 0, 0, 2,
		31, 1, 0, 0, 0, 4, 59, 1, 0, 0, 0, 6, 61, 1, 0, 0, 0, 8, 63, 1, 0, 0, 0,
		10, 65, 1, 0, 0, 0, 12, 76, 1, 0, 0, 0, 14, 78, 1, 0, 0, 0, 16, 89, 1,
		0, 0, 0, 18, 102, 1, 0, 0, 0, 20, 117, 1, 0, 0, 0, 22, 23, 3, 2, 1, 0,
		23, 24, 5, 0, 0, 1, 24, 1, 1, 0, 0, 0, 25, 26, 6, 1, -1, 0, 26, 27, 5,
		1, 0, 0, 27, 28, 3, 2, 1, 0, 28, 29, 5, 2, 0, 0, 29, 32, 1, 0, 0, 0, 30,
		32, 3, 4, 2, 0, 31, 25, 1, 0, 0, 0, 31, 30, 1, 0, 0, 0, 32, 39, 1, 0, 0,
		0, 33, 34, 10, 3, 0, 0, 34, 35, 3, 10, 5, 0, 35, 36, 3, 2, 1, 4, 36, 38,
		1, 0, 0, 0, 37, 33, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0,
		39, 40, 1, 0, 0, 0, 40, 3, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 43, 3, 6,
		3, 0, 43, 44, 5, 4, 0, 0, 44, 45, 5, 16, 0, 0, 45, 46, 3, 16, 8, 0, 46,
		47, 5, 22, 0, 0, 47, 60, 1, 0, 0, 0, 48, 49, 3, 6, 3, 0, 49, 50, 5, 4,
		0, 0, 50, 51, 5, 16, 0, 0, 51, 52, 5, 22, 0, 0, 52, 60, 1, 0, 0, 0, 53,
		54, 3, 8, 4, 0, 54, 55, 5, 4, 0, 0, 55, 56, 5, 16, 0, 0, 56, 57, 3, 12,
		6, 0, 57, 58, 5, 22, 0, 0, 58, 60, 1, 0, 0, 0, 59, 42, 1, 0, 0, 0, 59,
		48, 1, 0, 0, 0, 59, 53, 1, 0, 0, 0, 60, 5, 1, 0, 0, 0, 61, 62, 7, 0, 0,
		0, 62, 7, 1, 0, 0, 0, 63, 64, 5, 15, 0, 0, 64, 9, 1, 0, 0, 0, 65, 66, 5,
		3, 0, 0, 66, 11, 1, 0, 0, 0, 67, 68, 3, 14, 7, 0, 68, 69, 5, 25, 0, 0,
		69, 70, 3, 14, 7, 0, 70, 77, 1, 0, 0, 0, 71, 72, 5, 27, 0, 0, 72, 77, 3,
		14, 7, 0, 73, 74, 5, 26, 0, 0, 74, 77, 3, 14, 7, 0, 75, 77, 3, 14, 7, 0,
		76, 67, 1, 0, 0, 0, 76, 71, 1, 0, 0, 0, 76, 73, 1, 0, 0, 0, 76, 75, 1,
		0, 0, 0, 77, 13, 1, 0, 0, 0, 78, 79, 5, 28, 0, 0, 79, 15, 1, 0, 0, 0, 80,
		81, 6, 8, -1, 0, 81, 82, 5, 19, 0, 0, 82, 83, 5, 30, 0, 0, 83, 90, 5, 19,
		0, 0, 84, 85, 5, 1, 0, 0, 85, 86, 3, 16, 8, 0, 86, 87, 5, 2, 0, 0, 87,
		90, 1, 0, 0, 0, 88, 90, 3, 18, 9, 0, 89, 80, 1, 0, 0, 0, 89, 84, 1, 0,
		0, 0, 89, 88, 1, 0, 0, 0, 90, 99, 1, 0, 0, 0, 91, 92, 10, 3, 0, 0, 92,
		93, 3, 10, 5, 0, 93, 94, 3, 16, 8, 4, 94, 98, 1, 0, 0, 0, 95, 96, 10, 2,
		0, 0, 96, 98, 3, 16, 8, 3, 97, 91, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98,
		101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 17, 1, 0,
		0, 0, 101, 99, 1, 0, 0, 0, 102, 103, 6, 9, -1, 0, 103, 104, 5, 23, 0, 0,
		104, 109, 1, 0, 0, 0, 105, 106, 10, 2, 0, 0, 106, 108, 5, 23, 0, 0, 107,
		105, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110,
		1, 0, 0, 0, 110, 19, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 113, 6, 10,
		-1, 0, 113, 118, 5, 23, 0, 0, 114, 118, 5, 1, 0, 0, 115, 118, 5, 2, 0,
		0, 116, 118, 3, 10, 5, 0, 117, 112, 1, 0, 0, 0, 117, 114, 1, 0, 0, 0, 117,
		115, 1, 0, 0, 0, 117, 116, 1, 0, 0, 0, 118, 129, 1, 0, 0, 0, 119, 120,
		10, 8, 0, 0, 120, 128, 5, 23, 0, 0, 121, 122, 10, 7, 0, 0, 122, 128, 5,
		1, 0, 0, 123, 124, 10, 6, 0, 0, 124, 128, 5, 2, 0, 0, 125, 126, 10, 5,
		0, 0, 126, 128, 3, 10, 5, 0, 127, 119, 1, 0, 0, 0, 127, 121, 1, 0, 0, 0,
		127, 123, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129,
		127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 21, 1, 0, 0, 0, 131, 129, 1,
		0, 0, 0, 11, 31, 39, 59, 76, 89, 97, 99, 109, 117, 127, 129,
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

// VirgoQueryInit initializes any static state used to implement VirgoQuery. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVirgoQuery(). You can call this function if you wish to initialize the static state ahead
// of time.
func VirgoQueryInit() {
	staticData := &VirgoQueryParserStaticData
	staticData.once.Do(virgoqueryParserInit)
}

// NewVirgoQuery produces a new parser instance for the optional input antlr.TokenStream.
func NewVirgoQuery(input antlr.TokenStream) *VirgoQuery {
	VirgoQueryInit()
	this := new(VirgoQuery)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VirgoQueryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "VirgoQuery.g4"

	return this
}

// VirgoQuery tokens.
const (
	VirgoQueryEOF             = antlr.TokenEOF
	VirgoQueryLPAREN          = 1
	VirgoQueryRPAREN          = 2
	VirgoQueryBOOLEAN         = 3
	VirgoQueryCOLON           = 4
	VirgoQueryTITLE           = 5
	VirgoQueryJOURNAL_TITLE   = 6
	VirgoQueryAUTHOR          = 7
	VirgoQuerySUBJECT         = 8
	VirgoQueryKEYWORD         = 9
	VirgoQueryFULLTEXT        = 10
	VirgoQuerySERIES          = 11
	VirgoQueryPUBLISHED       = 12
	VirgoQueryIDENTIFIER      = 13
	VirgoQueryFILTER          = 14
	VirgoQueryDATE            = 15
	VirgoQueryLBRACE          = 16
	VirgoQueryWS1             = 17
	VirgoQueryERROR_CHARACTER = 18
	VirgoQueryQUOTE           = 19
	VirgoQueryLBRACKET        = 20
	VirgoQueryRBRACKET        = 21
	VirgoQueryRBRACE          = 22
	VirgoQuerySEARCH_WORD     = 23
	VirgoQueryWS2             = 24
	VirgoQueryTO              = 25
	VirgoQueryAFTER           = 26
	VirgoQueryBEFORE          = 27
	VirgoQueryDATE_STRING     = 28
	VirgoQueryWS3             = 29
	VirgoQueryQUOTE_STR       = 30
)

// VirgoQuery rules.
const (
	VirgoQueryRULE_query               = 0
	VirgoQueryRULE_query_parts         = 1
	VirgoQueryRULE_field_query         = 2
	VirgoQueryRULE_field_type          = 3
	VirgoQueryRULE_range_field_type    = 4
	VirgoQueryRULE_boolean_op          = 5
	VirgoQueryRULE_range_search_string = 6
	VirgoQueryRULE_date_string         = 7
	VirgoQueryRULE_search_string       = 8
	VirgoQueryRULE_search_part         = 9
	VirgoQueryRULE_quoted_search_part  = 10
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Query_parts() IQuery_partsContext
	EOF() antlr.TerminalNode

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Query_parts() IQuery_partsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuery_partsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.query_parts(0)
	}
	{
		p.SetState(23)
		p.Match(VirgoQueryEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQuery_partsContext is an interface to support dynamic dispatch.
type IQuery_partsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllQuery_parts() []IQuery_partsContext
	Query_parts(i int) IQuery_partsContext
	RPAREN() antlr.TerminalNode
	Field_query() IField_queryContext
	Boolean_op() IBoolean_opContext

	// IsQuery_partsContext differentiates from other interfaces.
	IsQuery_partsContext()
}

type Query_partsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuery_partsContext() *Query_partsContext {
	var p = new(Query_partsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_query_parts
	return p
}

func InitEmptyQuery_partsContext(p *Query_partsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_query_parts
}

func (*Query_partsContext) IsQuery_partsContext() {}

func NewQuery_partsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Query_partsContext {
	var p = new(Query_partsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_query_parts

	return p
}

func (s *Query_partsContext) GetParser() antlr.Parser { return s.parser }

func (s *Query_partsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryLPAREN, 0)
}

func (s *Query_partsContext) AllQuery_parts() []IQuery_partsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQuery_partsContext); ok {
			len++
		}
	}

	tst := make([]IQuery_partsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQuery_partsContext); ok {
			tst[i] = t.(IQuery_partsContext)
			i++
		}
	}

	return tst
}

func (s *Query_partsContext) Query_parts(i int) IQuery_partsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuery_partsContext); ok {
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

	return t.(IQuery_partsContext)
}

func (s *Query_partsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryRPAREN, 0)
}

func (s *Query_partsContext) Field_query() IField_queryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_queryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_queryContext)
}

func (s *Query_partsContext) Boolean_op() IBoolean_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VirgoQueryLPAREN:
		{
			p.SetState(26)
			p.Match(VirgoQueryLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.query_parts(0)
		}
		{
			p.SetState(28)
			p.Match(VirgoQueryRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VirgoQueryTITLE, VirgoQueryJOURNAL_TITLE, VirgoQueryAUTHOR, VirgoQuerySUBJECT, VirgoQueryKEYWORD, VirgoQueryFULLTEXT, VirgoQuerySERIES, VirgoQueryPUBLISHED, VirgoQueryIDENTIFIER, VirgoQueryFILTER, VirgoQueryDATE:
		{
			p.SetState(30)
			p.Field_query()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
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
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
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
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IField_queryContext is an interface to support dynamic dispatch.
type IField_queryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_type() IField_typeContext
	COLON() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	Search_string() ISearch_stringContext
	RBRACE() antlr.TerminalNode
	Range_field_type() IRange_field_typeContext
	Range_search_string() IRange_search_stringContext

	// IsField_queryContext differentiates from other interfaces.
	IsField_queryContext()
}

type Field_queryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_queryContext() *Field_queryContext {
	var p = new(Field_queryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_field_query
	return p
}

func InitEmptyField_queryContext(p *Field_queryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_field_query
}

func (*Field_queryContext) IsField_queryContext() {}

func NewField_queryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_queryContext {
	var p = new(Field_queryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_field_query

	return p
}

func (s *Field_queryContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_queryContext) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearch_stringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearch_stringContext)
}

func (s *Field_queryContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(VirgoQueryRBRACE, 0)
}

func (s *Field_queryContext) Range_field_type() IRange_field_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRange_field_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRange_field_typeContext)
}

func (s *Field_queryContext) Range_search_string() IRange_search_stringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRange_search_stringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Field_type()
		}
		{
			p.SetState(43)
			p.Match(VirgoQueryCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(VirgoQueryLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.search_string(0)
		}
		{
			p.SetState(46)
			p.Match(VirgoQueryRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(VirgoQueryLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(VirgoQueryRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Match(VirgoQueryLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Range_search_string()
		}
		{
			p.SetState(57)
			p.Match(VirgoQueryRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IField_typeContext is an interface to support dynamic dispatch.
type IField_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TITLE() antlr.TerminalNode
	JOURNAL_TITLE() antlr.TerminalNode
	AUTHOR() antlr.TerminalNode
	SUBJECT() antlr.TerminalNode
	KEYWORD() antlr.TerminalNode
	PUBLISHED() antlr.TerminalNode
	FULLTEXT() antlr.TerminalNode
	SERIES() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	FILTER() antlr.TerminalNode

	// IsField_typeContext differentiates from other interfaces.
	IsField_typeContext()
}

type Field_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_typeContext() *Field_typeContext {
	var p = new(Field_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_field_type
	return p
}

func InitEmptyField_typeContext(p *Field_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_field_type
}

func (*Field_typeContext) IsField_typeContext() {}

func NewField_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_typeContext {
	var p = new(Field_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_field_type

	return p
}

func (s *Field_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_typeContext) TITLE() antlr.TerminalNode {
	return s.GetToken(VirgoQueryTITLE, 0)
}

func (s *Field_typeContext) JOURNAL_TITLE() antlr.TerminalNode {
	return s.GetToken(VirgoQueryJOURNAL_TITLE, 0)
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

func (s *Field_typeContext) PUBLISHED() antlr.TerminalNode {
	return s.GetToken(VirgoQueryPUBLISHED, 0)
}

func (s *Field_typeContext) FULLTEXT() antlr.TerminalNode {
	return s.GetToken(VirgoQueryFULLTEXT, 0)
}

func (s *Field_typeContext) SERIES() antlr.TerminalNode {
	return s.GetToken(VirgoQuerySERIES, 0)
}

func (s *Field_typeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VirgoQueryIDENTIFIER, 0)
}

func (s *Field_typeContext) FILTER() antlr.TerminalNode {
	return s.GetToken(VirgoQueryFILTER, 0)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32736) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRange_field_typeContext is an interface to support dynamic dispatch.
type IRange_field_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATE() antlr.TerminalNode

	// IsRange_field_typeContext differentiates from other interfaces.
	IsRange_field_typeContext()
}

type Range_field_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_field_typeContext() *Range_field_typeContext {
	var p = new(Range_field_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_range_field_type
	return p
}

func InitEmptyRange_field_typeContext(p *Range_field_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_range_field_type
}

func (*Range_field_typeContext) IsRange_field_typeContext() {}

func NewRange_field_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_field_typeContext {
	var p = new(Range_field_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(VirgoQueryDATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolean_opContext is an interface to support dynamic dispatch.
type IBoolean_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN() antlr.TerminalNode

	// IsBoolean_opContext differentiates from other interfaces.
	IsBoolean_opContext()
}

type Boolean_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_opContext() *Boolean_opContext {
	var p = new(Boolean_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_boolean_op
	return p
}

func InitEmptyBoolean_opContext(p *Boolean_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_boolean_op
}

func (*Boolean_opContext) IsBoolean_opContext() {}

func NewBoolean_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_opContext {
	var p = new(Boolean_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(VirgoQueryBOOLEAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRange_search_stringContext is an interface to support dynamic dispatch.
type IRange_search_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDate_string() []IDate_stringContext
	Date_string(i int) IDate_stringContext
	TO() antlr.TerminalNode
	BEFORE() antlr.TerminalNode
	AFTER() antlr.TerminalNode

	// IsRange_search_stringContext differentiates from other interfaces.
	IsRange_search_stringContext()
}

type Range_search_stringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_search_stringContext() *Range_search_stringContext {
	var p = new(Range_search_stringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_range_search_string
	return p
}

func InitEmptyRange_search_stringContext(p *Range_search_stringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_range_search_string
}

func (*Range_search_stringContext) IsRange_search_stringContext() {}

func NewRange_search_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_search_stringContext {
	var p = new(Range_search_stringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_range_search_string

	return p
}

func (s *Range_search_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_search_stringContext) AllDate_string() []IDate_stringContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDate_stringContext); ok {
			len++
		}
	}

	tst := make([]IDate_stringContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDate_stringContext); ok {
			tst[i] = t.(IDate_stringContext)
			i++
		}
	}

	return tst
}

func (s *Range_search_stringContext) Date_string(i int) IDate_stringContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDate_stringContext); ok {
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
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Date_string()
		}
		{
			p.SetState(68)
			p.Match(VirgoQueryTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDate_stringContext is an interface to support dynamic dispatch.
type IDate_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATE_STRING() antlr.TerminalNode

	// IsDate_stringContext differentiates from other interfaces.
	IsDate_stringContext()
}

type Date_stringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_stringContext() *Date_stringContext {
	var p = new(Date_stringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_date_string
	return p
}

func InitEmptyDate_stringContext(p *Date_stringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_date_string
}

func (*Date_stringContext) IsDate_stringContext() {}

func NewDate_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_stringContext {
	var p = new(Date_stringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(VirgoQueryDATE_STRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISearch_stringContext is an interface to support dynamic dispatch.
type ISearch_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode
	QUOTE_STR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllSearch_string() []ISearch_stringContext
	Search_string(i int) ISearch_stringContext
	RPAREN() antlr.TerminalNode
	Search_part() ISearch_partContext
	Boolean_op() IBoolean_opContext

	// IsSearch_stringContext differentiates from other interfaces.
	IsSearch_stringContext()
}

type Search_stringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearch_stringContext() *Search_stringContext {
	var p = new(Search_stringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_search_string
	return p
}

func InitEmptySearch_stringContext(p *Search_stringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_search_string
}

func (*Search_stringContext) IsSearch_stringContext() {}

func NewSearch_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Search_stringContext {
	var p = new(Search_stringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *Search_stringContext) QUOTE_STR() antlr.TerminalNode {
	return s.GetToken(VirgoQueryQUOTE_STR, 0)
}

func (s *Search_stringContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryLPAREN, 0)
}

func (s *Search_stringContext) AllSearch_string() []ISearch_stringContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearch_stringContext); ok {
			len++
		}
	}

	tst := make([]ISearch_stringContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearch_stringContext); ok {
			tst[i] = t.(ISearch_stringContext)
			i++
		}
	}

	return tst
}

func (s *Search_stringContext) Search_string(i int) ISearch_stringContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearch_stringContext); ok {
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

	return t.(ISearch_stringContext)
}

func (s *Search_stringContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(VirgoQueryRPAREN, 0)
}

func (s *Search_stringContext) Search_part() ISearch_partContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearch_partContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearch_partContext)
}

func (s *Search_stringContext) Boolean_op() IBoolean_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VirgoQueryQUOTE:
		{
			p.SetState(81)
			p.Match(VirgoQueryQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(VirgoQueryQUOTE_STR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(VirgoQueryQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VirgoQueryLPAREN:
		{
			p.SetState(84)
			p.Match(VirgoQueryLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.search_string(0)
		}
		{
			p.SetState(86)
			p.Match(VirgoQueryRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VirgoQuerySEARCH_WORD:
		{
			p.SetState(88)
			p.search_part(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSearch_stringContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_search_string)
				p.SetState(91)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(92)
					p.Boolean_op()
				}
				{
					p.SetState(93)
					p.search_string(4)
				}

			case 2:
				localctx = NewSearch_stringContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_search_string)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(96)
					p.search_string(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISearch_partContext is an interface to support dynamic dispatch.
type ISearch_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEARCH_WORD() antlr.TerminalNode
	Search_part() ISearch_partContext

	// IsSearch_partContext differentiates from other interfaces.
	IsSearch_partContext()
}

type Search_partContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearch_partContext() *Search_partContext {
	var p = new(Search_partContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_search_part
	return p
}

func InitEmptySearch_partContext(p *Search_partContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_search_part
}

func (*Search_partContext) IsSearch_partContext() {}

func NewSearch_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Search_partContext {
	var p = new(Search_partContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VirgoQueryRULE_search_part

	return p
}

func (s *Search_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Search_partContext) SEARCH_WORD() antlr.TerminalNode {
	return s.GetToken(VirgoQuerySEARCH_WORD, 0)
}

func (s *Search_partContext) Search_part() ISearch_partContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearch_partContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(VirgoQuerySEARCH_WORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSearch_partContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_search_part)
			p.SetState(105)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(106)
				p.Match(VirgoQuerySEARCH_WORD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQuoted_search_partContext is an interface to support dynamic dispatch.
type IQuoted_search_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEARCH_WORD() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Boolean_op() IBoolean_opContext
	Quoted_search_part() IQuoted_search_partContext

	// IsQuoted_search_partContext differentiates from other interfaces.
	IsQuoted_search_partContext()
}

type Quoted_search_partContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuoted_search_partContext() *Quoted_search_partContext {
	var p = new(Quoted_search_partContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_quoted_search_part
	return p
}

func InitEmptyQuoted_search_partContext(p *Quoted_search_partContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VirgoQueryRULE_quoted_search_part
}

func (*Quoted_search_partContext) IsQuoted_search_partContext() {}

func NewQuoted_search_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Quoted_search_partContext {
	var p = new(Quoted_search_partContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_opContext)
}

func (s *Quoted_search_partContext) Quoted_search_part() IQuoted_search_partContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuoted_search_partContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VirgoQuerySEARCH_WORD:
		{
			p.SetState(113)
			p.Match(VirgoQuerySEARCH_WORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VirgoQueryLPAREN:
		{
			p.SetState(114)
			p.Match(VirgoQueryLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VirgoQueryRPAREN:
		{
			p.SetState(115)
			p.Match(VirgoQueryRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VirgoQueryBOOLEAN:
		{
			p.SetState(116)
			p.Boolean_op()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewQuoted_search_partContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_quoted_search_part)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(120)
					p.Match(VirgoQuerySEARCH_WORD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewQuoted_search_partContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_quoted_search_part)
				p.SetState(121)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(122)
					p.Match(VirgoQueryLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 3:
				localctx = NewQuoted_search_partContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_quoted_search_part)
				p.SetState(123)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(124)
					p.Match(VirgoQueryRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 4:
				localctx = NewQuoted_search_partContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VirgoQueryRULE_quoted_search_part)
				p.SetState(125)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(126)
					p.Boolean_op()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *VirgoQuery) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *Query_partsContext = nil
		if localctx != nil {
			t = localctx.(*Query_partsContext)
		}
		return p.Query_parts_Sempred(t, predIndex)

	case 8:
		var t *Search_stringContext = nil
		if localctx != nil {
			t = localctx.(*Search_stringContext)
		}
		return p.Search_string_Sempred(t, predIndex)

	case 9:
		var t *Search_partContext = nil
		if localctx != nil {
			t = localctx.(*Search_partContext)
		}
		return p.Search_part_Sempred(t, predIndex)

	case 10:
		var t *Quoted_search_partContext = nil
		if localctx != nil {
			t = localctx.(*Quoted_search_partContext)
		}
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
