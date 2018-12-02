package rom

import (
	"encoding/binary"
	"io"
	"strings"
)

// InternalSceneTableEntry is a single entry of the InternalSceneTable
// Sources:
// - https://wiki.cloudmodding.com/mm/Scene_Table
// binpacked, do not change struct size
type InternalSceneTableEntry struct {
	VROMStart          uint32
	VROMEnd            uint32
	EntranceMessageID  uint16
	_                  uint8
	SceneConfiguration uint8
	_                  uint32
}

// Scene holds a Scene headers and contents. A Scene contents size depends on
// where the 0x14 header end marker was found.
// Sources:
// - https://wiki.cloudmodding.com/mm/Scenes_and_Rooms#Header_Commands
type Scene struct {
	InternalSceneTableEntry
	LocationHeader

	Rooms []Room

	Name            string
	EntranceMessage string
	Valid           bool   // Is the scene valid (has data)
	DataStartOffset uint32 // ROM offset to the scene data

	data []byte
}

var sceneHeaderEndCommand byte = 0x14

func (s *Scene) load(r io.ReadSeeker, entry InternalSceneTableEntry) {
	s.InternalSceneTableEntry = entry
	if entry.VROMStart == 0 && entry.VROMEnd == 0 {
		return
	}

	s.Valid = true
	s.Name = FileNames[entry.VROMStart]

	s.DataStartOffset = s.LocationHeader.load(r, entry.VROMStart)

	size := entry.VROMEnd - s.DataStartOffset
	s.data = make([]byte, size, size)
	binary.Read(r, binary.BigEndian, s.data)
}

func (s *Scene) loadRooms(r io.ReadSeeker) {
	s.Rooms = make([]Room, s.RoomsCount, s.RoomsCount)
	if len(s.Rooms) <= 0 {
		return
	}

	listOffset := s.RoomsSegmentOffset & 0x00FFFFFF // ditch 0x02
	r.Seek(int64(s.VROMStart+listOffset), io.SeekStart)

	var start uint32
	for i := byte(0); i < s.RoomsCount; i++ {
		binary.Read(r, binary.BigEndian, &start)

		s.Rooms[i] = Room{
			ID:             i,
			VROMStart:      start,
			SceneName:      strings.Join([]string{s.Name, s.EntranceMessage}, " - "),
			SceneVROMStart: s.VROMStart,
		}
	}

	// Load them after because we need to keep our reader seek position.
	for k := range s.Rooms {
		s.Rooms[k].load(r)
	}
}
