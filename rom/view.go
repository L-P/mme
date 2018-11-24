package rom

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dustin/go-humanize"
)

// A View holds a ROM real data and accessors for dynamically placed data.
type View struct {
	Files  []File
	Scenes []Scene

	rom *ROM
	fd  *os.File
}

// NewView creates a new view from a ROM
func NewView(path string) (*View, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open ROM: %s", err)
	}

	stat, err := fd.Stat()
	if err != nil {
		fd.Close()
		return nil, err
	}

	if stat.Size() != Size {
		fd.Close()
		return nil, fmt.Errorf("expected %d bytes of ROM data, got %d", Size, stat.Size())
	}

	rom, err := New(fd)
	if err != nil {
		fd.Close()
		return nil, err
	}

	v := &View{
		rom:    rom,
		Scenes: make([]Scene, len(rom.InternalSceneTable), len(rom.InternalSceneTable)),
		Files:  make([]File, len(rom.DMAData), len(rom.DMAData)),
		fd:     fd,
	}

	if err := v.load(fd); err != nil {
		fd.Close()
		return nil, err
	}

	return v, nil
}

// Close implements io.Closer
func (v *View) Close() {
	v.fd.Close()
}

func (v *View) load(r io.ReadSeeker) error {
	if err := v.loadFiles(r); err != nil {
		return err
	}

	if err := v.loadScenes(r); err != nil {
		return err
	}

	return nil
}

func (v *View) loadFiles(r io.ReadSeeker) error {
	if len(v.Files) != len(v.rom.DMAData) {
		return errors.New("len(v.files) != len (v.rom.DMAData")
	}

	size := 0
	for k, entry := range v.rom.DMAData {
		v.Files[k].load(r, entry)
		size += len(v.Files[k].data)
	}

	log.Printf("Loaded %d Files (%s)", len(v.Files), humanize.IBytes(uint64(size)))

	return nil
}

func (v *View) loadScenes(r io.ReadSeeker) error {
	if len(v.Scenes) != len(v.rom.InternalSceneTable) {
		return errors.New("len(v.scenes) != len (v.rom.InternalSceneTable")
	}

	for k, entry := range v.rom.InternalSceneTable {
		v.Scenes[k].load(r, entry)
	}

	log.Printf("Loaded %d Scenes", len(v.Scenes))

	return nil
}

func (v *View) Read(p []byte) (n int, err error) {
	return v.fd.Read(p)
}

// Seek implements io.Seeker
func (v *View) Seek(offset int64, whence int) (int64, error) {
	return v.fd.Seek(offset, whence)
}
