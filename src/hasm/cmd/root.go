package cmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/leonhfr/nand2tetris/src/hasm/parser"
	"github.com/leonhfr/nand2tetris/src/hasm/symboltable"
)

var config = NewConfig()

// Executes executes the root command.
func Execute() error {
	if config.input == "" {
		return fmt.Errorf("input file not defined")
	}
	if config.output == "" {
		return fmt.Errorf("output file not defined")
	}
	bytes, err := ioutil.ReadFile(config.input)
	if err != nil {
		return err
	}
	content := string(bytes)

	parsed := parser.Parse(config.input, content)
	// pretty, err := json.MarshalIndent(parsed, "", "  ")
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(pretty))
	st := symboltable.New()

	err = parsed.FirstPass(st)
	if err != nil {
		return err
	}

	output, err := parsed.SecondPass(st)
	if err != nil {
		return err
	}

	// fmt.Println(output)

	writeFile(output)

	return nil
}

func writeFile(content string) error {
	f, err := os.Create(config.output)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
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
