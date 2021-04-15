package bytecode

type CommandType int

const (
	C_ADD CommandType = iota
	C_SUB
	C_NEG
	C_EQ
	C_GT
	C_LT
	C_AND
	C_OR
	C_NOT
	C_PUSH
	C_POP
	// C_LABEL
	// C_GOTO
	// C_IF
	// C_FUNCTION
	// C_RETURN
	// C_CALL
)

type Command struct {
	Original string
	Type     CommandType
	Arg1     string
	Arg2     int
}
