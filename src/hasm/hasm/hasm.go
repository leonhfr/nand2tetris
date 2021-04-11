package hasm

import (
	"fmt"

	"github.com/leonhfr/nand2tetris/src/hasm/symboltable"
)

type CommandType int

const (
	A_COMMAND CommandType = iota
	C_COMMAND
	L_COMMAND
)

type HasmFile struct {
	FileName string      `json:"fileName"`
	Commands CommandList `json:"commands"`
}

type Command interface {
	Handle(st *symboltable.SymbolTable) (string, error)
}

type CommandList []Command

func (h *HasmFile) FirstPass(st *symboltable.SymbolTable) error {
	for _, command := range h.Commands {
		switch command.(type) {
		case LCommand:
			_, err := command.Handle(st)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *HasmFile) SecondPass(st *symboltable.SymbolTable) (string, error) {
	var output string
	for _, command := range h.Commands {
		switch command.(type) {
		case ACommand:
			a, err := command.Handle(st)
			if err != nil {
				return output, err
			}
			output += fmt.Sprintln(a)
		case CCommand:
			c, err := command.Handle(st)
			if err != nil {
				return output, err
			}
			output += fmt.Sprintln(c)
		}
	}
	return output, nil
}
