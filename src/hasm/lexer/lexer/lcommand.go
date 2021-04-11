package lexer

import (
	"github.com/leonhfr/nand2tetris/src/hasm/lexer/errors"
	"github.com/leonhfr/nand2tetris/src/hasm/lexer/lexertoken"
)

// LexLeftParenthesis is a lexer function that emits a TOKEN_PARENTHESIS_LEFT
// and then returns the lexer for a section header.
func LexLeftParenthesis(lexer *Lexer) LexFn {
	lexer.Inc()
	lexer.Emit(lexertoken.TOKEN_PARENTHESIS_LEFT)
	return LexLSymbol
}

// LexLSymbol is a lexer function that emits a TOKEN_L_SYMBOL
func LexLSymbol(lexer *Lexer) LexFn {
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_MISSING_RIGHT_PARENTHESIS)
		}

		if lexer.IsNext(lexertoken.PARENTHESIS_RIGHT) {
			lexer.Emit(lexertoken.TOKEN_L_SYMBOL)
			return LexRightParenthesis
		}

		lexer.Inc()
	}
}

func LexRightParenthesis(lexer *Lexer) LexFn {
	lexer.Inc()
	lexer.Emit(lexertoken.TOKEN_PARENTHESIS_RIGHT)
	return LexStart
}
