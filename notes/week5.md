# Week 5: Computer Architecture

## Von Neumann Architecture

- Universality: same hardware can run many different software programs

Elements:
- CPU = Registers + ALU
- Memory = Data + Program

Information flows
- Data bus
- Address bus
- Control bus

## The Fetch-Execute Cycle

The basic CPU loop
- fetch an instruction from the program memory
- execute it

Fetching
- put the location of the next instruction into the address of the program memory
- get the instruction code itself by reading the memory contents at that location

The Program Counter
- normally, old value + 1
- jump achieved by loading another value

Executing
- the instruction code specifies what to do
  - which arithmetic or logical instruction
  - what memory to access (read/write)
  - if/where to jump
- often, different subsets of the buts control different aspects of the operation
- executing the operation involves also accessing registers and/or data memory

We have a clash between the fetch cycle and the execute cycle
- solution: multiplex

Simpler solution: Harvard Architecture
- variant of von Neumann architecture
- keep program and data in two separate memory modules
- complication avoided

## Central Processing Unit (CPU)

- centerpiece of computation
- seat of control

Hack CPU = black box that performs two things:
- 16-bit processor, can execute an instruction written in the Hack language
- figure out which instruction to execute next

Hack CPU interface