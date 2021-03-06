package pineapplepkg

import "errors"

func parseName(lexer *Lexer) (string, error) {
	_, name := lexer.NextTokenIs(TOKEN_NAME)
	return name, nil
}

func parseString(lexer *Lexer) (string, error) {
	str := ""
	switch lexer.LookAhead() {
	case TOKEN_DOUBLE_QUOTE:
		lexer.NextTokenIs(TOKEN_DOUBLE_QUOTE)
		lexer.LookAheadAndSkip(TOKEN_IGNORED)
		return str, nil
	case TOKEN_QUOTE:
		lexer.NextTokenIs(TOKEN_QUOTE)
		str = lexer.scanBeforeToken("\"")
		lexer.NextTokenIs(TOKEN_QUOTE)
		lexer.LookAheadAndSkip(TOKEN_IGNORED)
		return str, nil
	default:
		return "", errors.New("parseString error")
	}
}

func parseVariableStr(lexer *Lexer) (*Variable, error) {
	var variable Variable
	variable.isStrValue = true
	var err error

	lexer.NextTokenIs(TOKEN_VARIABLE_STR)
	if variable.name, err = parseName(lexer); err != nil {
		return nil, err
	}
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	return &variable, nil
}

func parseVariableInt(lexer *Lexer) (*Variable, error) {
	var variable Variable
	variable.isStrValue = false
	var err error
	lexer.NextTokenIs(TOKEN_VARIABLE_INT)
	if variable.name, err = parseName(lexer); err != nil {
		return nil, err
	}
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	return &variable, nil
}

func parseVariable(lexer *Lexer) (*Variable, error) {
	switch lexer.LookAhead() {
	case TOKEN_VARIABLE_STR:
		return parseVariableStr(lexer)
	case TOKEN_VARIABLE_INT:
		return parseVariableInt(lexer)
	default:
		return nil, errors.New("parseVariable error")
	}
}

func parseAssignment(lexer *Lexer) (*Assignment, error) {
	var assenment Assignment
	var err error

	if assenment.variable, err = parseVariable(lexer); err != nil {
		return nil, err
	}
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	lexer.NextTokenIs(TOKEN_EQUALTY)
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	if assenment.value, err = parseString(lexer); err != nil {
		return nil, err
	}
	return &assenment, nil
}

func parsePrint(lexer *Lexer) (*Print, error) {
	var print Print
	var err error

	lexer.NextTokenIs(TOKEN_PRINT)
	lexer.NextTokenIs(TOKEN_LEFT_PAIR)
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	if print.variable, err = parseVariable(lexer); err != nil {
		return nil, err
	}
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	lexer.NextTokenIs(TOKEN_RIGHT_PAIR)
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	return &print, err
}

var mathOperatorMap = map[int]int{
	TOKEN_ADD:  MATH_ADD,
	TOKEN_SUB:  MATH_SUB,
	TOKEN_MUIT: MATH_MUTL,
	TOKEN_DIV:  MATH_DIV,
}

//TODO
func parseMath(lexer *Lexer, which int) (*MathOperation, error) {
	var operator MathOperation
	var err error

	operator.which = mathOperatorMap[which]

	lexer.NextTokenIs(which)
	lexer.NextTokenIs(TOKEN_LEFT_PAIR)
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	if operator.lhs, err = parseVariable(lexer); err != nil {
		return nil, err
	}
	if operator.lhs.isStrValue {
		return nil, errors.New("add operation must be int value")
	}
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	lexer.NextTokenIs(TOKEN_SEPARATE)
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	if operator.rhs, err = parseVariable(lexer); err != nil {
		return nil, err
	}
	if operator.rhs.isStrValue {
		return nil, errors.New("add operation must be int value")
	}
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	lexer.NextTokenIs(TOKEN_RIGHT_PAIR)
	lexer.LookAheadAndSkip(TOKEN_IGNORED)
	return &operator, nil
}

/*
func parseAddVar(lexer *Lexer,add *Variable)(*Add, error) {
	var err error
	if add, err = parseVariable(lexer); err != nil {
		return nil, err
	}
	if add.isStrValue {
		return nil, errors.New("add operation must be int value")
	}
}
*/

//TODO support math operation
func parseStatement(lexer *Lexer) (Statement, error) {
	switch lexer.LookAhead() {
	case TOKEN_PRINT:
		return parsePrint(lexer)
	case TOKEN_ADD:
		return parseMath(lexer, TOKEN_ADD)
	case TOKEN_SUB:
		return parseMath(lexer, TOKEN_SUB)
	case TOKEN_MUIT:
		return parseMath(lexer, TOKEN_MUIT)
	case TOKEN_DIV:
		return parseMath(lexer, TOKEN_DIV)
	case TOKEN_VARIABLE_STR, TOKEN_VARIABLE_INT:
		return parseAssignment(lexer)
	default:
		return nil, errors.New("parseStatement(): unknown Statement.")
	}
}

func parseStatements(lexer *Lexer) ([]Statement, error) {
	var statements []Statement

	for !isSourceCodeEnd(lexer.LookAhead()) {
		var statement Statement
		var err error
		if statement, err = parseStatement(lexer); err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}
	return statements, nil
}

func isSourceCodeEnd(token int) bool {
	if token == TOKEN_EOF {
		return true
	}
	return false
}

func parseSourceCode(lexer *Lexer) (*SourceCode, error) {
	var sourceCode SourceCode
	var err error

	if sourceCode.Statements, err = parseStatements(lexer); err != nil {
		return nil, err
	}
	return &sourceCode, nil
}

func parse(code string) (*SourceCode, error) {
	var sourceCode *SourceCode
	var err error

	lexer := NewLexer(code)
	if sourceCode, err = parseSourceCode(lexer); err != nil {
		return nil, err
	}
	lexer.NextTokenIs(TOKEN_EOF)
	return sourceCode, nil
}
