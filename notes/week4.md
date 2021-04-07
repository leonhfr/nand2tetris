# Week 4: Machine Languages

## Overview

- Universality: the same hardware can run many different software programs
- Theory: Universal Turing Machine
- Practice: von Neumann architecture
  - Stored program computer (memory), different programs create different functionalities

- Machine language is a terrible language
- Humans code in high level languages
- Compiler translate high level languages to machine languages

Mnemonics:
- Instruction: 
```
0100010 0011 0010
ADD     R3   R2
```
- Interpretation 1: The symbolic form doesn't really exist but is just a convenient mnemonic to present machine language instructions to humans
- Interpretation 2: We will allow humans to write machine language instructions using this assembly language ans will have an Assembler program convert it to bit-form

Symbols: access memory locations, holding the index
```
ADD 1 index
```

here index means a fixed memory

## Elements

- Machine language is the most important interface in the whole world of computer science
- Interface between hardware and software; the way the software can control the hardware
- Specifies what operations the hardware performs, where does it get the data it operates on, what is the control of the operations, and so on
- Usually in close correspondence (almost 1:1) to actual hardware architecture
- Cost performance tradeoff:
  - the more sophisticated the operations we want to perform, the more the data structures are complex, the higher the cost
  - costs in term of silicon area
  - cost in time to complete instruction

Machine operations:
- usually correspond to what's implemented in the hardware
  - arithmetic operations: add, subtract
  - logical operations: and, or
  - flow control: goto, if
- differences between languages
  - different sets of operations (divisions, bulk copt, ...)
  - data types (width, floating point, ...)

Memory hierarchy:
- accessing a memory location is expensive
  - need to supply a long address
  - getting the memory contents into the CPU take time
- solution: memory hierarchy
  - instead of having one large block of memory, we have blocks of increasing sizes
  - first memory are small and easy/cheap to access, larger memory is harder to access (longer address)
  - smallest memory: registers (inside the CPU)
  - then cache
  - then RAM

Registers:
- CPU usually contain a few, easily accessed, registers
- their number and functions are a central part of the machine language
- data registers (ADD R1 R2) allows to perform quick computations
- address registers (STORE R1 @A) refers to an address in the larger memory

Addressing modes
- Register (ADD R1,R2): R2 <- R2 + R1
- Direct (ADD R1,M[200]): Mem[200] <- Mem[200] + R1
- Indirect (ADD R1,@A): Mem[A] <- Mem[A] + R1
- Immediate (ADD 73,R1): R1 <- R1 + 73

Input / Output
- many types of input and output devices (keyboard, mouse, camera, sensors, printers, screen, sound, ...)
- CPU needs some kind of protocol to talk to each of them, software drivers know these protocols
- One general method of interaction uses memory mapping
  - memory location 12345 holds the direction of the last movement of the mouse
  - memory location 45678 is not a real memory but a way to tell the printer which paper to use

Flow Control
- usually the CPU executes machine instructions in sequence
- sometimes we need to jump unconditionally to another location, eg so we cna jump
- sometimes we need to jump only if some condition is met

## The Hack Computer and Machine Language

- Particulars of our own Machine Language
- 16 bit machine: everything consists of 16 bits
- Computers with 3 main elements
  - data memory (RAM): a sequence of 16 bit registers (RAM[0], RAM[1], ...)
  - instruction memory (ROM): a sequence of 16 bit registers (ROM[0], ROM[1], ...)
  - central processing unit (CPU): performs 16-bit instructions
  - instruction bus / data bus / address bus

Software:
- how we control the computer
- hack machine language:
  - 16 bit A instructions
  - 16 bit C instructions
- hack program = sequence of instructions written in the hack machine language

Control:
- the ROM is loaded with a hack program
- the reset button is pushed
- the program starts running

Registers:
- the hack machine language recognizes 3 registers:
  - D holds a 16 bit value (in CPU)
  - A holds a 16 bit value (in CPU)
  - M represents the 16 bit RAM register addressed by A (in RAM)

The A instruction
- A for addressing
- Syntax: @value
- Where value is either:
  - a non negative decimal constant or
  - a symbol referring to such a constant (later)
- Semantics:
  - sets the A register to value
  - side effect: RAM[A] becomes the selected RAM register

```
// Set RAM[100] to -1
@100 // A=100
M=-1 // RAM[100]=-1
```

The C instruction
- C for compute
- Syntax: dest = comp ; jump (both dest and jump are optional)
- Where comp
  - is one of a predefined set of computation
- Where dest
  - is one of a set of 8 destinations
- Where jump
  - is one of a set of 8 possible conditions
  - they always compare the result of the computation to 0
- Semantics:
  - compute the value of comp
  - stores the result in dest
  - if the boolean expression (comp jump 0) is true, jumps to execute the instruction stored in ROM[A]

```
// Set the D register to -1
D=-1

// Set RAM[300] to the value of the D register minus 1
@300  // A=300
M=D-1 // RAM[300]=D-1

// If D-1==0, jump to execute instruction stored in ROM[56]
@56     // A=56
D-1;JEQ // if D-1==0 goto 56
```

Unconditional jump: 0;JEQ

## Hack Language Specification

Two ways to express the same semantics:
- binary code
- symbolic language

We need to translate symbolic language to binary code in order to execute it. We use an Assembler to do this.

The A instruction:
- value is either a non negative decimal constant <= 2^15-1
- a symbol referring to such a constant

```
// Symbolic syntax
@value     @21
// Binary syntax
0value     0000000000010101
```

The C instruction:

```
// Symbolic syntax
dest = comp ; jump
// Binary syntax
1 1 1 a c1 c2 c3 c4 c5 c6 d1 d2 d3 j1 j2 j3
// 1 = op code (defines C instruction)
// 1 1 = not used
// a c1 c2 c3 c4 c5 c6 = comp bits
// d1 d2 d3 = dest bits
// j1 j2 j3 = jump bits
```

Hack program:
- A hack program is a sequence of hack instructions
- white space is permitted
- comments are used to explain code
- we use an assembler to translate it to binary code

## Input / Output

- How we use machine language to control these I/O devices
- Peripheral I/O devices:
  - keyboard used to enter inputs
  - screen used to display outputs
- high level approach:
  - sophisticated software libraries enabling text, graphics, animation, audio, video, ...
- low level approach:
  - bits

### Output

Screen memory map:
- a designated memory area, dedicated to manage a display unit
- the physical display is continuously refreshed from the memory map, many times per second
- output is effected by writing code that manipulates the screen memory map

To set pixel (row,col) on/off:
1. word = Screen[32*row + col/16]
2. word = RAM[16384 + 32*row + col/16]
3. Set the (col % 16)th bit of word to 0 or 1
4. Commit word to the RAM

### Input

Keyboard memory map:
- same concept as the screen mm
- key scan code travels through the cord and is displayed in the keyboard mm
  
To check which key is currently pressed:
- probe the contents of the Keyboard chip
- in the Hack computer, probe the contents of RAM[24576]
- if the register contains 0, no key is pressed

## Hack Programming

### Working with registers and memory

- D = data register
- A = address/data register
- M = currently selected memory register M = RAM[A]

- Variables that store memory addresses are called pointers
- Hack pointer logic: whenever we have to access memory using a pointer, we need an instruction like A=M
- Typical pointer semantics: set the address register to the contents of some memory register