package translator

import (
	"fmt"

	"github.com/leonhfr/nand2tetris/src/vm/bytecode"
)

type Translator struct {
	filename string
	in       <-chan *bytecode.Command
	out      chan string
	errors   chan error
	labels   int
}

func New(filename string, in <-chan *bytecode.Command, out chan string, errors chan error) *Translator {
	labels := 0
	return &Translator{filename, in, out, errors, labels}
}

func (t *Translator) Translate() {
	for command := range t.in {
		t.translate(command)
	}
	close(t.out)
}

func (t *Translator) translate(command *bytecode.Command) {
	t.emitComment(command)
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
		t.emit("@R13")
		t.emit("A=M")
		t.emit("M=D")
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
		t.emit("@R13")
		t.emit("A=M")
		t.emit("M=D")
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
	t.emit("@SP")
	t.emit("A=M")
	t.emit("M=D")
}

func (t *Translator) emitComputeOneArg(operation string) {
	t.emit("@SP")
	t.emit("A=M-1")
	t.emit(fmt.Sprintf("M=%v", operation))
}

func (t *Translator) emitComputeTwoArgs(operation string) {
	t.emit("@SP")
	t.emit("AM=M-1")
	t.emit("D=M")
	t.emit("A=A-1")
	t.emit(fmt.Sprintf("M=%v", operation))
}

func (t *Translator) emitComputeComparable(operation string) {
	t.emitComputeTwoArgs("M-D")
	t.emit("D=M")
	t.emit(fmt.Sprintf("@%v.%v.%v", t.filename, operation, t.labels))
	t.emit(fmt.Sprintf("D;%v", operation))
	t.emit("@SP")
	t.emit("A=M-1")
	t.emit("M=0") // false
	t.emit(fmt.Sprintf("@%v.%v.%v.END", t.filename, operation, t.labels))
	t.emit("0;JMP")
	t.emit(fmt.Sprintf("(%v.%v.%v)", t.filename, operation, t.labels))
	t.emit("@SP")
	t.emit("A=M-1")
	t.emit("M=-1") // true
	t.emit(fmt.Sprintf("(%v.%v.%v.END)", t.filename, operation, t.labels))
	t.labels++
}

func (t *Translator) emitComment(command *bytecode.Command) {
	t.emit(fmt.Sprintf("// %v", command.Original))
}

func (t *Translator) emit(line string) {
	t.out <- line
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
