package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/leonhfr/nand2tetris/src/hasm/parser"
)

var config = NewConfig()

// Executes executes the root command.
func Execute() error {
	if config.input == "" {
		return fmt.Errorf("input file not defined")
	}
	bytes, err := ioutil.ReadFile(config.input)
	if err != nil {
		return err
	}
	content := string(bytes)

	parsed := parser.Parse(config.input, content)
	pretty, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(pretty))
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
