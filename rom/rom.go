package rom

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"unsafe"
)

var bigEndianROMHeader = [4]byte{0x80, 0x37, 0x12, 0x40}

const mmCRC1 = 0xDA6983E7
const mmCRC2 = 0x50674458

const romSize = 64 * 1024 * 1024

// ROM represents a decompressed TLoZ:MM NTSC 1.0
// Sources:
//   - https://github.com/mupen64plus/mupen64plus-core/blob/master/src/api/m64p_types.h
type ROM struct {
	Header         [4]byte   // 0x00
	ClockRate      uint32    // 0x04
	PC             uint32    // 0x08
	Release        uint32    // 0x0C
	CRC1           uint32    // 0x10
	CRC2           uint32    // 0x14
	_              [2]uint32 // 0x18
	Name           [20]byte  // 0x20
	_              uint32    // 0x34
	ManufacturerID uint32    // 0x38
	CartridgeID    uint16    // 0x3C - Game serial number
	CountryCode    uint16    // 0x3E

	_ [0x00C5A1E0 - 0x40]byte

	InternalSceneTable [113]InternalSceneTableEntry // 0x00C5A1E0 - 0x00C5A8F0

	_ [romSize - 0x00C5A8F0]byte
}

// New loads a new ROM from a file path
func New(r io.ReadSeeker) (*ROM, error) {
	rom := &ROM{}

	r.Seek(0, io.SeekStart)
	if err := binary.Read(r, binary.BigEndian, rom); err != nil {
		return nil, err
	}

	if err := rom.validate(); err != nil {
		return nil, err
	}

	if err := rom.read(); err != nil {
		return nil, nil
	}

	return rom, nil
}

func (r *ROM) read() error {
	return nil
}

func (r *ROM) validate() error {
	size := unsafe.Sizeof(*r)
	if size != romSize {
		return fmt.Errorf(
			"ROM struct size is %X, expected %X, this is either a programming error or the go compiler adding padding",
			size,
			romSize,
		)
	}

	if !bytes.Equal(r.Header[:], bigEndianROMHeader[:]) {
		return fmt.Errorf(
			"invalid header, expected 0x%04X got 0x%04X, a valid decompressed big-endian (z64) ROM is required",
			bigEndianROMHeader,
			r.Header,
		)
	}
	log.Printf("ROM is valid Nintendo®⁶⁴ big-endian ROM (z64) for %s", string(r.Name[:]))

	if r.CRC1 != mmCRC1 {
		return fmt.Errorf("CRC1 does not match, expected %04X got 0x%04X", mmCRC1, r.CRC1)
	}
	if r.CRC2 != mmCRC2 {
		return fmt.Errorf("CRC2 does not match, expected %04X got 0x%04X", mmCRC2, r.CRC2)
	}
	log.Printf("CRCs match MM NTSC 1.0")

	for k, v := range r.InternalSceneTable {
		if err := v.validate(); err != nil {
			return fmt.Errorf("IST entry #%d: %s", k, err)
		}
	}

	return nil
}
