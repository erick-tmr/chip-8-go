package main

import (
	"fmt"

	"github.com/erick-tmr/chip-8-go/pkg/emulator"
	"github.com/erick-tmr/chip-8-go/pkg/input"
	"github.com/erick-tmr/chip-8-go/pkg/screen"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	fmt.Println("CHIP-8 EMU MAIN EXECUTABLE")

	// Init emulator
	emulatorInstance := emulator.New()

	// Init screen
	screenInstance := screen.New()
	screenCleanup, err := screenInstance.Init()
	if err != nil {
		panic(err)
	}
	defer screenCleanup()

	err = screenInstance.Clear()
	if err != nil {
		panic(err)
	}

	// main loop
	for emulatorInstance.State != emulator.Quitted {
		handler := input.NewInputHandler()
		handler.HandleInput(&emulatorInstance)

		if emulatorInstance.State == emulator.Paused {
			continue
		}

		instruction_time := uint32(0) // calculate instruction time
		sdl.Delay(screen.DELAY_TIME - instruction_time)

		// clear screen and update
		err = screenInstance.Clear()
		if err != nil {
			panic(err)
		}
		err = screenInstance.Update()
		if err != nil {
			panic(err)
		}
	}
}
