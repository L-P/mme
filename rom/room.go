package rom

import (
	"encoding/binary"
	"io"
)

// A Room is a room within a Scene
type Room struct {
	ID              byte
	VROMStart       uint32
	DataStartOffset uint32 // VROM offset to the Room data
	LocationHeader

	SceneName      string
	SceneVROMStart uint32 // VROM offset the the Scene this Room belongs to

	data []byte
}

func (r *Room) load(rs io.ReadSeeker) {
	if r.VROMStart == 0 {
		return
	}

	r.DataStartOffset = r.LocationHeader.load(rs, r.VROMStart)
}

func (r *Room) loadData(rs io.ReadSeeker, end uint32) {
	size := end - r.DataStartOffset
	r.data = make([]byte, size, size)
	binary.Read(rs, binary.BigEndian, r.data)
}
