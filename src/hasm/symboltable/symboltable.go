package symboltable

import (
	"fmt"
)

var predefinedSymbols = map[string]int{
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16384,
	"KBD":    24576,
}

// SymbolTable the symbol table
type SymbolTable struct {
	entries map[string]int
	next    int
}

// New creates a new symbol table
func New() *SymbolTable {
	st := &SymbolTable{
		make(map[string]int),
		16,
	}
	for symbol, address := range predefinedSymbols {
		st.Insert(symbol, address)
	}
	return st
}

// Add adds a new symbol to the symbol table and returns its address
func (st *SymbolTable) Add(symbol string) (int, error) {
	ok := st.Has(symbol)
	if ok {
		return 0, fmt.Errorf("the symbol table already contains the given symbol")
	}
	index := st.next
	st.entries[symbol] = index
	st.next++
	return index, nil
}

// Insert adds a new entry to the symbol table
func (st *SymbolTable) Insert(symbol string, address int) error {
	ok := st.Has(symbol)
	if ok {
		return fmt.Errorf("the symbol table already contains the given symbol")
	}
	st.entries[symbol] = address
	return nil
}

// Has returns true if the symbol table already contains the symbol
func (st *SymbolTable) Has(symbol string) bool {
	_, ok := st.entries[symbol]
	return ok
}

// Get returns the address of the given symbol
func (st *SymbolTable) Get(symbol string) (int, error) {
	address, ok := st.entries[symbol]
	if !ok {
		return 0, fmt.Errorf("the symbol table does not contain the given symbol")
	}
	return address, nil
}

func (st *SymbolTable) String() string {
	formatted := ""
	formatted += fmt.Sprintln("Symbol table:")
	for k, v := range st.entries {
		formatted += fmt.Sprintln(k, ":", v)
	}
	formatted += fmt.Sprint("Next:", st.next)
	return formatted
}
