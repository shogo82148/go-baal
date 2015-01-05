
//line parser.go.y:2
package baal
import __yyfmt__ "fmt"
//line parser.go.y:2
		

//line parser.go.y:10
type yySymType struct{
	yys int
	namespaces    []Namespace
	namespace     Namespace
	tok           Token
	name          Name
	names         []Name
        declarations  []Declaration
        declaration   Declaration
	document      string
	bool          bool
	occurrence    int
	fields        []Field
	field         Field
	modified_type ModifiedType
	baal_type     Type
	iterations    []int
}

const IDENTIFIER = 57346
const QUOTED_STRING = 57347
const DOCUMENT_COMMENT = 57348
const NAMESPACE = 57349
const IMPORT = 57350
const ABSTRACT = 57351
const ENTITY = 57352
const SERVICE = 57353
const INCLUDES = 57354
const REQUIRED = 57355
const NULLABLE = 57356
const LIST_OF = 57357
const DICTIONARY_OF = 57358
const APPEND = 57359
const REMOVE = 57360
const ACCEPTS = 57361
const RETURNS = 57362
const AS = 57363
const BEGIN = 57364
const END = 57365
const CLOSE = 57366
const BOOLEAN = 57367
const INTEGER8 = 57368
const INTEGER16 = 57369
const INTEGER32 = 57370
const INTEGER64 = 57371
const FLOAT32 = 57372
const FLOAT64 = 57373
const DECIMAL32 = 57374
const DECIMAL64 = 57375
const DATE = 57376
const TIME = 57377
const TIMESTAMP = 57378
const STRING = 57379
const BINARY = 57380

var yyToknames = []string{
	"IDENTIFIER",
	"QUOTED_STRING",
	"DOCUMENT_COMMENT",
	"NAMESPACE",
	"IMPORT",
	"ABSTRACT",
	"ENTITY",
	"SERVICE",
	"INCLUDES",
	"REQUIRED",
	"NULLABLE",
	"LIST_OF",
	"DICTIONARY_OF",
	"APPEND",
	"REMOVE",
	"ACCEPTS",
	"RETURNS",
	"AS",
	"BEGIN",
	"END",
	"CLOSE",
	"BOOLEAN",
	"INTEGER8",
	"INTEGER16",
	"INTEGER32",
	"INTEGER64",
	"FLOAT32",
	"FLOAT64",
	"DECIMAL32",
	"DECIMAL64",
	"DATE",
	"TIME",
	"TIMESTAMP",
	"STRING",
	"BINARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:233


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	22, 11,
	-2, 4,
	-1, 34,
	22, 13,
	-2, 44,
}

const yyNprod = 46
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 67

var yyAct = []int{

	48, 49, 12, 8, 22, 20, 7, 45, 21, 21,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 15, 36, 17, 10, 5, 30,
	40, 10, 26, 16, 29, 32, 11, 24, 25, 38,
	31, 9, 43, 44, 14, 3, 39, 6, 27, 12,
	2, 46, 4, 50, 47, 41, 34, 35, 42, 23,
	37, 33, 28, 19, 18, 13, 1,
}
var yyPact = []int{

	38, 38, -1000, 43, -1000, -36, -1000, 19, 45, -1000,
	43, 43, -1000, 3, -1000, -35, -1000, -1000, -1000, -1000,
	28, -1000, -2, 22, -1000, -1000, 44, -1000, 7, 23,
	-1000, 43, 43, 2, -35, -36, -1000, -1000, 42, 9,
	29, -17, -1000, -1000, -1000, -1000, -15, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 66, 50, 6, 65, 64, 63, 62, 61, 60,
	5, 59, 58, 55, 54, 53, 24, 33, 51,
}
var yyR1 = []int{

	0, 1, 1, 2, 3, 3, 3, 4, 4, 5,
	6, 7, 7, 7, 8, 8, 9, 10, 10, 11,
	11, 13, 18, 18, 18, 12, 12, 14, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 16, 16, 17, 17,
}
var yyR2 = []int{

	0, 1, 2, 6, 0, 3, 3, 0, 2, 1,
	8, 0, 3, 3, 0, 2, 5, 0, 1, 0,
	1, 3, 0, 2, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 7, -2, -16, 4, -3, 39, 22,
	8, 17, 4, -4, -17, -16, -17, 23, -5, -6,
	-10, 6, 39, -11, 9, 40, 10, 4, -7, -3,
	22, 17, 12, -8, -16, -16, 23, -9, -10, 4,
	21, -13, -12, 13, 14, 24, -18, -14, 15, 16,
	-15, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 38,
}
var yyDef = []int{

	0, -2, 1, 0, 2, 4, 42, 0, 0, 7,
	0, 0, 43, 17, 5, 44, 6, 3, 8, 9,
	19, 18, 0, 0, 20, 45, 0, -2, 0, 0,
	14, 0, 0, 17, -2, 12, 10, 15, 0, 0,
	0, 0, 22, 25, 26, 16, 0, 21, 23, 24,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36,
	37, 38, 39, 40, 41,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 40, 3, 3, 3, 39,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:48
		{
			yyVAL.namespaces = []Namespace{yyS[yypt-0].namespace}
			yylex.(*Lexer).namespaces = yyVAL.namespaces
		}
	case 2:
		//line parser.go.y:53
		{
			yyVAL.namespaces = append(yyS[yypt-1].namespaces, yyS[yypt-0].namespace)
			yylex.(*Lexer).namespaces = yyVAL.namespaces
		}
	case 3:
		//line parser.go.y:60
		{
			yyVAL.namespace = Namespace{
				Name:         yyS[yypt-4].name,
				Imports:      yyS[yypt-3].names,
				Declarations: yyS[yypt-1].declarations,
			}
		}
	case 4:
		//line parser.go.y:70
		{
			yyVAL.names = []Name{}
		}
	case 5:
		//line parser.go.y:74
		{
			yyVAL.names = append(yyS[yypt-2].names, yyS[yypt-0].name)
		}
	case 6:
		//line parser.go.y:78
		{
			yyVAL.names = append(yyS[yypt-2].names, yyS[yypt-0].name)
		}
	case 7:
		//line parser.go.y:84
		{
			yyVAL.declarations = []Declaration{}
		}
	case 8:
		//line parser.go.y:88
		{
			yyVAL.declarations = append(yyS[yypt-1].declarations, yyS[yypt-0].declaration)
		}
	case 9:
		//line parser.go.y:94
		{
			yyVAL.declaration = yyS[yypt-0].declaration
		}
	case 10:
		//line parser.go.y:100
		{
			yyVAL.declaration = Entity{
				Document:   yyS[yypt-7].document,
				IsAbstract: yyS[yypt-6].bool,
				Name:       yyS[yypt-4].tok.Lit,
				Includes:   yyS[yypt-3].names,
				Fields:     yyS[yypt-1].fields,
			}
		}
	case 11:
		//line parser.go.y:112
		{
			yyVAL.names = []Name{}
		}
	case 12:
		//line parser.go.y:116
		{
			yyVAL.names = append(yyS[yypt-2].names, yyS[yypt-0].name)
		}
	case 13:
		//line parser.go.y:120
		{
			yyVAL.names = append(yyS[yypt-2].names, yyS[yypt-0].name)
		}
	case 14:
		//line parser.go.y:126
		{
			yyVAL.fields = []Field{}
		}
	case 15:
		//line parser.go.y:130
		{
			yyVAL.fields = append(yyS[yypt-1].fields, yyS[yypt-0].field)
		}
	case 16:
		//line parser.go.y:136
		{
			yyVAL.field = Field{
		        	Document:     yyS[yypt-4].document,
				Name:         yyS[yypt-3].tok.Lit,
				ModifiedType: yyS[yypt-1].modified_type,
			}
		}
	case 17:
		//line parser.go.y:146
		{
			yyVAL.document = ""
		}
	case 18:
		//line parser.go.y:150
		{
			yyVAL.document = yyS[yypt-0].tok.Lit
		}
	case 19:
		//line parser.go.y:156
		{
			yyVAL.bool = false
		}
	case 20:
		//line parser.go.y:160
		{
			yyVAL.bool = true
		}
	case 21:
		//line parser.go.y:166
		{
			yyVAL.modified_type = ModifiedType{
				Occurrence: yyS[yypt-2].occurrence,
				Iterations: yyS[yypt-1].iterations,
				Type:       yyS[yypt-0].baal_type,
			}
		}
	case 22:
		//line parser.go.y:176
		{
			yyVAL.iterations = []int{}
		}
	case 23:
		//line parser.go.y:180
		{
			yyVAL.iterations = append(yyS[yypt-1].iterations, LIST_OF)
		}
	case 24:
		//line parser.go.y:184
		{
			yyVAL.iterations = append(yyS[yypt-1].iterations, DICTIONARY_OF)
		}
	case 25:
		//line parser.go.y:189
		{ yyVAL.occurrence = REQUIRED }
	case 26:
		//line parser.go.y:190
		{ yyVAL.occurrence = NULLABLE }
	case 27:
		//line parser.go.y:194
		{
			yyVAL.baal_type = yyS[yypt-0].baal_type
		}
	case 28:
		//line parser.go.y:199
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 29:
		//line parser.go.y:200
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 30:
		//line parser.go.y:201
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 31:
		//line parser.go.y:202
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 32:
		//line parser.go.y:203
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 33:
		//line parser.go.y:204
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 34:
		//line parser.go.y:205
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 35:
		//line parser.go.y:206
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 36:
		//line parser.go.y:207
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 37:
		//line parser.go.y:208
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 38:
		//line parser.go.y:209
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 39:
		//line parser.go.y:210
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 40:
		//line parser.go.y:211
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 41:
		//line parser.go.y:212
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 42:
		//line parser.go.y:216
		{
			yyVAL.name = []string{yyS[yypt-0].tok.Lit}
		}
	case 43:
		//line parser.go.y:220
		{
			yyVAL.name = append(yyS[yypt-2].name, yyS[yypt-0].tok.Lit)
		}
	case 44:
		//line parser.go.y:226
		{
			yyVAL.name = yyS[yypt-0].name
		}
	case 45:
		//line parser.go.y:230
		{
			yyVAL.name = append(yyS[yypt-2].name, "*")
		}
	}
	goto yystack /* stack new state and value */
}
