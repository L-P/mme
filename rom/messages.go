package rom

import (
	"bytes"
	"encoding/binary"
	"io"
	"unicode/utf8"
)

// MessageEntry is a single entry in the message table
// Sources:
// - https://wiki.cloudmodding.com/mm/Text_Format#Message_Entry_Table
// binpacked, do not change struct size
type MessageEntry struct {
	ID     uint16
	_      uint16 // 0x0000
	Offset uint32 // prefixed with 0x08
}

// Message is text than can appear in a ingame textbox
type Message struct {
	MessageEntry
	MessageHeader
	VROMStart uint32

	String string
}

// MessageHeader is the standard header of every text message
type MessageHeader struct {
	TextBoxType       byte
	TextBoxPosition   byte
	Icon              byte
	NextMessageNumber uint16
	RupeeCost         uint16
	_                 uint32 // 0xFFFFFFFF 0xFFFFFFFF
}

// Note: everything here is ugly
func (f *Message) load(r io.ReadSeeker, entry MessageEntry) {
	f.MessageEntry = entry
	// Offset from start of message table
	f.VROMStart = 0x00AD1000 + (entry.Offset & 0xFFFFFF) // ditch first 0x08 byte
	r.Seek(int64(f.VROMStart), io.SeekStart)

	binary.Read(r, binary.BigEndian, &f.MessageHeader)

	var b byte
	buf := make([]byte, 0, 128)
	for b != 0xBF { // end marker
		binary.Read(r, binary.BigEndian, &b)
		buf = append(buf, b)
	}

	f.String = sanitizeString(buf)
}

func sanitizeString(src []byte) string {
	var buf bytes.Buffer
	var r rune
	skip := 0

	for _, b := range src {
		if skip > 0 {
			skip--
			continue
		}

		switch b {

		case 0x1B:
			fallthrough
		case 0x1C:
			fallthrough
		case 0x1D:
			fallthrough
		case 0x1E:
			fallthrough
		case 0x1F:
			skip = 2
			continue

		case 0x0A:
			skip = 1
			continue

		case 0xBF: // EOL
			continue

		case 0xB0:
			buf.WriteString("[A]")
			continue
		case 0xB1:
			buf.WriteString("[B]")
			continue
		case 0xB2:
			buf.WriteString("[C]")
			continue
		case 0xB3:
			buf.WriteString("[L]")
			continue
		case 0xB4:
			buf.WriteString("[R]")
			continue
		case 0xB5:
			buf.WriteString("[Z]")
			continue
		case 0xB6:
			buf.WriteString("[C Up]")
			continue
		case 0xB7:
			buf.WriteString("[C Down]")
			continue
		case 0xB8:
			buf.WriteString("[C Left]")
			continue
		case 0xB9:
			buf.WriteString("[C Right]")
			continue
		case 0xBA:
			buf.WriteString("▼")
			continue
		case 0xBB:
			buf.WriteString("[Control Stick]")
			continue

		case 0x10:
			fallthrough
		case 0x11:
			fallthrough
		case 0x12:
			fallthrough
		case 0x13:
			r = '\n'
		default:
			if b > 0x7E || b < 0x20 || !utf8.ValidRune(rune(b)) {
				continue
			}
			r = rune(b)
		}

		buf.WriteString(string(r))
	}

	return buf.String()
}
