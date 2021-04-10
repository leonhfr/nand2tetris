package command

// "github.com/leonhfr/nand2tetris/src/hasm/symboltable"

type CommandType int

const (
	A_COMMAND CommandType = iota
	C_COMMAND
	L_COMMAND
)

type Command struct {
	kind    CommandType
	literal string
}

func NewA(literal string) Command {
	kind := A_COMMAND
	return Command{kind, literal}
}

func NewC(literal string) Command {
	kind := C_COMMAND
	return Command{kind, literal}
}

func NewL(literal string) Command {
	kind := L_COMMAND
	return Command{kind, literal}
}

func (c *Command) Type() CommandType {
	return c.kind
}
