package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	name    = "circlecrossgame"
	version = "0.2.1"

	flagset   = flag.NewFlagSet(name, flag.ContinueOnError)
	isHelp    = flagset.Bool("help", false, "")
	isVersion = flagset.Bool("version", false, "")
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [OPTION]...
Play circle cross game on terminal.

Options:
	--help       show this help message
	--version    print the version
`[1:], name)
}

func printVersion() {
	fmt.Fprintln(os.Stderr, version)
}

func _main() error {
	flagset.SetOutput(ioutil.Discard)
	if err := flagset.Parse(os.Args[1:]); err != nil {
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
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
		os.Exit(1)
	}
}
