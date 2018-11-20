package main

import (
	"flag"
	"log"
	"os"

	"github.com/L-P/mme/rom"
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		os.Exit(1)
	}

	romPath := flag.Args()[0]

	_, err := rom.New(romPath)
	if err != nil {
		log.Fatal(err)
	}
}
