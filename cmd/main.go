package main

import (
	"log"

	"github.com/uvalib/virgo4-parser/v4parser"
)

// type v4Listener struct {
// 	*v4parser.BaseVirgoQueryListener
// 	searchStack []string
// }

// func (s *v4Listener) EnterSearch_string(ctx *v4parser.Search_stringContext) {
// 	log.Printf("START Search String")
// 	s.searchStack = append(s.searchStack, "")
// 	log.Printf("search stack size %d", len(s.searchStack))
// }
// func (s *v4Listener) ExitSearch_string(ctx *v4parser.Search_stringContext) {
// 	log.Printf("EXIT Search String")
// 	// pop last node from stack
// 	out := s.searchStack[len(s.searchStack)-1]
// 	s.searchStack = s.searchStack[:len(s.searchStack)-1]
// 	log.Printf("SEARCH STRING: %s", out)
// }
// func (s *v4Listener) ExitSearch_part(ctx *v4parser.Search_partContext) {
// 	st := ctx.GetStop()
// 	tt := st.GetTokenType()
// 	// log.Printf("START %s STOP %s", st.GetText(), ctx.GetStop().GetText())
// 	t := ctx.GetToken(tt, 0)
// 	val := t.GetText()
// 	qs := s.searchStack[len(s.searchStack)-1]
// 	s.searchStack = s.searchStack[:len(s.searchStack)-1]
// 	if val != `"` {
// 		if qs != "" {
// 			qs += " "
// 		}
// 		qs += val
// 	} else {
// 		qs = fmt.Sprintf(`"%s"`, qs)
// 	}
// 	s.searchStack = append(s.searchStack, qs)
// }

// func (s *v4Listener) ExitField_type(ctx *v4parser.Field_typeContext) {
// 	log.Printf("ExitField_type: %s", ctx.GetText())
// }
// func (s *v4Listener) ExitField_query(ctx *v4parser.Field_queryContext) {
// 	log.Printf("ExitField_query: %s", ctx.GetText())
// }
// func (s *v4Listener) ExitBoolean_op(ctx *v4parser.Boolean_opContext) {
// 	log.Printf("ExitBoolean_op: %s", ctx.GetText())
// }

/**
 * MAIN
 */
func main() {
	// test := `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	// /*
	// 	( ( ((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\" susan sontag \")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)") )  OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")
	// */
	// // test := "title: {bannanas}"
	// // bad := `( target : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	// log.Printf("Test virgo4 parser with good sample input stream %s", test)
	// is := antlr.NewInputStream(test)
	// lexer := v4parser.NewVirgoQueryLexer(is)
	// stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// // LISTENER
	// vq := v4parser.NewVirgoQuery(stream)
	// q := vq.Query()
	// antlr.ParseTreeWalkerDefault.Walk(&v4Listener{}, q)
	simple := "title: {bannanas}"
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
}
