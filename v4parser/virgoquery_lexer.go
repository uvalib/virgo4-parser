// Generated from VirgoQueryLexer.g4 by ANTLR 4.7.

package v4parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 161,
	8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6,
	4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12,
	9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9,
	17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22,
	4, 23, 9, 23, 3, 2, 6, 2, 50, 10, 2, 13, 2, 14, 2, 51, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 66, 10, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 149, 10, 21, 3, 21, 3, 21, 3, 22,
	6, 22, 154, 10, 22, 13, 22, 14, 22, 155, 3, 23, 3, 23, 3, 23, 3, 23, 2,
	2, 24, 4, 2, 6, 3, 8, 4, 10, 5, 12, 6, 14, 7, 16, 8, 18, 9, 20, 10, 22,
	11, 24, 12, 26, 13, 28, 14, 30, 2, 32, 2, 34, 15, 36, 16, 38, 2, 40, 17,
	42, 2, 44, 18, 46, 19, 4, 2, 3, 4, 5, 2, 11, 12, 15, 15, 34, 34, 8, 2,
	11, 12, 15, 15, 34, 34, 36, 36, 125, 125, 127, 127, 2, 164, 2, 6, 3, 2,
	2, 2, 2, 8, 3, 2, 2, 2, 2, 10, 3, 2, 2, 2, 2, 12, 3, 2, 2, 2, 2, 14, 3,
	2, 2, 2, 2, 16, 3, 2, 2, 2, 2, 18, 3, 2, 2, 2, 2, 20, 3, 2, 2, 2, 2, 22,
	3, 2, 2, 2, 2, 24, 3, 2, 2, 2, 2, 26, 3, 2, 2, 2, 3, 28, 3, 2, 2, 2, 3,
	30, 3, 2, 2, 2, 3, 32, 3, 2, 2, 2, 3, 34, 3, 2, 2, 2, 3, 36, 3, 2, 2, 2,
	3, 38, 3, 2, 2, 2, 3, 40, 3, 2, 2, 2, 3, 42, 3, 2, 2, 2, 3, 44, 3, 2, 2,
	2, 3, 46, 3, 2, 2, 2, 4, 49, 3, 2, 2, 2, 6, 53, 3, 2, 2, 2, 8, 55, 3, 2,
	2, 2, 10, 65, 3, 2, 2, 2, 12, 67, 3, 2, 2, 2, 14, 69, 3, 2, 2, 2, 16, 75,
	3, 2, 2, 2, 18, 82, 3, 2, 2, 2, 20, 90, 3, 2, 2, 2, 22, 98, 3, 2, 2, 2,
	24, 109, 3, 2, 2, 2, 26, 113, 3, 2, 2, 2, 28, 117, 3, 2, 2, 2, 30, 119,
	3, 2, 2, 2, 32, 123, 3, 2, 2, 2, 34, 127, 3, 2, 2, 2, 36, 129, 3, 2, 2,
	2, 38, 131, 3, 2, 2, 2, 40, 136, 3, 2, 2, 2, 42, 148, 3, 2, 2, 2, 44, 153,
	3, 2, 2, 2, 46, 157, 3, 2, 2, 2, 48, 50, 9, 2, 2, 2, 49, 48, 3, 2, 2, 2,
	50, 51, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 5, 3, 2,
	2, 2, 53, 54, 7, 42, 2, 2, 54, 7, 3, 2, 2, 2, 55, 56, 7, 43, 2, 2, 56,
	9, 3, 2, 2, 2, 57, 58, 7, 67, 2, 2, 58, 59, 7, 80, 2, 2, 59, 66, 7, 70,
	2, 2, 60, 61, 7, 81, 2, 2, 61, 66, 7, 84, 2, 2, 62, 63, 7, 80, 2, 2, 63,
	64, 7, 81, 2, 2, 64, 66, 7, 86, 2, 2, 65, 57, 3, 2, 2, 2, 65, 60, 3, 2,
	2, 2, 65, 62, 3, 2, 2, 2, 66, 11, 3, 2, 2, 2, 67, 68, 7, 60, 2, 2, 68,
	13, 3, 2, 2, 2, 69, 70, 7, 118, 2, 2, 70, 71, 7, 107, 2, 2, 71, 72, 7,
	118, 2, 2, 72, 73, 7, 110, 2, 2, 73, 74, 7, 103, 2, 2, 74, 15, 3, 2, 2,
	2, 75, 76, 7, 99, 2, 2, 76, 77, 7, 119, 2, 2, 77, 78, 7, 118, 2, 2, 78,
	79, 7, 106, 2, 2, 79, 80, 7, 113, 2, 2, 80, 81, 7, 116, 2, 2, 81, 17, 3,
	2, 2, 2, 82, 83, 7, 117, 2, 2, 83, 84, 7, 119, 2, 2, 84, 85, 7, 100, 2,
	2, 85, 86, 7, 108, 2, 2, 86, 87, 7, 103, 2, 2, 87, 88, 7, 101, 2, 2, 88,
	89, 7, 118, 2, 2, 89, 19, 3, 2, 2, 2, 90, 91, 7, 109, 2, 2, 91, 92, 7,
	103, 2, 2, 92, 93, 7, 123, 2, 2, 93, 94, 7, 121, 2, 2, 94, 95, 7, 113,
	2, 2, 95, 96, 7, 116, 2, 2, 96, 97, 7, 102, 2, 2, 97, 21, 3, 2, 2, 2, 98,
	99, 7, 107, 2, 2, 99, 100, 7, 102, 2, 2, 100, 101, 7, 103, 2, 2, 101, 102,
	7, 112, 2, 2, 102, 103, 7, 118, 2, 2, 103, 104, 7, 107, 2, 2, 104, 105,
	7, 104, 2, 2, 105, 106, 7, 107, 2, 2, 106, 107, 7, 103, 2, 2, 107, 108,
	7, 116, 2, 2, 108, 23, 3, 2, 2, 2, 109, 110, 7, 125, 2, 2, 110, 111, 3,
	2, 2, 2, 111, 112, 8, 12, 2, 2, 112, 25, 3, 2, 2, 2, 113, 114, 5, 4, 2,
	2, 114, 115, 3, 2, 2, 2, 115, 116, 8, 13, 3, 2, 116, 27, 3, 2, 2, 2, 117,
	118, 7, 36, 2, 2, 118, 29, 3, 2, 2, 2, 119, 120, 7, 42, 2, 2, 120, 121,
	3, 2, 2, 2, 121, 122, 8, 15, 4, 2, 122, 31, 3, 2, 2, 2, 123, 124, 7, 43,
	2, 2, 124, 125, 3, 2, 2, 2, 125, 126, 8, 16, 5, 2, 126, 33, 3, 2, 2, 2,
	127, 128, 7, 93, 2, 2, 128, 35, 3, 2, 2, 2, 129, 130, 7, 95, 2, 2, 130,
	37, 3, 2, 2, 2, 131, 132, 7, 125, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134,
	8, 19, 2, 2, 134, 135, 8, 19, 6, 2, 135, 39, 3, 2, 2, 2, 136, 137, 7, 127,
	2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 8, 20, 7, 2, 139, 41, 3, 2, 2, 2,
	140, 141, 7, 67, 2, 2, 141, 142, 7, 80, 2, 2, 142, 149, 7, 70, 2, 2, 143,
	144, 7, 81, 2, 2, 144, 149, 7, 84, 2, 2, 145, 146, 7, 80, 2, 2, 146, 147,
	7, 81, 2, 2, 147, 149, 7, 86, 2, 2, 148, 140, 3, 2, 2, 2, 148, 143, 3,
	2, 2, 2, 148, 145, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 151, 8, 21, 8,
	2, 151, 43, 3, 2, 2, 2, 152, 154, 10, 3, 2, 2, 153, 152, 3, 2, 2, 2, 154,
	155, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 45, 3,
	2, 2, 2, 157, 158, 5, 4, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 8, 23, 3,
	2, 160, 47, 3, 2, 2, 2, 8, 2, 3, 51, 65, 148, 155, 9, 7, 3, 2, 8, 2, 2,
	9, 3, 2, 9, 4, 2, 9, 12, 2, 6, 2, 2, 9, 5, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "SEARCH",
}

var lexerLiteralNames = []string{
	"", "", "", "", "':'", "'title'", "'author'", "'subject'", "'keyword'",
	"'identifier'", "'{'", "", "'\"'", "'['", "']'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "LPAREN", "RPAREN", "BOOLEAN", "COLON", "TITLE", "AUTHOR", "SUBJECT",
	"KEYWORD", "IDENTIFIER", "LBRACE", "WS1", "QUOTE", "LBRACKET", "RBRACKET",
	"RBRACE", "SEARCH_WORD", "WS2",
}

var lexerRuleNames = []string{
	"WS", "LPAREN", "RPAREN", "BOOLEAN", "COLON", "TITLE", "AUTHOR", "SUBJECT",
	"KEYWORD", "IDENTIFIER", "LBRACE", "WS1", "QUOTE", "LPAREN2", "RPAREN2",
	"LBRACKET", "RBRACKET", "LBRACE1", "RBRACE", "BOOLEAN1", "SEARCH_WORD",
	"WS2",
}

type VirgoQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewVirgoQueryLexer(input antlr.CharStream) *VirgoQueryLexer {

	l := new(VirgoQueryLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "VirgoQueryLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// VirgoQueryLexer tokens.
const (
	VirgoQueryLexerLPAREN      = 1
	VirgoQueryLexerRPAREN      = 2
	VirgoQueryLexerBOOLEAN     = 3
	VirgoQueryLexerCOLON       = 4
	VirgoQueryLexerTITLE       = 5
	VirgoQueryLexerAUTHOR      = 6
	VirgoQueryLexerSUBJECT     = 7
	VirgoQueryLexerKEYWORD     = 8
	VirgoQueryLexerIDENTIFIER  = 9
	VirgoQueryLexerLBRACE      = 10
	VirgoQueryLexerWS1         = 11
	VirgoQueryLexerQUOTE       = 12
	VirgoQueryLexerLBRACKET    = 13
	VirgoQueryLexerRBRACKET    = 14
	VirgoQueryLexerRBRACE      = 15
	VirgoQueryLexerSEARCH_WORD = 16
	VirgoQueryLexerWS2         = 17
)

// VirgoQueryLexerSEARCH is the VirgoQueryLexer mode.
const VirgoQueryLexerSEARCH = 1
