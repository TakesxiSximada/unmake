package main

import (
	"flag"
	"log"
	"os"
	unmake "github.com/TakesxiSximada/unmake/unmake-go"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("unmake: ")
	log.SetOutput(os.Stderr)

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		os.Exit(1)
	}

	for _, makeFilePath := range args {
		ParseInclude(makeFilePath)
	}
}
