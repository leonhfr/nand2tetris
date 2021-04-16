package translator

import (
	"fmt"

	"github.com/leonhfr/nand2tetris/src/vm/bytecode"
)

type Translator struct {
	filename string
	entry    bool
	in       <-chan *bytecode.Command
	out      chan string
	errors   chan error
	labels   int
}

func New(filename string, entry bool, in <-chan *bytecode.Command, out chan string, errors chan error) *Translator {
	labels := 0
	return &Translator{filename, entry, in, out, errors, labels}
}

func (t *Translator) Translate() {
	if t.entry {
		t.emitEntry()
	}
	for command := range t.in {
		t.translate(command)
	}
	close(t.out)
}

func (t *Translator) emitEntry() {
	t.emitComment("Entry commands")
	t.emitConstantToD(256)
	t.emitDToAddress("SP")
	t.emitCall("Sys.init", 0)
}

func (t *Translator) translate(command *bytecode.Command) {
	t.emitComment(command.Original)
	switch command.Type {
	case bytecode.C_ADD:
		t.emitComputeTwoArgs("D+M")
	case bytecode.C_SUB:
		t.emitComputeTwoArgs("M-D")
	case bytecode.C_NEG:
		t.emitComputeOneArg("-M")
	case bytecode.C_EQ:
		t.emitComputeComparable("JEQ")
	case bytecode.C_GT:
		t.emitComputeComparable("JGT")
	case bytecode.C_LT:
		t.emitComputeComparable("JLT")
	case bytecode.C_AND:
		t.emitComputeTwoArgs("D&M")
	case bytecode.C_OR:
		t.emitComputeTwoArgs("D|M")
	case bytecode.C_NOT:
		t.emitComputeOneArg("!M")
	case bytecode.C_PUSH:
		t.emitPush(command)
	case bytecode.C_POP:
		t.emitPop(command)
	case bytecode.C_LABEL:
		t.emitLabel(command.Arg1)
	case bytecode.C_GOTO:
		t.emitGoTo(command.Arg1)
	case bytecode.C_IF:
		t.emitIfGoTo(command.Arg1)
	case bytecode.C_FUNCTION:
		t.emitFunction(command.Arg1, command.Arg2)
	case bytecode.C_CALL:
		t.emitCall(command.Arg1, command.Arg2)
	case bytecode.C_RETURN:
		t.emitReturn()
	}
}

func (t *Translator) emitPush(command *bytecode.Command) {
	switch command.Arg1 {
	case "local", "argument", "this", "that":
		segment, err := segmentAddress(command.Arg1)
		if err != nil {
			t.errors <- err
		}
		// D = &(@segment + i)
		t.emit(fmt.Sprintf("@%v", segment))
		t.emit("D=M")
		t.emit(fmt.Sprintf("@%v", command.Arg2))
		t.emit("A=D+A")
		t.emit("D=M")
		t.emitDToStack()
		t.emitStackInc()
	case "constant":
		t.emitConstantToD(command.Arg2)
		t.emitDToStack()
		t.emitStackInc()
	case "static":
		// D = @<filename>.<i>
		t.emit(fmt.Sprintf("@%v.%v", t.filename, command.Arg2))
		t.emit("D=M")
		t.emitDToStack()
		t.emitStackInc()
	case "temp":
		// D = &(i + 5)
		t.emit(fmt.Sprintf("@%v", command.Arg2+5))
		t.emit("D=M")
		t.emitDToStack()
		t.emitStackInc()
	case "pointer":
		offset, err := pointerOffset(command.Arg2)
		if err != nil {
			t.errors <- err
		}
		t.emit(fmt.Sprintf("@%v", offset))
		t.emit("D=M")
		t.emitDToStack()
		t.emitStackInc()
	default:
		t.errors <- fmt.Errorf("unknown segment name %v", command.Arg1)
	}
}

func (t *Translator) emitPop(command *bytecode.Command) {
	switch command.Arg1 {
	case "local", "argument", "this", "that":
		segment, err := segmentAddress(command.Arg1)
		if err != nil {
			t.errors <- err
		}
		// R13 = @segment + i
		t.emitAddressToD(segment)
		t.emit(fmt.Sprintf("@%v", command.Arg2))
		t.emit("D=D+A")
		t.emitDToAddress("R13")
		// &R13 = D
		t.emitStackDec()
		t.emitStackToD()
		t.emits("@R13", "A=M", "M=D")
	case "constant":
		t.errors <- fmt.Errorf("pop constant is not supported")
	case "static":
		t.emitStackDec()
		t.emitStackToD()
		// @<filename>.<i> = D
		t.emit(fmt.Sprintf("@%v.%v", t.filename, command.Arg2))
		t.emit("M=D")
	case "temp":
		// R13 = 5 + i
		t.emitConstantToD(5)
		t.emit(fmt.Sprintf("@%v", command.Arg2))
		t.emit("D=D+A")
		t.emitDToAddress("R13")
		// &R13 = D
		t.emitStackDec()
		t.emitStackToD()
		t.emits("@R13", "A=M", "M=D")
	case "pointer":
		offset, err := pointerOffset(command.Arg2)
		if err != nil {
			t.errors <- err
		}
		t.emitStackDec()
		t.emitStackToD()
		t.emit(fmt.Sprintf("@%v", offset))
		t.emit("M=D")
	default:
		t.errors <- fmt.Errorf("unknown segment name %v", command.Arg1)
	}
}

func (t *Translator) emitLabel(label string) {
	t.emit(fmt.Sprintf("(%v)", label))
}

func (t *Translator) emitLabelAddress(label string) {
	t.emit(fmt.Sprintf("@%v", label))
}

func (t *Translator) emitGoTo(label string) {
	t.emit(fmt.Sprintf("@%v", label))
	t.emit("0;JMP")
}

func (t *Translator) emitIfGoTo(label string) {
	t.emitStackDec()
	t.emitStackToD()
	t.emit(fmt.Sprintf("@%v", label))
	t.emit("D;JNE")
}

func (t *Translator) emitFunction(name string, args int) {
	t.emitLabel(fmt.Sprintf("CALL_FUNCTION_%v", name))
	for i := args; i > 0; i-- {
		t.emitConstantToD(0)
		t.emitDToStack()
		t.emitStackInc()
	}
}

func (t *Translator) emitCall(name string, args int) {
	returnAddress := fmt.Sprintf("CALL_RETURN_%v_%v", name, t.labels)

	t.emitComment("call: store return address")
	t.emitLabelAddress(returnAddress)
	t.emit("D=A")
	t.emitDToStack()
	t.emitStackInc()

	t.emitComment("call: store call frame")
	t.emitAddressToD("LCL")
	t.emitDToStack()
	t.emitStackInc()
	t.emitAddressToD("ARG")
	t.emitDToStack()
	t.emitStackInc()
	t.emitAddressToD("THIS")
	t.emitDToStack()
	t.emitStackInc()
	t.emitAddressToD("THAT")
	t.emitDToStack()
	t.emitStackInc()

	t.emitComment("call: compute ARG")
	t.emitAddressToD("SP")
	t.emit("@5")
	t.emit("D=D-A")
	t.emit(fmt.Sprintf("@%v", args))
	t.emit("D=D-A")
	t.emitDToAddress("ARG")

	t.emitComment("call: set up LCL")
	t.emitAddressToD("SP")
	t.emitDToAddress("LCL")

	t.emitComment("call: jump to the function")
	t.emitLabelAddress(fmt.Sprintf("CALL_FUNCTION_%v", name))
	t.emit("0;JMP")

	t.emitComment("call: return address")
	t.emitLabel(returnAddress)

	t.labels++
}

func (t *Translator) emitReturn() {
	t.emitComment("return: store addresses of end of call frame and return")
	t.emitAddressToD("LCL")
	t.emitDToAddress("R13")
	t.emits("@5", "A=D-A", "D=M")
	t.emitDToAddress("R14")

	t.emitComment("return: pop the return value to the start of the frame")
	t.emitStackDec()
	t.emitStackToD()
	t.emits("@ARG", "A=M", "M=D")

	t.emitComment("return: set SP to parent")
	t.emits("@ARG", "D=M+1")
	t.emitDToAddress("SP")

	t.emitComment("return: restore the call frame")
	t.emitAddressToD("R13")
	t.emits("@1", "D=D-A", "A=D", "D=M", "@THAT", "M=D")

	t.emitAddressToD("R13")
	t.emits("@2", "D=D-A", "A=D", "D=M", "@THIS", "M=D")

	t.emitAddressToD("R13")
	t.emits("@3", "D=D-A", "A=D", "D=M", "@ARG", "M=D")

	t.emitAddressToD("R13")
	t.emits("@4", "D=D-A", "A=D", "D=M", "@LCL", "M=D")

	t.emitComment("return: jump to return address")
	t.emitAddressToD("R14")
	t.emit("A=D")
	t.emit("0;JMP")
}

func (t *Translator) emitStackInc() {
	t.emit("@SP")
	t.emit("M=M+1")
}

func (t *Translator) emitStackDec() {
	t.emit("@SP")
	t.emit("M=M-1")
}

func (t *Translator) emitStackToD() {
	t.emit("@SP")
	t.emit("A=M")
	t.emit("D=M")
}

func (t *Translator) emitAddressToD(address string) {
	t.emit(fmt.Sprintf("@%v", address))
	t.emit("D=M")
}

func (t *Translator) emitConstantToD(constant int) {
	t.emit(fmt.Sprintf("@%v", constant))
	t.emit("D=A")
}

func (t *Translator) emitDToAddress(address string) {
	t.emit(fmt.Sprintf("@%v", address))
	t.emit("M=D")
}

func (t *Translator) emitDToStack() {
	t.emits("@SP", "A=M", "M=D")
}

func (t *Translator) emitComputeOneArg(operation string) {
	t.emit("@SP")
	t.emit("A=M-1")
	t.emit(fmt.Sprintf("M=%v", operation))
}

func (t *Translator) emitComputeTwoArgs(operation string) {
	t.emits("@SP", "AM=M-1", "D=M", "A=A-1")
	t.emit(fmt.Sprintf("M=%v", operation))
}

func (t *Translator) emitComputeComparable(operation string) {
	t.emitComputeTwoArgs("M-D")
	t.emit("D=M")

	id := fmt.Sprintf("%v_%v_%v", operation, t.filename, t.labels)
	start := fmt.Sprintf("START_%v", id)
	end := fmt.Sprintf("END_%v", id)
	t.emitLabelAddress(start)
	t.emit(fmt.Sprintf("D;%v", operation))
	t.emit("@SP")
	t.emit("A=M-1")
	t.emit("M=0") // false
	t.emitLabelAddress(end)
	t.emit("0;JMP")

	t.emitLabel(start)
	t.emit("@SP")
	t.emit("A=M-1")
	t.emit("M=-1") // true
	t.emitLabel(end)

	t.labels++
}

func (t *Translator) emitComment(comment string) {
	t.emit(fmt.Sprintf("// %v", comment))
}

func (t *Translator) emit(line string) {
	t.out <- line
}

func (t *Translator) emits(lines ...string) {
	for _, line := range lines {
		t.out <- line
	}
}

func pointerOffset(offset int) (string, error) {
	switch offset {
	case 0:
		return "THIS", nil
	case 1:
		return "THAT", nil
	default:
		return "", fmt.Errorf("unknown pointer offset %v", offset)
	}
}

func segmentAddress(segment string) (string, error) {
	switch segment {
	case "local":
		return "LCL", nil
	case "argument":
		return "ARG", nil
	case "this":
		return "THIS", nil
	case "that":
		return "THAT", nil
	default:
		return "", fmt.Errorf("unknown segment name %v", segment)
	}
}
