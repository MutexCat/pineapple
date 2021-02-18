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
	} else if add, ok := statement.(*Add); ok {
		return resolveAdd(g, add)
	} else {
		return errors.New("resolveStatement(): undefined statement type.")
	}

}

func resolveAssignment(g *GlobalVariables, assignment *Assignment) error {
	varName := ""
	if varName = assignment.variable.name; varName == "" {
		return errors.New("resolveAssignment(): variable name can NOT be empty.")
	}
	varPair := NewVariablePair(assignment.variable.isStrValue, assignment.value)
	g.Variables[varName] = varPair
	return nil
}

func resolveAdd(g *GlobalVariables, add *Add) error {
	lhs := g.Variables[add.lhs.name]
	rhs := g.Variables[add.rhs.name]
	if lhs.isStrValue || rhs.isStrValue || lhs.value == "" || rhs.value == "" {
		return errors.New("resolveAdd():add operation needs int value or variable must not be empty")
	}
	lhsValue, _ := strconv.Atoi(lhs.value)
	rhsValue, _ := strconv.Atoi(rhs.value)
	totalValue := lhsValue + rhsValue
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
