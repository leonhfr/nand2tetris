package hasm

import (
	"fmt"

	"github.com/leonhfr/nand2tetris/src/hasm/symboltable"
)

type CCommand struct {
	Dest string `json:"dest"`
	Comp string `json:"comp"`
	Jump string `json:"jump"`
}

func NewC(dest, comp, jump string) CCommand {
	return CCommand{dest, comp, jump}
}

func (c CCommand) Handle(st *symboltable.SymbolTable) (string, error) {
	var output string = "111"

	if compA0, ok := compA0Table[c.Comp]; ok {
		output = fmt.Sprint(output, "0", compA0)
	} else if compA1, ok := compA1Table[c.Comp]; ok {
		output = fmt.Sprint(output, "1", compA1)
	} else {
		return output, fmt.Errorf("comparison symbol %v does not exist in the comp table", c.Comp)
	}

	if c.Dest == "" {
		output = fmt.Sprint(output, "000")
	} else if dest, ok := destTable[c.Dest]; ok {
		output = fmt.Sprint(output, dest)
	} else {
		return output, fmt.Errorf("destination symbol %v does not exist in the dest table", c.Dest)
	}

	if c.Jump == "" {
		output = fmt.Sprint(output, "000")
	} else if jump, ok := jumpTable[c.Jump]; ok {
		output = fmt.Sprint(output, jump)
	} else {
		return output, fmt.Errorf("jump symbol %v does not exist in the jump table", c.Dest)
	}

	return output, nil
}

var compA0Table = map[string]string{
	"0":   "101010",
	"1":   "111111",
	"-1":  "111010",
	"D":   "001100",
	"A":   "110000",
	"!D":  "001101",
	"!A":  "110001",
	"-D":  "001111",
	"-A":  "110011",
	"D+1": "011111",
	"A+1": "110111",
	"D-1": "001110",
	"A-1": "110010",
	"D+A": "000010",
	"D-A": "010011",
	"A-D": "000111",
	"D&A": "000000",
	"D|A": "010101",
}

var compA1Table = map[string]string{
	"M":   "110000",
	"!M":  "110001",
	"-M":  "110011",
	"M+1": "110111",
	"M-1": "110010",
	"D+M": "000010",
	"D-M": "010011",
	"M-D": "000111",
	"D&M": "000000",
	"D|M": "010101",
}

var destTable = map[string]string{
	"M":   "001",
	"D":   "010",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"AMD": "111",
}

var jumpTable = map[string]string{
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}
