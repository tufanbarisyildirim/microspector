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
	yyTabOfs   = -82
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
		57361: 0,   // ASSERT (52x)
		57364: 1,   // CMD (52x)
		57359: 2,   // DEBUG (52x)
		57366: 3,   // ECHO (52x)
		57360: 4,   // END (52x)
		57356: 5,   // HTTP (52x)
		57362: 6,   // INCLUDE (52x)
		57357: 7,   // MUST (52x)
		57355: 8,   // SET (52x)
		57358: 9,   // SHOULD (52x)
		57363: 10,  // SLEEP (52x)
		57344: 11,  // $end (51x)
		57365: 12,  // ASYNC (51x)
		57404: 13,  // IDENTIFIER (51x)
		123:   14,  // '{' (50x)
		36:    15,  // '$' (49x)
		40:    16,  // '(' (46x)
		57352: 17,  // FALSE (46x)
		57350: 18,  // FLOAT (46x)
		57349: 19,  // INTEGER (46x)
		57348: 20,  // STRING (46x)
		57351: 21,  // TRUE (46x)
		57394: 22,  // WHEN (45x)
		57403: 23,  // INTO (44x)
		57395: 24,  // AND (26x)
		57391: 25,  // CONTAIN (26x)
		57390: 26,  // CONTAINS (26x)
		57383: 27,  // EQUAL (26x)
		57382: 28,  // EQUALS (26x)
		57387: 29,  // GE (26x)
		57386: 30,  // GT (26x)
		57402: 31,  // IN (26x)
		57399: 32,  // IS (26x)
		57400: 33,  // ISNOT (26x)
		57389: 34,  // LE (26x)
		57388: 35,  // LT (26x)
		57398: 36,  // MATCH (26x)
		57397: 37,  // MATCHES (26x)
		57401: 38,  // NOT (26x)
		57385: 39,  // NOTEQUAL (26x)
		57384: 40,  // NOTEQUALS (26x)
		57396: 41,  // OR (26x)
		57392: 42,  // STARTSWITH (26x)
		57393: 43,  // STARTWITH (26x)
		57377: 44,  // BODY (21x)
		57378: 45,  // FOLLOW (21x)
		57376: 46,  // HEADER (21x)
		57381: 47,  // INSECURE (21x)
		57379: 48,  // NOFOLLOW (21x)
		57380: 49,  // SECURE (21x)
		57429: 50,  // variable (19x)
		57416: 51,  // expr (16x)
		57424: 52,  // operator (16x)
		41:    53,  // ')' (12x)
		125:   54,  // '}' (2x)
		57408: 55,  // assert_command (2x)
		57409: 56,  // cmd_command (2x)
		57411: 57,  // command (2x)
		57412: 58,  // command_cond (2x)
		57413: 59,  // debug_command (2x)
		57414: 60,  // echo_command (2x)
		57415: 61,  // end_command (2x)
		57417: 62,  // http_command (2x)
		57418: 63,  // http_command_param (2x)
		57421: 64,  // include_command (2x)
		57423: 65,  // must_command (2x)
		57426: 66,  // set_command (2x)
		57427: 67,  // should_command (2x)
		57428: 68,  // sleep_command (2x)
		57367: 69,  // GET (1x)
		57368: 70,  // HEAD (1x)
		57419: 71,  // http_command_params (1x)
		57420: 72,  // http_method (1x)
		57422: 73,  // microspector (1x)
		57373: 74,  // OPTIONS (1x)
		57375: 75,  // PATCH (1x)
		57369: 76,  // POST (1x)
		57370: 77,  // PUT (1x)
		57425: 78,  // run_comm (1x)
		57406: 79,  // $default (0x)
		37:    80,  // '%' (0x)
		38:    81,  // '&' (0x)
		42:    82,  // '*' (0x)
		43:    83,  // '+' (0x)
		44:    84,  // ',' (0x)
		45:    85,  // '-' (0x)
		47:    86,  // '/' (0x)
		91:    87,  // '[' (0x)
		93:    88,  // ']' (0x)
		124:   89,  // '|' (0x)
		57407: 90,  // array (0x)
		57410: 91,  // comma_separated_expressions (0x)
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
		"ASYNC",
		"IDENTIFIER",
		"'{'",
		"'$'",
		"'('",
		"FALSE",
		"FLOAT",
		"INTEGER",
		"STRING",
		"TRUE",
		"WHEN",
		"INTO",
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
		"'%'",
		"'&'",
		"'*'",
		"'+'",
		"','",
		"'-'",
		"'/'",
		"'['",
		"']'",
		"'|'",
		"array",
		"comma_separated_expressions",
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
		1:  {73, 0},
		2:  {73, 2},
		3:  {78, 1},
		4:  {78, 2},
		5:  {58, 3},
		6:  {58, 5},
		7:  {58, 3},
		8:  {58, 1},
		9:  {57, 1},
		10: {57, 1},
		11: {57, 1},
		12: {57, 1},
		13: {57, 1},
		14: {57, 1},
		15: {57, 1},
		16: {57, 1},
		17: {57, 1},
		18: {57, 1},
		19: {57, 1},
		20: {66, 3},
		21: {62, 3},
		22: {62, 4},
		23: {71, 1},
		24: {71, 2},
		25: {63, 2},
		26: {63, 2},
		27: {63, 1},
		28: {63, 1},
		29: {63, 1},
		30: {63, 1},
		31: {59, 2},
		32: {61, 2},
		33: {55, 2},
		34: {65, 2},
		35: {67, 2},
		36: {64, 2},
		37: {68, 2},
		38: {56, 2},
		39: {60, 2},
		40: {72, 1},
		41: {72, 1},
		42: {72, 1},
		43: {72, 1},
		44: {72, 1},
		45: {72, 1},
		46: {90, 2},
		47: {90, 3},
		48: {91, 1},
		49: {91, 3},
		50: {51, 3},
		51: {51, 1},
		52: {51, 1},
		53: {51, 1},
		54: {51, 1},
		55: {51, 1},
		56: {51, 1},
		57: {51, 3},
		58: {50, 5},
		59: {50, 2},
		60: {50, 1},
		61: {52, 1},
		62: {52, 1},
		63: {52, 1},
		64: {52, 1},
		65: {52, 1},
		66: {52, 1},
		67: {52, 1},
		68: {52, 1},
		69: {52, 1},
		70: {52, 1},
		71: {52, 1},
		72: {52, 1},
		73: {52, 1},
		74: {52, 1},
		75: {52, 1},
		76: {52, 1},
		77: {52, 1},
		78: {52, 1},
		79: {52, 1},
		80: {52, 1},
		81: {52, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [105][]uint16{
		// 0
		{81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 73: 83},
		{103, 108, 101, 109, 102, 100, 106, 104, 99, 105, 107, 82, 86, 55: 92, 97, 87, 85, 90, 98, 91, 89, 64: 95, 93, 88, 94, 96, 78: 84},
		{80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80},
		{79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79},
		{103, 108, 101, 109, 102, 100, 106, 104, 99, 105, 107, 55: 92, 97, 87, 186, 90, 98, 91, 89, 64: 95, 93, 88, 94, 96},
		// 5
		{74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 22: 180, 181},
		{73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 22: 73, 73},
		{72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 22: 72, 72},
		{71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 22: 71, 71},
		{70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 22: 70, 70},
		// 10
		{69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 22: 69, 69},
		{68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 22: 68, 68},
		{67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 22: 67, 67},
		{66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 22: 66, 66},
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 22: 65, 65},
		// 15
		{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 22: 64, 64},
		{63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 22: 63, 63},
		{13: 120, 118, 119, 50: 178},
		{69: 160, 162, 72: 159, 74: 163, 165, 161, 164},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 158},
		// 20
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 157},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 156},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 155},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 154},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 153},
		// 25
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 152},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 151},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 110},
		{43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 22: 141, 43, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 126},
		// 30
		{31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 22: 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 53: 31},
		{30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 22: 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 53: 30},
		{29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 22: 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 53: 29},
		{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 22: 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 53: 28},
		{27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 22: 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 53: 27},
		// 35
		{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 22: 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 53: 26},
		{14: 122},
		{13: 121},
		{22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 53: 22},
		{23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 53: 23},
		// 40
		{13: 123},
		{54: 124},
		{54: 125},
		{24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 53: 24},
		{22: 141, 24: 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128, 127},
		// 45
		{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 22: 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 53: 32},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 150},
		{13: 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{13: 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{13: 19, 19, 19, 19, 19, 19, 19, 19, 19},
		// 50
		{13: 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{13: 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{13: 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{13: 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{13: 14, 14, 14, 14, 14, 14, 14, 14, 14},
		// 55
		{13: 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{13: 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{13: 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{13: 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{13: 9, 9, 9, 9, 9, 9, 9, 9, 9},
		// 60
		{13: 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{13: 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{13: 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{13: 5, 5, 5, 5, 5, 5, 5, 5, 5},
		{13: 4, 4, 4, 4, 4, 4, 4, 4, 4},
		// 65
		{13: 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{13: 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{13: 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 22: 141, 25, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 25, 25, 25, 25, 25, 25, 52: 128, 25},
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 22: 141, 44, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		// 70
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 22: 141, 45, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		{46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 22: 141, 46, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		{47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 22: 141, 47, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 22: 141, 48, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 22: 141, 49, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		// 75
		{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 22: 141, 50, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		{51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 22: 141, 51, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 166},
		{13: 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{13: 41, 41, 41, 41, 41, 41, 41, 41, 41},
		// 80
		{13: 40, 40, 40, 40, 40, 40, 40, 40, 40},
		{13: 39, 39, 39, 39, 39, 39, 39, 39, 39},
		{13: 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{13: 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 22: 141, 61, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 170, 171, 169, 173, 172, 174, 52: 128, 63: 168, 71: 167},
		// 85
		{60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 22: 60, 60, 44: 170, 171, 169, 173, 172, 174, 63: 177},
		{59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 22: 59, 59, 44: 59, 59, 59, 59, 59, 59},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 176},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 175},
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 22: 55, 55, 44: 55, 55, 55, 55, 55, 55},
		// 90
		{54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 22: 54, 54, 44: 54, 54, 54, 54, 54, 54},
		{53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 22: 53, 53, 44: 53, 53, 53, 53, 53, 53},
		{52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 22: 52, 52, 44: 52, 52, 52, 52, 52, 52},
		{56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 22: 141, 56, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 56, 56, 56, 56, 56, 56, 52: 128},
		{57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 22: 141, 57, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 57, 57, 57, 57, 57, 57, 52: 128},
		// 95
		{58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 22: 58, 58, 44: 58, 58, 58, 58, 58, 58},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 179},
		{62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 22: 141, 62, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		{13: 120, 118, 119, 111, 117, 113, 114, 112, 116, 50: 115, 183},
		{13: 120, 118, 119, 50: 182},
		// 100
		{75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75},
		{77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 22: 141, 184, 142, 138, 137, 130, 129, 134, 133, 149, 146, 147, 136, 135, 145, 144, 148, 132, 131, 143, 139, 140, 52: 128},
		{13: 120, 118, 119, 50: 185},
		{76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76},
		{78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78},
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
			yyVAL.cmd = &AssertCommand{}
		}
	case 34:
		{
			yyVAL.cmd = &MustCommand{}
		}
	case 35:
		{
			yyVAL.cmd = &ShouldCommand{}
		}
	case 36:
		{
			yyVAL.cmd = &IncludeCommand{}
		}
	case 37:
		{
			yyVAL.cmd = &SleepCommand{}
		}
	case 38:
		{
			yyVAL.cmd = &CmdCommand{}
		}
	case 39:
		{
			yyVAL.cmd = &EchoCommand{}
		}
	case 46:
		{
			yyVAL.expressions = ExprArray{}
		}
	case 47:
		{
			yyVAL.expressions = yyS[yypt-1].expressions
		}
	case 48:
		{
			yyVAL.expressions.Values = append(yyVAL.expressions.Values, yyS[yypt-0].expression)
		}
	case 49:
		{
			yyVAL.expressions.Values = append(yyVAL.expressions.Values, yyS[yypt-0].expression)
		}
	case 50:
		{
			yyVAL.expression = yyS[yypt-1].expression
		}
	case 51:
		{
			yyVAL.expression = &ExprString{
				Val: yyS[yypt-0].val.(string),
			}
		}
	case 52:
		{
			yyVAL.expression = &ExprFloat{
				Val: yyS[yypt-0].val.(float64),
			}
		}
	case 53:
		{
			yyVAL.expression = &ExprInteger{
				Val: yyS[yypt-0].val.(int64),
			}
		}
	case 54:
		{

			yyVAL.expression = &ExprVariable{
				Name: yyS[yypt-0].variable.Name,
			}
		}
	case 55:
		{
			yyVAL.expression = &ExprBool{
				Val: true,
			}
		}
	case 56:
		{
			yyVAL.expression = &ExprBool{
				Val: false,
			}
		}
	case 57:
		{
			yyVAL.expression = &ExprPredicate{
				Left:     yyS[yypt-2].expression,
				Operator: yyS[yypt-1].val.(string),
				Right:    yyS[yypt-0].expression,
			}
		}
	case 58:
		{
			yyVAL.variable = ExprVariable{
				Name: yyS[yypt-2].val.(string),
			}
		}
	case 59:
		{
			yyVAL.variable = ExprVariable{
				Name: yyS[yypt-0].val.(string),
			}
		}
	case 60:
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
