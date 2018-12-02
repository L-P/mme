package rom

import (
	"encoding/binary"
	"io"
)

// Vec3 is a simple x/y/z vector
type Vec3 struct {
	X uint16
	Y uint16
	Z uint16
}

// ActorEntry are entries that point to the dynamic objects present in a Room.
type ActorEntry struct {
	ID                uint16
	SpawnTimeFlags    uint16
	SceneCommandIndex byte
	Initialization    uint16

	NoXRotation bool
	NoYRotation bool
	NoZRotation bool

	Position Vec3
	Rotation Vec3
}

func (a *ActorEntry) load(r io.Reader) {
	a.loadRotationFlagsAndID(r)                   // 2 bytes
	binary.Read(r, binary.BigEndian, &a.Position) // 6 bytes

	a.loadXRotationAndSpawnTimeFlags(r)                 // 2 bytes
	a.loadYRotationAndSceneCommandIndex(r)              // 2 bytes
	a.loadZRotationAndSpawnTimeFlags(r)                 // 2 bytes
	binary.Read(r, binary.BigEndian, &a.Initialization) // 2 bytes
}

func (a *ActorEntry) loadXRotationAndSpawnTimeFlags(r io.Reader) {
	var v uint16
	binary.Read(r, binary.BigEndian, &v)
	a.Rotation.X = (v & 0xFF80) >> 7
	a.SpawnTimeFlags |= (v & 0x0007) << 7
}

func (a *ActorEntry) loadYRotationAndSceneCommandIndex(r io.Reader) {
	var v uint16
	binary.Read(r, binary.BigEndian, &v)
	a.Rotation.Y = (v & 0xFF80) >> 7
	a.SceneCommandIndex = byte(v & 0x007F)
}

func (a *ActorEntry) loadZRotationAndSpawnTimeFlags(r io.Reader) {
	var v uint16
	binary.Read(r, binary.BigEndian, &v)
	a.Rotation.Z = (v & 0xFF80) >> 7
	a.SpawnTimeFlags |= v & 0x007F
}

func (a *ActorEntry) loadRotationFlagsAndID(r io.Reader) {
	var v uint16
	binary.Read(r, binary.BigEndian, &v)

	// First three bits determine rotation options
	a.NoXRotation = v&0x8000 > 0
	a.NoYRotation = v&0x4000 > 0
	a.NoZRotation = v&0x2000 > 0

	// Rest of the value is the ID
	a.ID = v & 0x0FFF
}
