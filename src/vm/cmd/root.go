package cmd

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/leonhfr/nand2tetris/src/vm/bytecode"
	"github.com/leonhfr/nand2tetris/src/vm/filesystem"
	"github.com/leonhfr/nand2tetris/src/vm/parser"
	"github.com/leonhfr/nand2tetris/src/vm/translator"
)

var config = NewConfig()

func Execute() error {
	if config.input == "" {
		return fmt.Errorf("input file not defined")
	}

	lines := make(chan string)
	commands := make(chan *bytecode.Command)
	asm := make(chan string)
	errors := make(chan error)
	done := make(chan bool)

	r := filesystem.NewReader(config.input, lines, errors)
	p := parser.New(lines, commands, errors)
	t := translator.New(config.filename, commands, asm, errors)
	w := filesystem.NewWriter(config.output, asm, errors, done)

	go r.Lines()
	go p.Parse()
	go t.Translate()
	go w.Lines()

	for {
		select {
		case err := <-errors:
			if err != nil {
				fmt.Println(err)
			}
		case <-done:
			return nil
		}
	}
}

func init() {
	flag.Usage = usage

	flag.StringVar(&config.input, "input", "", inputUsage)

	v := flag.Bool("version", versionDefault, versionUsage)
	flag.Parse()

	ext := path.Ext(config.input)
	base := filepath.Base(config.input)
	config.output = config.input[0:len(config.input)-len(ext)] + ".asm"
	config.filename = base[0 : len(base)-len(ext)]

	if *v {
		version()
	}
}

func usage() {
	fmt.Println("vm is a vm translator")
	fmt.Printf("Usage: %s [OPTIONS] <path>\n", os.Args[0])
	flag.PrintDefaults()
}

func version() {
	fmt.Printf("current vm version: %s\n", config.version)
	os.Exit(0)
}
