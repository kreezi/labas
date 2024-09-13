package types

type VarType int

const (
	Integer VarType = iota
	Float
	Null
)

type Variable struct {
	Name  string
	Vtype VarType
	Value float64
}

func NewVariable(name string, vtype VarType, value float64) *Variable {
	return &Variable{name, vtype, value}
}

const (
	NUMBER   = "NUMBER"
	IDENT    = "IDENT"
	FUNCTION = "FUNCTION"
	OPERATOR = "OPERATOR"
	LPAREN   = "LPAREN"
	RPAREN   = "RPAREN"
)

type Node struct {
	Type  string
	Value string
	Args  []*Node
	Left  *Node
	Right *Node
}

type Function struct {
	Name string
	Arg  []string
	Body *Node
}

func NewFunction(name string, body *Node, arg []string) *Function {
	return &Function{name, arg, body}
}
