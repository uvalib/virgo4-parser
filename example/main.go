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
	return ""
}

func (v *solrVisitor) visitQuery(query antlr.RuleNode) interface{} {
	log.Printf("VISIT QUERY")
	first := query.GetChild(0)
	result := v.Visit(first)
	return (result)
}

func (v *solrVisitor) visitQueryParts(ctx antlr.RuleNode) interface{} {
	log.Printf("VISIT QUERY PARTS")
	//  query_parts : query_parts boolean_op query_parts
	if ctx.GetChildCount() == 3 {
		switch ctx.(type) {
		case *v4parser.Boolean_opContext:
			log.Printf("... boolean")
			query1 := v.Visit(ctx.GetChild(0))
			queryop := v.Visit(ctx.GetChild(1))
			query2 := v.Visit(ctx.GetChild(2))
			out := fmt.Sprintf(" (%s%s%s) ", query1, queryop, query2)
			return out
		case antlr.TerminalNode:
			log.Printf("... terminal")
			query1 := v.Visit(ctx.GetChild(1))
			out := fmt.Sprintf(" (%s) ", query1)
			return out
		}
	}

	// query_parts: field_query
	log.Printf("... field")
	first := ctx.GetChild(0)
	result := v.Visit(first)
	return result
}

func (v *solrVisitor) visitFieldQuery(ctx antlr.RuleNode) interface{} {
	log.Printf("VISIT FIELD QUERY")
	// field_query : field_type COLON LBRACE search_string RBRACE
	//   (_query_:"{!edismax qf=$title_qf pf=$title_pf}(certification of teachers )"
	fieldType := v.Visit(ctx.GetChild(0))
	query := v.Visit(ctx.GetChild(3))
	out := v.expand("", fieldType.(string), query)
	return out
}

func (v *solrVisitor) visitFieldType(ctx antlr.RuleNode) interface{} {
	log.Printf("VISIT FIELD TYPE")
	// field_type : TITLE | AUTHOR | SUBJECT | KEYWORD
	childTree := ctx.GetChild(0)
	t := childTree.GetPayload().(*antlr.CommonToken)
	fieldType := t.GetText()
	if fieldType == "title" || fieldType == "subject" || fieldType == "author" {
		qf := " qf=$" + fieldType + "_qf"
		pf := " pf=$" + fieldType + "_pf"
		out := qf + pf
		log.Printf(" field type: %s", out)
		return out
	}
	return ""
}

func (v *solrVisitor) visitSearchString(ctx antlr.RuleNode) interface{} {
	log.Printf("VISIT SEARCH")
	return ""
}

func (v *solrVisitor) visitChildren(ctx antlr.RuleNode) interface{} {
	log.Printf("VISIT CHILDREN")
	// Value result = null;
	//      int n = node.getChildCount();
	//      for (int i=0; i<n; i++)
	//      {
	//          ParseTree c = node.getChild(i);
	//          Value childResult = this.visit(c);
	//          result = aggregateResult(result, childResult);
	//      }

	//      return result;
	return ""
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
		// Value[] parts = query.asArray();
		// sb.append("(");
		// expand(sb, fieldType, parts[0]);
		// sb.append(parts[1].asString());
		// expand(sb, fieldType, parts[2]);
		// sb.append(")");
		// TODO
		log.Printf("EXPAND TYPE %s", kind)
		return ""
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
	v := solrVisitor{}
	is := antlr.NewInputStream(simple)
	lexer := v4parser.NewVirgoQueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	queryTree := v4parser.NewVirgoQuery(stream)
	v.Visit(queryTree.Query())

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
