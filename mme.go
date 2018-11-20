package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		os.Exit(1)
	}

	romPath := flag.Args()[0]

	_, err := NewROM(romPath)
	if err != nil {
		log.Fatal(err)
	}
}
