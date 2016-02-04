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

func printErr(err interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
}

func main() {
	flagset.SetOutput(ioutil.Discard)
	if err := flagset.Parse(os.Args[1:]); err != nil {
		printErr(err)
		os.Exit(2)
	}
	switch {
	case *isHelp:
		printUsage()
		os.Exit(0)
	case *isVersion:
		printVersion()
		os.Exit(0)
	}

	g := NewGame()
	if err := g.Init(); err != nil {
		printErr(err)
		os.Exit(1)
	}
	if err := g.Main(); err != nil {
		printErr(err)
		os.Exit(1)
	}
}
