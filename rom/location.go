package rom

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

// LocationHeader holds the of both Scenes and Rooms as many of them are shared.
type LocationHeader struct {
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
	ActorsCount                               byte   // 0x01xx0000
	ActorsSegmentOffset                       uint32 // 0xyyyyyyyy
	CamerasCount                              byte   // 0x02xx0000
	CamerasSegmentOffset                      uint32 // 0xyyyyyyyy
	CollisionHeaderSegmentOffset              uint32 // 0x03000000 0xyyyyyyyy *
	RoomsCount                                byte   // 0x04xx0000 *
	WindDirectionX                            byte   // 0x05000000 0xxxyyzzww
	WindDirectionY                            byte   // ───────────────┴┘││││
	WindDirectionZ                            byte   // ─────────────────┴┘││
	WindStrength                              byte   // ───────────────────┴┘
	RoomsSegmentOffset                        uint32 // 0xyyyyyyyy
	EntrancesCount                            byte   // 0x06xx0000 *
	EntrancesSegmentOffset                    uint32 // 0xyyyyyyyy
	SpecialObjectsByte0                       byte   // 0x07??0000 unknown purpose *
	SpecialObjects                            uint16 // 0x0000xxxx
	RoomBehavior0                             byte   // 0x08xx0000 0x0000yyzz "affects Sun's Song, backflipping with A"
	RoomBehavior1                             byte   // ─────────────────┴┘││ unknown
	RoomBehavior2                             byte   // ───────────────────┴┘ animations/tunic
	MeshSegmentOffset                         uint32 // 0x0A000000 0xyyyyyyyy *
	ObjectsCount                              byte   // 0x0Bxx0000
	ObjectsSegmentOffset                      uint32 // 0xyyyyyyyy
	LightSettingsCount                        byte   // 0x0Cxx0000
	LightSettingsSegmentOffset                uint32 // 0xyyyyyyyy
	PathsSegmentOffset                        uint32 // 0x0D000000 0xyyyyyyyy
	ActorTransitionsCount                     byte   // 0x0Exx0000
	ActorTransitionsSegmentOffset             uint32 // 0xyyyyyyyy
	EnvironmentSettingsCount                  byte   // 0x0Fxx0000 *
	EnvironmentSettingsSegmentOffset          uint32 // 0xyyyyyyyy
	TimeStart                                 uint16 // 0x10000000 0xxxxxyy00 (0xFFFF = current time)
	TimeSpeed                                 byte   // ─────────────────┴┘   (defaults to 0x03)
	SkyboxNumber                              byte   // 0x11000000 0xxx0y0z00 *
	SkyboxCast                                byte   // ────────────────┘ │
	SkyboxFog                                 byte   // ──────────────────┘
	SkyboxDisable                             bool   // 0x12000000 0xxxyy0000 true if > 0
	SkyboxModifier                            byte   // ───────────────┴┘     unknown
	ExitsSegmentOffset                        uint32 // 0x13000000 0xyyyyyyyy
	SoundReverb                               byte   // 0x15xx0000 0x0000yyzz *
	SoundNightSFX                             byte   // ─────────────────┴┘││
	SoundBackgroundSequence                   byte   // ───────────────────┴┘
	SoundEcho                                 byte   // 0x16000000 0x000000xx *
	CutscenesCount                            byte   // 0x17xx0000
	CutscenesSegmentOffset                    uint32 // 0xyyyyyyyy
	AlternateHeadersSegmentOffset             uint32 // 0x18000000 0xxxxxxxxx
	IsWorldMapLocation                        bool   // 0x19000000 0x00000000 (presence = true)
	TextureAnimationsSegmentOffset            uint32 // 0x1A000000 0xxxxxxxxx *
	CamerasAndCutscenesForActorsCount         byte   // 0x1Bxx0000 *
	CamerasAndCutscenesForActorsSegmentOffset uint32 // 0xyyyyyyyy
	MinimapsSegmentOffset                     uint32 // 0x1C000000 0xxxxxxxxx *
	MapChestPositionsCount                    byte   // 0x1Exx0000
	MapChestPositionsSegmentOffset            uint32 // 0xyyyyyyyy
}

// Returns offset after reading header (actual data start)
func (l *LocationHeader) load(r io.ReadSeeker, start uint32) uint32 {
	offset := start

	r.Seek(int64(start), io.SeekStart)
	var a, b uint32
	for {
		binary.Read(r, binary.BigEndian, &a)
		binary.Read(r, binary.BigEndian, &b)
		offset += 8

		command := byte((a & 0xFF000000) >> 24)
		if command == sceneHeaderEndCommand {
			break
		}

		if err := l.loadHeader(command, a, b); err != nil {
			log.Printf(
				"ERROR at offset 0x%08X: %s",
				offset-8,
				err,
			)
		}
	}

	return offset
}

func (l *LocationHeader) loadHeader(command byte, a uint32, b uint32) error {
	switch command {
	case 0x14:
		return fmt.Errorf("loadHeader does not handle the header end command")
	default:
		// There's loads of Room headers that are not documented yet so silence the warning for now
		// TODO: understand them, maybe
		// return fmt.Errorf("unknown LocationHeader command 0x%02X (0x%08X 0x%08X)", command, a, b)
	case 0x00:
		l.StartPositionsCount = byte((a & 0x00FF0000) >> 16)
		l.StartPositionsSegmentOffset = b
	case 0x01:
		l.ActorsCount = byte((a & 0x00FF0000) >> 16)
		l.ActorsSegmentOffset = b
	case 0x02:
		l.CamerasCount = byte((a & 0x00FF0000) >> 16)
		l.CamerasSegmentOffset = b
	case 0x03:
		l.CollisionHeaderSegmentOffset = b
	case 0x04:
		l.RoomsCount = byte((a & 0x00FF0000) >> 16)
		l.RoomsSegmentOffset = b
	case 0x05:
		l.WindDirectionX = byte((b & 0xFF000000) >> 24)
		l.WindDirectionY = byte((b & 0x00FF0000) >> 16)
		l.WindDirectionZ = byte((b & 0x0000FF00) >> 8)
		l.WindStrength = byte(b & 0x000000FF)
	case 0x06:
		l.EntrancesCount = byte((a & 0x00FF0000) >> 16)
		l.EntrancesSegmentOffset = b
	case 0x07:
		l.SpecialObjectsByte0 = byte((a & 0x00FF0000) >> 16)
		l.SpecialObjects = uint16(b & 0x0000FFFF)
	case 0x08:
		l.RoomBehavior0 = byte((a & 0x00FF0000) >> 16)
		l.RoomBehavior1 = byte((b & 0x0000FF00) >> 8)
		l.RoomBehavior2 = byte(b & 0x000000FF)
	case 0x09:
		// "Has two instructions saving values to the stack, which aren't read before being overwritten."
	case 0x0A:
		l.MeshSegmentOffset = b
	case 0x0B:
		l.ObjectsCount = byte((a & 0x00FF0000) >> 16)
		l.ObjectsSegmentOffset = b
	case 0x0C:
		l.LightSettingsCount = byte((a & 0x00FF0000) >> 16)
		l.LightSettingsSegmentOffset = b
	case 0x0D:
		l.PathsSegmentOffset = b
	case 0x0E:
		l.ActorTransitionsCount = byte((a & 0x00FF0000) >> 16)
		l.ActorTransitionsSegmentOffset = b
	case 0x0F:
		l.EnvironmentSettingsCount = byte((a & 0x00FF0000) >> 16)
		l.EnvironmentSettingsSegmentOffset = b
	case 0x10:
		l.TimeStart = uint16((b & 0xFFFF0000) >> 16)
		l.TimeSpeed = byte((b & 0x0000FF00) >> 8)
	case 0x11:
		l.SkyboxNumber = byte((b & 0xFF000000) >> 24)
		l.SkyboxCast = byte((b & 0x000F0000) >> 16)
		l.SkyboxFog = byte((b & 0x00000F00) >> 8)
	case 0x12:
		l.SkyboxDisable = (b & 0xFF000000) > 0
		l.SkyboxModifier = byte((b & 0x00FF0000) >> 16)
	case 0x13:
		l.ExitsSegmentOffset = b
	case 0x15:
		l.SoundReverb = byte((a & 0x00FF0000) >> 16)
		l.SoundNightSFX = byte((b & 0x0000FF00) >> 8)
		l.SoundBackgroundSequence = byte(b & 0x000000FF)
	case 0x16:
		l.SoundEcho = byte(b & 0x000000FF)
	case 0x17:
		l.CutscenesCount = byte((a & 0x00FF0000) >> 16)
		l.CutscenesSegmentOffset = b
	case 0x18:
		l.AlternateHeadersSegmentOffset = b
	case 0x19:
		l.IsWorldMapLocation = true
	case 0x1A:
		l.TextureAnimationsSegmentOffset = b
	case 0x1B:
		l.CamerasAndCutscenesForActorsCount = byte((a & 0x00FF0000) >> 16)
		l.CamerasAndCutscenesForActorsSegmentOffset = b
	case 0x1C:
		l.MinimapsSegmentOffset = b
	case 0x1E:
		l.MapChestPositionsCount = byte((a & 0x00FF0000) >> 16)
		l.MapChestPositionsSegmentOffset = b
	}

	return nil
}
