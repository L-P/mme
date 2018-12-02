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
	Files    []File
	Scenes   []Scene
	Messages []Message

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
		rom:      rom,
		Scenes:   make([]Scene, len(rom.InternalSceneTable), len(rom.InternalSceneTable)),
		Files:    make([]File, len(rom.DMAData), len(rom.DMAData)),
		Messages: make([]Message, len(rom.MessageTable), len(rom.MessageTable)),
		fd:       fd,
	}

	if err := v.load(fd); err != nil {
		fd.Close()
		return nil, err
	}

	log.Print("ROM loaded.")

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

	if err := v.loadMessages(r); err != nil {
		return err
	}

	if err := v.loadScenes(r); err != nil {
		return err
	}

	if err := v.loadRoomData(r); err != nil {
		return err
	}

	v.mapFileTypes()

	return nil
}

// O(n²) deal with it
func (v *View) mapFileTypes() {
	mapped := 0

loop:
	for k := range v.Files {
		if v.Files[k].VROMStart == 0 {
			continue
		}

		for _, scene := range v.Scenes {
			if scene.VROMStart == v.Files[k].VROMStart {
				v.Files[k].Type = "scene"
				mapped++
				continue loop
			}

			for _, room := range scene.Rooms {
				if room.VROMStart == v.Files[k].VROMStart {
					v.Files[k].Type = "room"
					mapped++
					continue loop
				}
			}
		}
	}

	log.Printf("Mapped %d file types", mapped)
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

// loadRoomData sets the raw room data (without headers), this needs to be done
// separately because we don't have the room size without looking at the file
// table
func (v *View) loadRoomData(r io.ReadSeeker) error {
	for scene := range v.Scenes {
		for room := range v.Scenes[scene].Rooms {
			for _, file := range v.Files {
				if file.VROMStart == v.Scenes[scene].Rooms[room].VROMStart {
					v.Scenes[scene].Rooms[room].loadData(r, file.VROMEnd)
				}
			}
		}
	}

	return nil
}

func (v *View) loadScenes(r io.ReadSeeker) error {
	if len(v.Scenes) != len(v.rom.InternalSceneTable) {
		return errors.New("len(v.scenes) != len (v.rom.InternalSceneTable")
	}

	for k, entry := range v.rom.InternalSceneTable {
		v.Scenes[k].load(r, entry)

		// loadRoomData needs the entrance ID to give a room it's proper name,
		// do this first
		for _, msg := range v.Messages {
			if v.Scenes[k].EntranceMessageID == msg.ID {
				v.Scenes[k].EntranceMessage = msg.String
			}
		}

		v.Scenes[k].loadRooms(r)
	}

	log.Printf("Loaded %d Scenes", len(v.Scenes))

	return nil
}

func (v *View) loadMessages(r io.ReadSeeker) error {
	if len(v.Messages) != len(v.rom.MessageTable) {
		return errors.New("len(v.scenes) != len (v.rom.MessageTable")
	}

	for k, entry := range v.rom.MessageTable {
		v.Messages[k].load(r, entry)
	}

	log.Printf("Loaded %d Messages", len(v.Messages))

	return nil
}

func (v *View) Read(p []byte) (n int, err error) {
	return v.fd.Read(p)
}

// Seek implements io.Seeker
func (v *View) Seek(offset int64, whence int) (int64, error) {
	return v.fd.Seek(offset, whence)
}

// GetFileByVROMStart returns a File from a VROMStart
func (v *View) GetFileByVROMStart(start uint32) (*File, error) {
	for k := range v.Files {
		if v.Files[k].VROMStart == start {
			return &v.Files[k], nil
		}
	}
	return nil, errors.New("file not found")
}

// GetSceneByVROMStart returns a Scene from a VROMStart
func (v *View) GetSceneByVROMStart(start uint32) (*Scene, error) {
	for k := range v.Scenes {
		if v.Scenes[k].VROMStart == start {
			return &v.Scenes[k], nil
		}
	}
	return nil, errors.New("scene not found")
}

// GetRoomByVROMStart returns a Room from a VROMStart
func (v *View) GetRoomByVROMStart(start uint32) (*Room, error) {
	for scene := range v.Scenes {
		for room := range v.Scenes[scene].Rooms {
			if v.Scenes[scene].Rooms[room].VROMStart == start {
				return &v.Scenes[scene].Rooms[room], nil
			}
		}
	}
	return nil, errors.New("room not found")
}

// GetROM returns the raw ROM struct
func (v *View) GetROM() *ROM {
	return v.rom
}
