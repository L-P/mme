package main

import (
	"flag"
	"log"
	"os"

	"github.com/L-P/mme/colormap"
	"github.com/L-P/mme/rom"
)

const colorMapPath = "out.png"

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		os.Exit(1)
	}

	romPath := flag.Args()[0]

	view, err := rom.NewView(romPath)
	if err != nil {
		log.Fatal(err)
	}
	defer view.Close()

	if err := colormap.Generate(colorMapPath, view); err != nil {
		log.Fatal(err)
	}

	log.Printf("Generated color map in %s", colorMapPath)
}
