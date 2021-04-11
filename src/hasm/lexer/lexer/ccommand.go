package lexer

import (
	"github.com/leonhfr/nand2tetris/src/hasm/lexer/errors"
	"github.com/leonhfr/nand2tetris/src/hasm/lexer/lexertoken"
)

func LexDestOrComp(lexer *Lexer) LexFn {
	for {
		if lexer.IsNext(lexertoken.EQUAL_SIGN) {
			lexer.Emit(lexertoken.TOKEN_C_DEST)
			return LexEqualSign
		}

		if lexer.IsNext(lexertoken.SEMI_COLON) {
			lexer.Emit(lexertoken.TOKEN_C_COMP)
			return LexSemiColon
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}

func LexEqualSign(lexer *Lexer) LexFn {
	lexer.Inc()
	lexer.Emit(lexertoken.TOKEN_EQUAL_SIGN)
	return LexComp
}

func LexSemiColon(lexer *Lexer) LexFn {
	lexer.Inc()
	lexer.Emit(lexertoken.TOKEN_SEMICOLON)
	return LexJump
}

func LexComp(lexer *Lexer) LexFn {
	for {
		if lexer.Peek() == ' ' {
			lexer.Emit(lexertoken.TOKEN_C_COMP)
			return LexStart
		}

		if lexer.IsNext(lexertoken.SEMI_COLON) {
			lexer.Emit(lexertoken.TOKEN_C_COMP)
			return LexSemiColon
		}

		if lexer.IsNext(lexertoken.NEWLINE) || lexer.IsNext(lexertoken.CARRIAGE_RETURN) {
			lexer.Emit(lexertoken.TOKEN_C_COMP)
			return LexStart
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}

func LexJump(lexer *Lexer) LexFn {
	for {
		if lexer.Peek() == ' ' {
			lexer.Emit(lexertoken.TOKEN_C_JUMP)
			return LexStart
		}

		if lexer.IsNext(lexertoken.NEWLINE) || lexer.IsNext(lexertoken.CARRIAGE_RETURN) {
			lexer.Emit(lexertoken.TOKEN_C_JUMP)
			return LexStart
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}
