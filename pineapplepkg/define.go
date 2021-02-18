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

/*
type AssignmentStr struct {
	lineNum  int
	variable *Variable
	value    string
}

type AssignmentInt struct {
	lineNum  int
	variable *Variable
	value    int
}
*/
type Print struct {
	lineNum  int
	variable *Variable
}

type Statement interface{}

var _ Statement = (*Print)(nil)
var _ Statement = (*Assignment)(nil)

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
