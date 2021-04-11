package lexer

import (
	"github.com/leonhfr/nand2tetris/src/hasm/lexer/errors"
	"github.com/leonhfr/nand2tetris/src/hasm/lexer/lexertoken"
)

func LexSlash(lexer *Lexer) LexFn {
	lexer.Inc()
	return LexDoubleSlash
}

func LexDoubleSlash(lexer *Lexer) LexFn {
	if !lexer.IsNext(lexertoken.SLASH) {
		return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_SINGLE_SLASH)
	}
	lexer.Inc()
	lexer.Emit(lexertoken.TOKEN_DOUBLE_SLASH)
	return LexComment
}

func LexComment(lexer *Lexer) LexFn {
	for {
		if lexer.IsNext(lexertoken.NEWLINE) || lexer.IsNext(lexertoken.CARRIAGE_RETURN) {
			lexer.Emit(lexertoken.TOKEN_COMMENT)
			return LexStart
		}

		lexer.Inc()

		if lexer.IsEOF() {
			lexer.Emit(lexertoken.TOKEN_COMMENT)
			return LexStart
		}
	}
}
