package pineapplepkg

type Variable struct {
	lineNum int
	name    string
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

type Statement interface{}

var _ Statement = (*Print)(nil)
var _ Statement = (*Assignment)(nil)

type SourceCode struct {
	LineNum    int
	Statements []Statement
}

type GlobalVariables struct {
	Variables map[string]string
}

func NewGlobalVariables() *GlobalVariables {
	var g GlobalVariables
	g.Variables = make(map[string]string)
	return &g
}
