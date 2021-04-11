package lexer

import (
	"unicode"
	"unicode/utf8"

	"github.com/leonhfr/nand2tetris/src/hasm/lexer/lexertoken"
)

// Lexer contains the state of our parser and provides
// a stream for accepting tokens.
type Lexer struct {
	// TODO: read stream of characters instead of the whole string
	Name   string
	Input  string
	Tokens chan lexertoken.Token
	State  LexFn

	Start int
	Pos   int
}

// New returns a new lexer with a given input string.
// This returns the instance of the lexer and a channel of tokens.
func New(name, input string) *Lexer {
	return &Lexer{
		Name:   name,
		Input:  input,
		State:  LexStart,
		Tokens: make(chan lexertoken.Token, 2),
	}
}

// Emit puts a token onto the token channel. The value of this
// token is read from the input based on the current lexer position.
func (lexer *Lexer) Emit(tokenType lexertoken.TokenType) {
	lexer.Tokens <- lexertoken.Token{
		Type:  tokenType,
		Value: lexer.Input[lexer.Start:lexer.Pos],
	}
	lexer.Start = lexer.Pos
}

// Next reads the next rune from the input stream and advances
// the lexer position.
func (lexer *Lexer) Next() rune {
	if lexer.Pos >= utf8.RuneCountInString(lexer.Input) {
		return lexertoken.EOF
	}

	result, _ := utf8.DecodeRuneInString(lexer.Input[lexer.Pos : lexer.Pos+1])
	lexer.Pos++
	return result
}

func (lexer *Lexer) NextToken() lexertoken.Token {
	for {
		select {
		case token := <-lexer.Tokens:
			return token
		default:
			lexer.State = lexer.State(lexer)
		}
	}
}

// Peek returns the next rune in the stream without consuming it.
func (lexer *Lexer) Peek() rune {
	r := lexer.Next()
	lexer.Dec()
	return r
}

// Inc increments the position.
func (lexer *Lexer) Inc() {
	lexer.Pos++
	if lexer.Pos >= utf8.RuneCountInString(lexer.Input) {
		lexer.Emit(lexertoken.TOKEN_EOF)
	}
}

// Dec decrements the position.
func (lexer *Lexer) Dec() {
	lexer.Pos--
}

// IsEOF returns whether the lexer is at the end of the input stream.
func (lexer *Lexer) IsEOF() bool {
	return lexer.Pos >= len(lexer.Input)
}

func (lexer *Lexer) IsNext(r rune) bool {
	return lexer.Peek() == r
}

// SkipWhitespace skips whitespace until we get something meaningful.
func (lexer *Lexer) SkipWhitespace() {
	for {
		ch := lexer.Next()
		if !unicode.IsSpace(ch) {
			lexer.Dec()
			lexer.Start = lexer.Pos
			break
		}
		if ch == lexertoken.EOF {
			lexer.Emit(lexertoken.TOKEN_EOF)
			break
		}
	}
}
