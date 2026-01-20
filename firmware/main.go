package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/skip2/go-qrcode"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

// --- SoftSPI ---
type SoftSPI struct {
	SCK machine.Pin
	SDO machine.Pin
}

func (s *SoftSPI) Tx(w, r []byte) error {
	for _, b := range w {
		for i := 7; i >= 0; i-- {
			s.SCK.Low()
			if b&(1<<i) != 0 {
				s.SDO.High()
			} else {
				s.SDO.Low()
			}
			time.Sleep(10 * time.Microsecond)
			s.SCK.High()
			time.Sleep(10 * time.Microsecond)
		}
	}
	s.SCK.Low()
	return nil
}

// --- EPD Driver ---
type EPD struct {
	spi  *SoftSPI
	cs   machine.Pin
	dc   machine.Pin
	rst  machine.Pin
	busy machine.Pin
}

func NewEPD(spi *SoftSPI, cs, dc, rst, busy machine.Pin) *EPD {
	cs.Configure(machine.PinConfig{Mode: machine.PinOutput})
	dc.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rst.Configure(machine.PinConfig{Mode: machine.PinOutput})
	busy.Configure(machine.PinConfig{Mode: machine.PinInput})
	return &EPD{spi: spi, cs: cs, dc: dc, rst: rst, busy: busy}
}

func (d *EPD) SendCommand(cmd byte) {
	d.dc.Low()
	d.cs.Low()
	d.spi.Tx([]byte{cmd}, nil)
	d.cs.High()
}

func (d *EPD) SendData(data byte) {
	d.dc.High()
	d.cs.Low()
	d.spi.Tx([]byte{data}, nil)
	d.cs.High()
}

func (d *EPD) Reset() {
	d.rst.High()
	time.Sleep(20 * time.Millisecond)
	d.rst.Low()
	time.Sleep(20 * time.Millisecond)
	d.rst.High()
	time.Sleep(200 * time.Millisecond)
}

func (d *EPD) Init() {
	d.Reset()
	d.SendCommand(0x12)
	time.Sleep(100 * time.Millisecond)

	d.SendCommand(0x21)
	d.SendData(0x00)
	d.SendData(0x00)
	d.SendCommand(0x3C)
	d.SendData(0x05)
	d.SendCommand(0x11)
	d.SendData(0x03)
	d.SendCommand(0x44)
	d.SendData(0x00)
	d.SendData(0x31)
	d.SendCommand(0x45)
	d.SendData(0x00)
	d.SendData(0x00)
	d.SendData(0x2B)
	d.SendData(0x01)
	d.SetCursor(0, 0)
}

func (d *EPD) SetCursor(x, y int) {
	d.SendCommand(0x4E)
	d.SendData(byte(x >> 3))
	d.SendCommand(0x4F)
	d.SendData(byte(y))
	d.SendData(byte(y >> 8))
}

func (d *EPD) DisplayBuffer(buffer []byte) {
	// SAFE SINGLE CHANNEL:
	// 0x00 = White.
	// Write ONLY to 0x24.
	// Rely on HW Reset to keep 0x26 clean (Transparent).

	d.SetCursor(0, 0)
	d.SendCommand(0x24)
	for _, b := range buffer {
		d.SendData(b)
	}

	// Write to 0x26 (Red/Old buffer) to ensure it's clean
	d.SendCommand(0x26)
	for i := 0; i < len(buffer); i++ {
		d.SendData(0xFF)
	}

	d.SendCommand(0x22)
	d.SendData(0xF7)
	d.SendCommand(0x20)
	time.Sleep(5 * time.Second)
}

// TinyGo Adapter
type TinyGoDisplay struct {
	epd    *EPD
	buffer []byte
}

func (t *TinyGoDisplay) Size() (int16, int16) { return 400, 300 }
func (t *TinyGoDisplay) SetPixel(x, y int16, c color.RGBA) {
	if x < 0 || x >= 400 || y < 0 || y >= 300 {
		return
	}
	idx := (int(x) + int(y)*400) / 8
	bit := byte(0x80) >> (uint(x) % 8)

	// Logic: 0 -> Black, 1 -> White (Inverted from original assumption)
	if c.R > 127 { // White
		t.buffer[idx] |= bit // Set (1) -> White
	} else { // Black
		t.buffer[idx] &= ^bit // Clear (0) -> Black
	}
}
func (t *TinyGoDisplay) Display() error {
	t.epd.DisplayBuffer(t.buffer)
	return nil
}

func main() {
	led := machine.GP0
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for i := 0; i < 5; i++ {
		led.High()
		time.Sleep(100 * time.Millisecond)
		led.Low()
		time.Sleep(100 * time.Millisecond)
	}

	sck := machine.GP10
	sdo := machine.GP11
	cs := machine.GP9
	dc := machine.GP8
	rst := machine.GP12
	busy := machine.GP13

	sck.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sdo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	spi := &SoftSPI{SCK: sck, SDO: sdo}
	epd := NewEPD(spi, cs, dc, rst, busy)
	epd.Init()

	// NO CLEANING - Keep it simple

	// --- RENDER DATA ---
	buffer := make([]byte, 15000)
	for i := range buffer {
		buffer[i] = 0xFF
	} // BG = White (0xFF)
	display := &TinyGoDisplay{epd: epd, buffer: buffer}
	black := color.RGBA{0, 0, 0, 255}

	// Data
	name := "Colby"
	age := "Age: 5 mo"
	dob := "DOB: 07/22/25"
	sex := "Sex: Male"
	colorInfo := "Black & White"
	persHeader := "Personality:"
	persList := "Affectionate, curious,"
	persList2 := "& playful!"
	url := "https://adoption-os.com/adopt/colby"

	// Layout
	tinyfont.WriteLine(display, &freemono.Bold24pt7b, 10, 50, name, black)
	tinyfont.WriteLine(display, &freemono.Bold12pt7b, 10, 80, age, black)
	tinyfont.WriteLine(display, &freemono.Bold12pt7b, 10, 105, dob, black)
	tinyfont.WriteLine(display, &freemono.Bold12pt7b, 10, 130, sex, black)
	tinyfont.WriteLine(display, &freemono.Regular9pt7b, 10, 160, colorInfo, black)
	tinyfont.WriteLine(display, &freemono.Bold9pt7b, 10, 195, persHeader, black)
	tinyfont.WriteLine(display, &freemono.Regular9pt7b, 10, 215, persList, black)
	tinyfont.WriteLine(display, &freemono.Regular9pt7b, 10, 235, persList2, black)

	// QR Code
	qr, err := qrcode.New(url, qrcode.Medium)
	if err == nil {
		startX, startY := 260, 0
		qrBitmap := qr.Bitmap()
		scale := 4
		for y, row := range qrBitmap {
			for x, val := range row {
				if val {
					for dy := 0; dy < scale; dy++ {
						for dx := 0; dx < scale; dx++ {
							display.SetPixel(int16(startX+x*scale+dx), int16(startY+y*scale+dy), black)
						}
					}
				}
			}
		}
		tinyfont.WriteLine(display, &freemono.Bold12pt7b, 265, 160, "Adopt Me!", black)
	}

	display.Display()

	for {
		led.High()
		time.Sleep(1000 * time.Millisecond)
		led.Low()
		time.Sleep(1000 * time.Millisecond)
	}
}
