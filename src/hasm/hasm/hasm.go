package hasm

import "fmt"

type CommandType int

const (
	A_COMMAND CommandType = iota
	C_COMMAND
	L_COMMAND
	COMMENT
)

type HasmFile struct {
	FileName string      `json:"fileName"`
	Commands CommandList `json:"commands"`
}

type Command interface {
	Encode() string
}

type CommandList []Command

type ACommand struct {
	Symbol string `json:"symbol"`
}

func NewA(symbol string) ACommand {
	return ACommand{symbol}
}

func (a ACommand) Encode() string {
	return a.Symbol
}

type CCommand struct {
	Dest string `json:"dest"`
	Comp string `json:"comp"`
	Jump string `json:"jump"`
}

func NewC(dest, comp, jump string) CCommand {
	return CCommand{dest, comp, jump}
}

func (c CCommand) Encode() string {
	return fmt.Sprintf("%q=%q;%q", c.Dest, c.Comp, c.Jump)
}

type LCommand struct {
	Symbol string `json:"symbol"`
}

func NewL(symbol string) LCommand {
	return LCommand{symbol}
}

func (l LCommand) Encode() string {
	return l.Symbol
}
