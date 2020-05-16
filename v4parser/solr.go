package v4parser

import (
	"errors"
	"fmt"

	"log"
	"reflect"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const EscapeQuote = `\"`
const Quote = `"`
const LParen = `(`
const RParen = `)`

// SolrParser will parse the query string into a format compatable with a solr search
type SolrParser struct {
	// options
	Debug         bool
	DebugCallTree bool // Debug must also be enabled

	// map of collected field values as parsed from the query, can be useful to the caller
	FieldValues map[string][]string

	// internal values
	level int
}

func (v *SolrParser) debug(format string, args ...interface{}) {
	if v.Debug == false {
		return
	}

	line := strings.Repeat(" ", v.level) + format

	log.Printf(line, args...)
}

func (v *SolrParser) enterFunction(funcName string) {
	if v.DebugCallTree == true {
		v.debug("%s {", funcName)
	}

	v.level++
}

func (v *SolrParser) exitFunction(funcName string) {
	v.level--

	if v.DebugCallTree == true {
		v.debug("} // %s", funcName)
	}
}

func (v *SolrParser) visit(tree antlr.Tree) interface{} {
	funcName := "visit()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

	switch val := tree.(type) {
	case antlr.RuleNode:
		return v.visitRuleNode(val)
	case antlr.TerminalNode:
		return v.visitTerminal(val)
	}

	return nil
}

func (v *SolrParser) visitRuleNode(rule antlr.RuleNode) interface{} {
	funcName := "visitRuleNode()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

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
	funcName := "visitQuery()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

	first := query.GetChild(0)
	result := v.visit(first)

	return result
}

func (v *SolrParser) visitQueryParts(ctx antlr.RuleNode) interface{} {
	funcName := "visitQueryParts()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

	//  query_parts : query_parts boolean_op query_parts
	if ctx.GetChildCount() == 3 {
		switch ctx.GetChild(1).(type) {
		case *Boolean_opContext:
			query1 := v.visit(ctx.GetChild(0)).(string)
			queryop := v.visit(ctx.GetChild(1)).(string)
			query2 := v.visit(ctx.GetChild(2)).(string)

			out := fmt.Sprintf("(%s %s %s)", query1, queryop, query2)

			v.debug("query part (boolean): [%s]", out)

			return out
		}
	}

	// query_parts : LPAREN query_parts RPAREN
	if ctx.GetChildCount() == 3 {
		switch ctx.GetChild(0).(type) {
		case antlr.TerminalNode:
			query := v.visit(ctx.GetChild(1)).(string)

			out := fmt.Sprintf("(%s)", query)

			v.debug("query part (terminal): [%s]", out)

			return out
		}
	}

	// query_parts: field_query
	out := v.visit(ctx.GetChild(0))

	if out == nil {
		out = ""
	}

	v.debug("query part: [%v]", out)

	return out
}

func (v *SolrParser) visitFieldQuery(ctx antlr.RuleNode) interface{} {
	funcName := "visitFieldQuery()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

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
		out = v.expand("", fieldName, fieldType, "*")
	default:
		query := v.visit(ctx.GetChild(3))
		out = v.expand("", fieldName, fieldType, query)
	}

	v.debug("field query: [%s]", out)

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
	funcName := "expand()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

	rt := reflect.TypeOf(query)
	if rt == nil {
		v.debug("nil query")
		return ""
	}

	kind := rt.Kind()

	v.debug("[expand] cur string : [%s]", inStr)
	v.debug("[expand] new field  : [%s]", fieldName)
	v.debug("[expand] new query  : [%v]", query)
	v.debug("[expand] query type : [%s]", kind)

	if kind == reflect.Array || kind == reflect.Slice {
		parts := reflect.ValueOf(query)

		out := fmt.Sprintf("%s(", inStr)
		out = v.expand(out, fieldName, fieldType, parts.Index(0).Interface())
		out += fmt.Sprintf(" %s ", parts.Index(1))
		out = v.expand(out, fieldName, fieldType, parts.Index(2).Interface())
		out = fmt.Sprintf("%s)", out)

		v.debug("[expand] new string : [%s]", out)

		return out
	}

	val := fmt.Sprintf("%s", query)

	v.FieldValues[fieldName] = append(v.FieldValues[fieldName], val)

	out := fmt.Sprintf(`%s_query_:"{%s}(%s)"`, inStr, fieldType, val)

	v.debug("[expand] new string : [%s]", out)

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

	specialChars := []string{`+`, `-`, `&&`, `||`, `!`, `{`, `}`, `[`, `]`, `^`, `"`, `~`, `?`, `:`, `/`}

	escaped := strings.ReplaceAll(s, `\`, `\\`)
	for _, c := range specialChars {
		escaped = strings.ReplaceAll(escaped, c, `\\`+c)
	}

	v.debug("[escape] from: [%s]", s)
	v.debug("[escape]   to: [%s]", escaped)

	return escaped
}

func (v *SolrParser) visitFieldType(ctx antlr.RuleNode) interface{} {
	funcName := "visitFieldType()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

	// field_type : TITLE | AUTHOR | SUBJECT | KEYWORD | IDENTIFIER
	childTree := ctx.GetChild(0)
	t := childTree.GetPayload().(*antlr.CommonToken)
	fieldType := t.GetText()

	out := "!edismax"

	// keyword field type doesn't need qf/pf parameter
	if fieldType == "keyword" {
		return out
	}

	// all other field types need qf/pf parameters based on their names
	qf := "qf=$" + fieldType + "_qf"
	pf := "pf=$" + fieldType + "_pf"

	out = out + " " + qf + " " + pf

	v.debug("field type: [%s]", out)

	return out
}

func (v *SolrParser) visitRangeFieldType(ctx antlr.RuleNode) interface{} {
	funcName := "visitRangeFieldType()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

	// range_field_type : DATE
	childTree := ctx.GetChild(0)
	t := childTree.GetPayload().(*antlr.CommonToken)
	fieldType := t.GetText()

	out := ""

	if fieldType == "date" {
		df := "published_daterange"
		out = "!lucene df=" + df
	}

	return out
}

func (v *SolrParser) visitRangeSearchString(ctx antlr.RuleNode) interface{} {
	funcName := "visitRangeSearchString()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

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
			terminal := ctx.GetChild(0).(antlr.TerminalNode)
			if terminal.GetSymbol().GetTokenType() == VirgoQueryLexerBEFORE {
				rangeEnd = value
			} else {
				rangeStart = value
			}
		}
	}

	out := "[" + rangeStart + " TO " + rangeEnd + "]"

	return out
}

func (v *SolrParser) visitSearchString(ctx antlr.RuleNode) interface{} {
	funcName := "visitSearchString()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

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

			v.debug("search string (boolean): [%v]", out)

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

				v.debug("search string (parentheses): [%v]", out)

				return out
			}
		}
	}

	//               | search_string search_part
	//               | search_part
	out := ""
	quoted := false

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

		if str == Quote {
			quoted = true
		}

		out = v.conditionalPad(i, out, str)
	}

	/*
		if quoted == false {
			v.debug("quoting unquoted search string")
			out = Quote + out + Quote
		}
	*/

	v.debug("search string (quoted: %v): [%s]", quoted, out)

	return out
}

func (v *SolrParser) conditionalPad(i int, left string, right string) string {
	out := left

	/*
		if i > 0 &&
			// these checks break solr; spaces seem necessary between quotes/parens
			left != Quote && strings.HasSuffix(left, LParen) == false &&
			right != Quote && strings.HasPrefix(right, RParen) == false {
			out += " "
		}
	*/
	if i > 0 {
		out += " "
	}

	out += right

	v.debug("[pad] from: [%s]", left)
	v.debug("[pad]   to: [%s]", out)

	return out
}

func (v *SolrParser) visitChildren(node antlr.RuleNode) interface{} {
	funcName := "visitChildren()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

	out := ""

	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		str := v.visit(child).(string)

		v.debug("visit children result: [%s]", str)

		out = v.conditionalPad(i, out, str)

		v.debug("out now: [%s]", out)
	}

	return out
}

func (v *SolrParser) visitTerminal(terminal antlr.TerminalNode) interface{} {
	funcName := "visitTerminal()"
	v.enterFunction(funcName)
	defer v.exitFunction(funcName)

	if terminal.GetSymbol().GetTokenType() == VirgoQueryLexerQUOTE {
		return EscapeQuote
	} else if terminal.GetSymbol().GetTokenType() == VirgoQueryLexerBOOLEAN {
		return terminal.GetText()
	} else if terminal.GetSymbol().GetTokenType() == VirgoQueryLexerDATE_STRING {
		return fmt.Sprintf("%s", strings.ReplaceAll(terminal.GetText(), "/", "-"))
	}

	return v.escapeSolrSpecialCharacters(terminal.GetText())
}

// ConvertToSolrWithParser will convert a v4 query string into solr query string. The passed SolrPaser struct will be
// poulated with details about the items parsed into the resulting string
func ConvertToSolrWithParser(sp *SolrParser, src string) (string, error) {
	// EXAMPLE: `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	// SOLR: ( ( ((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\" susan sontag \")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)")
	//              AND _query_:"{!edismax}(Maunsell)") )  OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")

	sp.debug("Convert to Solr: %s", src)

	sp.FieldValues = make(map[string][]string)

	is := antlr.NewInputStream(src)

	lexer := NewVirgoQueryLexer(is)
	lexer.RemoveErrorListeners()

	lel := virgoErrorListener{}
	lel.valid = true
	lexer.AddErrorListener(&lel)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := NewVirgoQuery(stream)
	parser.RemoveErrorListeners()

	pel := virgoErrorListener{}
	pel.valid = true
	parser.AddErrorListener(&lel)

	raw := sp.visit(parser.Query())

	if lel.valid == false {
		err := errors.New(lel.Errors())
		sp.debug("ERROR: LEXER: %s", err.Error())
		return "", err
	}

	if pel.valid == false {
		err := errors.New(pel.Errors())
		sp.debug("ERROR: PARSER: %s", err.Error())
		return "", err
	}

	if raw == nil {
		err := errors.New("Empty query")
		sp.debug("ERROR: %s", err.Error())
		return "", err
	}

	out := raw.(string)

	sp.debug("SUCCESS: QUERY: %s", out)

	return out, nil
}

// ConvertToSolr will convert a v4 query string into solr query string.
func ConvertToSolr(src string) (string, error) {
	sp := SolrParser{}
	return ConvertToSolrWithParser(&sp, src)
}

// ParseTree will generate a pretty parse tree
func ParseTree(src string) string {
	is := antlr.NewInputStream(src)

	lexer := NewVirgoQueryLexer(is)
	lexer.RemoveErrorListeners()

	lel := virgoErrorListener{}
	lel.valid = true
	lexer.AddErrorListener(&lel)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := NewVirgoQuery(stream)
	parser.RemoveErrorListeners()

	pel := virgoErrorListener{}
	pel.valid = true
	parser.AddErrorListener(&lel)

	tree := parser.Query()

	return printSyntaxTree(parser, tree)
}

func printSyntaxTree(parser antlr.Parser, root antlr.Tree) string {
	return recursive(parser, root, 0)
}

func recursive(parser antlr.Parser, branch antlr.Tree, offset int) string {
	line := strings.Repeat("  ", offset)

	line += fmt.Sprintf("%s\n", antlr.TreesGetNodeText(branch, parser.GetRuleNames(), parser))

	switch branch.(type) {
	case antlr.ParserRuleContext:
		for _, child := range branch.GetChildren() {
			line += recursive(parser, child, offset+1)
		}
	}

	return line
}
