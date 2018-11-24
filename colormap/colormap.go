package colormap

import (
	"bufio"
	"encoding/binary"
	"image"
	"image/png"
	"io"
	"log"

	"github.com/L-P/mme/rom"
)

const side = 4096 // sqrt(64 MiB) = 4096² 4-byte words, 1 pixel per uint32

// Generate creates an color-mapped image of the binary.
// Magenta is unknown/unmapped data
// Blue is files
// Half-values (eg. dark blue) is zeroes
func Generate(w io.Writer, v *rom.View) error {
	img := image.NewNRGBA(image.Rect(0, 0, side, side))
	log.Print("Generating color map…")
	if err := fill(img, v); err != nil {
		return err
	}

	log.Print("Compressing color map…")
	enc := png.Encoder{CompressionLevel: png.BestSpeed}
	return enc.Encode(w, img)
}

func fill(img *image.NRGBA, v *rom.View) error {
	for i := 0; i < len(img.Pix); i += 4 { // pink for unknown
		img.Pix[i+0] = 255 // R
		img.Pix[i+2] = 255 // B
		img.Pix[i+3] = 255 // A, set only once
	}

	for _, file := range v.Files { // blue for unknown files
		for i := file.VROMStart; i < file.VROMEnd; i += 4 {
			img.Pix[i+0] = 0
			img.Pix[i+1] = 0
			img.Pix[i+2] = 255
		}
	}

	// Mark dmadata as known
	for i := 0x0001A500; i < 0x0020600; i += 4 {
		img.Pix[i+0] = 0
		img.Pix[i+1] = 255
		img.Pix[i+2] = 0
	}

	// Mark InternalSceneTable as known
	for i := 0x00C5A1E0; i < 0x00C5A8F0; i += 4 {
		img.Pix[i+0] = 0
		img.Pix[i+1] = 255
		img.Pix[i+2] = 0
	}

	v.Seek(0, io.SeekStart)
	buf := bufio.NewReader(v)
	var word uint32
	for i := 0; i < rom.Size; i += 4 { // dim zeroes
		binary.Read(buf, binary.BigEndian, &word)
		if word == 0x00000000 {
			img.Pix[i+0] /= 2
			img.Pix[i+1] /= 2
			img.Pix[i+2] /= 2
		}
	}

	return nil
}
