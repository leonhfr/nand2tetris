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

	r := filesystem.NewReader(config.input, config.directory, lines, errors)
	p := parser.New(lines, commands, errors)
	t := translator.New(config.filename, config.directory, commands, asm, errors)
	w := filesystem.NewWriter(config.output, asm, errors, done)

	go r.Read()
	go p.Parse()
	go t.Translate()
	go w.Write()

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

	if *v {
		version()
		return
	}

	s, err := os.Stat(config.input)
	if err != nil {
		panic(err)
	}

	ext := path.Ext(config.input)
	base := filepath.Base(config.input)
	config.filename = base[0 : len(base)-len(ext)]

	if s.IsDir() {
		config.directory = true
		config.output = path.Join(config.input, base+".asm")
		return
	}

	if ext != ".vm" {
		panic(fmt.Errorf("input is neither a directory nor a vm file"))
	}

	config.output = config.input[0:len(config.input)-len(ext)] + ".asm"
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
