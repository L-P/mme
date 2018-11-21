package rom

import "fmt"

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
