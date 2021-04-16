// Entry commands
@256
D=A
@SP
M=D
// call: store return address
@CALL_RETURN_Sys.init_0
D=A
@SP
A=M
M=D
@SP
M=M+1
// call: store call frame
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
// call: compute ARG
@SP
D=M
@5
D=D-A
@0
D=D-A
@ARG
M=D
// call: set up LCL
@SP
D=M
@LCL
M=D
// call: jump to the function
@CALL_FUNCTION_Sys.init
0;JMP
// call: return address
(CALL_RETURN_Sys.init_0)
// function Main.fibonacci 0
(CALL_FUNCTION_Main.fibonacci)
// push argument 0
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// push constant 2
@2
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
@START_JLT_FibonacciElement_1
D;JLT
@SP
A=M-1
M=0
@END_JLT_FibonacciElement_1
0;JMP
(START_JLT_FibonacciElement_1)
@SP
A=M-1
M=-1
(END_JLT_FibonacciElement_1)
// if-goto IF_TRUE
@SP
M=M-1
@SP
A=M
D=M
@IF_TRUE
D;JNE
// goto IF_FALSE
@IF_FALSE
0;JMP
// label IF_TRUE
(IF_TRUE)
// push argument 0
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// return
// return: store addresses of end of call frame and return
@LCL
D=M
@R13
M=D
@5
A=D-A
D=M
@R14
M=D
// return: pop the return value to the start of the frame
@SP
M=M-1
@SP
A=M
D=M
@ARG
A=M
M=D
// return: set SP to parent
@ARG
D=M+1
@SP
M=D
// return: restore the call frame
@R13
D=M
@1
D=D-A
A=D
D=M
@THAT
M=D
@R13
D=M
@2
D=D-A
A=D
D=M
@THIS
M=D
@R13
D=M
@3
D=D-A
A=D
D=M
@ARG
M=D
@R13
D=M
@4
D=D-A
A=D
D=M
@LCL
M=D
// return: jump to return address
@R14
D=M
A=D
0;JMP
// label IF_FALSE
(IF_FALSE)
// push argument 0
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// push constant 2
@2
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
// call Main.fibonacci 1
// call: store return address
@CALL_RETURN_Main.fibonacci_2
D=A
@SP
A=M
M=D
@SP
M=M+1
// call: store call frame
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
// call: compute ARG
@SP
D=M
@5
D=D-A
@1
D=D-A
@ARG
M=D
// call: set up LCL
@SP
D=M
@LCL
M=D
// call: jump to the function
@CALL_FUNCTION_Main.fibonacci
0;JMP
// call: return address
(CALL_RETURN_Main.fibonacci_2)
// push argument 0
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// push constant 1
@1
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
// call Main.fibonacci 1
// call: store return address
@CALL_RETURN_Main.fibonacci_3
D=A
@SP
A=M
M=D
@SP
M=M+1
// call: store call frame
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
// call: compute ARG
@SP
D=M
@5
D=D-A
@1
D=D-A
@ARG
M=D
// call: set up LCL
@SP
D=M
@LCL
M=D
// call: jump to the function
@CALL_FUNCTION_Main.fibonacci
0;JMP
// call: return address
(CALL_RETURN_Main.fibonacci_3)
// add
@SP
AM=M-1
D=M
A=A-1
M=D+M
// return
// return: store addresses of end of call frame and return
@LCL
D=M
@R13
M=D
@5
A=D-A
D=M
@R14
M=D
// return: pop the return value to the start of the frame
@SP
M=M-1
@SP
A=M
D=M
@ARG
A=M
M=D
// return: set SP to parent
@ARG
D=M+1
@SP
M=D
// return: restore the call frame
@R13
D=M
@1
D=D-A
A=D
D=M
@THAT
M=D
@R13
D=M
@2
D=D-A
A=D
D=M
@THIS
M=D
@R13
D=M
@3
D=D-A
A=D
D=M
@ARG
M=D
@R13
D=M
@4
D=D-A
A=D
D=M
@LCL
M=D
// return: jump to return address
@R14
D=M
A=D
0;JMP
// function Sys.init 0
(CALL_FUNCTION_Sys.init)
// push constant 4
@4
D=A
@SP
A=M
M=D
@SP
M=M+1
// call Main.fibonacci 1
// call: store return address
@CALL_RETURN_Main.fibonacci_4
D=A
@SP
A=M
M=D
@SP
M=M+1
// call: store call frame
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
// call: compute ARG
@SP
D=M
@5
D=D-A
@1
D=D-A
@ARG
M=D
// call: set up LCL
@SP
D=M
@LCL
M=D
// call: jump to the function
@CALL_FUNCTION_Main.fibonacci
0;JMP
// call: return address
(CALL_RETURN_Main.fibonacci_4)
// label WHILE
(WHILE)
// goto WHILE
@WHILE
0;JMP
