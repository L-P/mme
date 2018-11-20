package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

var bigEndianROMHeader = [4]byte{0x80, 0x37, 0x12, 0x40}

const romSize = 64 * 1024 * 1024

// ROM represents a decompressed TLoZ:MM NTSC 1.0
type ROM struct {
	Header              [4]byte  // 0x000000-0x000003
	_                   [4]byte  // 0x000004-0x000007
	CodeRAMSegmentStart uint32   // 0x000008-0x00000B
	_                   [20]byte // 0x00000C-0x00001F
	Name                [20]byte // 0x000020-0x000033

	_ [romSize - 0x34]byte
}

// NewROM creates a new ROM from a file path
func NewROM(path string) (*ROM, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open ROM: %s", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if stat.Size() != romSize {
		return nil, fmt.Errorf("expected %d bytes of ROM data, got %d", romSize, stat.Size())
	}

	rom := &ROM{}

	if err := binary.Read(file, binary.BigEndian, rom); err != nil {
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
	if !bytes.Equal(r.Header[:], bigEndianROMHeader[:]) {
		return fmt.Errorf(
			"invalid header, expected 0x%04X got 0x%04X, a valid decompressed big-endian (z64) ROM is required",
			bigEndianROMHeader,
			r.Header,
		)
	}

	log.Printf("ROM is valid Nintendo®⁶⁴ big-endian ROM (z64) for %s", string(r.Name[:]))
	log.Printf("Code RAM segment starts at 0x%04X", r.CodeRAMSegmentStart)

	return nil
}
