package v4parser

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// solrParser will parse the query string into a format compatable with a solr search
type solrParser struct{}

func (v *solrParser) visit(tree antlr.Tree) interface{} {
	switch val := tree.(type) {
	case antlr.RuleNode:
		return v.visitRuleNode(val)
	case antlr.TerminalNode:
		return v.visitTerminal(val)
	}
	return nil
}

func (v *solrParser) visitRuleNode(rule antlr.RuleNode) interface{} {
	switch rule.(type) {
	case *QueryContext:
		return v.visitQuery(rule)
	case *Query_partsContext:
		return v.visitQueryParts(rule)
	case *Field_queryContext:
		return v.visitFieldQuery(rule)
	case *Field_typeContext:
		return v.visitFieldType(rule)
	case *Search_stringContext:
		return v.visitSearchString(rule)
	default:
		return v.visitChildren(rule)
	}
}

func (v *solrParser) visitQuery(query antlr.RuleNode) interface{} {
	first := query.GetChild(0)
	result := v.visit(first)
	return (result)
}

func (v *solrParser) visitQueryParts(ctx antlr.RuleNode) interface{} {
	//  query_parts : query_parts boolean_op query_parts
	if ctx.GetChildCount() == 3 {
		switch ctx.GetChild(1).(type) {
		case *Boolean_opContext:
			// log.Printf("... boolean")
			query1 := v.visit(ctx.GetChild(0))
			queryop := v.visit(ctx.GetChild(1))
			query2 := v.visit(ctx.GetChild(2))
			out := fmt.Sprintf(" (%s%s%s) ", query1, queryop, query2)
			// log.Printf("   QP Bool: %s", out)
			return out
		}

		// query_parts : LPAREN query_parts RPAREN
		switch ctx.GetChild(0).(type) {
		case antlr.TerminalNode:
			// log.Printf("... terminal")
			query1 := v.visit(ctx.GetChild(1))
			out := fmt.Sprintf(" (%s) ", query1)
			// log.Printf("   QP terminal: %s", out)
			return out
		}
	}

	// query_parts: field_query
	first := ctx.GetChild(0)
	result := v.visit(first)
	// log.Printf("QueryPart: %s", result)
	return result
}

func (v *solrParser) visitFieldQuery(ctx antlr.RuleNode) interface{} {
	// field_query : field_type COLON LBRACE search_string RBRACE
	//   (_query_:"{!edismax qf=$title_qf pf=$title_pf}(certification of teachers )"
	fieldType := v.visit(ctx.GetChild(0))
	query := v.visit(ctx.GetChild(3))
	out := v.expand("", fieldType.(string), query)
	// log.Printf("Field Q: %s", out)
	return out
}

// This expands the search_string inside a field_query
// if there are no boolean operators in the search_string it merely takes the words and characters
// of the search_string and adds the proper query fragment syntax as required by solrParser.
// e.g. going from this:     title : {susan sontag}
//      to this:             _query_:"{!edismax qf=$title_qf pf=$title_pf}(susan sontag)"
// if there ARE boolean operators in the search_string they need to be (recursively) expanded and
// wrapped with parentheses to ensure proper operator precedence
// e.g. this field_query :   title : {"susan sontag" OR music title}
//expands to this:           (_query_:"{!edismax qf=$title_qf pf=$title_pf}(\" susan sontag \")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)")
func (v *solrParser) expand(inStr string, fieldType string, query interface{}) string {
	// log.Printf("==> Expand inStr %s for field %s, q: %s", inStr, fieldType, query)
	rt := reflect.TypeOf(query)
	kind := rt.Kind()
	if kind == reflect.Array || kind == reflect.Slice {
		parts := reflect.ValueOf(query)
		out := fmt.Sprintf("%s(", inStr)
		out = v.expand(out, fieldType, parts.Index(0))
		out += fmt.Sprintf("%s", parts.Index(1))
		out = v.expand(out, fieldType, parts.Index(2))
		out = fmt.Sprintf("%s)", out)
		// log.Printf("EXPAND TYPE %s to %s", kind, out)
		return out
	}

	out := fmt.Sprintf(`%s_query_:"{!edismax%s}(%s)"`, inStr, fieldType, query)
	// log.Printf("EXPANDED: %s", out)
	return out
}

func (v *solrParser) visitFieldType(ctx antlr.RuleNode) interface{} {
	// field_type : TITLE | AUTHOR | SUBJECT | KEYWORD
	childTree := ctx.GetChild(0)
	t := childTree.GetPayload().(*antlr.CommonToken)
	fieldType := t.GetText()
	if fieldType == "title" || fieldType == "subject" || fieldType == "author" {
		qf := " qf=$" + fieldType + "_qf"
		pf := " pf=$" + fieldType + "_pf"
		out := qf + pf
		// log.Printf("FieldType: %s", out)
		return out
	}
	return ""
}

func (v *solrParser) visitSearchString(ctx antlr.RuleNode) interface{} {
	// search_string : search_string boolean_op search_string
	// n.b.  this returns an array of three objects.
	// the first or third of the objects could be either a string or an array of three objects
	if ctx.GetChildCount() == 3 {
		switch ctx.GetChild(1).(type) {
		case *Boolean_opContext:
			// log.Printf("...boolen")
			var out []interface{}
			out = append(out, v.visit(ctx.GetChild(0)))
			out = append(out, v.visit(ctx.GetChild(1)))
			out = append(out, v.visit(ctx.GetChild(2)))
			// log.Printf("Bool Str: %s", out)
			return out
		}
	}

	// search_string : LPAREN search_string RPAREN
	//               | search_string search_part
	//               | search_part
	out := ""
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if i > 0 {
			out += " "
		}
		childResult := v.visit(child).(string)
		// log.Printf("   SEARCH Child %d:%s", i, childResult)
		out += childResult
	}
	// log.Printf("search string: [%s]", out)
	return out
}

func (v *solrParser) visitChildren(node antlr.RuleNode) interface{} {
	out := ""
	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		childResult := v.visit(child).(string)
		// log.Printf("    visit children result: [%s]", childResult)
		if out != "" && childResult != `\"` && out != `\"` {
			out += " "
		}
		out += childResult
	}

	return out
}

func (v *solrParser) visitTerminal(terminal antlr.TerminalNode) interface{} {
	if terminal.GetSymbol().GetTokenType() == VirgoQueryLexerQUOTE {
		return `\"`
	} else if terminal.GetSymbol().GetTokenType() == VirgoQueryLexerBOOLEAN {
		return fmt.Sprintf(" %s ", terminal.GetText())
	}
	return terminal.GetText()
}

// ConvertToSolr convert a v4 query string into solr
func ConvertToSolr(src string) (string, error) {
	// EXAMPLE: `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	// SOLR: ( ( ((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\" susan sontag \")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)")
	//              AND _query_:"{!edismax}(Maunsell)") )  OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")
	//
	log.Printf("Convert to Solr: %s", src)

	sp := solrParser{}
	is := antlr.NewInputStream(src)
	lexer := NewVirgoQueryLexer(is)
	lexer.RemoveErrorListeners()
	lel := lexerErrorLister{}
	lel.valid = true
	lexer.AddErrorListener(&lel)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	queryTree := NewVirgoQuery(stream)
	raw := sp.visit(queryTree.Query())

	if lel.valid == false {
		return "", errors.New(lel.Errors())
	}

	out := strings.TrimSpace(raw.(string))
	var re = regexp.MustCompile(`\s+`)
	out = re.ReplaceAllString(out, " ")
	re = regexp.MustCompile(`\s*\(\s*\(`)
	out = re.ReplaceAllString(out, "((")
	re = regexp.MustCompile(`\s*\)\s*\)`)
	out = re.ReplaceAllString(out, "))")
	return out, nil
}
