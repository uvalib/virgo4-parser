package v4parser

import (
	"errors"
	"fmt"

	"log"
	"reflect"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// SolrParser will parse the query string into a format compatable with a solr search
type SolrParser struct {
	// options
	Debug bool

	// basic info about the parsed query, can be useful to the caller
	Titles      []string
	Authors     []string
	Subjects    []string
	Keywords    []string
	Identifiers []string
}

func (v *SolrParser) debug(format string, args ...interface{}) {
	if v.Debug == false {
		return
	}

	log.Printf(format, args...)
}

func (v *SolrParser) visit(tree antlr.Tree) interface{} {
	switch val := tree.(type) {
	case antlr.RuleNode:
		return v.visitRuleNode(val)
	case antlr.TerminalNode:
		return v.visitTerminal(val)
	}

	return nil
}

func (v *SolrParser) visitRuleNode(rule antlr.RuleNode) interface{} {
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
	first := query.GetChild(0)
	result := v.visit(first)

	return result
}

func (v *SolrParser) visitQueryParts(ctx antlr.RuleNode) interface{} {
	//  query_parts : query_parts boolean_op query_parts
	if ctx.GetChildCount() == 3 {
		switch ctx.GetChild(1).(type) {
		case *Boolean_opContext:
			query1 := v.visit(ctx.GetChild(0))
			queryop := v.visit(ctx.GetChild(1))
			query2 := v.visit(ctx.GetChild(2))

			out := fmt.Sprintf("(%s%s%s)", query1, queryop, query2)

			v.debug("\tquery part (boolean): %s", out)

			return out
		}

		// query_parts : LPAREN query_parts RPAREN
		switch ctx.GetChild(0).(type) {
		case antlr.TerminalNode:
			query1 := v.visit(ctx.GetChild(1))

			out := fmt.Sprintf("(%s)", query1)

			v.debug("\tquery part (terminal): %s", out)

			return out
		}
	}

	// query_parts: field_query
	first := ctx.GetChild(0)
	result := v.visit(first)

	v.debug("\tquery part: %s", result)

	return result
}

func (v *SolrParser) visitFieldQuery(ctx antlr.RuleNode) interface{} {
	// field_query : field_type COLON LBRACE search_string RBRACE
	//             | range_field_type COLON LBRACE range_search_string RBRACE
	//   (_query_:"{!edismax qf=$title_qf pf=$title_pf}(certification of teachers )"

	// grey magic to get field name so we can populate per-field query slices in expand()

	// if this is nil, the query isn't going to parse anyway, so bail now to avoid crashing
	if ctx.GetChild(0).GetChild(0) == nil {
		return ""
	}

	fieldType := ctx.GetChild(0).GetChild(0).GetPayload().(*antlr.CommonToken).GetText()
	fieldAttr := v.visit(ctx.GetChild(0))

	v.debug("\tfield type: %s", fieldType)
	v.debug("\tfield attr: %s", fieldAttr)

	var out string

	switch ctx.GetChild(3).(type) {
	case antlr.TerminalNode:
		out = v.expand("", fieldType, fieldAttr.(string), "*")
	default:
		query := v.visit(ctx.GetChild(3))
		out = v.expand("", fieldType, fieldAttr.(string), query)
	}

	v.debug("\tfield query: %s", out)

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
func (v *SolrParser) expand(inStr string, fieldType string, fieldAttr string, query interface{}) string {
	v.debug("\t==> Expand inStr %s for field %s, q: %s", inStr, fieldType, query)

	rt := reflect.TypeOf(query)
	if rt == nil {
		return ""
	}

	kind := rt.Kind()

	if kind == reflect.Array || kind == reflect.Slice {
		parts := reflect.ValueOf(query)

		out := fmt.Sprintf("%s(", inStr)
		out = v.expand(out, fieldType, fieldAttr, parts.Index(0))
		out += fmt.Sprintf("%s", parts.Index(1))
		out = v.expand(out, fieldType, fieldAttr, parts.Index(2))
		out = fmt.Sprintf("%s)", out)

		v.debug("\texpand type %s to %s", kind, out)

		return out
	}

	val := fmt.Sprintf("%s", query)

	v.debug("\tfieldType = [%s]  query = [%s]  kind = [%s]  val = [%s]", fieldType, query, kind, val)

	switch fieldType {
	case "title":
		v.Titles = append(v.Titles, val)
	case "author":
		v.Authors = append(v.Authors, val)
	case "subject":
		v.Subjects = append(v.Subjects, val)
	case "keyword":
		v.Keywords = append(v.Keywords, val)
	case "identifier":
		v.Identifiers = append(v.Identifiers, val)
	}

	out := fmt.Sprintf(`%s_query_:"{%s}(%s)"`, inStr, fieldAttr, query)

	v.debug("\texpanded: %s", out)

	return out
}

func (v *SolrParser) visitFieldType(ctx antlr.RuleNode) interface{} {
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

	v.debug("\tfield type: %s", out)

	return out
}

func (v *SolrParser) visitRangeFieldType(ctx antlr.RuleNode) interface{} {
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

			v.debug("\tsearch string (boolean): %s", out)

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

				var out []interface{}

				out = append(out, v.visit(ctx.GetChild(1)))

				return out
			}
		}
	}

	//               | search_string search_part
	//               | search_part
	out := ""

	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)

		if i > 0 {
			out += " "
		}

		childResult := v.visit(child)

		str := ""

		switch t := childResult.(type) {
		case string:
			v.debug("\tchild %d type: string", i)
			str = t
		case interface{}:
			v.debug("\tchild %d type: interface{}", i)
			str = fmt.Sprintf("%v", t)
		default:
			v.debug("\tchild %d type: unknown", i)
		}

		v.debug("\tchild %d value: %s", i, str)

		out += str
	}

	v.debug("\tsearch string: %s", out)

	return out
}

func (v *SolrParser) visitChildren(node antlr.RuleNode) interface{} {
	out := ""

	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		childResult := v.visit(child).(string)

		v.debug("\tvisit children result: [%s]", childResult)

		if out != "" && childResult != `\"` && out != `\"` {
			out += " "
		}

		out += childResult
	}

	return out
}

func (v *SolrParser) visitTerminal(terminal antlr.TerminalNode) interface{} {
	if terminal.GetSymbol().GetTokenType() == VirgoQueryLexerQUOTE {
		return `\"`
	} else if terminal.GetSymbol().GetTokenType() == VirgoQueryLexerBOOLEAN {
		return fmt.Sprintf(" %s ", terminal.GetText())
	} else if terminal.GetSymbol().GetTokenType() == VirgoQueryLexerDATE_STRING {
		return fmt.Sprintf("%s", strings.ReplaceAll(terminal.GetText(), "/", "-"))
	}

	return terminal.GetText()
}

// ConvertToSolrWithParser will convert a v4 query string into solr query string. The passed SolrPaser struct will be
// poulated with details about the items parsed into the resulting string
func ConvertToSolrWithParser(sp *SolrParser, src string) (string, error) {
	// EXAMPLE: `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	// SOLR: ( ( ((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\" susan sontag \")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)")
	//              AND _query_:"{!edismax}(Maunsell)") )  OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")

	sp.debug("Convert to Solr: %s", src)

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
	sp := SolrParser{Debug: false}
	return ConvertToSolrWithParser(&sp, src)
}
