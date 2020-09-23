package v4parser

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/uvalib/antlr4/runtime/Go/antlr"
)

//virgoErrorListener implements the antlr.ErrorListener interface
//and is used by both the lexer and the parser
type virgoErrorListener struct {
	name     string
	timer    *time.Timer
	quiet    bool
	valid    bool
	errors   []string
	warnings []string
}

func (eh *virgoErrorListener) checkForTimeout() {
	if eh.timer == nil {
		return
	}

	select {
	case <-eh.timer.C:
		panic("timeout exceeded")
	default:
	}
}

func (eh *virgoErrorListener) LogError(msg string) {
	if eh.quiet == false {
		log.Printf("ERROR: [V4QUERY] (%s): %s", eh.name, msg)
	}

	eh.errors = append(eh.errors, msg)
	eh.valid = false

	eh.checkForTimeout()
}

func (eh *virgoErrorListener) LogWarning(msg string) {
	if eh.quiet == false {
		log.Printf("WARNING: [V4QUERY] (%s): %s", eh.name, msg)
	}

	eh.warnings = append(eh.warnings, msg)

	eh.checkForTimeout()
}

func (eh *virgoErrorListener) Errors() string {
	return strings.Join(eh.errors, "; ")
}

func (eh *virgoErrorListener) Warnings() string {
	return strings.Join(eh.warnings, "; ")
}

func (eh *virgoErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	eh.LogError(fmt.Sprintf("Line %d, Column %d: %s", line, column, msg))

	eh.checkForTimeout()
}

func (eh *virgoErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex,
	stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	eh.LogWarning("Ambiguous query")

	eh.checkForTimeout()
}

func (eh *virgoErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA,
	startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	eh.LogWarning("Lexer full context")

	eh.checkForTimeout()
}

func (eh *virgoErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA,
	startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	eh.LogWarning("Lexer context sensitivity")

	eh.checkForTimeout()
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

// validate will validate an input string and return true or false, with an optional timeout
func validate(src string, timeout int) (valid bool, errors string) {
	start := time.Now()

	v := validator{valid: true}
	l := virgoErrorListener{name: "lexer", valid: true}
	p := virgoErrorListener{name: "parser", valid: true}

	if timeout > 0 {
		l.timer = time.NewTimer(time.Duration(timeout) * time.Second)
		p.timer = time.NewTimer(time.Duration(timeout) * time.Second)
	}

	defer func() {
		elapsedMS := int64(time.Since(start) / time.Millisecond)

		if x := recover(); x != nil {
			valid = false
			errors = fmt.Sprintf("%v", x)
			log.Printf("ERROR: [V4QUERY] (recovered): %s", errors)
		} else {
			valid = v.valid && l.valid && p.valid

			if valid == false {
				errors = strings.Join([]string{"lexer: [" + l.Errors() + "]", "parser: [" + p.Errors() + "]"}, "; ")
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

	antlr.ParseTreeWalkerDefault.Walk(&v, parser.Query())

	return
}

// Validate will run validation with no timeout
func Validate(src string) (bool, string) {
	return validate(src, 0)
}

// ValidateWithTimeout will run validation with a timeout
func ValidateWithTimeout(src string, timeout int) (bool, string) {
	return validate(src, timeout)
}
