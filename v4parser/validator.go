package v4parser

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
)

// virgoErrorListener implements the antlr.ErrorListener interface
// and is used by both the lexer and the parser
type virgoErrorListener struct {
	name     string
	quiet    bool
	valid    bool
	errors   []string
	warnings []string
}

func (eh *virgoErrorListener) LogError(msg string) {
	if eh.quiet == false {
		log.Printf("ERROR: [V4QUERY] (%s): %s", eh.name, msg)
	}

	eh.errors = append(eh.errors, msg)
	eh.valid = false
}

func (eh *virgoErrorListener) LogWarning(msg string) {
	if eh.quiet == false {
		log.Printf("WARNING: [V4QUERY] (%s): %s", eh.name, msg)
	}

	eh.warnings = append(eh.warnings, msg)
}

func (eh *virgoErrorListener) Errors() string {
	return strings.Join(eh.errors, "; ")
}

func (eh *virgoErrorListener) Warnings() string {
	return strings.Join(eh.warnings, "; ")
}

func (eh *virgoErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any,
	line, column int, msg string, e antlr.RecognitionException) {
	log.Printf("INFO: [V4QUERY] (%s) Line %d, Column %d: %s", eh.name, line, column, msg)
}

func (eh *virgoErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex,
	stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	eh.LogWarning("Ambiguous query")
}

func (eh *virgoErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA,
	startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	eh.LogWarning("Lexer full context")
}

func (eh *virgoErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA,
	startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	eh.LogWarning("Lexer context sensitivity")
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

type validateResponse struct {
	valid  bool
	errors string
}

// validate will validate an input string and return true or false
func validate(src string) (res validateResponse) {
	start := time.Now()

	v := validator{valid: true}
	l := virgoErrorListener{name: "lexer", valid: true}
	p := virgoErrorListener{name: "parser", valid: true}

	defer func() {
		elapsedMS := int64(time.Since(start) / time.Millisecond)

		if x := recover(); x != nil {
			res.valid = false
			res.errors = fmt.Sprintf("%v", x)
			log.Printf("ERROR: [V4QUERY] (recovered): %s", res.errors)
		} else {
			res.valid = v.valid && l.valid && p.valid

			if res.valid == false {
				res.errors = strings.Join([]string{"lexer: [" + l.Errors() + "]", "parser: [" + p.Errors() + "]"}, "; ")
			}
		}

		log.Printf("INFO: [V4QUERY] Validation elapsed ms: %d", elapsedMS)
	}()

	log.Printf("INFO: [V4QUERY] Validate: %s", src)

	is := antlr.NewInputStream(src)

	lexer := NewVirgoQueryLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&l)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := NewVirgoQuery(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(&p)

	parser.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	antlr.ParseTreeWalkerDefault.Walk(&v, parser.Query())

	return
}

// Validate will run validation with no timeout
func Validate(src string) (bool, string) {
	res := validate(src)
	return res.valid, res.errors
}

// ValidateWithTimeout will run validation with a timeout
func ValidateWithTimeout(src string, timeout int) (bool, string) {
	ch := make(chan validateResponse, 1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- validate(src)
		}
	}()

	select {
	case res := <-ch:
		return res.valid, res.errors
	case <-ctx.Done():
		return false, ctx.Err().Error()
	}
}
