package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/leonhfr/nand2tetris/src/hasm/parser"
)

var config = NewConfig()

// Executes executes the root command.
func Execute() error {
	if config.input == "" {
		return fmt.Errorf("input file not defined")
	}
	p, err := parser.New(config.input)
	if err != nil {
		return err
	}
	err = p.Parse()
	if err != nil {
		return err
	}
	// st := symboltable.New()
	// comds := p.Commands()
	// for _, c := range  {
	// }
	return nil
}

func init() {
	flag.Usage = usage

	flag.StringVar(&config.input, "input", "", inputUsage)
	flag.StringVar(&config.output, "output", "", outputUsage)

	v := flag.Bool("version", versionDefault, versionUsage)
	flag.Parse()

	if *v {
		version()
	}
}

func usage() {
	fmt.Println("hasm is an hack machine language assembler")
	fmt.Printf("Usage: %s [OPTIONS] <path>\n", os.Args[0])
	flag.PrintDefaults()
}

func version() {
	fmt.Printf("current hasm version: %s\n", config.version)
	os.Exit(0)
}
