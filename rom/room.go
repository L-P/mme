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

	ActorList []ActorEntry

	data []byte
}

func (r *Room) load(rs io.ReadSeeker) {
	if r.VROMStart == 0 {
		return
	}

	r.DataStartOffset = r.LocationHeader.load(rs, r.VROMStart)
	r.loadActors(rs)
}

func (r *Room) loadActors(rs io.ReadSeeker) {
	r.ActorList = make([]ActorEntry, r.ActorsCount, r.ActorsCount)
	if r.ActorsCount <= 0 {
		return
	}

	listOffset := r.ActorsSegmentOffset & 0x00FFFFFF // ditch 0x03
	rs.Seek(int64(r.VROMStart+listOffset), io.SeekStart)

	for i := byte(0); i < r.ActorsCount; i++ {
		r.ActorList[i].load(rs)
	}
}

func (r *Room) loadData(rs io.ReadSeeker, end uint32) {
	size := end - r.DataStartOffset
	r.data = make([]byte, size, size)
	binary.Read(rs, binary.BigEndian, r.data)
}
