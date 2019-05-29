package main

import (
	"fmt"
	"log"

   "github.com/uvalib/virgo4-parser/v4parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	test := `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	log.Printf("Test virgo4 parser with sample input stream %s", test)
	is := antlr.NewInputStream(test)
	lexer := v4parser.NewVirgoQueryLexer(is)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
