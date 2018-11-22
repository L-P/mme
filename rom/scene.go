package rom

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

// InternalSceneTableEntry is a single entry of the InternalSceneTable
// Sources:
// - https://wiki.cloudmodding.com/mm/Scene_Table
type InternalSceneTableEntry struct {
	VROMStart          uint32
	VROMEnd            uint32
	EntranceMessageID  uint16
	Padding0           uint8
	SceneConfiguration uint8
	Padding1           uint32
}

func (e *InternalSceneTableEntry) validate() error {
	if e.Padding0 != 0 {
		return fmt.Errorf("Padding0 is not 0: %X", e.Padding0)
	}

	if e.Padding1 != 0 {
		return fmt.Errorf("Padding1 is not 0: %X", e.Padding1)
	}

	return nil
}

// Scene holds a Scene headers and contents. A Scene contents size depends on
// where the 0x14 header end marker was found.
// Sources:
// - https://wiki.cloudmodding.com/mm/Scenes_and_Rooms#Header_Commands
type Scene struct {
	// Headers are two uint32 with the first byte determining what the data is,
	// the data here is already interpreted, eg. for a header described as
	// "0x12xx5678 0x0000yyyy" and fed the data 0x123456789 0x123456789 its
	// first value would be 0x34 and its second one would be 0x6789
	//
	// Headers marked with an asterisk are mandatory.
	//
	//                     second header value ─────────────────────┐
	//                        first header value ───────────┬┐      │
	//                              header command byte ──┬┐││      │
	//                                                    ││││      │
	StartPositionsCount                       byte   // 0x00xx0000* │
	StartPositionsSegmentOffset               uint32 // 0xyyyyyyyy  ┘
	CamerasCount                              byte   // 0x02xx0000
	CamerasSegmentOffset                      uint32 // 0xyyyyyyyy
	CollisionHeaderSegmentOffset              uint32 // 0x03000000 0xyyyyyyyy *
	RoomsCount                                byte   // 0x04xx0000 *
	RoomsSegmentOffset                        uint32 // 0xyyyyyyyy
	EntrancesCount                            byte   // 0x06xx0000 *
	EntrancesSegmentOffset                    uint32 // 0xyyyyyyyy
	SpecialObjectsByte0                       byte   // 0x07??0000 unknown purpose *
	SpecialObjects                            uint16 // 0x0000xxxx
	LightSettingsCount                        byte   // 0x0Cxx0000
	LightSettingsSegmentOffset                uint32 // 0xyyyyyyyy
	PathsSegmentOffset                        uint32 // 0x0D000000 0xyyyyyyyy
	ActorTransitionsCount                     byte   // 0x0Exx0000
	ActorTransitionsSegmentOffset             uint32 // 0xyyyyyyyy
	EnvironmentSettingsCount                  byte   // 0x0Fxx0000 *
	EnvironmentSettingsSegmentOffset          uint32 // 0xyyyyyyyy
	SkyboxNumber                              byte   // 0x11000000 0xxx0y0z00 *
	SkyboxCast                                byte   // ────────────────┘ │
	SkyboxFog                                 byte   // ──────────────────┘
	ExitsSegmentOffset                        uint32 // 0x13000000 0xyyyyyyyy
	SoundReverb                               byte   // 0x15xx0000 0x0000yyzz *
	SoundNightSFX                             byte   // ─────────────────┴┘││
	SoundBackgroundSequence                   byte   // ───────────────────┴┘
	CutscenesCount                            byte   // 0x17xx0000
	CutscenesSegmentOffset                    uint32 // 0xyyyyyyyy
	AlternateHeadersSegmentOffset             uint32 // 0x18000000 0xxxxxxxxx
	WorldMapLocation                          bool   // 0x19000000 0x00000000 (presence = true)
	TextureAnimationsSegmentOffset            uint32 // 0x1A000000 0xxxxxxxxx *
	CamerasAndCutscenesForActorsCount         byte   // 0x1Bxx0000 *
	CamerasAndCutscenesForActorsSegmentOffset uint32 // 0xyyyyyyyy
	MinimapsSegmentOffset                     uint32 // 0x1C000000 0xxxxxxxxx *
	ChestPositionsCount                       byte   // 0x1Exx0000
	ChestPositionsSegmentOffset               uint32 // 0xyyyyyyyy

	Valid           bool  // Is the scene valid (has data)
	DataStartOffset int64 // ROM offset the scene data
	data            []byte
}

var sceneHeaderEndCommand byte = 0x14

func (s *Scene) load(r io.ReadSeeker, start uint32, end uint32) {
	if start == 0 && end == 0 {
		return
	}
	s.Valid = true

	r.Seek(int64(start), io.SeekStart)
	s.DataStartOffset = int64(start)

	var a, b uint32
	for {
		binary.Read(r, binary.BigEndian, &a)
		binary.Read(r, binary.BigEndian, &b)
		s.DataStartOffset += 8

		command := byte((a & 0xFF000000) >> 24)
		if command == sceneHeaderEndCommand {
			break
		}

		if err := s.loadHeader(command, a, b); err != nil {
			log.Printf(
				"At offset 0x%08X (Scene at 0x%08X): %s",
				s.DataStartOffset-8,
				start,
				err,
			)
		}
	}

	size := int64(end) - s.DataStartOffset
	s.data = make([]byte, size, size)
	binary.Read(r, binary.BigEndian, s.data)
}

func (s *Scene) loadHeader(command byte, a uint32, b uint32) error {
	switch command {
	default:
		return fmt.Errorf("Unknown Scene header command 0x%02X (0x%08X 0x%08X)", command, a, b)
	case 0x00:
		s.StartPositionsCount = byte((a & 0x00FF0000) >> 16)
		s.StartPositionsSegmentOffset = b
	case 0x02:
		s.CamerasCount = byte((a & 0x00FF0000) >> 16)
		s.CamerasSegmentOffset = b
	case 0x03:
		s.CollisionHeaderSegmentOffset = b
	case 0x04:
		s.RoomsCount = byte((a & 0x00FF0000) >> 16)
		s.RoomsSegmentOffset = b
	case 0x06:
		s.EntrancesCount = byte((a & 0x00FF0000) >> 16)
		s.EntrancesSegmentOffset = b
	case 0x07:
		s.SpecialObjectsByte0 = byte((a & 0x00FF0000) >> 16)
		s.SpecialObjects = uint16(b & 0x0000FFFF)
	case 0x0C:
		s.LightSettingsCount = byte((a & 0x00FF0000) >> 16)
		s.LightSettingsSegmentOffset = b
	case 0x0D:
		s.PathsSegmentOffset = b
	case 0x0E:
		s.ActorTransitionsCount = byte((a & 0x00FF0000) >> 16)
		s.ActorTransitionsSegmentOffset = b
	case 0x0F:
		s.EnvironmentSettingsCount = byte((a & 0x00FF0000) >> 16)
		s.EnvironmentSettingsSegmentOffset = b
	case 0x11:
		s.SkyboxNumber = byte((b & 0xFF000000) >> 24)
		s.SkyboxCast = byte((b & 0x000F0000) >> 16)
		s.SkyboxFog = byte((b & 0x00000F00) >> 8)
	case 0x13:
		s.ExitsSegmentOffset = b
	case 0x15:
		s.SoundReverb = byte((a & 0x00FF0000) >> 16)
		s.SoundNightSFX = byte((b & 0x0000FF00) >> 8)
		s.SoundBackgroundSequence = byte(b & 0x000000FF)
	case 0x17:
		s.CutscenesCount = byte((a & 0x00FF0000) >> 16)
		s.CutscenesSegmentOffset = b
	case 0x18:
		s.AlternateHeadersSegmentOffset = b
	case 0x19:
		s.WorldMapLocation = true
	case 0x1A:
		s.TextureAnimationsSegmentOffset = b
	case 0x1B:
		s.CamerasAndCutscenesForActorsCount = byte((a & 0x00FF0000) >> 16)
		s.CamerasAndCutscenesForActorsSegmentOffset = b
	case 0x1C:
		s.MinimapsSegmentOffset = b
	case 0x1E:
		s.ChestPositionsCount = byte((a & 0x00FF0000) >> 16)
		s.ChestPositionsSegmentOffset = b
	}

	return nil
}
