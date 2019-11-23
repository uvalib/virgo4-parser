package v4parser

import (
	"fmt"
	"log"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

//virgoErrorListener implements the antlr.ErrorListener interface
//and is used by both the lexer and the parser
type virgoErrorListener struct {
	valid  bool
	errors []string
}

func (eh *virgoErrorListener) Errors() string {
	return strings.Join(eh.errors, ", ")
}

func (eh *virgoErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	eh.valid = false
	eh.errors = append(eh.errors, fmt.Sprintf("Line %d, Column %d: %s", line, column, msg))
}

func (eh *virgoErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex,
	stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	msg := "Ambiguous query"
	log.Printf(msg)
//	eh.valid = false
	eh.errors = append(eh.errors, msg)
}

func (eh *virgoErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA,
	startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	log.Printf("LEXER FULL CONTEXT?")
//	eh.valid = false
}

func (eh *virgoErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA,
	startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	log.Printf("LEXER CONTEXT SENSITIVITY")
//	eh.valid = false
}

// validator is an implementation of the v4Parser that just checks for errors.
type validator struct {
	*BaseVirgoQueryListener
	valid bool
}

// VisitErrorNode is called when the parser encounters an error
func (v *validator) VisitErrorNode(node antlr.ErrorNode) {
	v.valid = false
}

// Validate will validate an input string and return true or false
func Validate(src string) (bool, string) {
	log.Printf("Validate %s", src)

	v := validator{}
	v.valid = true

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
	parser.AddErrorListener(&pel)

	antlr.ParseTreeWalkerDefault.Walk(&v, parser.Query())

	valid := v.valid && lel.valid && pel.valid
	errors := strings.Join([]string{ "lexer: [" + lel.Errors() + "]", "parser: [" + pel.Errors() + "]" }, "; ")

	return valid, errors
}
