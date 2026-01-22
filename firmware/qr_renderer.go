package main

import (
	"image/color"

	"github.com/skip2/go-qrcode"
)

// Displayer is the interface we need to draw pixels
type PixelSetter interface {
	SetPixel(x, y int16, c color.RGBA)
}

// DrawQRCode generates a QR code and draws it to the display at the given position.
// scale is the size multiplier for each module (pixel).
func DrawQRCode(display PixelSetter, content string, xOffset, yOffset int16, scale int16) error {
	// Generate QR Code with Low error correction to produce smaller matrix (fewer modules)
	// This helps fit the QR code better while maintaining scale.
	qr, err := qrcode.New(content, qrcode.Low)
	if err != nil {
		return err
	}

	// Disable border in the raw calculation if possible, or just handle it.
	// go-qrcode includes a quiet zone by default.
	qr.DisableBorder = true

	// Bitmap() returns [][]bool
	// true = black (foreground), false = white (background)
	modules := qr.Bitmap()

	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	for y, row := range modules {
		for x, val := range row {
			// Determine color
			c := white
			if val {
				c = black
			}

			// Draw scaled pixels
			startX := xOffset + int16(x)*scale
			startY := yOffset + int16(y)*scale

			for sy := int16(0); sy < scale; sy++ {
				for sx := int16(0); sx < scale; sx++ {
					display.SetPixel(startX+sx, startY+sy, c)
				}
			}
		}
	}

	return nil
}
