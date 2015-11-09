package main

import (
	"flag"
	"fmt"
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
0.1.1
`[1:])
}

func _main() error {
	isHelp := flag.Bool("help", false, "")
	isVersion := flag.Bool("version", false, "")
	flag.Usage = usage
	flag.Parse()
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
