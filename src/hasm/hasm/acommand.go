package hasm

import (
	"fmt"
	"strconv"

	"github.com/leonhfr/nand2tetris/src/hasm/symboltable"
	"github.com/leonhfr/nand2tetris/src/hasm/util"
)

type ACommand struct {
	Symbol string `json:"symbol"`
}

func NewA(symbol string) ACommand {
	return ACommand{symbol}
}

func (a ACommand) Handle(st *symboltable.SymbolTable) (string, error) {
	if decimal, err := strconv.Atoi(a.Symbol); err == nil {
		return encodeA(decimal), nil
	}

	if address, err := st.Get(a.Symbol); err == nil {
		return encodeA(address), nil
	}

	address, err := st.Add(a.Symbol)
	if err != nil {
		return "", err
	}

	return encodeA(address), nil
}

func encodeA(address int) string {
	binary := util.DecimalToBinary(address)
	return fmt.Sprint("0", util.PadZeroLeft(binary, 15))
}
