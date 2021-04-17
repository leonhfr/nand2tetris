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
// function Class1.set 0
(CALL_FUNCTION_Class1.set)
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
// pop static 0
@SP
M=M-1
@SP
A=M
D=M
@StaticsTest.0
M=D
// push argument 1
@ARG
D=M
@1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// pop static 1
@SP
M=M-1
@SP
A=M
D=M
@StaticsTest.1
M=D
// push constant 0
@0
D=A
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
// function Class1.get 0
(CALL_FUNCTION_Class1.get)
// push static 0
@StaticsTest.0
D=M
@SP
A=M
M=D
@SP
M=M+1
// push static 1
@StaticsTest.1
D=M
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
// function Class2.set 0
(CALL_FUNCTION_Class2.set)
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
// pop static 0
@SP
M=M-1
@SP
A=M
D=M
@StaticsTest.0
M=D
// push argument 1
@ARG
D=M
@1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// pop static 1
@SP
M=M-1
@SP
A=M
D=M
@StaticsTest.1
M=D
// push constant 0
@0
D=A
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
// function Class2.get 0
(CALL_FUNCTION_Class2.get)
// push static 0
@StaticsTest.0
D=M
@SP
A=M
M=D
@SP
M=M+1
// push static 1
@StaticsTest.1
D=M
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
// push constant 6
@6
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 8
@8
D=A
@SP
A=M
M=D
@SP
M=M+1
// call Class1.set 2
// call: store return address
@CALL_RETURN_Class1.set_1
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
@2
D=D-A
@ARG
M=D
// call: set up LCL
@SP
D=M
@LCL
M=D
// call: jump to the function
@CALL_FUNCTION_Class1.set
0;JMP
// call: return address
(CALL_RETURN_Class1.set_1)
// pop temp 0
@5
D=A
@0
D=D+A
@R13
M=D
@SP
M=M-1
@SP
A=M
D=M
@R13
A=M
M=D
// push constant 23
@23
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 15
@15
D=A
@SP
A=M
M=D
@SP
M=M+1
// call Class2.set 2
// call: store return address
@CALL_RETURN_Class2.set_2
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
@2
D=D-A
@ARG
M=D
// call: set up LCL
@SP
D=M
@LCL
M=D
// call: jump to the function
@CALL_FUNCTION_Class2.set
0;JMP
// call: return address
(CALL_RETURN_Class2.set_2)
// pop temp 0
@5
D=A
@0
D=D+A
@R13
M=D
@SP
M=M-1
@SP
A=M
D=M
@R13
A=M
M=D
// call Class1.get 0
// call: store return address
@CALL_RETURN_Class1.get_3
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
@CALL_FUNCTION_Class1.get
0;JMP
// call: return address
(CALL_RETURN_Class1.get_3)
// call Class2.get 0
// call: store return address
@CALL_RETURN_Class2.get_4
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
@CALL_FUNCTION_Class2.get
0;JMP
// call: return address
(CALL_RETURN_Class2.get_4)
// label WHILE
(WHILE)
// goto WHILE
@WHILE
0;JMP
