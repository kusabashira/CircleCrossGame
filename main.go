package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func printUsage() {
	os.Stderr.WriteString(`
Usage: circlecrossgame [OPTION]...
Play circle cross game on terminal.

Options:
	--help       show this help message
	--version    print the version
`[1:])
}

func printVersion() {
	os.Stderr.WriteString(`
0.2.1
`[1:])
}

func _main() error {
	f := flag.NewFlagSet("circlecrossgame:", flag.ContinueOnError)
	f.SetOutput(ioutil.Discard)

	isHelp := f.Bool("help", false, "")
	isVersion := f.Bool("version", false, "")
	if err := f.Parse(os.Args[1:]); err != nil {
		return err
	}
	switch {
	case *isHelp:
		printUsage()
		return nil
	case *isVersion:
		printVersion()
		return nil
	}

	g := NewGame()
	if err := g.Init(); err != nil {
		return err
	}
	return g.Main()
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintln(os.Stderr, "circlecrossgame:", err)
		os.Exit(1)
	}
}
