package emulator

import (
	"errors"
	"fmt"

	"github.com/erick-tmr/chip-8-go/pkg/cpu"
	"github.com/erick-tmr/chip-8-go/pkg/memory"
	"github.com/erick-tmr/chip-8-go/pkg/screen"
	"github.com/veandco/go-sdl2/sdl"
)

// Sentinel errors
var ErrEmulator error = errors.New("emulator: error in emulator package")
var ErrEmulatorNew error = fmt.Errorf("emulator/New: error creating new emulator / %v", ErrEmulator)
var ErrEmulatorRun error = fmt.Errorf("emulator/Run: error running emulator / %v", ErrEmulator)

var font = [5 * 16]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

type EmuState int

const (
	Running EmuState = iota
	Paused
	Quitted
)

type Emulator struct {
	State  EmuState
	Memory memory.Memory
	CPU    cpu.CPU
	Screen screen.Screen
}

func New() (emulatorInstance Emulator, cleanup func(), err error) {
	emulatorInstance = Emulator{
		State:  Running,
		Memory: memory.New(),
		CPU:    cpu.New(),
	}

	defer func() {
		if err != nil {
			emulatorInstance.State = Quitted
			err = fmt.Errorf("%v / %v", ErrEmulatorNew, err)
		}
	}()

	// Init screen
	emulatorInstance.Screen = screen.New()
	screenCleanup, err := emulatorInstance.Screen.Init()
	if err != nil {
		return emulatorInstance, nil, err
	}
	cleanup = func() {
		screenCleanup()
	}

	return emulatorInstance, cleanup, nil
}

func (e Emulator) Run() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%v / %v", ErrEmulatorNew, err)
		}
	}()

	// Initial screen clear
	err = e.Screen.Clear()
	if err != nil {
		return err
	}

	screen.SetHelloWorld(&e.Screen)

	// main loop
	for e.State != Quitted {
		handler := NewInputHandler()
		handler.HandleInput(&e)

		if e.State == Paused {
			continue
		}

		instruction_time := uint32(0) // calculate instruction time
		sdl.Delay(screen.DELAY_TIME - instruction_time)

		err = e.Screen.Update()
		if err != nil {
			return err
		}
	}

	return nil
}
