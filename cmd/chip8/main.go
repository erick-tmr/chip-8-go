package main

import (
	"fmt"

	"github.com/erick-tmr/chip-8-go/pkg/emulator"
)

func main() {
	fmt.Println("CHIP-8 EMU MAIN EXECUTABLE")

	// Init emulator
	emulatorInstance, emulatorCleanup, err := emulator.New()
	if err != nil {
		panic(err)
	}
	defer emulatorCleanup()

	err = emulatorInstance.Run()
	if err != nil {
		panic(err)
	}
}
