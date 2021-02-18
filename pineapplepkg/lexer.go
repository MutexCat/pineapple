package pineapplepkg

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	TOKEN_EOF          = iota //EOF
	TOKEN_NAME                //name
	TOKEN_QUOTE               // "
	TOKEN_LEFT_PAIR           // (
	TOKEN_RIGHT_PAIR          // )
	TOKEN_DOUBLE_QUOTE        // ""
	TOKEN_EQUALTY             // =
	TOKEN_VARIABLE_STR        //$
	TOKEN_VARIABLE_INT        //@
	TOKEN_IGNORED             // Ignored
	TOKEN_PRINT
	TOKEN_ADD
	TOKEN_SUB
	TOKEN_MUIT
	TOKEN_DIV
)

var tokenNameMap = map[int]string{}

var keywords = map[string]int{
	"print": TOKEN_PRINT,
	"add":   TOKEN_ADD,
	"sub":   TOKEN_SUB,
	"mutl":  TOKEN_MUIT,
	"div":   TOKEN_DIV,
}

type Lexer struct {
	SourceCode    string
	LineNum       int
	NextToken     string
	NextTokenType int
	NextLineNUm   int
}

func NewLexer(code string) *Lexer {
	return &Lexer{
		SourceCode:    code,
		LineNum:       1,
		NextToken:     "",
		NextTokenType: 0,
		NextLineNUm:   0,
	}
}

func (lexet *Lexer) scanBeforeToken(token string) string {
	str := strings.Split(lexet.SourceCode, token)
	if len(str) < 2 {
		panic("unreachable")
	}
	lexet.Skip(len(str[0]))
	return str[0]
}

func (lexer *Lexer) LookAhead() int {
	// lexer.nextToken* already setted
	if lexer.NextLineNUm > 0 {
		return lexer.NextTokenType
	}
	// set it
	nowLineNum := lexer.LineNum
	lineNum, tokenType, token := lexer.GetNextToken()
	lexer.LineNum = nowLineNum
	lexer.NextLineNUm = lineNum
	lexer.NextTokenType = tokenType
	lexer.NextToken = token
	return tokenType
}

func (lexer *Lexer) LookAheadAndSkip(expectedType int) {
	// get next token
	nowLineNum := lexer.LineNum
	lineNum, tokenType, token := lexer.GetNextToken()
	// not is expected type, reverse cursor
	if tokenType != expectedType {
		lexer.LineNum = nowLineNum
		lexer.NextLineNUm = lineNum
		lexer.NextTokenType = tokenType
		lexer.NextToken = token
	}
}

func (lexer *Lexer) NextTokenIs(tokenType int) (lineNum int, token string) {
	nowLineNum, nowTokenType, nowToken := lexer.GetNextToken()
	if tokenType != nowTokenType {
		err := fmt.Sprintf("NextTokenIs(): syntax error near '%s', expected token: {%s} but got {%s}.",
			tokenNameMap[nowTokenType], tokenNameMap[tokenType], tokenNameMap[nowTokenType])
		panic(err)
	}
	return nowLineNum, nowToken
}

func (lexer *Lexer) GetNextToken() (lineNum int, tokenType int, token string) {
	// next token already loaded
	if lexer.NextLineNUm > 0 {
		lineNum = lexer.NextLineNUm
		tokenType = lexer.NextTokenType
		token = lexer.NextToken
		lexer.LineNum = lexer.NextLineNUm
		lexer.NextLineNUm = 0
		return
	}
	return lexer.MatchToken()
}

func (lexer *Lexer) MatchToken() (int, int, string) {

	if lexer.Ignored() {
		return lexer.LineNum, TOKEN_IGNORED, ""
	}

	if len(lexer.SourceCode) == 0 {
		return lexer.LineNum, TOKEN_EOF, ""
	}

	switch lexer.SourceCode[0] {
	case '$':
		lexer.Skip(1)
		return lexer.LineNum, TOKEN_VARIABLE_STR, "$"
	case '@':
		lexer.Skip(1)
		return lexer.LineNum, TOKEN_VARIABLE_INT, "@"
	case '(':
		lexer.Skip(1)
		return lexer.LineNum, TOKEN_LEFT_PAIR, "("
	case ')':
		lexer.Skip(1)
		return lexer.LineNum, TOKEN_RIGHT_PAIR, ")"
	case '=':
		lexer.Skip(1)
		return lexer.LineNum, TOKEN_EQUALTY, "="
	case '"':
		if lexer.NextLetterIs("\"\"") {
			lexer.Skip(2)
			return lexer.LineNum, TOKEN_DOUBLE_QUOTE, "\"\""
		}
		lexer.Skip(1)
		return lexer.LineNum, TOKEN_QUOTE, "\""
	}

	if lexer.SourceCode[0] == '_' || isLetter(lexer.SourceCode[0]) {
		str := lexer.scanName()
		if token, match := keywords[str]; match {
			//lexer.Skip(len(str))
			return lexer.LineNum, token, str
		} else {
			//lexer.Skip(len(str))
			return lexer.LineNum, TOKEN_NAME, str
		}
	}
	err := fmt.Sprintf("MatchToken error Line <%d>", lexer.LineNum)
	panic(err)
}

func (lexer *Lexer) Ignored() bool {
	isIgnored := false
	// target pattern
	isNewLine := func(c byte) bool {
		return c == '\r' || c == '\n'
	}
	isWhiteSpace := func(c byte) bool {
		switch c {
		case '\t', '\n', '\v', '\f', '\r', ' ':
			return true
		}
		return false
	}
	// matching
	for len(lexer.SourceCode) > 0 {
		if lexer.NextLetterIs("\r\n") || lexer.NextLetterIs("\n\r") {
			lexer.Skip(2)
			lexer.LineNum += 1
			isIgnored = true
		} else if isNewLine(lexer.SourceCode[0]) {
			lexer.Skip(1)
			lexer.LineNum += 1
			isIgnored = true
		} else if isWhiteSpace(lexer.SourceCode[0]) {
			lexer.Skip(1)
			isIgnored = true
		} else {
			break
		}
	}
	return isIgnored
}

func (lexer *Lexer) scanName() string {
	var reg = regexp.MustCompile(`^[_\w\d]+`)
	if token := reg.FindString(lexer.SourceCode); token != "" {
		lexer.Skip(len(token))
		return token
	}
	panic("unreachable scanName")
	return ""
}

func isLetter(c byte) bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z'
}

func (lexer *Lexer) NextLetterIs(letter string) bool {
	return strings.HasPrefix(lexer.SourceCode, letter)
}

func (lexer *Lexer) Skip(n int) {
	lexer.SourceCode = lexer.SourceCode[n:]
}
