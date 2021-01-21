// Code generated from ValidationExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ValidationExpr

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 35, 240,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 62, 10, 2, 3, 3, 3, 3, 5, 3, 66, 10, 3, 3, 4,
	3, 4, 3, 4, 7, 4, 71, 10, 4, 12, 4, 14, 4, 74, 11, 4, 3, 5, 3, 5, 3, 5,
	7, 5, 79, 10, 5, 12, 5, 14, 5, 82, 11, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 99, 10,
	7, 12, 7, 14, 7, 102, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	7, 9, 111, 10, 9, 12, 9, 14, 9, 114, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 7, 10, 121, 10, 10, 12, 10, 14, 10, 124, 11, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 133, 10, 11, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 7, 15, 149, 10, 15, 12, 15, 14, 15, 152, 11, 15, 3, 15, 3, 15, 5, 15,
	156, 10, 15, 3, 16, 3, 16, 5, 16, 160, 10, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 5, 17, 168, 10, 17, 3, 18, 3, 18, 3, 19, 3, 19, 5, 19,
	174, 10, 19, 3, 20, 6, 20, 177, 10, 20, 13, 20, 14, 20, 178, 3, 20, 3,
	20, 3, 21, 6, 21, 184, 10, 21, 13, 21, 14, 21, 185, 3, 21, 3, 21, 6, 21,
	190, 10, 21, 13, 21, 14, 21, 191, 3, 21, 3, 21, 6, 21, 196, 10, 21, 13,
	21, 14, 21, 197, 3, 22, 3, 22, 3, 22, 6, 22, 203, 10, 22, 13, 22, 14, 22,
	204, 3, 22, 6, 22, 208, 10, 22, 13, 22, 14, 22, 209, 3, 22, 6, 22, 213,
	10, 22, 13, 22, 14, 22, 214, 3, 22, 5, 22, 218, 10, 22, 3, 23, 6, 23, 221,
	10, 23, 13, 23, 14, 23, 222, 3, 23, 6, 23, 226, 10, 23, 13, 23, 14, 23,
	227, 3, 23, 3, 23, 6, 23, 232, 10, 23, 13, 23, 14, 23, 233, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 25, 2, 2, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 5, 3, 2, 22, 23,
	3, 2, 24, 25, 3, 2, 29, 30, 2, 253, 2, 61, 3, 2, 2, 2, 4, 65, 3, 2, 2,
	2, 6, 67, 3, 2, 2, 2, 8, 75, 3, 2, 2, 2, 10, 83, 3, 2, 2, 2, 12, 85, 3,
	2, 2, 2, 14, 103, 3, 2, 2, 2, 16, 105, 3, 2, 2, 2, 18, 115, 3, 2, 2, 2,
	20, 132, 3, 2, 2, 2, 22, 134, 3, 2, 2, 2, 24, 136, 3, 2, 2, 2, 26, 140,
	3, 2, 2, 2, 28, 155, 3, 2, 2, 2, 30, 159, 3, 2, 2, 2, 32, 167, 3, 2, 2,
	2, 34, 169, 3, 2, 2, 2, 36, 173, 3, 2, 2, 2, 38, 176, 3, 2, 2, 2, 40, 183,
	3, 2, 2, 2, 42, 217, 3, 2, 2, 2, 44, 231, 3, 2, 2, 2, 46, 235, 3, 2, 2,
	2, 48, 237, 3, 2, 2, 2, 50, 51, 7, 3, 2, 2, 51, 52, 5, 4, 3, 2, 52, 53,
	7, 4, 2, 2, 53, 54, 5, 4, 3, 2, 54, 55, 7, 5, 2, 2, 55, 56, 5, 4, 3, 2,
	56, 57, 7, 2, 2, 3, 57, 62, 3, 2, 2, 2, 58, 59, 5, 4, 3, 2, 59, 60, 7,
	2, 2, 3, 60, 62, 3, 2, 2, 2, 61, 50, 3, 2, 2, 2, 61, 58, 3, 2, 2, 2, 62,
	3, 3, 2, 2, 2, 63, 66, 5, 6, 4, 2, 64, 66, 5, 8, 5, 2, 65, 63, 3, 2, 2,
	2, 65, 64, 3, 2, 2, 2, 66, 5, 3, 2, 2, 2, 67, 72, 5, 8, 5, 2, 68, 69, 7,
	6, 2, 2, 69, 71, 5, 8, 5, 2, 70, 68, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72,
	70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 7, 3, 2, 2, 2, 74, 72, 3, 2, 2,
	2, 75, 80, 5, 10, 6, 2, 76, 77, 7, 7, 2, 2, 77, 79, 5, 10, 6, 2, 78, 76,
	3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2,
	81, 9, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 84, 5, 12, 7, 2, 84, 11, 3,
	2, 2, 2, 85, 100, 5, 14, 8, 2, 86, 87, 7, 8, 2, 2, 87, 99, 5, 14, 8, 2,
	88, 89, 7, 9, 2, 2, 89, 99, 5, 14, 8, 2, 90, 91, 7, 10, 2, 2, 91, 99, 5,
	14, 8, 2, 92, 93, 7, 11, 2, 2, 93, 99, 5, 14, 8, 2, 94, 95, 7, 12, 2, 2,
	95, 99, 5, 14, 8, 2, 96, 97, 7, 13, 2, 2, 97, 99, 5, 14, 8, 2, 98, 86,
	3, 2, 2, 2, 98, 88, 3, 2, 2, 2, 98, 90, 3, 2, 2, 2, 98, 92, 3, 2, 2, 2,
	98, 94, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3,
	2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 13, 3, 2, 2, 2, 102, 100, 3, 2, 2,
	2, 103, 104, 5, 16, 9, 2, 104, 15, 3, 2, 2, 2, 105, 112, 5, 18, 10, 2,
	106, 107, 7, 14, 2, 2, 107, 111, 5, 18, 10, 2, 108, 109, 7, 15, 2, 2, 109,
	111, 5, 18, 10, 2, 110, 106, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 114,
	3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 17, 3, 2,
	2, 2, 114, 112, 3, 2, 2, 2, 115, 122, 5, 20, 11, 2, 116, 117, 7, 16, 2,
	2, 117, 121, 5, 20, 11, 2, 118, 119, 7, 17, 2, 2, 119, 121, 5, 20, 11,
	2, 120, 116, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122,
	120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 19, 3, 2, 2, 2, 124, 122, 3,
	2, 2, 2, 125, 133, 5, 24, 13, 2, 126, 133, 5, 26, 14, 2, 127, 133, 5, 48,
	25, 2, 128, 133, 5, 32, 17, 2, 129, 130, 5, 22, 12, 2, 130, 131, 5, 20,
	11, 2, 131, 133, 3, 2, 2, 2, 132, 125, 3, 2, 2, 2, 132, 126, 3, 2, 2, 2,
	132, 127, 3, 2, 2, 2, 132, 128, 3, 2, 2, 2, 132, 129, 3, 2, 2, 2, 133,
	21, 3, 2, 2, 2, 134, 135, 7, 18, 2, 2, 135, 23, 3, 2, 2, 2, 136, 137, 7,
	19, 2, 2, 137, 138, 5, 4, 3, 2, 138, 139, 7, 20, 2, 2, 139, 25, 3, 2, 2,
	2, 140, 141, 7, 32, 2, 2, 141, 142, 5, 28, 15, 2, 142, 27, 3, 2, 2, 2,
	143, 156, 7, 34, 2, 2, 144, 145, 7, 19, 2, 2, 145, 150, 5, 30, 16, 2, 146,
	147, 7, 21, 2, 2, 147, 149, 5, 30, 16, 2, 148, 146, 3, 2, 2, 2, 149, 152,
	3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 153, 3, 2,
	2, 2, 152, 150, 3, 2, 2, 2, 153, 154, 7, 20, 2, 2, 154, 156, 3, 2, 2, 2,
	155, 143, 3, 2, 2, 2, 155, 144, 3, 2, 2, 2, 156, 29, 3, 2, 2, 2, 157, 160,
	5, 48, 25, 2, 158, 160, 5, 32, 17, 2, 159, 157, 3, 2, 2, 2, 159, 158, 3,
	2, 2, 2, 160, 31, 3, 2, 2, 2, 161, 168, 3, 2, 2, 2, 162, 168, 5, 34, 18,
	2, 163, 168, 5, 36, 19, 2, 164, 168, 5, 46, 24, 2, 165, 168, 5, 42, 22,
	2, 166, 168, 5, 44, 23, 2, 167, 161, 3, 2, 2, 2, 167, 162, 3, 2, 2, 2,
	167, 163, 3, 2, 2, 2, 167, 164, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167,
	166, 3, 2, 2, 2, 168, 33, 3, 2, 2, 2, 169, 170, 9, 2, 2, 2, 170, 35, 3,
	2, 2, 2, 171, 174, 7, 31, 2, 2, 172, 174, 5, 38, 20, 2, 173, 171, 3, 2,
	2, 2, 173, 172, 3, 2, 2, 2, 174, 37, 3, 2, 2, 2, 175, 177, 7, 31, 2, 2,
	176, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178,
	179, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 181, 9, 3, 2, 2, 181, 39, 3,
	2, 2, 2, 182, 184, 7, 31, 2, 2, 183, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2,
	2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187,
	189, 7, 26, 2, 2, 188, 190, 7, 31, 2, 2, 189, 188, 3, 2, 2, 2, 190, 191,
	3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 195, 3, 2,
	2, 2, 193, 194, 7, 26, 2, 2, 194, 196, 7, 31, 2, 2, 195, 193, 3, 2, 2,
	2, 196, 197, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198,
	41, 3, 2, 2, 2, 199, 218, 7, 33, 2, 2, 200, 202, 7, 27, 2, 2, 201, 203,
	7, 32, 2, 2, 202, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 202, 3, 2,
	2, 2, 204, 205, 3, 2, 2, 2, 205, 212, 3, 2, 2, 2, 206, 208, 7, 21, 2, 2,
	207, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209,
	210, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 213, 7, 32, 2, 2, 212, 207,
	3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2,
	2, 2, 215, 216, 3, 2, 2, 2, 216, 218, 7, 28, 2, 2, 217, 199, 3, 2, 2, 2,
	217, 200, 3, 2, 2, 2, 218, 43, 3, 2, 2, 2, 219, 221, 5, 36, 19, 2, 220,
	219, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 222, 223,
	3, 2, 2, 2, 223, 225, 3, 2, 2, 2, 224, 226, 7, 21, 2, 2, 225, 224, 3, 2,
	2, 2, 226, 227, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2,
	228, 229, 3, 2, 2, 2, 229, 230, 5, 36, 19, 2, 230, 232, 3, 2, 2, 2, 231,
	220, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 233, 234,
	3, 2, 2, 2, 234, 45, 3, 2, 2, 2, 235, 236, 7, 32, 2, 2, 236, 47, 3, 2,
	2, 2, 237, 238, 9, 4, 2, 2, 238, 49, 3, 2, 2, 2, 29, 61, 65, 72, 80, 98,
	100, 110, 112, 120, 122, 132, 150, 155, 159, 167, 173, 178, 185, 191, 197,
	204, 209, 214, 217, 222, 227, 233,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'if'", "':'", "'else'", "'or'", "'and'", "'='", "'!='", "'<'", "'>'",
	"'<='", "'>='", "'+'", "'-'", "'*'", "'/'", "'not'", "'('", "')'", "','",
	"'true'", "'false'", "'ms'", "'s'", "'.'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "LOCAL_VARIABLE", "CONTEXT_VARIABLE",
	"INTEGER", "STR", "NONE", "NIL", "WS",
}

var ruleNames = []string{
	"selectionExpr", "validationExpr", "validationOrExpr", "validationAndExpr",
	"valueLogical", "relationalExpression", "numericExpression", "additiveExpression",
	"multiplicativeExpression", "unaryExpression", "unaryOperator", "brackettedExpression",
	"functionCall", "functionArgumentList", "functionAttribute", "literal",
	"boolenLiteral", "numericLiteral", "timeDurationLiteral", "versionNumberLiteral",
	"arrayLiteral", "rangeLiteral", "stringLiteral", "variable",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ValidationExprParser struct {
	*antlr.BaseParser
}

func NewValidationExprParser(input antlr.TokenStream) *ValidationExprParser {
	this := new(ValidationExprParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ValidationExpr.g4"

	return this
}

// ValidationExprParser tokens.
const (
	ValidationExprParserEOF              = antlr.TokenEOF
	ValidationExprParserT__0             = 1
	ValidationExprParserT__1             = 2
	ValidationExprParserT__2             = 3
	ValidationExprParserT__3             = 4
	ValidationExprParserT__4             = 5
	ValidationExprParserT__5             = 6
	ValidationExprParserT__6             = 7
	ValidationExprParserT__7             = 8
	ValidationExprParserT__8             = 9
	ValidationExprParserT__9             = 10
	ValidationExprParserT__10            = 11
	ValidationExprParserT__11            = 12
	ValidationExprParserT__12            = 13
	ValidationExprParserT__13            = 14
	ValidationExprParserT__14            = 15
	ValidationExprParserT__15            = 16
	ValidationExprParserT__16            = 17
	ValidationExprParserT__17            = 18
	ValidationExprParserT__18            = 19
	ValidationExprParserT__19            = 20
	ValidationExprParserT__20            = 21
	ValidationExprParserT__21            = 22
	ValidationExprParserT__22            = 23
	ValidationExprParserT__23            = 24
	ValidationExprParserT__24            = 25
	ValidationExprParserT__25            = 26
	ValidationExprParserLOCAL_VARIABLE   = 27
	ValidationExprParserCONTEXT_VARIABLE = 28
	ValidationExprParserINTEGER          = 29
	ValidationExprParserSTR              = 30
	ValidationExprParserNONE             = 31
	ValidationExprParserNIL              = 32
	ValidationExprParserWS               = 33
)

// ValidationExprParser rules.
const (
	ValidationExprParserRULE_selectionExpr            = 0
	ValidationExprParserRULE_validationExpr           = 1
	ValidationExprParserRULE_validationOrExpr         = 2
	ValidationExprParserRULE_validationAndExpr        = 3
	ValidationExprParserRULE_valueLogical             = 4
	ValidationExprParserRULE_relationalExpression     = 5
	ValidationExprParserRULE_numericExpression        = 6
	ValidationExprParserRULE_additiveExpression       = 7
	ValidationExprParserRULE_multiplicativeExpression = 8
	ValidationExprParserRULE_unaryExpression          = 9
	ValidationExprParserRULE_unaryOperator            = 10
	ValidationExprParserRULE_brackettedExpression     = 11
	ValidationExprParserRULE_functionCall             = 12
	ValidationExprParserRULE_functionArgumentList     = 13
	ValidationExprParserRULE_functionAttribute        = 14
	ValidationExprParserRULE_literal                  = 15
	ValidationExprParserRULE_boolenLiteral            = 16
	ValidationExprParserRULE_numericLiteral           = 17
	ValidationExprParserRULE_timeDurationLiteral      = 18
	ValidationExprParserRULE_versionNumberLiteral     = 19
	ValidationExprParserRULE_arrayLiteral             = 20
	ValidationExprParserRULE_rangeLiteral             = 21
	ValidationExprParserRULE_stringLiteral            = 22
	ValidationExprParserRULE_variable                 = 23
)

// ISelectionExprContext is an interface to support dynamic dispatch.
type ISelectionExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionExprContext differentiates from other interfaces.
	IsSelectionExprContext()
}

type SelectionExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionExprContext() *SelectionExprContext {
	var p = new(SelectionExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_selectionExpr
	return p
}

func (*SelectionExprContext) IsSelectionExprContext() {}

func NewSelectionExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionExprContext {
	var p = new(SelectionExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_selectionExpr

	return p
}

func (s *SelectionExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionExprContext) AllValidationExpr() []IValidationExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValidationExprContext)(nil)).Elem())
	var tst = make([]IValidationExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValidationExprContext)
		}
	}

	return tst
}

func (s *SelectionExprContext) ValidationExpr(i int) IValidationExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValidationExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValidationExprContext)
}

func (s *SelectionExprContext) EOF() antlr.TerminalNode {
	return s.GetToken(ValidationExprParserEOF, 0)
}

func (s *SelectionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterSelectionExpr(s)
	}
}

func (s *SelectionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitSelectionExpr(s)
	}
}

func (p *ValidationExprParser) SelectionExpr() (localctx ISelectionExprContext) {
	localctx = NewSelectionExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ValidationExprParserRULE_selectionExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ValidationExprParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(ValidationExprParserT__0)
		}
		{
			p.SetState(49)
			p.ValidationExpr()
		}
		{
			p.SetState(50)
			p.Match(ValidationExprParserT__1)
		}
		{
			p.SetState(51)
			p.ValidationExpr()
		}
		{
			p.SetState(52)
			p.Match(ValidationExprParserT__2)
		}
		{
			p.SetState(53)
			p.ValidationExpr()
		}
		{
			p.SetState(54)
			p.Match(ValidationExprParserEOF)
		}

	case ValidationExprParserEOF, ValidationExprParserT__3, ValidationExprParserT__4, ValidationExprParserT__5, ValidationExprParserT__6, ValidationExprParserT__7, ValidationExprParserT__8, ValidationExprParserT__9, ValidationExprParserT__10, ValidationExprParserT__11, ValidationExprParserT__12, ValidationExprParserT__13, ValidationExprParserT__14, ValidationExprParserT__15, ValidationExprParserT__16, ValidationExprParserT__19, ValidationExprParserT__20, ValidationExprParserT__24, ValidationExprParserLOCAL_VARIABLE, ValidationExprParserCONTEXT_VARIABLE, ValidationExprParserINTEGER, ValidationExprParserSTR, ValidationExprParserNONE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.ValidationExpr()
		}
		{
			p.SetState(57)
			p.Match(ValidationExprParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValidationExprContext is an interface to support dynamic dispatch.
type IValidationExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidationExprContext differentiates from other interfaces.
	IsValidationExprContext()
}

type ValidationExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidationExprContext() *ValidationExprContext {
	var p = new(ValidationExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_validationExpr
	return p
}

func (*ValidationExprContext) IsValidationExprContext() {}

func NewValidationExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidationExprContext {
	var p = new(ValidationExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_validationExpr

	return p
}

func (s *ValidationExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidationExprContext) ValidationOrExpr() IValidationOrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValidationOrExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValidationOrExprContext)
}

func (s *ValidationExprContext) ValidationAndExpr() IValidationAndExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValidationAndExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValidationAndExprContext)
}

func (s *ValidationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidationExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidationExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterValidationExpr(s)
	}
}

func (s *ValidationExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitValidationExpr(s)
	}
}

func (p *ValidationExprParser) ValidationExpr() (localctx IValidationExprContext) {
	localctx = NewValidationExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ValidationExprParserRULE_validationExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(61)
			p.ValidationOrExpr()
		}

	case 2:
		{
			p.SetState(62)
			p.ValidationAndExpr()
		}

	}

	return localctx
}

// IValidationOrExprContext is an interface to support dynamic dispatch.
type IValidationOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidationOrExprContext differentiates from other interfaces.
	IsValidationOrExprContext()
}

type ValidationOrExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidationOrExprContext() *ValidationOrExprContext {
	var p = new(ValidationOrExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_validationOrExpr
	return p
}

func (*ValidationOrExprContext) IsValidationOrExprContext() {}

func NewValidationOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidationOrExprContext {
	var p = new(ValidationOrExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_validationOrExpr

	return p
}

func (s *ValidationOrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidationOrExprContext) AllValidationAndExpr() []IValidationAndExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValidationAndExprContext)(nil)).Elem())
	var tst = make([]IValidationAndExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValidationAndExprContext)
		}
	}

	return tst
}

func (s *ValidationOrExprContext) ValidationAndExpr(i int) IValidationAndExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValidationAndExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValidationAndExprContext)
}

func (s *ValidationOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidationOrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidationOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterValidationOrExpr(s)
	}
}

func (s *ValidationOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitValidationOrExpr(s)
	}
}

func (p *ValidationExprParser) ValidationOrExpr() (localctx IValidationOrExprContext) {
	localctx = NewValidationOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ValidationExprParserRULE_validationOrExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.ValidationAndExpr()
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ValidationExprParserT__3 {
		{
			p.SetState(66)
			p.Match(ValidationExprParserT__3)
		}
		{
			p.SetState(67)
			p.ValidationAndExpr()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IValidationAndExprContext is an interface to support dynamic dispatch.
type IValidationAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidationAndExprContext differentiates from other interfaces.
	IsValidationAndExprContext()
}

type ValidationAndExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidationAndExprContext() *ValidationAndExprContext {
	var p = new(ValidationAndExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_validationAndExpr
	return p
}

func (*ValidationAndExprContext) IsValidationAndExprContext() {}

func NewValidationAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidationAndExprContext {
	var p = new(ValidationAndExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_validationAndExpr

	return p
}

func (s *ValidationAndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidationAndExprContext) AllValueLogical() []IValueLogicalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueLogicalContext)(nil)).Elem())
	var tst = make([]IValueLogicalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueLogicalContext)
		}
	}

	return tst
}

func (s *ValidationAndExprContext) ValueLogical(i int) IValueLogicalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueLogicalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueLogicalContext)
}

func (s *ValidationAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidationAndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidationAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterValidationAndExpr(s)
	}
}

func (s *ValidationAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitValidationAndExpr(s)
	}
}

func (p *ValidationExprParser) ValidationAndExpr() (localctx IValidationAndExprContext) {
	localctx = NewValidationAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ValidationExprParserRULE_validationAndExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.ValueLogical()
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ValidationExprParserT__4 {
		{
			p.SetState(74)
			p.Match(ValidationExprParserT__4)
		}
		{
			p.SetState(75)
			p.ValueLogical()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IValueLogicalContext is an interface to support dynamic dispatch.
type IValueLogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueLogicalContext differentiates from other interfaces.
	IsValueLogicalContext()
}

type ValueLogicalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueLogicalContext() *ValueLogicalContext {
	var p = new(ValueLogicalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_valueLogical
	return p
}

func (*ValueLogicalContext) IsValueLogicalContext() {}

func NewValueLogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueLogicalContext {
	var p = new(ValueLogicalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_valueLogical

	return p
}

func (s *ValueLogicalContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueLogicalContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *ValueLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueLogicalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterValueLogical(s)
	}
}

func (s *ValueLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitValueLogical(s)
	}
}

func (p *ValidationExprParser) ValueLogical() (localctx IValueLogicalContext) {
	localctx = NewValueLogicalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ValidationExprParserRULE_valueLogical)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.RelationalExpression()
	}

	return localctx
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_relationalExpression
	return p
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) AllNumericExpression() []INumericExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumericExpressionContext)(nil)).Elem())
	var tst = make([]INumericExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumericExpressionContext)
		}
	}

	return tst
}

func (s *RelationalExpressionContext) NumericExpression(i int) INumericExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumericExpressionContext)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (p *ValidationExprParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ValidationExprParserRULE_relationalExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.NumericExpression()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ValidationExprParserT__5)|(1<<ValidationExprParserT__6)|(1<<ValidationExprParserT__7)|(1<<ValidationExprParserT__8)|(1<<ValidationExprParserT__9)|(1<<ValidationExprParserT__10))) != 0 {
		p.SetState(96)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ValidationExprParserT__5:
			{
				p.SetState(84)
				p.Match(ValidationExprParserT__5)
			}
			{
				p.SetState(85)
				p.NumericExpression()
			}

		case ValidationExprParserT__6:
			{
				p.SetState(86)
				p.Match(ValidationExprParserT__6)
			}
			{
				p.SetState(87)
				p.NumericExpression()
			}

		case ValidationExprParserT__7:
			{
				p.SetState(88)
				p.Match(ValidationExprParserT__7)
			}
			{
				p.SetState(89)
				p.NumericExpression()
			}

		case ValidationExprParserT__8:
			{
				p.SetState(90)
				p.Match(ValidationExprParserT__8)
			}
			{
				p.SetState(91)
				p.NumericExpression()
			}

		case ValidationExprParserT__9:
			{
				p.SetState(92)
				p.Match(ValidationExprParserT__9)
			}
			{
				p.SetState(93)
				p.NumericExpression()
			}

		case ValidationExprParserT__10:
			{
				p.SetState(94)
				p.Match(ValidationExprParserT__10)
			}
			{
				p.SetState(95)
				p.NumericExpression()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INumericExpressionContext is an interface to support dynamic dispatch.
type INumericExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericExpressionContext differentiates from other interfaces.
	IsNumericExpressionContext()
}

type NumericExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericExpressionContext() *NumericExpressionContext {
	var p = new(NumericExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_numericExpression
	return p
}

func (*NumericExpressionContext) IsNumericExpressionContext() {}

func NewNumericExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericExpressionContext {
	var p = new(NumericExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_numericExpression

	return p
}

func (s *NumericExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *NumericExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterNumericExpression(s)
	}
}

func (s *NumericExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitNumericExpression(s)
	}
}

func (p *ValidationExprParser) NumericExpression() (localctx INumericExpressionContext) {
	localctx = NewNumericExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ValidationExprParserRULE_numericExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.AdditiveExpression()
	}

	return localctx
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_additiveExpression
	return p
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplicativeExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplicativeExpressionContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (p *ValidationExprParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ValidationExprParserRULE_additiveExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.MultiplicativeExpression()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ValidationExprParserT__11 || _la == ValidationExprParserT__12 {
		p.SetState(108)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ValidationExprParserT__11:
			{
				p.SetState(104)
				p.Match(ValidationExprParserT__11)
			}
			{
				p.SetState(105)
				p.MultiplicativeExpression()
			}

		case ValidationExprParserT__12:
			{
				p.SetState(106)
				p.Match(ValidationExprParserT__12)
			}
			{
				p.SetState(107)
				p.MultiplicativeExpression()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_multiplicativeExpression
	return p
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) AllUnaryExpression() []IUnaryExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem())
	var tst = make([]IUnaryExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnaryExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) UnaryExpression(i int) IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (p *ValidationExprParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ValidationExprParserRULE_multiplicativeExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.UnaryExpression()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ValidationExprParserT__13 || _la == ValidationExprParserT__14 {
		p.SetState(118)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ValidationExprParserT__13:
			{
				p.SetState(114)
				p.Match(ValidationExprParserT__13)
			}
			{
				p.SetState(115)
				p.UnaryExpression()
			}

		case ValidationExprParserT__14:
			{
				p.SetState(116)
				p.Match(ValidationExprParserT__14)
			}
			{
				p.SetState(117)
				p.UnaryExpression()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_unaryExpression
	return p
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) BrackettedExpression() IBrackettedExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBrackettedExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBrackettedExpressionContext)
}

func (s *UnaryExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *UnaryExpressionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *UnaryExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *UnaryExpressionContext) UnaryOperator() IUnaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (p *ValidationExprParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ValidationExprParserRULE_unaryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.BrackettedExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.FunctionCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Variable()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
			p.Literal()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)
			p.UnaryOperator()
		}
		{
			p.SetState(128)
			p.UnaryExpression()
		}

	}

	return localctx
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_unaryOperator
	return p
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (p *ValidationExprParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ValidationExprParserRULE_unaryOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(ValidationExprParserT__15)
	}

	return localctx
}

// IBrackettedExpressionContext is an interface to support dynamic dispatch.
type IBrackettedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBrackettedExpressionContext differentiates from other interfaces.
	IsBrackettedExpressionContext()
}

type BrackettedExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBrackettedExpressionContext() *BrackettedExpressionContext {
	var p = new(BrackettedExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_brackettedExpression
	return p
}

func (*BrackettedExpressionContext) IsBrackettedExpressionContext() {}

func NewBrackettedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BrackettedExpressionContext {
	var p = new(BrackettedExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_brackettedExpression

	return p
}

func (s *BrackettedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BrackettedExpressionContext) ValidationExpr() IValidationExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValidationExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValidationExprContext)
}

func (s *BrackettedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BrackettedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BrackettedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterBrackettedExpression(s)
	}
}

func (s *BrackettedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitBrackettedExpression(s)
	}
}

func (p *ValidationExprParser) BrackettedExpression() (localctx IBrackettedExpressionContext) {
	localctx = NewBrackettedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ValidationExprParserRULE_brackettedExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(ValidationExprParserT__16)
	}
	{
		p.SetState(135)
		p.ValidationExpr()
	}
	{
		p.SetState(136)
		p.Match(ValidationExprParserT__17)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) STR() antlr.TerminalNode {
	return s.GetToken(ValidationExprParserSTR, 0)
}

func (s *FunctionCallContext) FunctionArgumentList() IFunctionArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgumentListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *ValidationExprParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ValidationExprParserRULE_functionCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(ValidationExprParserSTR)
	}
	{
		p.SetState(139)
		p.FunctionArgumentList()
	}

	return localctx
}

// IFunctionArgumentListContext is an interface to support dynamic dispatch.
type IFunctionArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgumentListContext differentiates from other interfaces.
	IsFunctionArgumentListContext()
}

type FunctionArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgumentListContext() *FunctionArgumentListContext {
	var p = new(FunctionArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_functionArgumentList
	return p
}

func (*FunctionArgumentListContext) IsFunctionArgumentListContext() {}

func NewFunctionArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgumentListContext {
	var p = new(FunctionArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_functionArgumentList

	return p
}

func (s *FunctionArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgumentListContext) NIL() antlr.TerminalNode {
	return s.GetToken(ValidationExprParserNIL, 0)
}

func (s *FunctionArgumentListContext) AllFunctionAttribute() []IFunctionAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionAttributeContext)(nil)).Elem())
	var tst = make([]IFunctionAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionAttributeContext)
		}
	}

	return tst
}

func (s *FunctionArgumentListContext) FunctionAttribute(i int) IFunctionAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionAttributeContext)
}

func (s *FunctionArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterFunctionArgumentList(s)
	}
}

func (s *FunctionArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitFunctionArgumentList(s)
	}
}

func (p *ValidationExprParser) FunctionArgumentList() (localctx IFunctionArgumentListContext) {
	localctx = NewFunctionArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ValidationExprParserRULE_functionArgumentList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ValidationExprParserNIL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(ValidationExprParserNIL)
		}

	case ValidationExprParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Match(ValidationExprParserT__16)
		}
		{
			p.SetState(143)
			p.FunctionAttribute()
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ValidationExprParserT__18 {
			{
				p.SetState(144)
				p.Match(ValidationExprParserT__18)
			}
			{
				p.SetState(145)
				p.FunctionAttribute()
			}

			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(151)
			p.Match(ValidationExprParserT__17)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionAttributeContext is an interface to support dynamic dispatch.
type IFunctionAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionAttributeContext differentiates from other interfaces.
	IsFunctionAttributeContext()
}

type FunctionAttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionAttributeContext() *FunctionAttributeContext {
	var p = new(FunctionAttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_functionAttribute
	return p
}

func (*FunctionAttributeContext) IsFunctionAttributeContext() {}

func NewFunctionAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionAttributeContext {
	var p = new(FunctionAttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_functionAttribute

	return p
}

func (s *FunctionAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionAttributeContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FunctionAttributeContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *FunctionAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterFunctionAttribute(s)
	}
}

func (s *FunctionAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitFunctionAttribute(s)
	}
}

func (p *ValidationExprParser) FunctionAttribute() (localctx IFunctionAttributeContext) {
	localctx = NewFunctionAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ValidationExprParserRULE_functionAttribute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ValidationExprParserLOCAL_VARIABLE, ValidationExprParserCONTEXT_VARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Variable()
		}

	case ValidationExprParserT__17, ValidationExprParserT__18, ValidationExprParserT__19, ValidationExprParserT__20, ValidationExprParserT__24, ValidationExprParserINTEGER, ValidationExprParserSTR, ValidationExprParserNONE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) BoolenLiteral() IBoolenLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolenLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolenLiteralContext)
}

func (s *LiteralContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *LiteralContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LiteralContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *LiteralContext) RangeLiteral() IRangeLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *ValidationExprParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ValidationExprParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.BoolenLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
			p.NumericLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(162)
			p.StringLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(163)
			p.ArrayLiteral()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(164)
			p.RangeLiteral()
		}

	}

	return localctx
}

// IBoolenLiteralContext is an interface to support dynamic dispatch.
type IBoolenLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolenLiteralContext differentiates from other interfaces.
	IsBoolenLiteralContext()
}

type BoolenLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolenLiteralContext() *BoolenLiteralContext {
	var p = new(BoolenLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_boolenLiteral
	return p
}

func (*BoolenLiteralContext) IsBoolenLiteralContext() {}

func NewBoolenLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolenLiteralContext {
	var p = new(BoolenLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_boolenLiteral

	return p
}

func (s *BoolenLiteralContext) GetParser() antlr.Parser { return s.parser }
func (s *BoolenLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolenLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolenLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterBoolenLiteral(s)
	}
}

func (s *BoolenLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitBoolenLiteral(s)
	}
}

func (p *ValidationExprParser) BoolenLiteral() (localctx IBoolenLiteralContext) {
	localctx = NewBoolenLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ValidationExprParserRULE_boolenLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ValidationExprParserT__19 || _la == ValidationExprParserT__20) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_numericLiteral
	return p
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ValidationExprParserINTEGER, 0)
}

func (s *NumericLiteralContext) TimeDurationLiteral() ITimeDurationLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeDurationLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeDurationLiteralContext)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (p *ValidationExprParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ValidationExprParserRULE_numericLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Match(ValidationExprParserINTEGER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.TimeDurationLiteral()
		}

	}

	return localctx
}

// ITimeDurationLiteralContext is an interface to support dynamic dispatch.
type ITimeDurationLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeDurationLiteralContext differentiates from other interfaces.
	IsTimeDurationLiteralContext()
}

type TimeDurationLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeDurationLiteralContext() *TimeDurationLiteralContext {
	var p = new(TimeDurationLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_timeDurationLiteral
	return p
}

func (*TimeDurationLiteralContext) IsTimeDurationLiteralContext() {}

func NewTimeDurationLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeDurationLiteralContext {
	var p = new(TimeDurationLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_timeDurationLiteral

	return p
}

func (s *TimeDurationLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeDurationLiteralContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(ValidationExprParserINTEGER)
}

func (s *TimeDurationLiteralContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(ValidationExprParserINTEGER, i)
}

func (s *TimeDurationLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeDurationLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeDurationLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterTimeDurationLiteral(s)
	}
}

func (s *TimeDurationLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitTimeDurationLiteral(s)
	}
}

func (p *ValidationExprParser) TimeDurationLiteral() (localctx ITimeDurationLiteralContext) {
	localctx = NewTimeDurationLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ValidationExprParserRULE_timeDurationLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ValidationExprParserINTEGER {
		{
			p.SetState(173)
			p.Match(ValidationExprParserINTEGER)
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(178)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ValidationExprParserT__21 || _la == ValidationExprParserT__22) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IVersionNumberLiteralContext is an interface to support dynamic dispatch.
type IVersionNumberLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVersionNumberLiteralContext differentiates from other interfaces.
	IsVersionNumberLiteralContext()
}

type VersionNumberLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionNumberLiteralContext() *VersionNumberLiteralContext {
	var p = new(VersionNumberLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_versionNumberLiteral
	return p
}

func (*VersionNumberLiteralContext) IsVersionNumberLiteralContext() {}

func NewVersionNumberLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionNumberLiteralContext {
	var p = new(VersionNumberLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_versionNumberLiteral

	return p
}

func (s *VersionNumberLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionNumberLiteralContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(ValidationExprParserINTEGER)
}

func (s *VersionNumberLiteralContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(ValidationExprParserINTEGER, i)
}

func (s *VersionNumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionNumberLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionNumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterVersionNumberLiteral(s)
	}
}

func (s *VersionNumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitVersionNumberLiteral(s)
	}
}

func (p *ValidationExprParser) VersionNumberLiteral() (localctx IVersionNumberLiteralContext) {
	localctx = NewVersionNumberLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ValidationExprParserRULE_versionNumberLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ValidationExprParserINTEGER {
		{
			p.SetState(180)
			p.Match(ValidationExprParserINTEGER)
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(185)
		p.Match(ValidationExprParserT__23)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ValidationExprParserINTEGER {
		{
			p.SetState(186)
			p.Match(ValidationExprParserINTEGER)
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ValidationExprParserT__23 {
		{
			p.SetState(191)
			p.Match(ValidationExprParserT__23)
		}
		{
			p.SetState(192)
			p.Match(ValidationExprParserINTEGER)
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) NONE() antlr.TerminalNode {
	return s.GetToken(ValidationExprParserNONE, 0)
}

func (s *ArrayLiteralContext) AllSTR() []antlr.TerminalNode {
	return s.GetTokens(ValidationExprParserSTR)
}

func (s *ArrayLiteralContext) STR(i int) antlr.TerminalNode {
	return s.GetToken(ValidationExprParserSTR, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *ValidationExprParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ValidationExprParserRULE_arrayLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(215)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ValidationExprParserNONE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.Match(ValidationExprParserNONE)
		}

	case ValidationExprParserT__24:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Match(ValidationExprParserT__24)
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ValidationExprParserSTR {
			{
				p.SetState(199)
				p.Match(ValidationExprParserSTR)
			}

			p.SetState(202)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ValidationExprParserT__18 {
			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == ValidationExprParserT__18 {
				{
					p.SetState(204)
					p.Match(ValidationExprParserT__18)
				}

				p.SetState(207)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(209)
				p.Match(ValidationExprParserSTR)
			}

			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(214)
			p.Match(ValidationExprParserT__25)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRangeLiteralContext is an interface to support dynamic dispatch.
type IRangeLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeLiteralContext differentiates from other interfaces.
	IsRangeLiteralContext()
}

type RangeLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeLiteralContext() *RangeLiteralContext {
	var p = new(RangeLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_rangeLiteral
	return p
}

func (*RangeLiteralContext) IsRangeLiteralContext() {}

func NewRangeLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeLiteralContext {
	var p = new(RangeLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_rangeLiteral

	return p
}

func (s *RangeLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeLiteralContext) AllNumericLiteral() []INumericLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem())
	var tst = make([]INumericLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumericLiteralContext)
		}
	}

	return tst
}

func (s *RangeLiteralContext) NumericLiteral(i int) INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *RangeLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterRangeLiteral(s)
	}
}

func (s *RangeLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitRangeLiteral(s)
	}
}

func (p *ValidationExprParser) RangeLiteral() (localctx IRangeLiteralContext) {
	localctx = NewRangeLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ValidationExprParserRULE_rangeLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ValidationExprParserINTEGER {
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ValidationExprParserINTEGER {
			{
				p.SetState(217)
				p.NumericLiteral()
			}

			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ValidationExprParserT__18 {
			{
				p.SetState(222)
				p.Match(ValidationExprParserT__18)
			}

			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(227)
			p.NumericLiteral()
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) STR() antlr.TerminalNode {
	return s.GetToken(ValidationExprParserSTR, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (p *ValidationExprParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ValidationExprParserRULE_stringLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(ValidationExprParserSTR)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValidationExprParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValidationExprParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) LOCAL_VARIABLE() antlr.TerminalNode {
	return s.GetToken(ValidationExprParserLOCAL_VARIABLE, 0)
}

func (s *VariableContext) CONTEXT_VARIABLE() antlr.TerminalNode {
	return s.GetToken(ValidationExprParserCONTEXT_VARIABLE, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValidationExprListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *ValidationExprParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ValidationExprParserRULE_variable)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ValidationExprParserLOCAL_VARIABLE || _la == ValidationExprParserCONTEXT_VARIABLE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
