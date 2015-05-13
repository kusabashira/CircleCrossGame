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
v0.1.0
`[1:])
}

func _main() (int, error) {
	isHelp := flag.Bool("help", false, "")
	isVersion := flag.Bool("version", false, "")
	flag.Usage = usage
	flag.Parse()
	switch {
	case *isHelp:
		usage()
		return 2, nil
	case *isVersion:
		version()
		return 2, nil
	}

	g := NewGame()
	if err := g.Init(); err != nil {
		return 1, err
	}
	if err := g.Main(); err != nil {
		return 1, err
	}
	return 0, nil
}

func main() {
	e, err := _main()
	if err != nil {
		fmt.Fprintln(os.Stderr, "circlecrossgame:", err)
	}
	os.Exit(e)
}
