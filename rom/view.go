package rom

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// A View holds a ROM real data and accessors for dynamically placed data.
type View struct {
	rom    *ROM
	scenes []Scene
}

// NewView creates a new view from a ROM
func NewView(path string) (*View, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open ROM: %s", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if stat.Size() != romSize {
		return nil, fmt.Errorf("expected %d bytes of ROM data, got %d", romSize, stat.Size())
	}

	rom, err := New(file)
	if err != nil {
		return nil, err
	}

	v := &View{
		rom:    rom,
		scenes: make([]Scene, len(rom.InternalSceneTable), len(rom.InternalSceneTable)),
	}

	if err := v.load(file); err != nil {
		return nil, err
	}

	return v, nil
}

func (v *View) load(r io.ReadSeeker) error {
	if err := v.loadScenes(r); err != nil {
		return err
	}

	return nil
}

func (v *View) loadScenes(r io.ReadSeeker) error {
	if len(v.scenes) != len(v.rom.InternalSceneTable) {
		return errors.New("len(v.scenes) != len (v.rom.InternalSceneTable")
	}

	for k, entry := range v.rom.InternalSceneTable {
		v.scenes[k].load(r, entry.VROMStart, entry.VROMEnd)
	}

	log.Printf("Loaded %d Scenes", len(v.scenes))

	return nil
}
