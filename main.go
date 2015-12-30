package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func usage() {
	os.Stderr.WriteString(`
Usage: circlecrossgame [OPTION]...
Play circle cross game on terminal.

Options:
	--help       show this help message
	--version    print the version
`[1:])
}

func version() {
	os.Stderr.WriteString(`
0.2.0
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
		usage()
		return nil
	case *isVersion:
		version()
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
