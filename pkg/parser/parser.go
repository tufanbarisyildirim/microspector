// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

import (
	"log"
	"strconv"
	"strings"
)

var GlobalVars = map[string]interface{}{}

type yySymType struct {
	yys      int
	val      interface{}
	vals     []interface{}
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
	yyDefault  = 57385
	yyEofCode  = 57344
	AND        = 57381
	ASSERT     = 57360
	CONNECT    = 57366
	CONTAINS   = 57378
	DEBUG      = 57358
	DELETE     = 57365
	END        = 57359
	EOF        = 57347
	EOL        = 57346
	EQUALS     = 57372
	FALSE      = 57352
	FLOAT      = 57350
	GE         = 57375
	GET        = 57361
	GT         = 57374
	HEAD       = 57362
	HEADER     = 57370
	HTTP       = 57355
	IDENTIFIER = 57384
	INTEGER    = 57349
	INTO       = 57383
	KEYWORD    = 57353
	LE         = 57377
	LT         = 57376
	MUST       = 57356
	NOTEQUALS  = 57373
	OPTIONS    = 57367
	OR         = 57382
	PATCH      = 57369
	POST       = 57363
	PUT        = 57364
	QUERY      = 57371
	SET        = 57354
	SHOULD     = 57357
	STARTSWITH = 57379
	STRING     = 57348
	TRACE      = 57368
	TRUE       = 57351
	WHEN       = 57380
	yyErrCode  = 57345

	yyMaxDepth = 200
	yyTabOfs   = -62
)

var (
	yyPrec = map[int]int{}

	yyXLAT = map[int]int{
		123:   0,  // '{' (45x)
		57348: 1,  // STRING (44x)
		57344: 2,  // $end (43x)
		57360: 3,  // ASSERT (43x)
		57358: 4,  // DEBUG (43x)
		57359: 5,  // END (43x)
		57355: 6,  // HTTP (43x)
		57356: 7,  // MUST (43x)
		57354: 8,  // SET (43x)
		57357: 9,  // SHOULD (43x)
		57380: 10, // WHEN (38x)
		57383: 11, // INTO (37x)
		57350: 12, // FLOAT (32x)
		57349: 13, // INTEGER (32x)
		57381: 14, // AND (23x)
		57382: 15, // OR (23x)
		57404: 16, // variable (17x)
		57403: 17, // string_or_var (15x)
		57387: 18, // any_value (14x)
		41:    19, // ')' (13x)
		40:    20, // '(' (12x)
		57352: 21, // FALSE (12x)
		57351: 22, // TRUE (12x)
		57389: 23, // boolean_exp (11x)
		57370: 24, // HEADER (9x)
		57371: 25, // QUERY (9x)
		57378: 26, // CONTAINS (8x)
		57372: 27, // EQUALS (8x)
		57375: 28, // GE (8x)
		57374: 29, // GT (8x)
		57377: 30, // LE (8x)
		57376: 31, // LT (8x)
		57373: 32, // NOTEQUALS (8x)
		57379: 33, // STARTSWITH (8x)
		125:   34, // '}' (2x)
		57395: 35, // http_command_param (2x)
		57400: 36, // operator (2x)
		57386: 37, // any_command (1x)
		57388: 38, // assert_command (1x)
		57390: 39, // command (1x)
		57391: 40, // command_with_condition_opt (1x)
		57366: 41, // CONNECT (1x)
		57392: 42, // debug_command (1x)
		57365: 43, // DELETE (1x)
		57393: 44, // end_command (1x)
		57361: 45, // GET (1x)
		57362: 46, // HEAD (1x)
		57394: 47, // http_command (1x)
		57396: 48, // http_command_params (1x)
		57397: 49, // http_method (1x)
		57384: 50, // IDENTIFIER (1x)
		57398: 51, // multi_any_value (1x)
		57399: 52, // must_command (1x)
		57367: 53, // OPTIONS (1x)
		57369: 54, // PATCH (1x)
		57363: 55, // POST (1x)
		57364: 56, // PUT (1x)
		57401: 57, // set_command (1x)
		57402: 58, // should_command (1x)
		57368: 59, // TRACE (1x)
		57385: 60, // $default (0x)
		57347: 61, // EOF (0x)
		57346: 62, // EOL (0x)
		57345: 63, // error (0x)
		57353: 64, // KEYWORD (0x)
	}

	yySymNames = []string{
		"'{'",
		"STRING",
		"$end",
		"ASSERT",
		"DEBUG",
		"END",
		"HTTP",
		"MUST",
		"SET",
		"SHOULD",
		"WHEN",
		"INTO",
		"FLOAT",
		"INTEGER",
		"AND",
		"OR",
		"variable",
		"string_or_var",
		"any_value",
		"')'",
		"'('",
		"FALSE",
		"TRUE",
		"boolean_exp",
		"HEADER",
		"QUERY",
		"CONTAINS",
		"EQUALS",
		"GE",
		"GT",
		"LE",
		"LT",
		"NOTEQUALS",
		"STARTSWITH",
		"'}'",
		"http_command_param",
		"operator",
		"any_command",
		"assert_command",
		"command",
		"command_with_condition_opt",
		"CONNECT",
		"debug_command",
		"DELETE",
		"end_command",
		"GET",
		"HEAD",
		"http_command",
		"http_command_params",
		"http_method",
		"IDENTIFIER",
		"multi_any_value",
		"must_command",
		"OPTIONS",
		"PATCH",
		"POST",
		"PUT",
		"set_command",
		"should_command",
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
		1:  {37, 0},
		2:  {37, 2},
		3:  {40, 5},
		4:  {40, 3},
		5:  {40, 3},
		6:  {40, 1},
		7:  {39, 1},
		8:  {39, 1},
		9:  {39, 1},
		10: {39, 1},
		11: {39, 1},
		12: {39, 1},
		13: {39, 1},
		14: {42, 2},
		15: {44, 3},
		16: {44, 2},
		17: {44, 1},
		18: {38, 2},
		19: {52, 2},
		20: {58, 2},
		21: {57, 3},
		22: {57, 3},
		23: {47, 4},
		24: {47, 3},
		25: {48, 1},
		26: {48, 2},
		27: {35, 2},
		28: {35, 2},
		29: {49, 1},
		30: {49, 1},
		31: {49, 1},
		32: {49, 1},
		33: {49, 1},
		34: {49, 1},
		35: {49, 1},
		36: {49, 1},
		37: {49, 1},
		38: {51, 1},
		39: {51, 2},
		40: {18, 1},
		41: {18, 1},
		42: {18, 1},
		43: {17, 1},
		44: {17, 1},
		45: {16, 5},
		46: {36, 1},
		47: {36, 1},
		48: {36, 1},
		49: {36, 1},
		50: {36, 1},
		51: {36, 1},
		52: {36, 1},
		53: {36, 1},
		54: {36, 1},
		55: {36, 1},
		56: {23, 1},
		57: {23, 1},
		58: {23, 3},
		59: {23, 3},
		60: {23, 3},
		61: {23, 3},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [86][]uint16{
		// 0
		{2: 61, 61, 61, 61, 61, 61, 61, 61, 37: 63},
		{2: 62, 75, 73, 74, 79, 76, 78, 77, 38: 70, 65, 64, 42: 68, 44: 69, 47: 67, 52: 71, 57: 66, 72},
		{2: 60, 60, 60, 60, 60, 60, 60, 60},
		{2: 56, 56, 56, 56, 56, 56, 56, 56, 143, 142},
		{2: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55},
		// 5
		{2: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54},
		{2: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53},
		{2: 52, 52, 52, 52, 52, 52, 52, 52, 52, 52},
		{2: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51},
		{2: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50},
		// 10
		{2: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49},
		{93, 92, 12: 109, 110, 16: 91, 108, 140, 51: 139},
		{93, 92, 45, 45, 45, 45, 45, 45, 45, 45, 136, 45, 109, 110, 16: 91, 108, 115, 20: 113, 112, 111, 137},
		{93, 92, 12: 109, 110, 16: 91, 108, 115, 20: 113, 112, 111, 135},
		{93, 92, 12: 109, 110, 16: 91, 108, 115, 20: 113, 112, 111, 134},
		// 15
		{93, 92, 12: 109, 110, 16: 91, 108, 115, 20: 113, 112, 111, 133},
		{93, 16: 105},
		{41: 86, 43: 85, 45: 81, 82, 49: 80, 53: 87, 89, 83, 84, 59: 88},
		{93, 92, 16: 91, 90},
		{33, 33},
		// 20
		{32, 32},
		{31, 31},
		{30, 30},
		{29, 29},
		{28, 28},
		// 25
		{27, 27},
		{26, 26},
		{25, 25},
		{2: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 24: 100, 101, 35: 99, 48: 98},
		{19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19: 19, 24: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		// 30
		{18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 19: 18, 24: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{94},
		{50: 95},
		{34: 96},
		{34: 97},
		// 35
		{17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 19: 17, 17, 17, 17, 24: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{2: 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 24: 100, 101, 35: 104},
		{2: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 24: 37, 37},
		{1: 103},
		{1: 102},
		// 40
		{2: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 24: 34, 34},
		{2: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 24: 35, 35},
		{2: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 24: 36, 36},
		{93, 92, 12: 109, 110, 16: 91, 108, 106, 20: 113, 112, 111, 107},
		{2: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 14: 124, 125, 26: 122, 116, 119, 118, 121, 120, 117, 123, 36: 126},
		// 45
		{2: 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 14: 128, 129},
		{22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 19: 22, 26: 22, 22, 22, 22, 22, 22, 22, 22},
		{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 19: 21, 26: 21, 21, 21, 21, 21, 21, 21, 21},
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 19: 20, 26: 20, 20, 20, 20, 20, 20, 20, 20},
		{2: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 14: 6, 6, 19: 6},
		// 50
		{2: 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 14: 5, 5, 19: 5},
		{93, 92, 12: 109, 110, 16: 91, 108, 115, 20: 113, 112, 111, 114},
		{14: 128, 129, 19: 130},
		{14: 124, 125, 26: 122, 116, 119, 118, 121, 120, 117, 123, 36: 126},
		{16, 16, 12: 16, 16},
		// 55
		{15, 15, 12: 15, 15},
		{14, 14, 12: 14, 14},
		{13, 13, 12: 13, 13},
		{12, 12, 12: 12, 12},
		{11, 11, 12: 11, 11},
		// 60
		{10, 10, 12: 10, 10},
		{9, 9, 12: 9, 9},
		{8, 8, 12: 8, 8},
		{7, 7, 12: 7, 7},
		{93, 92, 12: 109, 110, 16: 91, 108, 127},
		// 65
		{2: 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 14: 1, 1, 19: 1},
		{93, 92, 12: 109, 110, 16: 91, 108, 115, 20: 113, 112, 111, 132},
		{93, 92, 12: 109, 110, 16: 91, 108, 115, 20: 113, 112, 111, 131},
		{2: 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 14: 2, 2, 19: 2},
		{2: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 14: 128, 129, 19: 3},
		// 70
		{2: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 14: 128, 129, 19: 4},
		{2: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 14: 128, 129},
		{2: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 14: 128, 129},
		{2: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 14: 128, 129},
		{93, 92, 12: 109, 110, 16: 91, 108, 115, 20: 113, 112, 111, 138},
		// 75
		{2: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 14: 128, 129},
		{2: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 14: 128, 129},
		{93, 92, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 109, 110, 16: 91, 108, 141},
		{24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24},
		{23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		// 80
		{93, 16: 145},
		{93, 92, 12: 109, 110, 16: 91, 108, 115, 20: 113, 112, 111, 144},
		{2: 57, 57, 57, 57, 57, 57, 57, 57, 14: 128, 129},
		{2: 58, 58, 58, 58, 58, 58, 58, 58, 146},
		{93, 92, 12: 109, 110, 16: 91, 108, 115, 20: 113, 112, 111, 147},
		// 85
		{2: 59, 59, 59, 59, 59, 59, 59, 59, 14: 128, 129},
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
	const yyError = 63

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
				GlobalVars[yyS[yypt-2].variable.name] = yyS[yypt-4].cmd.Run()
			}
		}
	case 4:
		{
			//command INTO variable
			if strings.Contains(yyS[yypt-0].variable.name, ".") {
				yylex.Error("nested variables are not supported yet")
			}
			GlobalVars[yyS[yypt-0].variable.name] = yyS[yypt-2].cmd.Run()
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
			yyVAL.cmd = &DebugCommand{
				Values: yyS[yypt-0].vals,
			}
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
			return -1
		}
	case 18:
		{
			yyVAL.cmd = &AssertCommand{}
		}
	case 19:
		{
			//if $2 is false, fail
			yyVAL.cmd = &MustCommand{}
		}
	case 20:
		{
			//if $2 is false, write a warning
			yyVAL.cmd = &ShouldCommand{}
		}
	case 21:
		{
			//GlobalVars[$2.name] = $3
			yyVAL.cmd = &SetCommand{
				Name:  yyS[yypt-1].variable.name,
				Value: yyS[yypt-0].val,
			}
		}
	case 22:
		{
			//GlobalVars[$2.name] = $3
			yyVAL.cmd = &SetCommand{
				Name:  yyS[yypt-1].variable.name,
				Value: yyS[yypt-0].boolean,
			}
		}
	case 23:
		{
			//call http with header here.
			yyVAL.cmd = &HttpCommand{
				Method:        yyS[yypt-2].val.(string),
				CommandParams: yyS[yypt-0].http_command_params,
				Url:           yyS[yypt-1].val.(string),
			}
		}
	case 24:
		{
			//simple http command
			yyVAL.cmd = &HttpCommand{
				Method: yyS[yypt-1].val.(string),
				Url:    yyS[yypt-0].val.(string),
			}
		}
	case 25:
		{
			if yyVAL.http_command_params == nil {
				yyVAL.http_command_params = make([]HttpCommandParam, 0)
			}
			yyVAL.http_command_params = append(yyVAL.http_command_params, yyS[yypt-0].http_command_param)
		}
	case 26:
		{
			if yyVAL.http_command_params == nil {
				yyVAL.http_command_params = make([]HttpCommandParam, 0)
			}

			yyVAL.http_command_params = append(yyVAL.http_command_params, yyS[yypt-0].http_command_param)
		}
	case 27:
		{
			//addin header
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-1].val.(string),
				ParamValue: yyS[yypt-0].val.(string),
			}
		}
	case 28:
		{
			//adding query param
			yyVAL.http_command_param = HttpCommandParam{
				ParamName:  yyS[yypt-1].val.(string),
				ParamValue: yyS[yypt-0].val.(string),
			}
		}
	case 29:
		{
			//http get
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
			//getting a single value from multi_value exp
			yyVAL.vals = append(yyVAL.vals, yyS[yypt-0].val)
		}
	case 39:
		{
			//multi value
			yyVAL.vals = append(yyVAL.vals, yyS[yypt-0].val)
		}
	case 40:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 41:
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 42:
		{
			yyVAL.val, _ = strconv.Atoi(yyS[yypt-0].val.(string))
		}
	case 43:
		{
			//found variable
			switch yyS[yypt-0].variable.value.(type) {
			case string:
				if isTemplate(yyS[yypt-0].variable.value.(string)) {
					yyVAL.val, _ = executeTemplate(yyS[yypt-0].variable.value.(string), GlobalVars)
				} else {
					yyVAL.val = yyS[yypt-0].variable.value
				}
			default:
				yyVAL.val = yyS[yypt-0].variable.value
			}

		}
	case 44:
		{
			//found string
			if isTemplate(yyS[yypt-0].val.(string)) {
				yyVAL.val, _ = executeTemplate(yyS[yypt-0].val.(string), GlobalVars)
			} else {

			}
		}
	case 45:
		{
			//getting variable
			yyVAL.variable.name = yyS[yypt-2].val.(string)
			yyVAL.variable.value = query(yyS[yypt-2].val.(string), GlobalVars)
		}
	case 56:
		{
			yyVAL.boolean = true
		}
	case 57:
		{
			yyVAL.boolean = false
		}
	case 58:
		{
			yyVAL.boolean = IsTrue(yyS[yypt-2].boolean) && IsTrue(yyS[yypt-0].boolean)
		}
	case 59:
		{
			yyVAL.boolean = IsTrue(yyS[yypt-2].boolean) || IsTrue(yyS[yypt-0].boolean)
		}
	case 60:
		{
			yyVAL.boolean = yyS[yypt-1].boolean
		}
	case 61:
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

func Parse(text string) *lex {
	s := NewScanner(strings.NewReader(strings.TrimSpace(text)))
	tokens := make(Tokens, 0)

	for {
		token := s.Scan()
		//fmt.Println(token)
		if token.Type == EOF || token.Type == -1 {
			break
		}
		tokens = append(tokens, token)
	}

	l := &lex{tokens}

	return l
}

func Reset() {
	GlobalVars = map[string]interface{}{}
}

func Run(l *lex) {
	yyParse(l)
}
