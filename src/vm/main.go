package main

import (
	"fmt"
	"os"

	"github.com/leonhfr/nand2tetris/src/vm/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		if s := err.Error(); s != "" {
			fmt.Printf("vm: %s\n", s)
		}
		os.Exit(1)
	}
}

// open file
// test if dir, error
// open output file
// read file line bine line, stream to a channel
// lexer/parser receives the channel, outputs formatted objects to another channel
// dont stream if comment
// codewriter receives the objects and emits the translated commands
// code writer receives the objects and write to a file
// also write original command as comment
// close channels
