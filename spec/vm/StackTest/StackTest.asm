// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
AM=M-1
D=M
A=A-1
M=M-D
D=M
@StackTest.JEQ.0
D;JEQ
@SP
A=M-1
M=0
@StackTest.JEQ.0.END
0;JMP
(StackTest.JEQ.0)
@SP
A=M-1
M=-1
(StackTest.JEQ.0.END)
// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
AM=M-1
D=M
A=A-1
M=M-D
D=M
@StackTest.JEQ.1
D;JEQ
@SP
A=M-1
M=0
@StackTest.JEQ.1.END
0;JMP
(StackTest.JEQ.1)
@SP
A=M-1
M=-1
(StackTest.JEQ.1.END)
// push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
AM=M-1
D=M
A=A-1
M=M-D
D=M
@StackTest.JEQ.2
D;JEQ
@SP
A=M-1
M=0
@StackTest.JEQ.2.END
0;JMP
(StackTest.JEQ.2)
@SP
A=M-1
M=-1
(StackTest.JEQ.2.END)
// push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
AM=M-1
D=M
A=A-1
M=M-D
D=M
@StackTest.JLT.3
D;JLT
@SP
A=M-1
M=0
@StackTest.JLT.3.END
0;JMP
(StackTest.JLT.3)
@SP
A=M-1
M=-1
(StackTest.JLT.3.END)
// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
AM=M-1
D=M
A=A-1
M=M-D
D=M
@StackTest.JLT.4
D;JLT
@SP
A=M-1
M=0
@StackTest.JLT.4.END
0;JMP
(StackTest.JLT.4)
@SP
A=M-1
M=-1
(StackTest.JLT.4.END)
// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@SP
AM=M-1
D=M
A=A-1
M=M-D
D=M
@StackTest.JLT.5
D;JLT
@SP
A=M-1
M=0
@StackTest.JLT.5.END
0;JMP
(StackTest.JLT.5)
@SP
A=M-1
M=-1
(StackTest.JLT.5.END)
// push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
AM=M-1
D=M
A=A-1
M=M-D
D=M
@StackTest.JGT.6
D;JGT
@SP
A=M-1
M=0
@StackTest.JGT.6.END
0;JMP
(StackTest.JGT.6)
@SP
A=M-1
M=-1
(StackTest.JGT.6.END)
// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
AM=M-1
D=M
A=A-1
M=M-D
D=M
@StackTest.JGT.7
D;JGT
@SP
A=M-1
M=0
@StackTest.JGT.7.END
0;JMP
(StackTest.JGT.7)
@SP
A=M-1
M=-1
(StackTest.JGT.7.END)
// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@SP
AM=M-1
D=M
A=A-1
M=M-D
D=M
@StackTest.JGT.8
D;JGT
@SP
A=M-1
M=0
@StackTest.JGT.8.END
0;JMP
(StackTest.JGT.8)
@SP
A=M-1
M=-1
(StackTest.JGT.8.END)
// push constant 57
@57
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 31
@31
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 53
@53
D=A
@SP
A=M
M=D
@SP
M=M+1
// add
@SP
AM=M-1
D=M
A=A-1
M=D+M
// push constant 112
@112
D=A
@SP
A=M
M=D
@SP
M=M+1
// sub
@SP
AM=M-1
D=M
A=A-1
M=M-D
// neg
@SP
A=M-1
M=-M
// and
@SP
AM=M-1
D=M
A=A-1
M=D&M
// push constant 82
@82
D=A
@SP
A=M
M=D
@SP
M=M+1
// or
@SP
AM=M-1
D=M
A=A-1
M=D|M
// not
@SP
A=M-1
M=!M
