package lexertoken

import (
	"fmt"
)

type TokenType int

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_EOF

	TOKEN_AT_SIGN
	TOKEN_EQUAL_SIGN
	TOKEN_SEMICOLON
	TOKEN_PARENTHESIS_LEFT
	TOKEN_PARENTHESIS_RIGHT
	TOKEN_DOUBLE_SLASH
	TOKEN_A_SYMBOL
	TOKEN_L_SYMBOL
	TOKEN_C_DEST
	TOKEN_C_COMP
	TOKEN_C_JUMP
	TOKEN_COMMENT
)

const (
	EOF               rune = 0
	NEWLINE           rune = '\n'
	CARRIAGE_RETURN   rune = '\r'
	AT_SIGN           rune = '@'
	EQUAL_SIGN        rune = '='
	SEMI_COLON        rune = ';'
	PARENTHESIS_LEFT  rune = '('
	PARENTHESIS_RIGHT rune = ')'
	SLASH             rune = '/'
)

type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) String() string {
	switch t.Type {
	case TOKEN_EOF:
		return "EOF"
	case TOKEN_ERROR:
		return t.Value
	default:
		return fmt.Sprintf("%q", t.Value)
	}
}
