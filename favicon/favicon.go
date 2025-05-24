package favicon

import (
	"encoding/binary"
	"log"
	"math/rand"
	"net/http"
)

var ico = make([]byte, 6+16+40+1024)

func init() {
	// ICO header
	binary.LittleEndian.PutUint16(ico[0:2], 0) // Reserved
	binary.LittleEndian.PutUint16(ico[2:4], 1) // Type: ICO
	binary.LittleEndian.PutUint16(ico[4:6], 1) // Count: 1 image

	// ICO directory entry
	ico[6] = 16                                        // Width: 16px
	ico[7] = 16                                        // Height: 16px
	ico[8] = 0                                         // Colors: 0 (no palette)
	ico[9] = 0                                         // Reserved
	binary.LittleEndian.PutUint16(ico[10:12], 1)       // Planes
	binary.LittleEndian.PutUint16(ico[12:14], 32)      // Bits per pixel
	binary.LittleEndian.PutUint32(ico[14:18], 40+1024) // Size
	binary.LittleEndian.PutUint32(ico[18:22], 22)      // Offset

	// BMP header
	binary.LittleEndian.PutUint32(ico[22:26], 40) // Header size
	binary.LittleEndian.PutUint32(ico[26:30], 16) // Width
	binary.LittleEndian.PutUint32(ico[30:34], 32) // Height (double for ICO)
	binary.LittleEndian.PutUint16(ico[34:36], 1)  // Planes
	binary.LittleEndian.PutUint16(ico[36:38], 32) // Bits per pixel

	// Fill with color #0a0a0a (BGR format with alpha)
	for i := 62; i < 62+1024; i += 4 {
		ico[i] = 0x0a   // Blue
		ico[i+1] = 0x0a // Green
		ico[i+2] = 0x0a // Red
		ico[i+3] = 0xff // Alpha
	}
}

func Handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")

	for i := 62; i < 62+1024; i += 4 {
		ico[i+3] = uint8(rand.Intn(256))
	}

	if _, err := w.Write(ico); err != nil {
		log.Printf("Error writing favicon: %v", err)
	}
}
