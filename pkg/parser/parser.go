// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"log"
	"strconv"
	"strings"
)

var globalvars = map[string]interface{}{}

type yySymType struct {
	yys      int
	val      interface{}
	str      string
	integer  int
	boolean  bool
	flt      int64
	bytes    []byte
	cmd      Command
	variable struct {
		name  string
		value interface{}
	}
	http_command_params []HttpCommandParam
	http_command_param  HttpCommandParam
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault  = 57383
	yyEofCode  = 57344
	AND        = 57379
	ASSERT     = 57360
	CONNECT    = 57366
	CONTAINS   = 57376
	DEBUG      = 57358
	DELETE     = 57365
	END        = 57359
	EOF        = 57347
	EOL        = 57346
	EQUALS     = 57372
	FALSE      = 57352
	FLOAT      = 57350
	GET        = 57361
	GT         = 57374
	HEAD       = 57362
	HEADER     = 57370
	HTTP       = 57355
	IDENTIFIER = 57382
	INTEGER    = 57349
	INTO       = 57381
	KEYWORD    = 57353
	LT         = 57375
	MUST       = 57356
	NOTEQUALS  = 57373
	OPTIONS    = 57367
	OR         = 57380
	PATCH      = 57369
	POST       = 57363
	PUT        = 57364
	QUERY      = 57371
	SET        = 57354
	SHOULD     = 57357
	STARTSWITH = 57377
	STRING     = 57348
	TRACE      = 57368
	TRUE       = 57351
	WHEN       = 57378
	yyErrCode  = 57345

	yyMaxDepth = 200
	yyTabOfs   = -56
)

var (
	yyPrec = map[int]int{}

	yyXLAT = map[int]int{
		57360: 0,  // ASSERT (42x)
		57358: 1,  // DEBUG (42x)
		57359: 2,  // END (42x)
		57355: 3,  // HTTP (42x)
		57356: 4,  // MUST (42x)
		57354: 5,  // SET (42x)
		57357: 6,  // SHOULD (42x)
		57344: 7,  // $end (41x)
		57378: 8,  // WHEN (37x)
		123:   9,  // '{' (35x)
		57381: 10, // INTO (35x)
		57348: 11, // STRING (32x)
		57376: 12, // CONTAINS (23x)
		57372: 13, // EQUALS (23x)
		57374: 14, // GT (23x)
		57375: 15, // LT (23x)
		57373: 16, // NOTEQUALS (23x)
		57377: 17, // STARTSWITH (23x)
		40:    18, // '(' (20x)
		57379: 19, // AND (20x)
		57352: 20, // FALSE (20x)
		57350: 21, // FLOAT (20x)
		57349: 22, // INTEGER (20x)
		57380: 23, // OR (20x)
		57351: 24, // TRUE (20x)
		57401: 25, // variable (18x)
		41:    26, // ')' (13x)
		57385: 27, // any_value (13x)
		57387: 28, // boolean_exp (13x)
		57370: 29, // HEADER (9x)
		57371: 30, // QUERY (9x)
		57397: 31, // operator (4x)
		57400: 32, // string_or_var (3x)
		125:   33, // '}' (2x)
		57386: 34, // assert_command (2x)
		57388: 35, // command (2x)
		57389: 36, // command_with_condition_opt (2x)
		57390: 37, // debug_command (2x)
		57391: 38, // end_command (2x)
		57392: 39, // http_command (2x)
		57393: 40, // http_command_param (2x)
		57396: 41, // must_command (2x)
		57398: 42, // set_command (2x)
		57399: 43, // should_command (2x)
		57384: 44, // any_command (1x)
		57366: 45, // CONNECT (1x)
		57365: 46, // DELETE (1x)
		57361: 47, // GET (1x)
		57362: 48, // HEAD (1x)
		57394: 49, // http_command_params (1x)
		57395: 50, // http_method (1x)
		57382: 51, // IDENTIFIER (1x)
		57367: 52, // OPTIONS (1x)
		57369: 53, // PATCH (1x)
		57363: 54, // POST (1x)
		57364: 55, // PUT (1x)
		57368: 56, // TRACE (1x)
		57383: 57, // $default (0x)
		57347: 58, // EOF (0x)
		57346: 59, // EOL (0x)
		57345: 60, // error (0x)
		57353: 61, // KEYWORD (0x)
	}

	yySymNames = []string{
		"ASSERT",
		"DEBUG",
		"END",
		"HTTP",
		"MUST",
		"SET",
		"SHOULD",
		"$end",
		"WHEN",
		"'{'",
		"INTO",
		"STRING",
		"CONTAINS",
		"EQUALS",
		"GT",
		"LT",
		"NOTEQUALS",
		"STARTSWITH",
		"'('",
		"AND",
		"FALSE",
		"FLOAT",
		"INTEGER",
		"OR",
		"TRUE",
		"variable",
		"')'",
		"any_value",
		"boolean_exp",
		"HEADER",
		"QUERY",
		"operator",
		"string_or_var",
		"'}'",
		"assert_command",
		"command",
		"command_with_condition_opt",
		"debug_command",
		"end_command",
		"http_command",
		"http_command_param",
		"must_command",
		"set_command",
		"should_command",
		"any_command",
		"CONNECT",
		"DELETE",
		"GET",
		"HEAD",
		"http_command_params",
		"http_method",
		"IDENTIFIER",
		"OPTIONS",
		"PATCH",
		"POST",
		"PUT",
		"TRACE",
		"$default",
		"EOF",
		"EOL",
		"error",
		"KEYWORD",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {44, 1},
		2:  {44, 2},
		3:  {36, 5},
		4:  {36, 3},
		5:  {36, 3},
		6:  {36, 1},
		7:  {35, 1},
		8:  {35, 1},
		9:  {35, 1},
		10: {35, 1},
		11: {35, 1},
		12: {35, 1},
		13: {35, 1},
		14: {37, 2},
		15: {38, 3},
		16: {38, 2},
		17: {34, 2},
		18: {41, 2},
		19: {43, 2},
		20: {42, 3},
		21: {39, 4},
		22: {39, 3},
		23: {49, 1},
		24: {49, 2},
		25: {40, 2},
		26: {40, 2},
		27: {50, 1},
		28: {50, 1},
		29: {50, 1},
		30: {50, 1},
		31: {50, 1},
		32: {50, 1},
		33: {50, 1},
		34: {50, 1},
		35: {50, 1},
		36: {27, 1},
		37: {27, 1},
		38: {27, 1},
		39: {27, 1},
		40: {27, 1},
		41: {32, 1},
		42: {32, 1},
		43: {25, 5},
		44: {31, 1},
		45: {31, 1},
		46: {31, 1},
		47: {31, 1},
		48: {31, 1},
		49: {31, 1},
		50: {28, 1},
		51: {28, 1},
		52: {28, 3},
		53: {28, 3},
		54: {28, 3},
		55: {28, 3},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [82][]uint16{
		// 0
		{69, 67, 68, 73, 70, 72, 71, 34: 64, 59, 58, 62, 63, 61, 41: 65, 60, 66, 57},
		{69, 67, 68, 73, 70, 72, 71, 56, 34: 64, 59, 137, 62, 63, 61, 41: 65, 60, 66},
		{55, 55, 55, 55, 55, 55, 55, 55},
		{50, 50, 50, 50, 50, 50, 50, 50, 132, 10: 131},
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 10: 49},
		// 5
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 10: 48},
		{47, 47, 47, 47, 47, 47, 47, 47, 47, 10: 47},
		{46, 46, 46, 46, 46, 46, 46, 46, 46, 10: 46},
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 10: 45},
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 10: 44},
		// 10
		{43, 43, 43, 43, 43, 43, 43, 43, 43, 10: 43},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 130, 104},
		{8: 127, 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 110, 128},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 110, 126},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 110, 125},
		// 15
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 110, 124},
		{9: 87, 25: 99},
		{45: 80, 79, 75, 76, 50: 74, 52: 81, 83, 77, 78, 82},
		{9: 87, 11: 86, 25: 85, 32: 84},
		{9: 29, 11: 29},
		// 20
		{9: 28, 11: 28},
		{9: 27, 11: 27},
		{9: 26, 11: 26},
		{9: 25, 11: 25},
		{9: 24, 11: 24},
		// 25
		{9: 23, 11: 23},
		{9: 22, 11: 22},
		{9: 21, 11: 21},
		{34, 34, 34, 34, 34, 34, 34, 34, 34, 10: 34, 29: 94, 95, 40: 93, 49: 92},
		{15, 15, 15, 15, 15, 15, 15, 15, 15, 10: 15, 29: 15, 15},
		// 30
		{14, 14, 14, 14, 14, 14, 14, 14, 14, 10: 14, 29: 14, 14},
		{9: 88},
		{51: 89},
		{33: 90},
		{33: 91},
		// 35
		{13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 26: 13, 29: 13, 13},
		{35, 35, 35, 35, 35, 35, 35, 35, 35, 10: 35, 29: 94, 95, 40: 98},
		{33, 33, 33, 33, 33, 33, 33, 33, 33, 10: 33, 29: 33, 33},
		{9: 87, 11: 86, 25: 85, 32: 97},
		{9: 87, 11: 86, 25: 85, 32: 96},
		// 40
		{30, 30, 30, 30, 30, 30, 30, 30, 30, 10: 30, 29: 30, 30},
		{31, 31, 31, 31, 31, 31, 31, 31, 31, 10: 31, 29: 31, 31},
		{32, 32, 32, 32, 32, 32, 32, 32, 32, 10: 32, 29: 32, 32},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 100, 104},
		{36, 36, 36, 36, 36, 36, 36, 36, 36, 10: 36, 12: 115, 111, 113, 114, 112, 116, 31: 117},
		// 45
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 10: 20, 12: 20, 20, 20, 20, 20, 20, 19: 20, 23: 20, 26: 20},
		{19, 19, 19, 19, 19, 19, 19, 19, 19, 10: 19, 12: 19, 19, 19, 19, 19, 19, 19: 19, 23: 19, 26: 19},
		{18, 18, 18, 18, 18, 18, 18, 18, 18, 10: 18, 12: 18, 18, 18, 18, 18, 18, 19: 18, 23: 18, 26: 18},
		{17, 17, 17, 17, 17, 17, 17, 17, 17, 10: 17, 12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120, 26: 17},
		{16, 16, 16, 16, 16, 16, 16, 16, 16, 10: 16, 12: 16, 16, 16, 16, 16, 16, 19: 16, 23: 16, 26: 16},
		// 50
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 10: 6, 12: 6, 6, 6, 6, 6, 6, 19: 6, 23: 6, 26: 6},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 10: 5, 12: 5, 5, 5, 5, 5, 5, 19: 5, 23: 5, 26: 5},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 110, 109},
		{12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120, 26: 121},
		{12: 115, 111, 113, 114, 112, 116, 31: 117},
		// 55
		{9: 12, 11: 12, 18: 12, 20: 12, 12, 12, 24: 12},
		{9: 11, 11: 11, 18: 11, 20: 11, 11, 11, 24: 11},
		{9: 10, 11: 10, 18: 10, 20: 10, 10, 10, 24: 10},
		{9: 9, 11: 9, 18: 9, 20: 9, 9, 9, 24: 9},
		{9: 8, 11: 8, 18: 8, 20: 8, 8, 8, 24: 8},
		// 60
		{9: 7, 11: 7, 18: 7, 20: 7, 7, 7, 24: 7},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 118, 104},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 10: 1, 12: 115, 111, 113, 114, 112, 116, 19: 1, 23: 1, 26: 1, 31: 117},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 110, 123},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 110, 122},
		// 65
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 10: 2, 12: 2, 2, 2, 2, 2, 2, 19: 2, 23: 2, 26: 2},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 10: 3, 12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120, 26: 3},
		{4, 4, 4, 4, 4, 4, 4, 4, 4, 10: 4, 12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120, 26: 4},
		{37, 37, 37, 37, 37, 37, 37, 37, 37, 10: 37, 12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120},
		{38, 38, 38, 38, 38, 38, 38, 38, 38, 10: 38, 12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120},
		// 70
		{39, 39, 39, 39, 39, 39, 39, 39, 39, 10: 39, 12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 110, 129},
		{40, 40, 40, 40, 40, 40, 40, 40, 40, 10: 40, 12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120},
		{41, 41, 41, 41, 41, 41, 41, 41, 41, 10: 41, 12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120},
		{42, 42, 42, 42, 42, 42, 42, 42, 42, 10: 42, 12: 115, 111, 113, 114, 112, 116, 31: 117},
		// 75
		{9: 87, 25: 134},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 110, 133},
		{51, 51, 51, 51, 51, 51, 51, 51, 12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120},
		{52, 52, 52, 52, 52, 52, 52, 52, 135},
		{9: 87, 11: 101, 18: 108, 20: 107, 102, 103, 24: 106, 105, 27: 110, 136},
		// 80
		{53, 53, 53, 53, 53, 53, 53, 53, 12: 17, 17, 17, 17, 17, 17, 19: 119, 23: 120},
		{54, 54, 54, 54, 54, 54, 54, 54},
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
	const yyError = 60

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
			//run command put result into variable WHEN boolean_exp is true

			if strings.Contains(yyS[yypt-2].variable.name, ".") {
				yylex.Error("nested variables are not supported yet")
			}

			if yyS[yypt-0].boolean {
				globalvars[yyS[yypt-2].variable.name] = yyS[yypt-4].cmd.Run()
			}
		}
	case 4:
		{
			if strings.Contains(yyS[yypt-0].variable.name, ".") {
				yylex.Error("nested variables are not supported yet")
			}
			yyS[yypt-0].variable.value = yyS[yypt-2].cmd.Run()
		}
	case 5:
		{
			//run the command only if boolean_exp is true
			if yyS[yypt-0].boolean {
				yyS[yypt-2].cmd.Run()
			}
		}
	case 6:
		{
			//just run the command
			yyS[yypt-0].cmd.Run()
			//run command without condition

		}
	case 14:
		{
			fmt.Printf("%+v\n", yyS[yypt-0].val)
			yyVAL.cmd = &DebugCommand{}
		}
	case 15:
		{
			if yyS[yypt-0].boolean {
				return -1
			}
			yyVAL.cmd = &EndCommand{}
		}
	case 16:
		{
			if yyS[yypt-0].boolean {
				return -1
			}

			yyVAL.cmd = &EndCommand{}
		}
	case 17:
		{
			yyVAL.cmd = &AssertCommand{}
		}
	case 18:
		{
			//if $2 is false, fail
			yyVAL.cmd = &MustCommand{}
		}
	case 19:
		{
			//if $2 is false, write a warning
			yyVAL.cmd = &ShouldCommand{}
		}
	case 20:
		{
			//globalvars[$2.name] = $3
			yyVAL.cmd = &SetCommand{
				Name:  yyS[yypt-1].variable.name,
				Value: yyS[yypt-0].val,
			}
		}
	case 21:
		{
			//call http with header here.
			yyVAL.cmd = &HttpCommand{}
		}
	case 22:
		{
			yyVAL.cmd = &HttpCommand{}
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
				ParamValue: yyS[yypt-0].val.(string),
			}
		}
	case 26:
		{
			//adding query param
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-1].val.(string),
				ParamValue: yyS[yypt-0].val.(string),
			}
		}
	case 27:
		{
			//http get
			yyVAL.val = yyS[yypt-0].val
		}
	case 28:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 29:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 30:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 31:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 32:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 33:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 34:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 35:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 36:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 37:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 38:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 39:
		{
			yyVAL.val = yyS[yypt-0].boolean
		}
	case 40:
		{
			yyVAL.val = yyS[yypt-0].variable.value
		}
	case 41:
		{
			//found variable
			yyVAL.val = yyS[yypt-0].variable.value
		}
	case 42:
		{
			//found string
			yyVAL.val = yyS[yypt-0].val
		}
	case 43:
		{
			//getting variable
			yyVAL.variable.name = yyS[yypt-2].val.(string)
			yyVAL.variable.value = query(yyS[yypt-2].val.(string), globalvars)
		}
	case 50:
		{
			yyVAL.boolean = true
		}
	case 51:
		{
			yyVAL.boolean = false
		}
	case 52:
		{
			yyVAL.boolean = yyS[yypt-2].boolean && yyS[yypt-0].boolean
		}
	case 53:
		{
			yyVAL.boolean = yyS[yypt-2].boolean || yyS[yypt-0].boolean
		}
	case 54:
		{
			yyVAL.boolean = yyS[yypt-1].boolean
		}
	case 55:
		{
			//what should we do here?
			operator_result := runop(yyS[yypt-2].val, yyS[yypt-1].val, yyS[yypt-0].val)
			yyVAL.boolean = operator_result
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}

func runop(left, operator, right interface{}) bool {
	switch operator {
	case "EQUALS":
		return left == right
	case "NOTEQUALS":
		return left != right
	case "CONTAINS":
		return strings.Contains(fmt.Sprintf("%s", left), fmt.Sprintf("%s", right))
	case "LT", "GT":
		ll, err := strconv.Atoi(left.(string))
		if err != nil {
			fmt.Println(err)
		}
		rr, err := strconv.Atoi(right.(string))
		if err != nil {
			fmt.Println(err)
		}

		if operator == "GT" {
			return ll > rr
		} else {
			return ll < rr
		}

	}

	return false
}

func query(fieldPath string, thevars map[string]interface{}) interface{} {
	b, _ := json.Marshal(thevars)
	return gojsonq.New().JSONString(string(b)).Find(fieldPath)
}

type lex struct {
	tokens []Token
}

func (l *lex) Lex(lval *yySymType) int {
	if len(l.tokens) == 0 {
		return 0
	}

	v := l.tokens[0]
	l.tokens = l.tokens[1:]
	lval.val = v.Text
	return v.Type
}

func (l *lex) Error(e string) {
	log.Fatal(e)
}

func Parse(text string) {
	s := NewScanner(strings.NewReader(strings.TrimSpace(text)))
	tokens := make(Tokens, 0)

	for {
		token := s.Scan()
		if token.Type == EOF || token.Type == -1 {
			break
		}
		tokens = append(tokens, token)
	}

	l := &lex{tokens}

	yyParse(l)
	fmt.Println(globalvars)
}
