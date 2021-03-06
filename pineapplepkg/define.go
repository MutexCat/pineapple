package pineapplepkg

type Variable struct {
	lineNum    int
	isStrValue bool
	name       string
}

type Assignment struct {
	lineNum  int
	variable *Variable
	value    string
}

type Print struct {
	lineNum  int
	variable *Variable
}

const (
	MATH_ADD = iota
	MATH_SUB
	MATH_MUTL
	MATH_DIV
)

//TODO
type MathOperation struct {
	which   int
	lineNUm int
	lhs     *Variable
	rhs     *Variable
}

type Statement interface{}

var _ Statement = (*Print)(nil)
var _ Statement = (*Assignment)(nil)
var _ Statement = (*MathOperation)(nil)

type SourceCode struct {
	LineNum    int
	Statements []Statement
}

func NewVariablePair(isStr bool, value string) *VariablePair {
	return &VariablePair{
		isStrValue: isStr,
		value:      value,
	}
}

type VariablePair struct {
	isStrValue bool
	value      string
}

type GlobalVariables struct {
	Variables map[string]*VariablePair
}

func NewGlobalVariables() *GlobalVariables {
	var g GlobalVariables
	g.Variables = make(map[string]*VariablePair)
	return &g
}
