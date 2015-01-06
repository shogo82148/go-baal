
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
	methods       []Method
	method        Method
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

//line parser.go.y:298


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 29,
	22, 12,
	-2, 4,
	-1, 41,
	22, 14,
	-2, 54,
}

const yyNprod = 56
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 87

var yyAct = []int{

	15, 54, 8, 22, 5, 7, 62, 59, 21, 21,
	21, 34, 30, 51, 47, 53, 50, 6, 20, 27,
	12, 56, 57, 16, 3, 44, 37, 17, 64, 65,
	25, 48, 24, 43, 14, 32, 41, 42, 68, 69,
	70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
	80, 81, 39, 58, 10, 60, 26, 10, 6, 46,
	29, 36, 82, 11, 28, 12, 35, 2, 9, 4,
	61, 67, 66, 63, 52, 49, 55, 23, 38, 33,
	45, 40, 31, 19, 18, 13, 1,
}
var yyPact = []int{

	17, 17, -1000, 54, -1000, -37, -1000, 46, 61, -1000,
	54, 54, -1000, 4, -1000, -36, -1000, -1000, -1000, -1000,
	21, -1000, 16, 9, 60, -1000, -1000, 56, -10, -1000,
	-1000, -11, 49, 3, -1000, 54, 54, -1000, -1000, 29,
	2, -36, -37, -7, -1000, -1000, 27, -3, -8, -5,
	8, 8, -17, 8, -1000, -1000, -1000, -1000, -18, -1000,
	-1000, 13, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -37,
}
var yyPgo = []int{

	0, 86, 67, 5, 85, 84, 83, 82, 81, 80,
	79, 78, 18, 77, 76, 1, 75, 74, 73, 72,
	71, 0, 23, 70,
}
var yyR1 = []int{

	0, 1, 1, 2, 3, 3, 3, 4, 4, 5,
	6, 6, 7, 7, 7, 8, 8, 9, 10, 10,
	11, 16, 16, 17, 17, 12, 12, 13, 13, 15,
	23, 23, 23, 14, 14, 18, 18, 19, 19, 19,
	19, 19, 19, 19, 19, 19, 19, 19, 19, 19,
	19, 20, 21, 21, 22, 22,
}
var yyR2 = []int{

	0, 1, 2, 6, 0, 3, 3, 0, 2, 1,
	8, 6, 0, 3, 3, 0, 2, 5, 0, 2,
	6, 0, 2, 0, 2, 0, 1, 0, 1, 3,
	0, 2, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 7, -2, -21, 4, -3, 39, 22,
	8, 17, 4, -4, -22, -21, -22, 23, -5, -6,
	-12, 6, 39, -13, 11, 9, 40, 10, 4, 4,
	22, -7, -3, -10, 22, 17, 12, 23, -11, -12,
	-8, -21, -21, 4, 23, -9, -12, 21, 4, -16,
	19, 21, -17, 20, -15, -14, 13, 14, -15, 24,
	-15, -23, 24, -18, 15, 16, -19, -20, 25, 26,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36,
	37, 38, -21,
}
var yyDef = []int{

	0, -2, 1, 0, 2, 4, 52, 0, 0, 7,
	0, 0, 53, 25, 5, 54, 6, 3, 8, 9,
	27, 26, 0, 0, 0, 28, 55, 0, 0, -2,
	18, 0, 0, 25, 15, 0, 0, 11, 19, 0,
	25, -2, 13, 0, 10, 16, 0, 21, 0, 23,
	0, 0, 0, 0, 22, 30, 33, 34, 0, 20,
	24, 0, 17, 29, 31, 32, 35, 36, 37, 38,
	39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51,
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
		//line parser.go.y:52
		{
			yyVAL.namespaces = []Namespace{yyS[yypt-0].namespace}
			yylex.(*Lexer).namespaces = yyVAL.namespaces
		}
	case 2:
		//line parser.go.y:57
		{
			yyVAL.namespaces = append(yyS[yypt-1].namespaces, yyS[yypt-0].namespace)
			yylex.(*Lexer).namespaces = yyVAL.namespaces
		}
	case 3:
		//line parser.go.y:64
		{
			yyVAL.namespace = Namespace{
				Name:         yyS[yypt-4].name,
				Imports:      yyS[yypt-3].names,
				Declarations: yyS[yypt-1].declarations,
			}
		}
	case 4:
		//line parser.go.y:74
		{
			yyVAL.names = []Name{}
		}
	case 5:
		//line parser.go.y:78
		{
			yyVAL.names = append(yyS[yypt-2].names, yyS[yypt-0].name)
		}
	case 6:
		//line parser.go.y:82
		{
			yyVAL.names = append(yyS[yypt-2].names, yyS[yypt-0].name)
		}
	case 7:
		//line parser.go.y:88
		{
			yyVAL.declarations = []Declaration{}
		}
	case 8:
		//line parser.go.y:92
		{
			yyVAL.declarations = append(yyS[yypt-1].declarations, yyS[yypt-0].declaration)
		}
	case 9:
		//line parser.go.y:98
		{
			yyVAL.declaration = yyS[yypt-0].declaration
		}
	case 10:
		//line parser.go.y:104
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
		//line parser.go.y:114
		{
			yyVAL.declaration = Service{
				Document: yyS[yypt-5].document,
				Name:     yyS[yypt-3].tok.Lit,
				Methods:  yyS[yypt-1].methods,
			}
		}
	case 12:
		//line parser.go.y:124
		{
			yyVAL.names = []Name{}
		}
	case 13:
		//line parser.go.y:128
		{
			yyVAL.names = append(yyS[yypt-2].names, yyS[yypt-0].name)
		}
	case 14:
		//line parser.go.y:132
		{
			yyVAL.names = append(yyS[yypt-2].names, yyS[yypt-0].name)
		}
	case 15:
		//line parser.go.y:138
		{
			yyVAL.fields = []Field{}
		}
	case 16:
		//line parser.go.y:142
		{
			yyVAL.fields = append(yyS[yypt-1].fields, yyS[yypt-0].field)
		}
	case 17:
		//line parser.go.y:148
		{
			yyVAL.field = Field{
		        	Document:     yyS[yypt-4].document,
				Name:         yyS[yypt-3].tok.Lit,
				ModifiedType: yyS[yypt-1].modified_type,
			}
		}
	case 18:
		//line parser.go.y:158
		{
			yyVAL.methods = []Method{}
		}
	case 19:
		//line parser.go.y:162
		{
			yyVAL.methods = append(yyS[yypt-1].methods, yyS[yypt-0].method)
		}
	case 20:
		//line parser.go.y:168
		{
			yyVAL.method = Method{
				Document: yyS[yypt-5].document,
				Name:     yyS[yypt-4].tok.Lit,
	                        Request:  yyS[yypt-2].modified_type,
	                        Response: yyS[yypt-1].modified_type,
			}
		}
	case 21:
		//line parser.go.y:179
		{
			yyVAL.modified_type = ModifiedType{}
		}
	case 22:
		//line parser.go.y:183
		{
			yyVAL.modified_type = yyS[yypt-0].modified_type
		}
	case 23:
		//line parser.go.y:189
		{
			yyVAL.modified_type = ModifiedType{}
		}
	case 24:
		//line parser.go.y:193
		{
			yyVAL.modified_type = yyS[yypt-0].modified_type
		}
	case 25:
		//line parser.go.y:199
		{
			yyVAL.document = ""
		}
	case 26:
		//line parser.go.y:203
		{
			yyVAL.document = yyS[yypt-0].tok.Lit
		}
	case 27:
		//line parser.go.y:209
		{
			yyVAL.bool = false
		}
	case 28:
		//line parser.go.y:213
		{
			yyVAL.bool = true
		}
	case 29:
		//line parser.go.y:219
		{
			yyVAL.modified_type = ModifiedType{
				Occurrence: yyS[yypt-2].occurrence,
				Iterations: yyS[yypt-1].iterations,
				Type:       yyS[yypt-0].baal_type,
			}
		}
	case 30:
		//line parser.go.y:229
		{
			yyVAL.iterations = []int{}
		}
	case 31:
		//line parser.go.y:233
		{
			yyVAL.iterations = append(yyS[yypt-1].iterations, LIST_OF)
		}
	case 32:
		//line parser.go.y:237
		{
			yyVAL.iterations = append(yyS[yypt-1].iterations, DICTIONARY_OF)
		}
	case 33:
		//line parser.go.y:242
		{ yyVAL.occurrence = REQUIRED }
	case 34:
		//line parser.go.y:243
		{ yyVAL.occurrence = NULLABLE }
	case 35:
		//line parser.go.y:247
		{
			yyVAL.baal_type = yyS[yypt-0].baal_type
		}
	case 36:
		//line parser.go.y:251
		{
			yyVAL.baal_type = yyS[yypt-0].baal_type
		}
	case 37:
		//line parser.go.y:256
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 38:
		//line parser.go.y:257
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 39:
		//line parser.go.y:258
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 40:
		//line parser.go.y:259
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 41:
		//line parser.go.y:260
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 42:
		//line parser.go.y:261
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 43:
		//line parser.go.y:262
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 44:
		//line parser.go.y:263
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 45:
		//line parser.go.y:264
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 46:
		//line parser.go.y:265
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 47:
		//line parser.go.y:266
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 48:
		//line parser.go.y:267
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 49:
		//line parser.go.y:268
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 50:
		//line parser.go.y:269
		{ yyVAL.baal_type = PrimitiveType{ Tok: yyS[yypt-0].tok.Tok } }
	case 51:
		//line parser.go.y:273
		{
			yyVAL.baal_type = ReferenceType{
				Name: yyS[yypt-0].name,
			}
		}
	case 52:
		//line parser.go.y:281
		{
			yyVAL.name = []string{yyS[yypt-0].tok.Lit}
		}
	case 53:
		//line parser.go.y:285
		{
			yyVAL.name = append(yyS[yypt-2].name, yyS[yypt-0].tok.Lit)
		}
	case 54:
		//line parser.go.y:291
		{
			yyVAL.name = yyS[yypt-0].name
		}
	case 55:
		//line parser.go.y:295
		{
			yyVAL.name = append(yyS[yypt-2].name, "*")
		}
	}
	goto yystack /* stack new state and value */
}
