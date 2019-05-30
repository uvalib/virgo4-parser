package main

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/uvalib/virgo4-parser/v4parser"
)

type v4Listener struct {
	*v4parser.BaseVirgoQueryListener
}

func (s *v4Listener) EnterSearch_string(ctx *v4parser.Search_stringContext) {
	log.Printf("START Search String %s", ctx.GetPayload())
}
func (s *v4Listener) ExitSearch_string(ctx *v4parser.Search_stringContext) {
	log.Printf("EXIT Search String %s", ctx.GetPayload())
}
func (s *v4Listener) ExitSearch_part(ctx *v4parser.Search_partContext) {
	// out
	//  t := ctx.GetTokens(ctx.Get
	// x := ctx.GetRuleIndex()
	// t := ctx.GetChildren()
	// log.Printf("COUNT %d", ctx.GetChildCount())
	st := ctx.GetStop()
	tt := st.GetTokenType()
	// log.Printf("START %s STOP %s", st.GetText(), ctx.GetStop().GetText())
	t := ctx.GetToken(tt, 0)
	log.Printf("TOKEN %s:%s", t.GetText())
	// ct
	// for i := 0; i < ctx.GetChildCount(); i++ {
	// 	if i > 0 {
	// 		out += " "
	// 	}
	// 	out += ctx.GetChild(i)
	// }
	log.Printf("ExitSearch_part: %s", ctx.GetText())
}

// func (s *v4Listener) ExitField_type(ctx *v4parser.Field_typeContext) {
// 	log.Printf("ExitField_type: %s", ctx.GetText())
// }
// func (s *v4Listener) ExitField_query(ctx *v4parser.Field_queryContext) {
// 	log.Printf("ExitField_query: %s", ctx.GetText())
// }
// func (s *v4Listener) ExitBoolean_op(ctx *v4parser.Boolean_opContext) {
// 	log.Printf("ExitBoolean_op: %s", ctx.GetText())
// }
// ExitEveryRule is called when any rule is exited.
func (s *v4Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// log.Printf("ExitEveryRule: %+v:%s", ctx.GetRuleContext().GetRuleIndex(), ctx.GetText())
}

/**
 * MAIN
 */
func main() {
	test := `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	/*
		( ( ((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\" susan sontag \")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)") )  OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")
	*/
	// test := "title: {bannanas}"
	// bad := `( target : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	log.Printf("Test virgo4 parser with good sample input stream %s", test)
	is := antlr.NewInputStream(test)
	lexer := v4parser.NewVirgoQueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// LISTENER
	vq := v4parser.NewVirgoQuery(stream)
	q := vq.Query()
	antlr.ParseTreeWalkerDefault.Walk(&v4Listener{}, q)

}
