package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tombell/moniker"
)

const helpText = `usage: moniker [options] <directory>

Format options:
  --format   specify format of the file name to be changed to

Special options:
  --help     show this message, then exit
  --version  show the version number, then exit
`

var (
	format  = flag.String("format", "{artist} - {name}", "")
	version = flag.Bool("version", false, "")
)

func usage() {
	fmt.Fprintf(os.Stderr, helpText)
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "moniker %s (%s)\n", Version, Commit)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}

	if err := moniker.Run(args[0], *format); err != nil {
		fmt.Println(err)
	}
}
