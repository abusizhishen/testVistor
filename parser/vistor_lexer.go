// Code generated from Vistor.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 87, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 6,
	11, 60, 10, 11, 13, 11, 14, 11, 61, 3, 12, 6, 12, 65, 10, 12, 13, 12, 14,
	12, 66, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 2,
	2, 17, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 3, 2, 4, 3, 2, 50, 59, 5, 2,
	11, 12, 15, 15, 34, 34, 2, 88, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 3, 33, 3, 2, 2, 2, 5, 35, 3, 2, 2, 2, 7, 37, 3, 2,
	2, 2, 9, 39, 3, 2, 2, 2, 11, 41, 3, 2, 2, 2, 13, 43, 3, 2, 2, 2, 15, 48,
	3, 2, 2, 2, 17, 54, 3, 2, 2, 2, 19, 56, 3, 2, 2, 2, 21, 59, 3, 2, 2, 2,
	23, 64, 3, 2, 2, 2, 25, 70, 3, 2, 2, 2, 27, 73, 3, 2, 2, 2, 29, 78, 3,
	2, 2, 2, 31, 83, 3, 2, 2, 2, 33, 34, 7, 45, 2, 2, 34, 4, 3, 2, 2, 2, 35,
	36, 7, 47, 2, 2, 36, 6, 3, 2, 2, 2, 37, 38, 7, 44, 2, 2, 38, 8, 3, 2, 2,
	2, 39, 40, 7, 49, 2, 2, 40, 10, 3, 2, 2, 2, 41, 42, 7, 61, 2, 2, 42, 12,
	3, 2, 2, 2, 43, 44, 7, 118, 2, 2, 44, 45, 7, 116, 2, 2, 45, 46, 7, 119,
	2, 2, 46, 47, 7, 103, 2, 2, 47, 14, 3, 2, 2, 2, 48, 49, 7, 104, 2, 2, 49,
	50, 7, 99, 2, 2, 50, 51, 7, 110, 2, 2, 51, 52, 7, 117, 2, 2, 52, 53, 7,
	103, 2, 2, 53, 16, 3, 2, 2, 2, 54, 55, 7, 125, 2, 2, 55, 18, 3, 2, 2, 2,
	56, 57, 7, 127, 2, 2, 57, 20, 3, 2, 2, 2, 58, 60, 9, 2, 2, 2, 59, 58, 3,
	2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62,
	22, 3, 2, 2, 2, 63, 65, 9, 3, 2, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2,
	2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69,
	8, 12, 2, 2, 69, 24, 3, 2, 2, 2, 70, 71, 7, 107, 2, 2, 71, 72, 7, 104,
	2, 2, 72, 26, 3, 2, 2, 2, 73, 74, 7, 103, 2, 2, 74, 75, 7, 110, 2, 2, 75,
	76, 7, 107, 2, 2, 76, 77, 7, 104, 2, 2, 77, 28, 3, 2, 2, 2, 78, 79, 7,
	103, 2, 2, 79, 80, 7, 110, 2, 2, 80, 81, 7, 117, 2, 2, 81, 82, 7, 103,
	2, 2, 82, 30, 3, 2, 2, 2, 83, 84, 7, 103, 2, 2, 84, 85, 7, 112, 2, 2, 85,
	86, 7, 102, 2, 2, 86, 32, 3, 2, 2, 2, 5, 2, 61, 66, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "';'", "'true'", "'false'", "'{'", "'}'",
	"", "", "'if'", "'elif'", "'else'", "'end'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "NUM", "WS", "IF", "ELIF", "ELSE",
	"END",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"NUM", "WS", "IF", "ELIF", "ELSE", "END",
}

type VistorLexer struct {
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

func NewVistorLexer(input antlr.CharStream) *VistorLexer {

	l := new(VistorLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Vistor.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// VistorLexer tokens.
const (
	VistorLexerT__0 = 1
	VistorLexerT__1 = 2
	VistorLexerT__2 = 3
	VistorLexerT__3 = 4
	VistorLexerT__4 = 5
	VistorLexerT__5 = 6
	VistorLexerT__6 = 7
	VistorLexerT__7 = 8
	VistorLexerT__8 = 9
	VistorLexerNUM  = 10
	VistorLexerWS   = 11
	VistorLexerIF   = 12
	VistorLexerELIF = 13
	VistorLexerELSE = 14
	VistorLexerEND  = 15
)