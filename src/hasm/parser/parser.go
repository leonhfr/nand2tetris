package parser

import (
	"bufio"
	"fmt"
	"os"

	"github.com/leonhfr/nand2tetris/src/hasm/command"
)

const (
	START = iota
	SLASH
	COMMENT
	C_COMMAND
	A_COMMAND
	L_COMMAND
)

type Parser struct {
	reader   bufio.Reader
	commands []command.Command
}

func New(path string) (*Parser, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	return &Parser{*reader, make([]command.Command, 0)}, nil
}

func (p *Parser) Commands() []command.Command {
	return p.commands
}

func (p *Parser) Parse() error {
	state := 0
	cache := ""
	charcode, err := p.reader.ReadByte()
	for err == nil {
		character := string(charcode)
		switch state {
		case START:
			cache = ""
			switch character {
			case "/":
				state = SLASH
			case " ":
				state = START
			case "\r":
				state = START
			case "\n":
				state = START
			case "@":
				state = A_COMMAND
			case "(":
				state = L_COMMAND
			default:
				state = C_COMMAND
				cache = character
			}
		case SLASH:
			switch character {
			case "/":
				state = COMMENT
			default:
				return fmt.Errorf("unexpected single slash")
			}
		case COMMENT:
			switch character {
			case "\r":
				state = START
			case "\n":
				state = START
			}
		case A_COMMAND:
			switch character {
			case " ":
				state = A_COMMAND
			case "/":
				a := command.NewA(cache)
				p.commands = append(p.commands, a)
				state = SLASH
			case "\r":
				a := command.NewA(cache)
				p.commands = append(p.commands, a)
				state = START
			case "\n":
				a := command.NewA(cache)
				p.commands = append(p.commands, a)
				state = START
			default:
				state = A_COMMAND
				cache = cache + character
			}
		case C_COMMAND:
			switch character {
			case " ":
				state = C_COMMAND
			case "/":
				c := command.NewC(cache)
				p.commands = append(p.commands, c)
				state = SLASH
			case "\r":
				c := command.NewC(cache)
				p.commands = append(p.commands, c)
				state = START
			case "\n":
				c := command.NewC(cache)
				p.commands = append(p.commands, c)
				state = START
			default:
				cache = cache + character
			}
		case L_COMMAND:
			switch character {
			case " ":
				state = L_COMMAND
			case "\r":
				return fmt.Errorf("unexpected missing )")
			case "\n":
				return fmt.Errorf("unexpected missing )")
			case ")":
				l := command.NewL(cache)
				p.commands = append(p.commands, l)
				state = START
			default:
				cache = cache + character
			}
		}

		charcode, err = p.reader.ReadByte()
	}

	fmt.Println(p)
	return nil
}

func ParseC(command string) (string, string, string) {
	var dest, comp, jump, cache string
	state := 0
	for _, c := range command {
		switch state {
		case 0:
			// start, dest or comp
			switch c {
			case '=':
				dest, cache = cache, ""
				state = 1
			case ';':
				comp, cache = cache, ""
				state = 2
			default:
				cache = cache + string(c)
			}
		case 1:
			// comp
			switch c {
			case ';':
				comp, cache = cache, ""
				state = 2
			default:
				cache = cache + string(c)
			}
		case 2:
			// dest
			cache = cache + string(c)
		}
	}
	switch state {
	case 1:
		// cache to comp
		comp = cache
	case 2:
		// cache to dest
		dest = cache
	}
	return dest, comp, jump
}

func (p Parser) String() string {
	s := ""
	for _, c := range p.commands {
		s = s + fmt.Sprintln(c)
	}
	return s
}
