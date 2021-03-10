package v4parser

import (
	"context"
	"errors"
	"fmt"

	"log"
	"reflect"
	"strings"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const escapeQuote = `\"`

// SolrParser will parse the query string into a format compatable with a solr search
type SolrParser struct {
	Debug         bool                // print debugging info
	DebugCallTree bool                // if Debug enabled, also print call tree
	FieldValues   map[string][]string // map of values seen per field type
	callStack     []string            // internal call stack (len == level)
	escapeText    bool                // internal escaping flag
}

// internal parser info
type parseContext struct {
	lexerErrorListener  virgoErrorListener
	parserErrorListener virgoErrorListener
	parser              *VirgoQuery
	tree                antlr.ParserRuleContext
}

func (v *SolrParser) debug(format string, args ...interface{}) {
	if v.Debug == false {
		return
	}

	line := strings.Repeat(" ", len(v.callStack)) + format

	log.Printf("INFO: [V4QUERY] "+line, args...)
}

func (v *SolrParser) enterFunction(funcName string) {
	if v.DebugCallTree == true {
		v.debug("%s {", funcName)
	}

	v.callStack = append(v.callStack, funcName)
}

func (v *SolrParser) exitFunction() {
	funcName := v.callStack[len(v.callStack)-1]
	v.callStack = v.callStack[:len(v.callStack)-1]

	if v.DebugCallTree == true {
		v.debug("} // %s", funcName)
	}
}

func (v *SolrParser) visit(tree antlr.Tree) interface{} {
	v.enterFunction("visit()")
	defer v.exitFunction()

	switch val := tree.(type) {
	case antlr.RuleNode:
		return v.visitRuleNode(val)
	case antlr.TerminalNode:
		return v.visitTerminal(val)
	}

	return nil
}

func (v *SolrParser) visitRuleNode(rule antlr.RuleNode) interface{} {
	v.enterFunction("visitRuleNode()")
	defer v.exitFunction()

	switch rule.(type) {
	case *QueryContext:
		return v.visitQuery(rule)
	case *Query_partsContext:
		return v.visitQueryParts(rule)
	case *Field_queryContext:
		return v.visitFieldQuery(rule)
	case *Field_typeContext:
		return v.visitFieldType(rule)
	case *Range_field_typeContext:
		return v.visitRangeFieldType(rule)
	case *Range_search_stringContext:
		return v.visitRangeSearchString(rule)
	case *Search_stringContext:
		return v.visitSearchString(rule)
	default:
		return v.visitChildren(rule)
	}
}

func (v *SolrParser) visitQuery(query antlr.RuleNode) interface{} {
	v.enterFunction("visitQuery()")
	defer v.exitFunction()

	first := query.GetChild(0)
	result := v.visit(first)

	return result
}

func (v *SolrParser) visitQueryParts(ctx antlr.RuleNode) interface{} {
	v.enterFunction("visitQueryParts()")
	defer v.exitFunction()

	//  query_parts : query_parts boolean_op query_parts
	if ctx.GetChildCount() == 3 {
		switch ctx.GetChild(1).(type) {
		case *Boolean_opContext:
			query1 := v.visit(ctx.GetChild(0)).(string)
			queryop := v.visit(ctx.GetChild(1)).(string)
			query2 := v.visit(ctx.GetChild(2)).(string)

			out := fmt.Sprintf("(%s %s %s)", query1, queryop, query2)

			v.debug("[GRAMMAR] query_parts (boolean): [%s]", out)

			return out
		}
	}

	// query_parts : LPAREN query_parts RPAREN
	if ctx.GetChildCount() == 3 {
		switch ctx.GetChild(0).(type) {
		case antlr.TerminalNode:
			query := v.visit(ctx.GetChild(1)).(string)

			out := fmt.Sprintf("(%s)", query)

			v.debug("[GRAMMAR] query_parts (parentheses): [%s]", out)

			return out
		}
	}

	// query_parts: field_query
	out := v.visit(ctx.GetChild(0))

	if out == nil {
		out = ""
	}

	v.debug("[GRAMMAR] query_parts (field_query): [%v]", out)

	return out
}

func (v *SolrParser) visitFieldQuery(ctx antlr.RuleNode) interface{} {
	v.enterFunction("visitFieldQuery()")
	defer v.exitFunction()

	// field_query : field_type COLON LBRACE search_string RBRACE
	//             | range_field_type COLON LBRACE range_search_string RBRACE
	//   (_query_:"{!edismax qf=$title_qf pf=$title_pf}(certification of teachers )"

	// grey magic to get field name so we can populate per-field query slices in expand()

	// if this is nil, the query isn't going to parse anyway, so bail now to avoid crashing
	if ctx.GetChild(0).GetChild(0) == nil {
		return ""
	}

	fieldName := ctx.GetChild(0).GetChild(0).GetPayload().(*antlr.CommonToken).GetText()
	fieldType := v.visit(ctx.GetChild(0)).(string)

	v.debug("field name: [%s]", fieldName)
	v.debug("field type: [%s]", fieldType)

	var out string

	switch ctx.GetChild(3).(type) {
	case antlr.TerminalNode:
		// no query supplied; default to '*' query
		out = v.expand("", fieldName, fieldType, "*")
	default:
		// set escaping flag.  don't want to escape filters
		esc := v.escapeText
		v.escapeText = true
		if fieldName == "filter" {
			v.escapeText = false
		}

		query := v.visit(ctx.GetChild(3))
		out = v.expand("", fieldName, fieldType, query)

		// restore escape flag
		v.escapeText = esc
	}

	v.debug("[GRAMMAR] field_query: [%s]", out)

	return out
}

// This expands the search_string inside a field_query
// if there are no boolean operators in the search_string it merely takes the words and characters
// of the search_string and adds the proper query fragment syntax as required by SolrParser.
// e.g. going from this:     title : {susan sontag}
//      to this:             _query_:"{!edismax qf=$title_qf pf=$title_pf}(susan sontag)"
// if there ARE boolean operators in the search_string they need to be (recursively) expanded and
// wrapped with parentheses to ensure proper operator precedence
// e.g. this field_query :   title : {"susan sontag" OR music title}
// expands to this:           (_query_:"{!edismax qf=$title_qf pf=$title_pf}(\" susan sontag \")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)")
func (v *SolrParser) expand(inStr string, fieldName string, fieldType string, query interface{}) string {
	v.enterFunction("expand()")
	defer v.exitFunction()

	rt := reflect.TypeOf(query)
	if rt == nil {
		v.debug("nil query")
		return ""
	}

	kind := rt.Kind()

	v.debug("[EXPAND] cur string : [%s]", inStr)
	v.debug("[EXPAND] new field  : [%s]", fieldName)
	v.debug("[EXPAND] new query  : [%v]", query)
	v.debug("[EXPAND] query type : [%s]", kind)

	if kind == reflect.Array || kind == reflect.Slice {
		parts := reflect.ValueOf(query)

		out := fmt.Sprintf("%s(", inStr)
		out = v.expand(out, fieldName, fieldType, parts.Index(0).Interface())
		out += fmt.Sprintf(" %s ", parts.Index(1))
		out = v.expand(out, fieldName, fieldType, parts.Index(2).Interface())
		out = fmt.Sprintf("%s)", out)

		v.debug("[EXPAND] new string : [%s]", out)

		return out
	}

	val := fmt.Sprintf("%s", query)

	v.FieldValues[fieldName] = append(v.FieldValues[fieldName], val)

	out := fmt.Sprintf(`%s_query_:"%s(%s)"`, inStr, fieldType, val)

	v.debug("[EXPAND] new string : [%s]", out)

	return out
}

func (v *SolrParser) escapeSolrSpecialCharacters(s string) string {
	// https://lucene.apache.org/solr/guide/7_7/the-standard-query-parser.html#TheStandardQueryParser-EscapingSpecialCharacters
	//
	// escape characters in the list above, with the following modifications:
	//
	// * escape any existing backslashes first to sanitize user input
	// * do not escape `*` as it is allowed as a wildcard
	// * do not escape `(` or `)` as they are allowed for ordering

	// NOTE: strings passed to this function appear in the quoted part of
	// each solr query fragment, therefore they must be doubly-escaped.

	specialChars := []string{`+`, `-`, `&&`, `||`, `!`, `{`, `}`, `[`, `]`, `^`, `"`, `~`, `?`, `:`, `/`}

	escaped := strings.ReplaceAll(s, `\`, `\\\\`)
	for _, c := range specialChars {
		escaped = strings.ReplaceAll(escaped, c, `\\`+c)
	}

	v.debug("[ESCAPE] from: [%s]", s)
	v.debug("[ESCAPE]   to: [%s]", escaped)

	return escaped
}

func (v *SolrParser) visitFieldType(ctx antlr.RuleNode) interface{} {
	v.enterFunction("visitFieldType()")
	defer v.exitFunction()

	// field_type : TITLE | AUTHOR | SUBJECT | KEYWORD | PUBLISHED | FILTER | IDENTIFIER
	childTree := ctx.GetChild(0)
	t := childTree.GetPayload().(*antlr.CommonToken)
	fieldType := t.GetText()

	out := "!edismax"

	// filter field type needs special handling
	if fieldType == "filter" {
		return ""
	}

	// keyword field type doesn't need qf/pf parameter
	if fieldType == "keyword" {
		return "{" + out + "}"
	}

	// all other field types need qf/pf parameters based on their names
	qf := "qf=$" + fieldType + "_qf"
	pf := "pf=$" + fieldType + "_pf"

	out = "{" + out + " " + qf + " " + pf + "}"

	v.debug("[GRAMMAR] field_type: [%s]", out)

	return out
}

func (v *SolrParser) visitRangeFieldType(ctx antlr.RuleNode) interface{} {
	v.enterFunction("visitRangeFieldType()")
	defer v.exitFunction()

	// range_field_type : DATE
	childTree := ctx.GetChild(0)
	t := childTree.GetPayload().(*antlr.CommonToken)
	fieldType := t.GetText()

	out := ""

	if fieldType == "date" {
		df := "published_daterange"
		out = "!lucene df=" + df
	}

	v.debug("[GRAMMAR] range_field_type: [%s]", out)

	return "{" + out + "}"
}

func (v *SolrParser) getRangeBeforeAfterValues(date string) (string, string) {
	// calculate the appropriate values for dates in BEFORE/AFTER clauses

	// handle YYYY case
	if t, err := time.Parse("2006", date); err == nil {
		b := t.AddDate(-1, 0, 0)
		a := t.AddDate(1, 0, 0)
		return b.Format("2006"), a.Format("2006")
	}

	// handle YYYY-MM case
	if t, err := time.Parse("2006-1", date); err == nil {
		b := t.AddDate(0, -1, 0)
		a := t.AddDate(0, 1, 0)
		return b.Format("2006-01"), a.Format("2006-01")
	}

	// handle YYYY-MM-DD case
	if t, err := time.Parse("2006-1-2", date); err == nil {
		b := t.AddDate(0, 0, -1)
		a := t.AddDate(0, 0, 1)
		return b.Format("2006-01-02"), a.Format("2006-01-02")
	}

	// unhandled date format; just use the date that was passed
	return date, date
}

func (v *SolrParser) visitRangeSearchString(ctx antlr.RuleNode) interface{} {
	v.enterFunction("visitRangeSearchString()")
	defer v.exitFunction()

	//    range_search_string : date_string TO date_string
	//                        | BEFORE date_string
	//                        | AFTER date_string
	//                        | date_string

	if ctx.GetChildCount() == 1 {
		return v.visit(ctx.GetChild(0))
	}

	rangeStart := "*"
	rangeEnd := "*"

	if ctx.GetChildCount() == 3 {
		rangeStart = v.visit(ctx.GetChild(0)).(string)
		rangeEnd = v.visit(ctx.GetChild(2)).(string)
	} else if ctx.GetChildCount() == 2 {
		value := v.visit(ctx.GetChild(1)).(string)

		switch ctx.GetChild(0).(type) {
		case antlr.TerminalNode:
			before, after := v.getRangeBeforeAfterValues(value)

			terminal := ctx.GetChild(0).(antlr.TerminalNode)
			if terminal.GetSymbol().GetTokenType() == VirgoQueryLexerBEFORE {
				rangeEnd = before
			} else {
				rangeStart = after
			}
		}
	}

	out := "[" + rangeStart + " TO " + rangeEnd + "]"

	v.debug("[GRAMMAR] range_search_string: [%s]", out)

	return out
}

func (v *SolrParser) visitSearchString(ctx antlr.RuleNode) interface{} {
	v.enterFunction("visitSearchString()")
	defer v.exitFunction()

	// search_string : search_string boolean_op search_string
	// n.b.  this returns an array of three objects.
	// the first or third of the objects could be either a string or an array of three objects

	if ctx.GetChildCount() == 3 {
		switch ctx.GetChild(1).(type) {
		case *Boolean_opContext:
			var out []interface{}

			out = append(out, v.visit(ctx.GetChild(0)))
			out = append(out, v.visit(ctx.GetChild(1)))
			out = append(out, v.visit(ctx.GetChild(2)))

			v.debug("[GRAMMAR] search_string (boolean): [%v]", out)

			return out
		}
	}

	// search_string : LPAREN search_string RPAREN

	if ctx.GetChildCount() == 3 {
		switch child0 := ctx.GetChild(0).(type) {
		case antlr.TerminalNode:
			if child0.GetSymbol().GetTokenType() != VirgoQueryLexerLPAREN {
				break
			}

			switch child2 := ctx.GetChild(2).(type) {
			case antlr.TerminalNode:
				if child2.GetSymbol().GetTokenType() != VirgoQueryLexerRPAREN {
					break
				}

				var out interface{}

				out = v.visit(ctx.GetChild(1))

				v.debug("[GRAMMAR] search_string (parentheses): [%v]", out)

				return out
			}
		}
	}

	//               | search_string search_part
	//               | search_part

	out := ""

	for i := 0; i < ctx.GetChildCount(); i++ {
		childResult := v.visit(ctx.GetChild(i))

		rt := reflect.TypeOf(childResult)
		if rt == nil {
			continue
		}

		kind := rt.Kind()

		v.debug("search string component of type %s: [%v]", kind, childResult)

		str := ""

		if kind == reflect.Array || kind == reflect.Slice {
			parts := reflect.ValueOf(childResult)

			str = fmt.Sprintf("%s %s %s", parts.Index(0), parts.Index(1), parts.Index(2))
			v.debug("arr str = [%s]", str)
		} else {
			str = childResult.(string)
			v.debug("str str = [%s]", str)
		}

		// log.Printf("I %d, STR [%s]", i, str)
		if i > 0 && str != "\\\"" && out != "\\\"" {
			// log.Printf("ADD SPACE TO %s", out)
			out += " "
		}

		out += str
	}

	v.debug("[GRAMMAR] search_string: [%s]", out)

	return out
}

func (v *SolrParser) visitChildren(node antlr.RuleNode) interface{} {
	v.enterFunction("visitChildren()")
	defer v.exitFunction()

	out := ""

	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		str := v.visit(child).(string)

		v.debug("[CHILDREN] got child: [%s]", str)

		if i > 0 {
			out += " "
		}

		out += str
	}

	v.debug("[CHILDREN] all children: [%s]", out)

	return out
}

func (v *SolrParser) visitTerminal(terminal antlr.TerminalNode) interface{} {
	v.enterFunction("visitTerminal()")
	defer v.exitFunction()

	out := ""

	switch terminal.GetSymbol().GetTokenType() {
	case VirgoQueryLexerQUOTE:
		// singly-escaped quote character that begins/ends a quoted string
		// within a query fragment, which itself is quoted.
		out = escapeQuote

	case VirgoQueryLexerBOOLEAN:
		// do not escape this; it is an explicit boolean operator
		out = terminal.GetText()

	case VirgoQueryLexerDATE_STRING:
		// do not escape this; it contains some solr special chars/keywords
		out = strings.ReplaceAll(terminal.GetText(), "/", "-")

	default:
		out = terminal.GetText()

		if v.escapeText == true {
			out = v.escapeSolrSpecialCharacters(out)
		}
	}

	v.debug("[GRAMMAR] terminal: [%s]", out)

	return out
}

func (v *SolrParser) initializeParser(query string) (parseCtx parseContext, err error) {
	parseCtx = parseContext{
		lexerErrorListener:  virgoErrorListener{name: "lexer", valid: true, quiet: true},
		parserErrorListener: virgoErrorListener{name: "parser", valid: true, quiet: true},
	}

	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
			log.Printf("ERROR: [V4QUERY] (recovered): %s", err.Error())
		} else {
			if parseCtx.lexerErrorListener.valid == false {
				err = errors.New(parseCtx.lexerErrorListener.Errors())
				v.debug("ERROR: [V4QUERY] LEXER: %s", err.Error())
				return
			}

			if parseCtx.parserErrorListener.valid == false {
				err = errors.New(parseCtx.parserErrorListener.Errors())
				v.debug("ERROR: PARSER: %s", err.Error())
				return
			}
		}
	}()

	inputStream := antlr.NewInputStream(query)

	lexer := NewVirgoQueryLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&parseCtx.lexerErrorListener)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parseCtx.parser = NewVirgoQuery(tokenStream)
	parseCtx.parser.RemoveErrorListeners()
	parseCtx.parser.AddErrorListener(&parseCtx.parserErrorListener)

	parseCtx.parser.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	parseCtx.tree = parseCtx.parser.Query()

	return
}

type convertResponse struct {
	out string
	err error
}

func convert(v *SolrParser, query string) (res convertResponse) {
	start := time.Now()

	v.FieldValues = make(map[string][]string)

	parseCtx, parseErr := v.initializeParser(query)

	if parseErr != nil {
		res.err = parseErr
		return
	}

	v.debug("Convert to Solr: %s", query)

	raw := v.visit(parseCtx.tree)

	if raw == nil {
		res.err = errors.New("Empty query")
		v.debug("ERROR: %s", res.err.Error())
		return
	}

	res.out = raw.(string)

	v.debug("SUCCESS: QUERY: %s", res.out)

	elapsedMS := int64(time.Since(start) / time.Millisecond)

	log.Printf("INFO: [V4QUERY] Conversion elapsed ms: %d", elapsedMS)

	return
}

func convertWithTimeout(v *SolrParser, query string, timeout int) (string, error) {
	ch := make(chan convertResponse, 1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- convert(v, query)
		}
	}()

	select {
	case res := <-ch:
		return res.out, res.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// ConvertToSolr will convert a v4 query string into solr query string.
func ConvertToSolr(query string) (string, error) {
	v := SolrParser{}
	res := convert(&v, query)
	return res.out, res.err
}

// ConvertToSolrWithParser will convert a v4 query string into solr query string.
// The passed SolrPaser struct will be poulated with details about the items parsed.
func ConvertToSolrWithParser(v *SolrParser, query string) (string, error) {
	res := convert(v, query)
	return res.out, res.err
}

// ConvertToSolrWithTimeout will convert a v4 query string into solr query string with a timeout.
func ConvertToSolrWithTimeout(query string, timeout int) (string, error) {
	v := SolrParser{}
	return convertWithTimeout(&v, query, timeout)
}

// ConvertToSolrWithParserAndTimeout will convert a v4 query string into solr query string with a timeout.
// The passed SolrPaser  struct will be poulated with details about the items parsed.
func ConvertToSolrWithParserAndTimeout(v *SolrParser, query string, timeout int) (string, error) {
	return convertWithTimeout(v, query, timeout)
}

// ParseTree will generate a pretty parse tree
func ParseTree(query string) string {
	v := SolrParser{}

	parseCtx, _ := v.initializeParser(query)

	return printSyntaxTree(parseCtx.parser, parseCtx.tree)
}

func printSyntaxTree(parser antlr.Parser, root antlr.Tree) string {
	return printSyntaxTreeAtBranch(parser, root, 0)
}

func printSyntaxTreeAtBranch(parser antlr.Parser, branch antlr.Tree, offset int) string {
	line := strings.Repeat("  ", offset)

	line += fmt.Sprintf("%s\n", antlr.TreesGetNodeText(branch, parser.GetRuleNames(), parser))

	switch branch.(type) {
	case antlr.ParserRuleContext:
		for _, child := range branch.GetChildren() {
			line += printSyntaxTreeAtBranch(parser, child, offset+1)
		}
	}

	return line
}
