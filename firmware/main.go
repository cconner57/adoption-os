package main

import (
	"machine"
	"time"
)

func main() {
	// 1. PIN SETUP for Pico 2 W
	// SPI1 SCK: GP10
	// SPI1 TX:  GP11
	// CS:       GP9
	// DC:       GP8
	// RST:      GP12
	// BUSY:     GP13

	// Configure SPI
	spi := machine.SPI1
	spi.Configure(machine.SPIConfig{
		Frequency: 4000000,
		SCK:       machine.GP10,
		SDO:       machine.GP11,
		SDI:       machine.GP12,
		Mode:      0,
	})

	config := Config{
		SPI:  spi,
		CS:   machine.GP9,
		DC:   machine.GP8,
		RST:  machine.GP12,
		BUSY: machine.GP13,
	}

	display := NewEPD4in2V2(config)
	// Initialize with V2 sequence
	display.Init()

	// Clear display once at startup
	display.ClearBuffer()
	display.ClearDisplay()

	// Create Kennel Card Data
	card := KennelCard{
		Name:        "Colby",
		Age:         "5 mo",
		DOB:         "07/22/25",
		Sex:         "Male",
		Breed:       "Black & White\nTuxedo",
		Personality: "Affectionate, curious,\nand playful!",
		ProfileURL:  "https://www.idohr.app/adopt/colby-bdb24",
	}

	// Draw the card
	card.Draw(display)

	// Update screen
	display.Display()

	// Power down display
	display.Sleep()

	// Blink LED to indicate completion
	led := machine.GP0
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		time.Sleep(1000 * time.Millisecond)
		led.Low()
		time.Sleep(1000 * time.Millisecond)
	}
}
