package parser

import (
	"github.com/leonhfr/nand2tetris/src/hasm/hasm"
	"github.com/leonhfr/nand2tetris/src/hasm/lexer/lexer"
	"github.com/leonhfr/nand2tetris/src/hasm/lexer/lexertoken"
)

func isEOF(token lexertoken.Token) bool {
	return token.Type == lexertoken.TOKEN_EOF
}

func Parse(fileName, input string) hasm.HasmFile {
	output := hasm.HasmFile{
		FileName: fileName,
		Commands: make([]hasm.Command, 0),
	}

	l := lexer.New(fileName, input)

	var token lexertoken.Token
	var dest, comp string
	for {
		token = l.NextToken()

		if isEOF(token) {
			break
		}

		switch token.Type {
		case lexertoken.TOKEN_A_SYMBOL:
			if dest != "" && comp != "" {
				i := hasm.NewC(dest, comp, "")
				output.Commands = append(output.Commands, i)
				dest, comp = "", ""
			}
			i := hasm.NewA(token.Value)
			output.Commands = append(output.Commands, i)
		case lexertoken.TOKEN_C_DEST:
			if dest != "" && comp != "" {
				i := hasm.NewC(dest, comp, "")
				output.Commands = append(output.Commands, i)
				dest, comp = "", ""
			}
			dest = token.Value
		case lexertoken.TOKEN_C_COMP:
			if dest != "" && comp != "" {
				i := hasm.NewC(dest, comp, "")
				output.Commands = append(output.Commands, i)
				dest, comp = "", ""
			}
			comp = token.Value
		case lexertoken.TOKEN_C_JUMP:
			i := hasm.NewC(dest, comp, token.Value)
			output.Commands = append(output.Commands, i)
			dest, comp = "", ""
		case lexertoken.TOKEN_L_SYMBOL:
			if dest != "" && comp != "" {
				i := hasm.NewC(dest, comp, "")
				output.Commands = append(output.Commands, i)
				dest, comp = "", ""
			}
			i := hasm.NewL(token.Value)
			output.Commands = append(output.Commands, i)
		}
	}

	return output
}
