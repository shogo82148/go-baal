%{
package baal

%}

%token<tok> IDENTIFIER QUOTED_STRING DOCUMENT_COMMENT
%token NAMESPACE IMPORT ABSTRACT ENTITY SERVICE INCLUDES REQUIRED NULLABLE LIST_OF DICTIONARY_OF APPEND REMOVE ACCEPTS RETURNS AS BEGIN END CLOSE
%token<tok> BOOLEAN INTEGER8 INTEGER16 INTEGER32 INTEGER64 FLOAT32 FLOAT64 DECIMAL32 DECIMAL64 DATE TIME TIMESTAMP STRING BINARY

%union{
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
%type<namespaces>namespaces
%type<namespace>namespace
%type<names>imports
%type<declarations>declarations
%type<declaration>declaration
%type<declaration>entity
%type<names>includes
%type<fields>fields
%type<field>field
%type<methods>methods
%type<method>method
%type<document>document
%type<bool>abstract
%type<occurrence>occurrence
%type<modified_type>modified_type request response
%type<baal_type>baal_type primitive_type reference_type pseudo_type
%type<name>qualified_name qualified_name_with_wildcard
%type<iterations>iterations

%%

namespaces
	: namespace
	{
		$$ = []Namespace{$1}
		yylex.(*Lexer).namespaces = $$
	}
	| namespaces namespace
	{
		$$ = append($1, $2)
		yylex.(*Lexer).namespaces = $$
	}

namespace
	: NAMESPACE qualified_name imports BEGIN declarations END
	{
		$$ = Namespace{
			Name:         $2,
			Imports:      $3,
			Declarations: $5,
		}
	}

imports
	:
	{
		$$ = []Name{}
	}
	| imports IMPORT qualified_name_with_wildcard
	{
		$$ = append($1, $3)
	}
	| imports APPEND qualified_name_with_wildcard
	{
		$$ = append($1, $3)
	}

declarations
	:
	{
		$$ = []Declaration{}
	}
	| declarations declaration
	{
		$$ = append($1, $2)
	}

declaration
	: entity
	{
		$$ = $1
	}

entity
	: document abstract ENTITY IDENTIFIER includes BEGIN fields END
	{
		$$ = Entity{
			Document:   $1,
			IsAbstract: $2,
			Name:       $4.Lit,
			Includes:   $5,
			Fields:     $7,
		}
	}
	| document SERVICE IDENTIFIER BEGIN methods END
	{
		$$ = Service{
			Document: $1,
			Name:     $3.Lit,
			Methods:  $5,
		}
	}

includes
	:
	{
		$$ = []Name{}
	}
	| imports INCLUDES qualified_name
	{
		$$ = append($1, $3)
	}
	| imports APPEND qualified_name
	{
		$$ = append($1, $3)
	}

fields
	:
	{
		$$ = []Field{}
	}
	| fields field
	{
		$$ = append($1, $2)
	}

field
	: document IDENTIFIER AS modified_type CLOSE
	{
		$$ = Field{
	        	Document:     $1,
			Name:         $2.Lit,
			ModifiedType: $4,
		}
	}

methods
	:
	{
		$$ = []Method{}
	}
	| methods method
	{
		$$ = append($1, $2)
	}

method
	: document IDENTIFIER AS request response CLOSE
	{
		$$ = Method{
			Document: $1,
			Name:     $2.Lit,
                        Request:  $4,
                        Response: $5,
		}
	}

request
	:
	{
		$$ = ModifiedType{}
	}
	| ACCEPTS modified_type
	{
		$$ = $2
	}

response
	:
	{
		$$ = ModifiedType{}
	}
	| RETURNS modified_type
	{
		$$ = $2
	}

document
	:
	{
		$$ = ""
	}
	| DOCUMENT_COMMENT
	{
		$$ = $1.Lit
	}

abstract
	:
	{
		$$ = false
	}
	| ABSTRACT
	{
		$$ = true
	}

modified_type
	: occurrence iterations baal_type
	{
		$$ = ModifiedType{
			Occurrence: $1,
			Iterations: $2,
			Type:       $3,
		}
	}

iterations
	:
	{
		$$ = []int{}
	}
	| iterations LIST_OF
	{
		$$ = append($1, LIST_OF)
	}
	| iterations DICTIONARY_OF
	{
		$$ = append($1, DICTIONARY_OF)
	}

occurrence
        : REQUIRED { $$ = REQUIRED }
	| NULLABLE { $$ = NULLABLE }

baal_type
	: primitive_type
	{
		$$ = $1
	}
	| reference_type
	{
		$$ = $1
	}
	| pseudo_type
	{
		$$ = $1
	}

primitive_type
	: BOOLEAN   { $$ = PrimitiveType{ Tok: $1.Tok } }
	| INTEGER8  { $$ = PrimitiveType{ Tok: $1.Tok } }
	| INTEGER16 { $$ = PrimitiveType{ Tok: $1.Tok } }
	| INTEGER32 { $$ = PrimitiveType{ Tok: $1.Tok } }
	| INTEGER64 { $$ = PrimitiveType{ Tok: $1.Tok } }
	| FLOAT32   { $$ = PrimitiveType{ Tok: $1.Tok } }
	| FLOAT64   { $$ = PrimitiveType{ Tok: $1.Tok } }
	| DECIMAL32 { $$ = PrimitiveType{ Tok: $1.Tok } }
	| DECIMAL64 { $$ = PrimitiveType{ Tok: $1.Tok } }
	| DATE      { $$ = PrimitiveType{ Tok: $1.Tok } }
	| TIME      { $$ = PrimitiveType{ Tok: $1.Tok } }
	| TIMESTAMP { $$ = PrimitiveType{ Tok: $1.Tok } }
	| STRING    { $$ = PrimitiveType{ Tok: $1.Tok } }
	| BINARY    { $$ = PrimitiveType{ Tok: $1.Tok } }

reference_type
	: qualified_name
	{
		$$ = ReferenceType{
			Name: $1,
		}
	}

pseudo_type
	: QUOTED_STRING
	{
		$$ = PseudoType{
			Name: $1.Lit,
		}
	}

qualified_name
	: IDENTIFIER
	{
		$$ = []string{$1.Lit}
	}
	| qualified_name '.' IDENTIFIER
	{
		$$ = append($1, $3.Lit)
	}

qualified_name_with_wildcard
	: qualified_name
	{
		$$ = $1
	}
	| qualified_name '.' '*'
	{
		$$ = append($1, "*")
	}
%%
