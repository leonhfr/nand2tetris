package lexer

import (
	"fmt"

	"github.com/leonhfr/nand2tetris/src/hasm/lexer/lexertoken"
)

// LexFn is the lexer state function type that processes tokens.
type LexFn func(*Lexer) LexFn

// LexStart is a lexer function to start everything off.
func LexStart(lexer *Lexer) LexFn {
	lexer.SkipWhitespace()
	switch {
	case lexer.IsNext(lexertoken.AT_SIGN):
		return LexAtSign
	case lexer.IsNext(lexertoken.PARENTHESIS_LEFT):
		return LexLeftParenthesis
	case lexer.IsNext(lexertoken.SLASH):
		return LexSlash
	default:
		return LexDestOrComp
	}
}

// Errorf returns a token with error information.
func (lexer *Lexer) Errorf(format string, args ...interface{}) LexFn {
	lexer.Tokens <- lexertoken.Token{
		Type:  lexertoken.TOKEN_ERROR,
		Value: fmt.Sprintf(format, args...),
	}
	return nil
}
