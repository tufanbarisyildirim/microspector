// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import (
	"strings"
	"sync"
)

type yySymType struct {
	yys                 int
	expression          Expression
	expressions         ExprArray
	val                 interface{}
	vals                []interface{}
	str                 ExprString
	integer             ExprInteger
	boolean             ExprBool
	bytes               []byte
	cmd                 Command
	variable            ExprVariable
	http_command_params []HttpCommandParam
	http_command_param  HttpCommandParam
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault  = 57406
	yyEofCode  = 57344
	AND        = 57395
	ASSERT     = 57361
	ASYNC      = 57365
	BODY       = 57377
	CMD        = 57364
	CONNECT    = 57372
	CONTAIN    = 57391
	CONTAINS   = 57390
	DEBUG      = 57359
	DELETE     = 57371
	ECHO       = 57366
	END        = 57360
	EOF        = 57347
	EOL        = 57346
	EQUAL      = 57383
	EQUALS     = 57382
	FALSE      = 57352
	FLOAT      = 57350
	FOLLOW     = 57378
	GE         = 57387
	GET        = 57367
	GT         = 57386
	HEAD       = 57368
	HEADER     = 57376
	HTTP       = 57356
	IDENTIFIER = 57404
	IN         = 57402
	INCLUDE    = 57362
	INSECURE   = 57381
	INTEGER    = 57349
	INTO       = 57403
	IS         = 57399
	ISNOT      = 57400
	KEYWORD    = 57354
	LE         = 57389
	LT         = 57388
	MATCH      = 57398
	MATCHES    = 57397
	MUST       = 57357
	NOFOLLOW   = 57379
	NOT        = 57401
	NOTEQUAL   = 57385
	NOTEQUALS  = 57384
	NULL       = 57353
	OPTIONS    = 57373
	OR         = 57396
	PATCH      = 57375
	POST       = 57369
	PUT        = 57370
	SECURE     = 57380
	SET        = 57355
	SHOULD     = 57358
	SLEEP      = 57363
	STARTSWITH = 57392
	STARTWITH  = 57393
	STRING     = 57348
	TRACE      = 57374
	TRUE       = 57351
	TYPE       = 57405
	WHEN       = 57394
	yyErrCode  = 57345

	yyMaxDepth = 200
	yyTabOfs   = -91
)

var (
	yyPrec = map[int]int{
		'|': 0,
		'&': 1,
		'+': 2,
		'-': 2,
		'*': 3,
		'/': 3,
		'%': 3,
	}

	yyXLAT = map[int]int{
		37:    0,   // '%' (85x)
		45:    1,   // '-' (85x)
		57404: 2,   // IDENTIFIER (59x)
		123:   3,   // '{' (58x)
		57361: 4,   // ASSERT (58x)
		57364: 5,   // CMD (58x)
		57359: 6,   // DEBUG (58x)
		57366: 7,   // ECHO (58x)
		57360: 8,   // END (58x)
		57356: 9,   // HTTP (58x)
		57362: 10,  // INCLUDE (58x)
		57357: 11,  // MUST (58x)
		57355: 12,  // SET (58x)
		57358: 13,  // SHOULD (58x)
		57363: 14,  // SLEEP (58x)
		57344: 15,  // $end (57x)
		36:    16,  // '$' (57x)
		57365: 17,  // ASYNC (57x)
		57349: 18,  // INTEGER (56x)
		40:    19,  // '(' (54x)
		91:    20,  // '[' (54x)
		57352: 21,  // FALSE (54x)
		57350: 22,  // FLOAT (54x)
		57348: 23,  // STRING (54x)
		57351: 24,  // TRUE (54x)
		57394: 25,  // WHEN (54x)
		57403: 26,  // INTO (50x)
		42:    27,  // '*' (34x)
		43:    28,  // '+' (34x)
		47:    29,  // '/' (34x)
		57395: 30,  // AND (34x)
		57391: 31,  // CONTAIN (34x)
		57390: 32,  // CONTAINS (34x)
		57383: 33,  // EQUAL (34x)
		57382: 34,  // EQUALS (34x)
		57387: 35,  // GE (34x)
		57386: 36,  // GT (34x)
		57402: 37,  // IN (34x)
		57399: 38,  // IS (34x)
		57400: 39,  // ISNOT (34x)
		57389: 40,  // LE (34x)
		57388: 41,  // LT (34x)
		57398: 42,  // MATCH (34x)
		57397: 43,  // MATCHES (34x)
		57401: 44,  // NOT (34x)
		57385: 45,  // NOTEQUAL (34x)
		57384: 46,  // NOTEQUALS (34x)
		57396: 47,  // OR (34x)
		57392: 48,  // STARTSWITH (34x)
		57393: 49,  // STARTWITH (34x)
		57377: 50,  // BODY (26x)
		57378: 51,  // FOLLOW (26x)
		57376: 52,  // HEADER (26x)
		57381: 53,  // INSECURE (26x)
		57379: 54,  // NOFOLLOW (26x)
		57380: 55,  // SECURE (26x)
		57429: 56,  // variable (22x)
		93:    57,  // ']' (20x)
		44:    58,  // ',' (19x)
		57407: 59,  // array (19x)
		57416: 60,  // expr (19x)
		57424: 61,  // operator (19x)
		41:    62,  // ')' (17x)
		125:   63,  // '}' (2x)
		57408: 64,  // assert_command (2x)
		57409: 65,  // cmd_command (2x)
		57411: 66,  // command (2x)
		57412: 67,  // command_cond (2x)
		57413: 68,  // debug_command (2x)
		57414: 69,  // echo_command (2x)
		57415: 70,  // end_command (2x)
		57417: 71,  // http_command (2x)
		57418: 72,  // http_command_param (2x)
		57421: 73,  // include_command (2x)
		57423: 74,  // must_command (2x)
		57426: 75,  // set_command (2x)
		57427: 76,  // should_command (2x)
		57428: 77,  // sleep_command (2x)
		57410: 78,  // comma_separated_expressions (1x)
		57367: 79,  // GET (1x)
		57368: 80,  // HEAD (1x)
		57419: 81,  // http_command_params (1x)
		57420: 82,  // http_method (1x)
		57422: 83,  // microspector (1x)
		57373: 84,  // OPTIONS (1x)
		57375: 85,  // PATCH (1x)
		57369: 86,  // POST (1x)
		57370: 87,  // PUT (1x)
		57425: 88,  // run_comm (1x)
		57406: 89,  // $default (0x)
		38:    90,  // '&' (0x)
		124:   91,  // '|' (0x)
		57372: 92,  // CONNECT (0x)
		57371: 93,  // DELETE (0x)
		57347: 94,  // EOF (0x)
		57346: 95,  // EOL (0x)
		57345: 96,  // error (0x)
		57354: 97,  // KEYWORD (0x)
		57353: 98,  // NULL (0x)
		57374: 99,  // TRACE (0x)
		57405: 100, // TYPE (0x)
	}

	yySymNames = []string{
		"'%'",
		"'-'",
		"IDENTIFIER",
		"'{'",
		"ASSERT",
		"CMD",
		"DEBUG",
		"ECHO",
		"END",
		"HTTP",
		"INCLUDE",
		"MUST",
		"SET",
		"SHOULD",
		"SLEEP",
		"$end",
		"'$'",
		"ASYNC",
		"INTEGER",
		"'('",
		"'['",
		"FALSE",
		"FLOAT",
		"STRING",
		"TRUE",
		"WHEN",
		"INTO",
		"'*'",
		"'+'",
		"'/'",
		"AND",
		"CONTAIN",
		"CONTAINS",
		"EQUAL",
		"EQUALS",
		"GE",
		"GT",
		"IN",
		"IS",
		"ISNOT",
		"LE",
		"LT",
		"MATCH",
		"MATCHES",
		"NOT",
		"NOTEQUAL",
		"NOTEQUALS",
		"OR",
		"STARTSWITH",
		"STARTWITH",
		"BODY",
		"FOLLOW",
		"HEADER",
		"INSECURE",
		"NOFOLLOW",
		"SECURE",
		"variable",
		"']'",
		"','",
		"array",
		"expr",
		"operator",
		"')'",
		"'}'",
		"assert_command",
		"cmd_command",
		"command",
		"command_cond",
		"debug_command",
		"echo_command",
		"end_command",
		"http_command",
		"http_command_param",
		"include_command",
		"must_command",
		"set_command",
		"should_command",
		"sleep_command",
		"comma_separated_expressions",
		"GET",
		"HEAD",
		"http_command_params",
		"http_method",
		"microspector",
		"OPTIONS",
		"PATCH",
		"POST",
		"PUT",
		"run_comm",
		"$default",
		"'&'",
		"'|'",
		"CONNECT",
		"DELETE",
		"EOF",
		"EOL",
		"error",
		"KEYWORD",
		"NULL",
		"TRACE",
		"TYPE",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {83, 0},
		2:  {83, 2},
		3:  {88, 1},
		4:  {88, 2},
		5:  {67, 3},
		6:  {67, 5},
		7:  {67, 3},
		8:  {67, 1},
		9:  {66, 1},
		10: {66, 1},
		11: {66, 1},
		12: {66, 1},
		13: {66, 1},
		14: {66, 1},
		15: {66, 1},
		16: {66, 1},
		17: {66, 1},
		18: {66, 1},
		19: {66, 1},
		20: {75, 3},
		21: {71, 3},
		22: {71, 4},
		23: {81, 1},
		24: {81, 2},
		25: {72, 2},
		26: {72, 2},
		27: {72, 1},
		28: {72, 1},
		29: {72, 1},
		30: {72, 1},
		31: {68, 2},
		32: {70, 2},
		33: {70, 3},
		34: {64, 2},
		35: {74, 2},
		36: {76, 2},
		37: {73, 2},
		38: {77, 2},
		39: {65, 2},
		40: {69, 2},
		41: {82, 1},
		42: {82, 1},
		43: {82, 1},
		44: {82, 1},
		45: {82, 1},
		46: {82, 1},
		47: {59, 2},
		48: {59, 3},
		49: {78, 1},
		50: {78, 3},
		51: {60, 3},
		52: {60, 1},
		53: {60, 1},
		54: {60, 1},
		55: {60, 2},
		56: {60, 2},
		57: {60, 1},
		58: {60, 1},
		59: {60, 1},
		60: {60, 3},
		61: {60, 1},
		62: {56, 5},
		63: {56, 2},
		64: {56, 1},
		65: {61, 1},
		66: {61, 1},
		67: {61, 1},
		68: {61, 1},
		69: {61, 1},
		70: {61, 1},
		71: {61, 1},
		72: {61, 1},
		73: {61, 1},
		74: {61, 1},
		75: {61, 1},
		76: {61, 1},
		77: {61, 1},
		78: {61, 1},
		79: {61, 1},
		80: {61, 1},
		81: {61, 1},
		82: {61, 1},
		83: {61, 1},
		84: {61, 1},
		85: {61, 1},
		86: {61, 1},
		87: {61, 1},
		88: {61, 1},
		89: {61, 1},
		90: {61, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [124][]uint16{
		// 0
		{4: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 17: 90, 83: 92},
		{4: 112, 117, 110, 118, 111, 109, 115, 113, 108, 114, 116, 91, 17: 95, 64: 101, 106, 96, 94, 99, 107, 100, 98, 73: 104, 102, 97, 103, 105, 88: 93},
		{4: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 17: 89},
		{4: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 17: 88},
		{4: 112, 117, 110, 118, 111, 109, 115, 113, 108, 114, 116, 64: 101, 106, 96, 214, 99, 107, 100, 98, 73: 104, 102, 97, 103, 105},
		// 5
		{4: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 17: 83, 25: 208, 209},
		{4: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 17: 82, 25: 82, 82},
		{4: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 17: 81, 25: 81, 81},
		{4: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 17: 80, 25: 80, 80},
		{4: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 17: 79, 25: 79, 79},
		// 10
		{4: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 17: 78, 25: 78, 78},
		{4: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 17: 77, 25: 77, 77},
		{4: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 17: 76, 25: 76, 76},
		{4: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 17: 75, 25: 75, 75},
		{4: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 17: 74, 25: 74, 74},
		// 15
		{4: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 17: 73, 25: 73, 73},
		{4: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 17: 72, 25: 72, 72},
		{2: 133, 131, 16: 132, 56: 206},
		{79: 188, 190, 82: 187, 84: 191, 193, 189, 192},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 186},
		// 20
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 184, 56: 127, 59: 130, 183},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 182},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 181},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 180},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 179},
		// 25
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 178},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 177},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 119},
		{169, 168, 4: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 17: 51, 25: 156, 51, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 171, 59: 130, 173, 78: 172},
		// 30
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 141},
		{39, 39, 4: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 17: 39, 25: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 57: 39, 39, 62: 39},
		{38, 38, 4: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 17: 38, 25: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 57: 38, 38, 62: 38},
		{37, 37, 4: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 17: 37, 25: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 57: 37, 37, 62: 37},
		{18: 140},
		// 35
		{18: 139},
		{34, 34, 4: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 17: 34, 25: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 57: 34, 34, 62: 34},
		{33, 33, 4: 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 17: 33, 25: 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 57: 33, 33, 62: 33},
		{32, 32, 4: 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 17: 32, 25: 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 57: 32, 32, 62: 32},
		{30, 30, 4: 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 17: 30, 25: 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 57: 30, 30, 62: 30},
		// 40
		{3: 135},
		{2: 134},
		{27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 57: 27, 27, 62: 27},
		{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 57: 28, 28, 62: 28},
		{2: 136},
		// 45
		{63: 137},
		{63: 138},
		{29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 57: 29, 29, 62: 29},
		{35, 35, 4: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 17: 35, 25: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 57: 35, 35, 62: 35},
		{36, 36, 4: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 17: 36, 25: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 57: 36, 36, 62: 36},
		// 50
		{169, 168, 25: 156, 27: 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143, 142},
		{40, 40, 4: 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 17: 40, 25: 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 57: 40, 40, 62: 40},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 170},
		{26, 26, 26, 26, 16: 26, 18: 26, 26, 26, 26, 26, 26, 26},
		{25, 25, 25, 25, 16: 25, 18: 25, 25, 25, 25, 25, 25, 25},
		// 55
		{24, 24, 24, 24, 16: 24, 18: 24, 24, 24, 24, 24, 24, 24},
		{23, 23, 23, 23, 16: 23, 18: 23, 23, 23, 23, 23, 23, 23},
		{22, 22, 22, 22, 16: 22, 18: 22, 22, 22, 22, 22, 22, 22},
		{21, 21, 21, 21, 16: 21, 18: 21, 21, 21, 21, 21, 21, 21},
		{20, 20, 20, 20, 16: 20, 18: 20, 20, 20, 20, 20, 20, 20},
		// 60
		{19, 19, 19, 19, 16: 19, 18: 19, 19, 19, 19, 19, 19, 19},
		{18, 18, 18, 18, 16: 18, 18: 18, 18, 18, 18, 18, 18, 18},
		{17, 17, 17, 17, 16: 17, 18: 17, 17, 17, 17, 17, 17, 17},
		{16, 16, 16, 16, 16: 16, 18: 16, 16, 16, 16, 16, 16, 16},
		{15, 15, 15, 15, 16: 15, 18: 15, 15, 15, 15, 15, 15, 15},
		// 65
		{14, 14, 14, 14, 16: 14, 18: 14, 14, 14, 14, 14, 14, 14},
		{13, 13, 13, 13, 16: 13, 18: 13, 13, 13, 13, 13, 13, 13},
		{12, 12, 12, 12, 16: 12, 18: 12, 12, 12, 12, 12, 12, 12},
		{11, 11, 11, 11, 16: 11, 18: 11, 11, 11, 11, 11, 11, 11},
		{10, 10, 10, 10, 16: 10, 18: 10, 10, 10, 10, 10, 10, 10},
		// 70
		{9, 9, 9, 9, 16: 9, 18: 9, 9, 9, 9, 9, 9, 9},
		{8, 8, 8, 8, 16: 8, 18: 8, 8, 8, 8, 8, 8, 8},
		{7, 7, 7, 7, 16: 7, 18: 7, 7, 7, 7, 7, 7, 7},
		{6, 6, 6, 6, 16: 6, 18: 6, 6, 6, 6, 6, 6, 6},
		{5, 5, 5, 5, 16: 5, 18: 5, 5, 5, 5, 5, 5, 5},
		// 75
		{4, 4, 4, 4, 16: 4, 18: 4, 4, 4, 4, 4, 4, 4},
		{3, 3, 3, 3, 16: 3, 18: 3, 3, 3, 3, 3, 3, 3},
		{2, 2, 2, 2, 16: 2, 18: 2, 2, 2, 2, 2, 2, 2},
		{1, 1, 1, 1, 16: 1, 18: 1, 1, 1, 1, 1, 1, 1},
		{169, 168, 4: 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 17: 31, 25: 156, 31, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 31, 31, 31, 31, 31, 31, 57: 31, 31, 61: 143, 31},
		// 80
		{44, 44, 4: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 17: 44, 25: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 57: 44, 44, 62: 44},
		{57: 174, 175},
		{169, 168, 25: 156, 27: 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 57: 42, 42, 61: 143},
		{43, 43, 4: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 17: 43, 25: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 57: 43, 43, 62: 43},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 176},
		// 85
		{169, 168, 25: 156, 27: 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 57: 41, 41, 61: 143},
		{169, 168, 4: 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 17: 52, 25: 156, 52, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		{169, 168, 4: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 17: 53, 25: 156, 53, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		{169, 168, 4: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 17: 54, 25: 156, 54, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		{169, 168, 4: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 17: 55, 25: 156, 55, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		// 90
		{169, 168, 4: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 17: 56, 25: 156, 56, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		{169, 168, 4: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 17: 57, 25: 156, 57, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		{169, 168, 4: 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 17: 59, 25: 156, 59, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 185},
		{169, 168, 4: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 17: 58, 25: 156, 58, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		// 95
		{169, 168, 4: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 17: 60, 25: 156, 60, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 194},
		{50, 50, 50, 50, 16: 50, 18: 50, 50, 50, 50, 50, 50, 50},
		{49, 49, 49, 49, 16: 49, 18: 49, 49, 49, 49, 49, 49, 49},
		{48, 48, 48, 48, 16: 48, 18: 48, 48, 48, 48, 48, 48, 48},
		// 100
		{47, 47, 47, 47, 16: 47, 18: 47, 47, 47, 47, 47, 47, 47},
		{46, 46, 46, 46, 16: 46, 18: 46, 46, 46, 46, 46, 46, 46},
		{45, 45, 45, 45, 16: 45, 18: 45, 45, 45, 45, 45, 45, 45},
		{169, 168, 4: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 17: 70, 25: 156, 70, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 198, 199, 197, 201, 200, 202, 61: 143, 72: 196, 81: 195},
		{4: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 17: 69, 25: 69, 69, 50: 198, 199, 197, 201, 200, 202, 72: 205},
		// 105
		{4: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 17: 68, 25: 68, 68, 50: 68, 68, 68, 68, 68, 68},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 204},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 203},
		{4: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 17: 64, 25: 64, 64, 50: 64, 64, 64, 64, 64, 64},
		{4: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 17: 63, 25: 63, 63, 50: 63, 63, 63, 63, 63, 63},
		// 110
		{4: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 17: 62, 25: 62, 62, 50: 62, 62, 62, 62, 62, 62},
		{4: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 17: 61, 25: 61, 61, 50: 61, 61, 61, 61, 61, 61},
		{169, 168, 4: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 17: 65, 25: 156, 65, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 65, 65, 65, 65, 65, 65, 61: 143},
		{169, 168, 4: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 17: 66, 25: 156, 66, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 66, 66, 66, 66, 66, 66, 61: 143},
		{4: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 17: 67, 25: 67, 67, 50: 67, 67, 67, 67, 67, 67},
		// 115
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 207},
		{169, 168, 4: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 17: 71, 25: 156, 71, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		{125, 126, 133, 131, 16: 132, 18: 124, 121, 120, 129, 123, 122, 128, 56: 127, 59: 130, 211},
		{2: 133, 131, 16: 132, 56: 210},
		{4: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 17: 84},
		// 120
		{169, 168, 4: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 17: 86, 25: 156, 212, 165, 167, 166, 157, 153, 152, 145, 144, 149, 148, 164, 161, 162, 151, 150, 160, 159, 163, 147, 146, 158, 154, 155, 61: 143},
		{2: 133, 131, 16: 132, 56: 213},
		{4: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 17: 85},
		{4: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 17: 87},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 96

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 3:
		{
			yylex.(*Lexer).wg.Add(1)
			yyS[yypt-0].cmd.Run(yylex.(*Lexer))
		}
	case 4:
		{
			yylex.(*Lexer).wg.Add(1)
			go yyS[yypt-0].cmd.Run(yylex.(*Lexer))
		}
	case 5:
		{
			yyVAL.cmd = yyS[yypt-2].cmd
			yyVAL.cmd.SetWhen(yyS[yypt-0].expression)
		}
	case 6:
		{

			yyS[yypt-4].cmd.SetWhen(yyS[yypt-2].expression)
			//TODO: check if it compatible with SetInto
			yyS[yypt-4].cmd.(*HttpCommand).SetInto(yyS[yypt-0].variable.Name)
			yyVAL.cmd = yyS[yypt-4].cmd
		}
	case 7:
		{
			//TODO: check if it compatible with SetInto
			yyS[yypt-2].cmd.(*HttpCommand).SetInto(yyS[yypt-0].variable.Name)
			yyVAL.cmd = yyS[yypt-2].cmd
		}
	case 8:
		{
			yyVAL.cmd = yyS[yypt-0].cmd
		}
	case 20:
		{
			yyVAL.cmd = &SetCommand{
				Name:  yyS[yypt-1].variable.Name,
				Value: yyS[yypt-0].expression,
			}
		}
	case 21:
		{
			yyVAL.cmd = &HttpCommand{
				Method: yyS[yypt-1].val.(string),
				Url:    yyS[yypt-0].expression,
			}
		}
	case 22:
		{
			yyVAL.cmd = &HttpCommand{
				Method:        yyS[yypt-2].val.(string),
				Url:           yyS[yypt-1].expression,
				CommandParams: yyS[yypt-0].http_command_params,
			}
		}
	case 23:
		{
			if yyVAL.http_command_params == nil {
				yyVAL.http_command_params = make([]HttpCommandParam, 0)
			}
			yyVAL.http_command_params = append(yyVAL.http_command_params, yyS[yypt-0].http_command_param)
		}
	case 24:
		{
			if yyVAL.http_command_params == nil {
				yyVAL.http_command_params = make([]HttpCommandParam, 0)
			}

			yyVAL.http_command_params = append(yyVAL.http_command_params, yyS[yypt-0].http_command_param)
		}
	case 25:
		{
			//addin header
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-1].val.(string),
				ParamValue: yyS[yypt-0].expression,
			}
		}
	case 26:
		{
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-1].val.(string),
				ParamValue: yyS[yypt-0].expression,
			}
		}
	case 27:
		{
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-0].val.(string),
				ParamValue: &ExprBool{Val: true},
			}
		}
	case 28:
		{
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  "FOLLOW",
				ParamValue: &ExprBool{Val: false},
			}
		}
	case 29:
		{
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  "SECURE",
				ParamValue: &ExprBool{Val: false},
			}
		}
	case 30:
		{
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  "SECURE",
				ParamValue: &ExprBool{Val: true},
			}
		}
	case 31:
		{
			yyVAL.cmd = &DebugCommand{
				Values: yyS[yypt-0].expression,
			}
		}
	case 32:
		{
			yyVAL.cmd = &EndCommand{}
		}
	case 33:
		{
			yyVAL.cmd = &EndCommand{}
		}
	case 34:
		{
			yyVAL.cmd = &AssertCommand{}
		}
	case 35:
		{
			yyVAL.cmd = &MustCommand{}
		}
	case 36:
		{
			yyVAL.cmd = &ShouldCommand{}
		}
	case 37:
		{
			yyVAL.cmd = &IncludeCommand{}
		}
	case 38:
		{
			yyVAL.cmd = &SleepCommand{}
		}
	case 39:
		{
			yyVAL.cmd = &CmdCommand{}
		}
	case 40:
		{
			yyVAL.cmd = &EchoCommand{}
		}
	case 47:
		{
			yyVAL.expressions = ExprArray{}
		}
	case 48:
		{
			yyVAL.expressions = yyS[yypt-1].expressions
		}
	case 49:
		{
			yyVAL.expressions.Values = append(yyVAL.expressions.Values, yyS[yypt-0].expression)
		}
	case 50:
		{
			yyVAL.expressions.Values = append(yyVAL.expressions.Values, yyS[yypt-0].expression)
		}
	case 51:
		{
			yyVAL.expression = yyS[yypt-1].expression
		}
	case 52:
		{
			yyVAL.expression = &ExprString{
				Val: yyS[yypt-0].val.(string),
			}
		}
	case 53:
		{
			yyVAL.expression = &ExprFloat{
				Val: yyS[yypt-0].val.(float64),
			}
		}
	case 54:
		{
			yyVAL.expression = &ExprInteger{
				Val: yyS[yypt-0].val.(int64),
			}
		}
	case 55:
		{
			yyVAL.expression = &ExprInteger{
				Val: 1 / yyS[yypt-1].val.(int64),
			}
		}
	case 56:
		{
			yyVAL.expression = &ExprInteger{
				Val: -yyS[yypt-1].val.(int64),
			}
		}
	case 57:
		{

			yyVAL.expression = &ExprVariable{
				Name: yyS[yypt-0].variable.Name,
			}
		}
	case 58:
		{
			yyVAL.expression = &ExprBool{
				Val: true,
			}
		}
	case 59:
		{
			yyVAL.expression = &ExprBool{
				Val: false,
			}
		}
	case 60:
		{
			yyVAL.expression = &ExprPredicate{
				Left:     yyS[yypt-2].expression,
				Operator: yyS[yypt-1].val.(string),
				Right:    yyS[yypt-0].expression,
			}
		}
	case 61:
		{
			yyVAL.expression = &ExprArray{Values: yyS[yypt-0].expressions.Values}
		}
	case 62:
		{
			yyVAL.variable = ExprVariable{
				Name: yyS[yypt-2].val.(string),
			}
		}
	case 63:
		{
			yyVAL.variable = ExprVariable{
				Name: yyS[yypt-0].val.(string),
			}
		}
	case 64:
		{
			yyVAL.variable = ExprVariable{
				Name: yyS[yypt-0].val.(string),
			}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}

func Parse(text string) *Lexer {

	l := &Lexer{
		tokens:     make(chan Token),
		State:      NewStats(),
		GlobalVars: map[string]interface{}{},
		wg:         &sync.WaitGroup{},
	}

	l.GlobalVars["State"] = &l.State

	if Verbose {
		yyDebug = 3
	}

	go func() {
		s := NewScanner(strings.NewReader(strings.TrimSpace(text)))
		for {
			l.tokens <- s.Scan()
		}
	}()

	return l
}

func Run(l *Lexer) {
	yyParse(l)
	l.wg.Wait()
}
