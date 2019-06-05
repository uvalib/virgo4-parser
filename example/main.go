package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/uvalib/virgo4-parser/v4parser"
)

type solrVisitor struct{}

func (v *solrVisitor) Visit(tree antlr.Tree) interface{} {
	switch val := tree.(type) {
	case antlr.RuleNode:
		return v.visitRuleNode(val)
	case antlr.TerminalNode:
		return v.visitTerminal(val)
	}
	return nil
}

func (v *solrVisitor) visitRuleNode(rule antlr.RuleNode) interface{} {
	switch rule.(type) {
	case *v4parser.QueryContext:
		return v.visitQuery(rule)
	case *v4parser.Query_partsContext:
		return v.visitQueryParts(rule)
	case *v4parser.Field_queryContext:
		return v.visitFieldQuery(rule)
	case *v4parser.Field_typeContext:
		return v.visitFieldType(rule)
	case *v4parser.Search_stringContext:
		return v.visitSearchString(rule)
	default:
		return v.visitChildren(rule)
	}
}

func (v *solrVisitor) visitQuery(query antlr.RuleNode) interface{} {
	first := query.GetChild(0)
	result := v.Visit(first)
	return (result)
}

func (v *solrVisitor) visitQueryParts(ctx antlr.RuleNode) interface{} {
	//  query_parts : query_parts boolean_op query_parts
	if ctx.GetChildCount() == 3 {
		switch ctx.GetChild(1).(type) {
		case *v4parser.Boolean_opContext:
			log.Printf("... boolean")
			query1 := v.Visit(ctx.GetChild(0))
			queryop := v.Visit(ctx.GetChild(1))
			query2 := v.Visit(ctx.GetChild(2))
			out := fmt.Sprintf(" (%s%s%s) ", query1, queryop, query2)
			log.Printf("   QP Bool: %s", out)
			return out
		}

		// query_parts : LPAREN query_parts RPAREN
		switch ctx.GetChild(0).(type) {
		case antlr.TerminalNode:
			log.Printf("... terminal")
			query1 := v.Visit(ctx.GetChild(1))
			out := fmt.Sprintf(" (%s) ", query1)
			log.Printf("   QP terminal: %s", out)
			return out
		}
	}

	// query_parts: field_query
	first := ctx.GetChild(0)
	result := v.Visit(first)
	log.Printf("QueryPart: %s", result)
	return result
}

func (v *solrVisitor) visitFieldQuery(ctx antlr.RuleNode) interface{} {
	// field_query : field_type COLON LBRACE search_string RBRACE
	//   (_query_:"{!edismax qf=$title_qf pf=$title_pf}(certification of teachers )"
	fieldType := v.Visit(ctx.GetChild(0))
	query := v.Visit(ctx.GetChild(3))
	out := v.expand("", fieldType.(string), query)
	log.Printf("Field Q: %s", out)
	return out
}

func (v *solrVisitor) visitFieldType(ctx antlr.RuleNode) interface{} {
	// field_type : TITLE | AUTHOR | SUBJECT | KEYWORD
	childTree := ctx.GetChild(0)
	t := childTree.GetPayload().(*antlr.CommonToken)
	fieldType := t.GetText()
	if fieldType == "title" || fieldType == "subject" || fieldType == "author" {
		qf := " qf=$" + fieldType + "_qf"
		pf := " pf=$" + fieldType + "_pf"
		out := qf + pf
		log.Printf("FieldType: %s", out)
		return out
	}
	return ""
}

func (v *solrVisitor) visitSearchString(ctx antlr.RuleNode) interface{} {
	log.Printf("VISIT SEARCH")
	// search_string : search_string boolean_op search_string
	// n.b.  this returns an array of three objects.
	// the first or third of the objects could be either a string or an array of three objects
	if ctx.GetChildCount() == 3 {
		switch ctx.GetChild(1).(type) {
		case *v4parser.Boolean_opContext:
			log.Printf("...boolen")
			var out []interface{}
			out = append(out, v.Visit(ctx.GetChild(0)))
			out = append(out, v.Visit(ctx.GetChild(1)))
			out = append(out, v.Visit(ctx.GetChild(2)))
			log.Printf("Bool Str: %s", out)
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
		childResult := v.Visit(child).(string)
		log.Printf("   SEARCH Child %d:%s", i, childResult)
		out += childResult
	}
	log.Printf("search string: [%s]", out)
	return out
}

func (v *solrVisitor) visitChildren(node antlr.RuleNode) interface{} {
	out := ""
	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		childResult := v.Visit(child).(string)
		log.Printf("    visit children result: [%s]", childResult)
		if out != "" && childResult != `"` && out != `"` {
			out += " "
		}
		out += childResult
	}

	return out
}

func (v *solrVisitor) visitTerminal(terminal antlr.TerminalNode) interface{} {
	if terminal.GetSymbol().GetTokenType() == v4parser.VirgoQueryLexerQUOTE {
		return `"`
	} else if terminal.GetSymbol().GetTokenType() == v4parser.VirgoQueryLexerBOOLEAN {
		return fmt.Sprintf(" %s ", terminal.GetText())
	}
	return terminal.GetText()
}

// This expands the search_string inside a field_query
// if there are no boolean operators in the search_string it merely takes the words and characters
// of the search_string and adds the proper query fragment syntax as required by Solr.
// e.g. going from this:     title : {susan sontag}
//      to this:             _query_:"{!edismax qf=$title_qf pf=$title_pf}(susan sontag)"
// if there ARE boolean operators in the search_string they need to be (recursively) expanded and
// wrapped with parentheses to ensure proper operator precedence
// e.g. this field_query :   title : {"susan sontag" OR music title}
//expands to this:           (_query_:"{!edismax qf=$title_qf pf=$title_pf}(\" susan sontag \")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)")
func (v *solrVisitor) expand(inStr string, fieldType string, query interface{}) string {
	rt := reflect.TypeOf(query)
	kind := rt.Kind()
	if kind == reflect.Array || kind == reflect.Slice {
		parts := reflect.ValueOf(query)
		out := fmt.Sprintf("%s(", inStr)
		out = v.expand(out, fieldType, parts.Index(0))
		out += fmt.Sprintf("%s", parts.Index(1))
		out = v.expand(out, fieldType, parts.Index(2))
		out = fmt.Sprintf("%s)", out)
		log.Printf("EXPAND TYPE %s to %s", kind, out)
		return out
	}

	out := fmt.Sprintf(`%s_query_:"{!edismax%s}(%s)"`, inStr, fieldType, query)
	log.Printf("EXPANDED: %s", out)
	return out
}

/**
 * MAIN
 */
func main() {
	log.Printf("Testing out teh validtaion behavior...")
	simple := "title: {bannanas}"
	hard := `title : {"susan sontag" OR music title}`
	v := solrVisitor{}
	is := antlr.NewInputStream(hard)
	lexer := v4parser.NewVirgoQueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	queryTree := v4parser.NewVirgoQuery(stream)
	out := v.Visit(queryTree.Query())
	log.Printf("RESULT: %s", out.(string))
	log.Printf(" ==================================================== ")
	log.Printf(" ==================================================== ")

	validator := v4parser.Validator{}
	valid, errors := validator.Validate(simple)
	if valid == false {
		log.Printf("ERROR: [%s] is not valid, but is should be: %s", simple, errors)
	} else {
		log.Printf("SUCCESS: [%s] is valid", simple)
	}

	bad := "alligator: {bannana"
	valid, errors = validator.Validate(bad)
	if valid == false {
		log.Printf("SUCCESS: [%s] is not valid: %s", bad, errors)
	} else {
		log.Printf("ERROR: [%s] is valid, but it is not", bad)
	}

	test := `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	valid, errors = validator.Validate(test)
	if valid == false {
		log.Printf("ERROR: [%s] is not valid, but should be: %s", test, errors)
	} else {
		log.Printf("SUCCESS: [%s] is valid", test)
	}

	log.Printf("Testing out SOLR conversion behavior...")
	solr := v4parser.Solr{}
	solrQ, err := solr.Convert(simple)
	if err != nil {
		log.Printf("ERROR: Unable to convert [%s] : %s", simple, err.Error())
	} else {
		log.Printf("Converted to %s", solrQ)
	}
}
