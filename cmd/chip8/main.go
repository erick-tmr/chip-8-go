package main

import (
	"fmt"
	"os"

	"github.com/erick-tmr/chip-8-go/pkg/emulator"
)

func main() {
	fmt.Println("CHIP-8 EMU MAIN EXECUTABLE")

	// Check if ROM path is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./chip8 <path/to/rom>")
		os.Exit(1)
	}

	// Read ROM into memory
	romPath := os.Args[1]
	romData, err := os.ReadFile(romPath)
	if err != nil {
		fmt.Printf("Error reading ROM file: %v\n", err)
		os.Exit(1)
	}

	// Init emulator
	emulatorInstance, emulatorCleanup, err := emulator.New()
	if err != nil {
		panic(err)
	}
	defer emulatorCleanup()

	// Load ROM into memory starting at address 0x200 (512) which is the
	// standard starting point for CHIP-8 programs
	err = emulatorInstance.Memory.Copy(romData, 0x200)
	if err != nil {
		fmt.Printf("Error loading ROM into memory: %v\n", err)
		os.Exit(1)
	}

	err = emulatorInstance.Run()
	if err != nil {
		panic(err)
	}
}
