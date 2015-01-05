package baal

type Position struct {
	line   int
	column int
}

type Token struct {
	Pos Position
	Tok int
	Lit string
}

type Name []string

type Namespace struct {
	Name         Name
	Imports      []Name
	Declarations []Declaration
}

type Declaration interface{}

type Entity struct {
	Document   string
	IsAbstract bool
	Name       string
	Includes   []Name
	Fields     []Field
}

type Field struct {
	Document     string
	Name         string
	ModifiedType ModifiedType
}

type ModifiedType struct {
	Occurrence int
	Iterations []int
	Type       Type
}

type Type interface{}

type PrimitiveType struct {
	Tok int
}

type Service struct {
}
