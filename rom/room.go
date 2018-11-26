package rom

import (
	"io"
)

// A Room is a room within a Scene
type Room struct {
	ID              byte
	VROMStart       uint32
	DataStartOffset uint32 // ROM offset to the Room data
	LocationHeader

	data []byte
}

func (r *Room) load(rs io.ReadSeeker) {
	if r.VROMStart == 0 {
		return
	}

	r.DataStartOffset = r.LocationHeader.load(rs, r.VROMStart)

	/*
		size := entry.VROMEnd - s.DataStartOffset
		s.data = make([]byte, size, size)
		binary.Read(r, binary.BigEndian, s.data)
	*/
}
