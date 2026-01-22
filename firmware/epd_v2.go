package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers"
)

// EPD4in2V2 driver for Waveshare 4.2 inch V2 e-Paper module
// Resolution: 400x300
type EPD4in2V2 struct {
	spi    drivers.SPI
	cs     machine.Pin
	dc     machine.Pin
	rst    machine.Pin
	busy   machine.Pin
	buffer []byte
	width  int16
	height int16
}

// Config defines the pins and SPI interface
type Config struct {
	SPI  drivers.SPI
	CS   machine.Pin
	DC   machine.Pin
	RST  machine.Pin
	BUSY machine.Pin
}

func NewEPD4in2V2(c Config) *EPD4in2V2 {
	c.CS.Configure(machine.PinConfig{Mode: machine.PinOutput})
	c.DC.Configure(machine.PinConfig{Mode: machine.PinOutput})
	c.RST.Configure(machine.PinConfig{Mode: machine.PinOutput})
	c.BUSY.Configure(machine.PinConfig{Mode: machine.PinInput})

	// Buffer for 400x300 B/W: 400*300/8 = 15000 bytes
	return &EPD4in2V2{
		spi:    c.SPI,
		cs:     c.CS,
		dc:     c.DC,
		rst:    c.RST,
		busy:   c.BUSY,
		width:  400,
		height: 300,
		buffer: make([]byte, 15000),
	}
}

// Init initializes the display with the exact V2 C sequence
func (d *EPD4in2V2) Init() {
	d.Reset()
	d.WaitBusy() // Wait AFTER reset

	// 1. SW Reset
	d.SendCommand(0x12)
	d.WaitBusy()

	// 2. Display Update Control 1
	d.SendCommand(0x21)
	d.SendData(0x40) // CORRECTED: 0x40 (Ram source) instead of 0x00
	d.SendData(0x00)

	// 3. Border Waveform
	d.SendCommand(0x3C)
	d.SendData(0x05)

	// 4. Data Entry Mode
	d.SendCommand(0x11)
	d.SendData(0x03) // X inc, Y inc

	// 5. Set Windows to full screen
	d.SetWindows(0, 0, uint16(d.width)-1, uint16(d.height)-1)

	// 6. Set Cursor to 0,0
	d.SetCursor(0, 0)

	d.WaitBusy()
}

func (d *EPD4in2V2) Reset() {
	d.rst.High()
	time.Sleep(200 * time.Millisecond)
	d.rst.Low()
	time.Sleep(5 * time.Millisecond) // C Code: Delay_ms(2)
	d.rst.High()
	time.Sleep(200 * time.Millisecond)
}

func (d *EPD4in2V2) WaitBusy() {
	// Waveshare V2: BUSY pin HIGH when busy
	// Wait while HIGH
	// Add timeout to avoid infinite hang
	timeout := time.Now().Add(30 * time.Second) // Generous timeout for slow EPDs
	for d.busy.Get() {
		if time.Now().After(timeout) {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (d *EPD4in2V2) SendCommand(cmd uint8) {
	d.dc.Low()
	d.cs.Low()
	d.spi.Tx([]byte{cmd}, nil)
	d.cs.High()
}

func (d *EPD4in2V2) SendData(val uint8) {
	d.dc.High()
	d.cs.Low()
	d.spi.Tx([]byte{val}, nil)
	d.cs.High()
}

func (d *EPD4in2V2) SetWindows(xStart, yStart, xEnd, yEnd uint16) {
	d.SendCommand(0x44) // Set RAM X - Start End
	d.SendData(uint8((xStart >> 3) & 0xFF))
	d.SendData(uint8((xEnd >> 3) & 0xFF))

	d.SendCommand(0x45) // Set RAM Y - Start End
	d.SendData(uint8(yStart & 0xFF))
	d.SendData(uint8((yStart >> 8) & 0xFF))
	d.SendData(uint8(yEnd & 0xFF))
	d.SendData(uint8((yEnd >> 8) & 0xFF))
}

func (d *EPD4in2V2) SetCursor(x, y uint16) {
	d.SendCommand(0x4E) // Set RAM X Counter
	d.SendData(uint8((x >> 3) & 0xFF))

	d.SendCommand(0x4F) // Set RAM Y Counter
	d.SendData(uint8(y & 0xFF))
	d.SendData(uint8((y >> 8) & 0xFF))
}

// ClearBuffer sets the buffer to white (0xFF)
func (d *EPD4in2V2) ClearBuffer() {
	for i := range d.buffer {
		d.buffer[i] = 0xFF
	}
}

// ClearDisplay clears the screen physically
func (d *EPD4in2V2) ClearDisplay() {
	// To clear, we write white to both "Old" (Red) and "New" (BW) RAM

	// Write to BW (0x24)
	d.SendCommand(0x24)
	for i := 0; i < len(d.buffer); i++ {
		d.SendData(0xFF)
	}

	// Write to Red (0x26)
	d.SendCommand(0x26)
	for i := 0; i < len(d.buffer); i++ {
		d.SendData(0xFF)
	}
	d.Update()
}

// Display sends the current buffer and updates
func (d *EPD4in2V2) Display() error {
	// Write buffer to BW RAM
	d.SendCommand(0x24)
	for _, b := range d.buffer {
		d.SendData(b)
	}

	// Usually for partial updates or clean updates, we might copy to 0x26 too
	// But standard Display() usually just updates 0x24.
	// However, if the "Old" data isn't correct, ghosting/static happens.
	// For "Standard" refresh (0xF7), it usually diffs against old RAM.
	// Let's copy buffer to 0x26 as well to be safe, or write all White to it?
	// C driver doesn't explicitly write 0x26 in Display().
	// But it does in Clear().
	// For now, let's just write 0x24.

	d.Update()
	return nil
}

func (d *EPD4in2V2) Update() {
	d.SendCommand(0x22)
	d.SendData(0xF7)
	d.SendCommand(0x20)
	d.WaitBusy()
}

func (d *EPD4in2V2) Sleep() {
	d.SendCommand(0x10)
	d.SendData(0x01)
}

// tinyfont.Displayer interface
func (d *EPD4in2V2) Size() (int16, int16) {
	return d.width, d.height
}

func (d *EPD4in2V2) SetPixel(x, y int16, c color.RGBA) {
	if x < 0 || x >= d.width || y < 0 || y >= d.height {
		return
	}
	idx := (int(x) + int(y)*int(d.width)) / 8
	bit := byte(0x80) >> (uint(x) % 8)

	// EPD: 1 = White, 0 = Black.
	// If color is bright, set bit (White). Else clear bit (Black).
	if c.R > 127 || c.G > 127 || c.B > 127 {
		d.buffer[idx] |= bit
	} else {
		d.buffer[idx] &= ^bit
	}
}
