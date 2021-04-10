package main

import (
	"fmt"
	"os"

	"github.com/leonhfr/nand2tetris/src/hasm/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		if s := err.Error(); s != "" {
			fmt.Printf("hasm: %s\n", s)
		}
		os.Exit(1)
	}
}
