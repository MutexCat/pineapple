package pineapplepkg

import (
	"errors"
	"fmt"
	"strconv"
)

func Execute(code string) {
	var ast *SourceCode
	var err error

	g := NewGlobalVariables()

	// parse
	if ast, err = parse(code); err != nil {
		panic(err)
	}

	// resolve
	if err = resolveAST(g, ast); err != nil {
		panic(err)
	}
}

func resolveAST(g *GlobalVariables, ast *SourceCode) error {
	if len(ast.Statements) == 0 {
		return errors.New("resolveAST(): no code to execute, please check your input.")
	}
	for _, statement := range ast.Statements {
		if err := resolveStatement(g, statement); err != nil {
			return err
		}
	}
	return nil
}

func resolveStatement(g *GlobalVariables, statement Statement) error {
	if assignment, ok := statement.(*Assignment); ok {
		return resolveAssignment(g, assignment)
	} else if print, ok := statement.(*Print); ok {
		return resolvePrint(g, print)
	} else if operator, ok := statement.(*MathOperation); ok {
		return resolveMath(g, operator)
	} else {
		return errors.New("resolveStatement(): undefined statement type.")
	}

}

//TODO
func resolveAssignment(g *GlobalVariables, assignment *Assignment) error {
	varName := ""
	if varName = assignment.variable.name; varName == "" {
		return errors.New("resolveAssignment(): variable name can NOT be empty.")
	}
	varPair := NewVariablePair(assignment.variable.isStrValue, assignment.value)
	g.Variables[varName] = varPair
	return nil
}

//TODO
func resolveMath(g *GlobalVariables, operator *MathOperation) error {
	lhs := g.Variables[operator.lhs.name]
	rhs := g.Variables[operator.rhs.name]
	if lhs.isStrValue || rhs.isStrValue || lhs.value == "" || rhs.value == "" {
		return errors.New("resolveAdd():add operation needs int value or variable must not be empty")
	}
	lhsValue, _ := strconv.Atoi(lhs.value)
	rhsValue, _ := strconv.Atoi(rhs.value)
	resultFunc := func(lvalue, rvalue int, operator *MathOperation) int {
		switch operator.which {
		case MATH_ADD:
			return lvalue + rvalue
		case MATH_SUB:
			return lvalue - rvalue
		case MATH_MUTL:
			return lvalue * rvalue
		case MATH_DIV:
			return lvalue / rvalue
		default:
			return 0
		}
	}
	totalValue := resultFunc(lhsValue, rhsValue, operator)
	totalValueStr := strconv.Itoa(totalValue)
	lhs.value = totalValueStr
	return nil
}

func resolvePrint(g *GlobalVariables, print *Print) error {
	varName := ""
	if varName = print.variable.name; varName == "" {
		return errors.New("resolvePrint(): variable name can NOT be empty")
	}

	valuePair := NewVariablePair(true, "")
	ok := false
	if valuePair, ok = g.Variables[varName]; !ok {
		return errors.New(fmt.Sprintf("resolvePrint(): variable '$%s'not found.", varName))
	}

	if valuePair.isStrValue {
		fmt.Println(valuePair.value)
	} else {
		if value, err := strconv.Atoi(valuePair.value); err != err {
			return err
		} else {
			fmt.Println(value)
		}
	}

	return nil
}
