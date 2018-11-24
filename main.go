package main

import (
	"flag"
	"log"
	"os"

	"github.com/L-P/mme/rom"
	"github.com/L-P/mme/server"
)

const colorMapPath = "out.png"

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Printf("Usage: mme ROM")
		os.Exit(1)
	}

	romPath := flag.Args()[0]

	view, err := rom.NewView(romPath)
	if err != nil {
		log.Fatal(err)
	}
	defer view.Close()

	server := server.New(view)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
