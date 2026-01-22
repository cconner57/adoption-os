package main

import (
	"image/color"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

type KennelCard struct {
	Name        string
	Age         string
	DOB         string
	Sex         string
	Breed       string
	Personality string
	ProfileURL  string
}

// Helper to draw horizontal line
func DrawHLine(display *EPD4in2V2, x, y, width int16, thickness int16) {
	black := color.RGBA{0, 0, 0, 255}
	for t := int16(0); t < thickness; t++ {
		for i := int16(0); i < width; i++ {
			display.SetPixel(x+i, y+t, black)
		}
	}
}

// Helper to draw a box (border)
func DrawBox(display *EPD4in2V2, x, y, w, h int16) {
	// Top
	DrawHLine(display, x, y, w, 2)
	// Bottom
	DrawHLine(display, x, y+h-2, w, 2)

	// Left & Right (manual pixel loop for VLine)
	black := color.RGBA{0, 0, 0, 255}
	for i := int16(0); i < h; i++ {
		display.SetPixel(x, y+i, black)
		display.SetPixel(x+1, y+i, black)

		display.SetPixel(x+w-1, y+i, black)
		display.SetPixel(x+w-2, y+i, black)
	}
}

// Simple Pixel Art Cat Face (with Smile!)
func DrawCatFace(display *EPD4in2V2, x, y int16) {
	black := color.RGBA{0, 0, 0, 255}
	// Simple outline logic or hardcoded bitmap-like strokes

	// Head Outline (approx 30x24)
	// Top head line
	DrawHLine(display, x+6, y, 18, 2)
	// Ears
	// Left Ear
	display.SetPixel(x+5, y-1, black)
	display.SetPixel(x+4, y-2, black)
	display.SetPixel(x+3, y-3, black)
	display.SetPixel(x+2, y-2, black)
	display.SetPixel(x+1, y-1, black)
	display.SetPixel(x, y, black)
	display.SetPixel(x, y+1, black)

	// Right Ear
	display.SetPixel(x+24, y-1, black)
	display.SetPixel(x+25, y-2, black)
	display.SetPixel(x+26, y-3, black)
	display.SetPixel(x+27, y-2, black)
	display.SetPixel(x+28, y-1, black)
	display.SetPixel(x+29, y, black)
	display.SetPixel(x+29, y+1, black)

	// Side faces
	for i := int16(0); i < 15; i++ {
		display.SetPixel(x, y+2+i, black)    // Left
		display.SetPixel(x+29, y+2+i, black) // Right
	}
	// Chin
	DrawHLine(display, x+4, y+18, 22, 2)
	// Chin curves
	display.SetPixel(x+1, y+17, black)
	display.SetPixel(x+2, y+18, black)
	display.SetPixel(x+3, y+18, black)
	display.SetPixel(x+28, y+17, black)
	display.SetPixel(x+27, y+18, black)
	display.SetPixel(x+26, y+18, black)

	// Eyes (x+8, y+8) and (x+18, y+8)
	DrawBox(display, x+7, y+7, 4, 4)
	DrawBox(display, x+19, y+7, 4, 4)

	// Nose (Triangle-ish at x+14, y+12)
	DrawHLine(display, x+13, y+12, 4, 2)
	display.SetPixel(x+14, y+14, black)
	display.SetPixel(x+15, y+14, black)

	// Whiskers
	// Left
	DrawHLine(display, x-2, y+12, 4, 1)
	DrawHLine(display, x-2, y+14, 4, 1)
	// Right
	DrawHLine(display, x+28, y+12, 4, 1)
	DrawHLine(display, x+28, y+14, 4, 1)

	// Smile!
	// x+13, x+14 are center.
	// Curve up at ends.
	// Center mouth
	display.SetPixel(x+14, y+16, black)
	display.SetPixel(x+15, y+16, black)
	// Left side smile
	display.SetPixel(x+13, y+15, black)
	display.SetPixel(x+12, y+15, black)
	// Right side smile
	display.SetPixel(x+16, y+15, black)
	display.SetPixel(x+17, y+15, black)
}

func (k *KennelCard) Draw(display *EPD4in2V2) {
	black := color.RGBA{0, 0, 0, 255}

	// 1. Header Section
	// Name "Colby"
	tinyfont.WriteLine(display, &freesans.Bold24pt7b, 10, 45, k.Name, black)

	// Cat Face Icon (Right side of header)
	DrawCatFace(display, 340, 25)

	// Divider Line under Name to separate Header
	DrawHLine(display, 10, 55, 380, 4)

	// 2. Middle Section (Details + QR)
	// Left: Details
	startX := int16(10)
	startY := int16(90)
	lineHeight := int16(26)

	tinyfont.WriteLine(display, &freesans.Bold12pt7b, startX, startY, "Age: "+k.Age, black)
	tinyfont.WriteLine(display, &freesans.Bold12pt7b, startX, startY+lineHeight, "DOB: "+k.DOB, black)
	tinyfont.WriteLine(display, &freesans.Bold12pt7b, startX, startY+lineHeight*2, "Sex: "+k.Sex, black)

	// Breed (Multi-line support)
	var breed1, breed2 string
	nlIdx := -1
	for i := 0; i < len(k.Breed); i++ {
		if k.Breed[i] == '\n' {
			nlIdx = i
			break
		}
	}
	if nlIdx != -1 {
		breed1 = k.Breed[:nlIdx]
		if nlIdx+1 < len(k.Breed) {
			breed2 = k.Breed[nlIdx+1:]
		}
	} else {
		breed1 = k.Breed
	}

	tinyfont.WriteLine(display, &freesans.Bold12pt7b, startX, startY+lineHeight*3+6, breed1, black)
	if breed2 != "" {
		tinyfont.WriteLine(display, &freesans.Bold12pt7b, startX, startY+lineHeight*4+6, breed2, black)
	}

	// Right: QR Code
	// QR Scale 4, Low Correction
	qrX := int16(285)
	qrY := int16(65)
	DrawQRCode(display, k.ProfileURL, qrX, qrY, 4)

	// "Adopt Me!" text
	// Center under QR.
	// Text "Adopt Me!" (with exclamation) is wider.
	// Center logic: 285 is QR center.
	// Y=205 to move down further out of way of QR.
	btnY := int16(205)
	tinyfont.WriteLine(display, &freesans.Bold12pt7b, 280, btnY, "Adopt Me!", black)

	// 3. Personality Section (Footer)
	// Horizontal Divider separating Footer
	DrawHLine(display, 10, 210, 380, 2)

	// "Personality:" Label
	persY := int16(240)
	tinyfont.WriteLine(display, &freesans.Bold12pt7b, startX, persY, "Personality:", black)

	// Text wrapping logic
	var line1, line2 string
	newlineIdx := -1
	for i := 0; i < len(k.Personality); i++ {
		if k.Personality[i] == '\n' {
			newlineIdx = i
			break
		}
	}

	if newlineIdx != -1 {
		line1 = k.Personality[:newlineIdx]
		if newlineIdx+1 < len(k.Personality) {
			line2 = k.Personality[newlineIdx+1:]
		}
	} else {
		line1 = k.Personality
	}

	tinyfont.WriteLine(display, &freesans.Bold12pt7b, startX, persY+25, line1, black)
	if line2 != "" {
		tinyfont.WriteLine(display, &freesans.Bold12pt7b, startX, persY+50, line2, black)
	}
}
