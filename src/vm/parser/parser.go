package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leonhfr/nand2tetris/src/vm/bytecode"
)

type Parser struct {
	in     <-chan string
	out    chan *bytecode.Command
	errors chan error
}

func New(in <-chan string, out chan *bytecode.Command, errors chan error) *Parser {
	return &Parser{in, out, errors}
}

func (p *Parser) Parse() {
	for line := range p.in {
		line = strings.Split(line, "//")[0]
		line = strings.TrimSpace(line)

		if line != "" {
			p.parseCommand(line)
		}
	}
	close(p.out)
}

func (p *Parser) parseCommand(line string) {
	parts := strings.Split(line, " ")
	c := parts[0]

	bc, isArithmetic := arithmetic[c]
	if isArithmetic {
		p.out <- &bytecode.Command{
			Original: line,
			Type:     bc,
		}
		return
	}

	arg1 := parts[1]
	arg2, _ := strconv.Atoi(parts[2])

	switch c {
	case "push":
		p.out <- &bytecode.Command{
			Original: line,
			Type:     bytecode.C_PUSH,
			Arg1:     arg1,
			Arg2:     arg2,
		}
	case "pop":
		p.out <- &bytecode.Command{
			Original: line,
			Type:     bytecode.C_POP,
			Arg1:     arg1,
			Arg2:     arg2,
		}
	default:
		p.errors <- fmt.Errorf("unexpected command: %v", line)
	}
}

var arithmetic = map[string]bytecode.CommandType{
	"add": bytecode.C_ADD,
	"sub": bytecode.C_SUB,
	"neg": bytecode.C_NEG,
	"eq":  bytecode.C_EQ,
	"gt":  bytecode.C_GT,
	"lt":  bytecode.C_LT,
	"and": bytecode.C_AND,
	"or":  bytecode.C_OR,
	"not": bytecode.C_NOT,
}
