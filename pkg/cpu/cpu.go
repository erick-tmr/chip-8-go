package cpu

import "github.com/erick-tmr/chip-8-go/pkg/memory"

type CPU struct {
	I  uint16
	pc uint16
	V  [16]byte
}

func New() CPU {
	return CPU{
		pc: memory.LastProtectedAddress + 1, // Points to memory location after protected area (CHIP8 interpreter)
	}
}
