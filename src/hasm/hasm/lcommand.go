package hasm

import (
	"fmt"

	"github.com/leonhfr/nand2tetris/src/hasm/symboltable"
)

type LCommand struct {
	Symbol string `json:"symbol"`
}

func NewL(symbol string) LCommand {
	return LCommand{symbol}
}

func (l LCommand) Handle(st *symboltable.SymbolTable) (string, error) {
	address, err := st.Add(l.Symbol)
	return fmt.Sprint(address), err
}
