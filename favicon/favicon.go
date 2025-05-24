package favicon

import (
	"encoding/binary"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func createBaseICO() []byte {
	ico := make([]byte, 6+16+40+1024)

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

	// Draw main square
	for y := range 16 {
		for x := range 16 {
			if x >= 4 && x < 12 && y >= 4 && y < 12 {
				offset := 62 + ((15-y)*16+x)*4
				ico[offset] = 0x0a   // Blue
				ico[offset+1] = 0x0a // Green
				ico[offset+2] = 0x0a // Red
				ico[offset+3] = 0xff // Alpha
			}
		}
	}

	return ico
}

func Handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")

	ico := createBaseICO()

	// Random 2x2 square position within the main square (6x6 area)
	startX := 4 + rng.Intn(6)
	startY := 4 + rng.Intn(6)

	// Random alpha: either fully opaque or transparent
	alpha := byte(0)
	if rng.Intn(2) == 1 {
		alpha = 0xff
	}

	// Draw 2x2 square
	for dy := 0; dy < 2 && startY+dy < 12; dy++ {
		for dx := 0; dx < 2 && startX+dx < 12; dx++ {
			y := startY + dy
			x := startX + dx
			offset := 62 + ((15-y)*16+x)*4
			ico[offset+3] = alpha
		}
	}

	if _, err := w.Write(ico); err != nil {
		log.Printf("Error writing favicon: %v", err)
	}
}
