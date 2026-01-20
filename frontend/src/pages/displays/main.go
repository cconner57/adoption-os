package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/waveshare-epd/epd2in13"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

func main() {
	// 1. Configure the SPI pins (Standard SPI0 on Pico)
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		SCK:       machine.GP18,
		SDO:       machine.GP19,
		SDI:       machine.GP16, // Not strictly used for E-ink write, but required for config
	})

	// 2. Configure Control Pins
	display := epd2in13.New(
		machine.SPI0,
		machine.GP17, // CS
		machine.GP16, // DC
		machine.GP15, // RST
		machine.GP14, // BUSY
	)

	// 3. Configure the display
	display.Configure(epd2in13.Config{})

	// 4. Clear the screen (White)
	// This might take a few seconds - e-ink is slow!
	display.ClearBuffer()
	display.ClearDisplay()

	// 5. Draw some text
	// 0,0 is usually top-left. Color 1 is Black, 0 is White for this driver.
	black := color.RGBA{1, 1, 1, 255}

	tinyfont.WriteLine(&display, &freemono.Bold12pt7b, 10, 30, "Hello World!", black)
	tinyfont.WriteLine(&display, &freemono.Regular9pt7b, 10, 60, "Powered by", black)
	tinyfont.WriteLine(&display, &freemono.Bold12pt7b, 10, 85, "TinyGo & Pico", black)

	// 6. Update the display (Flush buffer to hardware)
	// IMPORTANT: You won't see anything until this runs
	display.Display()

	// 7. Put display to sleep to save power
	display.WaitUntilIdle()
	display.DeepSleep()

	// Loop forever just to keep main running
	for {
		time.Sleep(time.Hour)
	}
}
