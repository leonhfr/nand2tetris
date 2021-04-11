package lexer

import (
	"github.com/leonhfr/nand2tetris/src/hasm/lexer/errors"
	"github.com/leonhfr/nand2tetris/src/hasm/lexer/lexertoken"
)

func LexAtSign(lexer *Lexer) LexFn {
	lexer.Inc()
	lexer.Emit(lexertoken.TOKEN_AT_SIGN)
	return LexASymbol
}

func LexASymbol(lexer *Lexer) LexFn {
	for {
		if lexer.Peek() == ' ' {
			lexer.Emit(lexertoken.TOKEN_A_SYMBOL)
			return LexStart
		}

		if lexer.IsNext(lexertoken.NEWLINE) || lexer.IsNext(lexertoken.CARRIAGE_RETURN) {
			lexer.Emit(lexertoken.TOKEN_A_SYMBOL)
			return LexStart
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}
