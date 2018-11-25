package rom

import (
	"encoding/binary"
	"io"
)

// DMAEntry is a single entry of the filesystem table
// binpacked, do not change struct size
type DMAEntry struct {
	// Virtual (or physical when uncompressed)
	VROMStart uint32
	VROMEnd   uint32

	// Physical (when compressed)
	PROMStart uint32
	PROMEnd   uint32
}

// A File is anything referenced in the damadata section of the rom.
type File struct {
	DMAEntry

	Name  string
	Valid bool
	data  []byte
}

// Data returns the RAW file data
func (f *File) Data() []byte {
	return f.data
}

// Size returns the raw file data size
func (f *File) Size() int {
	return len(f.data)
}

func (f *File) load(r io.ReadSeeker, entry DMAEntry) {
	if entry.PROMStart == 0xFFFFFFFF || entry.PROMEnd == 0xFFFFFFFF {
		return
	}

	f.Name, _ = FileNames[entry.VROMStart]
	f.Valid = true

	f.DMAEntry = entry
	size := int64(f.VROMEnd - f.VROMStart)
	f.data = make([]byte, size, size)

	// As we're working on a manually decompressed ROM, PROMStart _should_
	// always be the same as VROMStart, but just in case, do the right thing
	// and start at the advertised physical ROM offset.
	// Compression is pretty straightforward, working on a compressed ROM
	// should not be too difficult…
	r.Seek(int64(f.PROMStart), io.SeekStart)
	binary.Read(r, binary.BigEndian, &f.data)
}
