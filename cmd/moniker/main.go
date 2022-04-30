package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tombell/moniker"
)

// TODO: add verbose flag for more logging output during a run...

const helpText = `usage: moniker [options] <directory>

Format options:
  --format   specify format of the file name to be changed to

    the available formatting tokens are:

     - {album}
     - {artist}
     - {genre}
     - {title}
     - {year}

Special options:
  --help     show this message, then exit
  --version  show the version number, then exit`

var (
	format = flag.String("format", "{artist} - {title}", "")
	vrsn   = flag.Bool("version", false, "")
)

func usage() {
	fmt.Fprintln(os.Stderr, helpText)
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *vrsn {
		fmt.Fprintf(os.Stdout, "moniker %s (%s)\n", version, commit)
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
