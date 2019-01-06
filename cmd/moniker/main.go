package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tombell/moniker"
)

// TODO: add details about available format values to helpText...
// TODO: add verbose flag for more logging output during a run...

const helpText = `usage: moniker [options] <directory>

Format options:
  --title    specify that format values should be title casing
  --format   specify format of the file name to be changed to

Special options:
  --help     show this message, then exit
  --version  show the version number, then exit
`

var (
	title   = flag.Bool("title", false, "")
	format  = flag.String("format", "{artist} - {title}", "")
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
