package emulator

type EmuState int

const (
	Running EmuState = iota
	Paused
	Quitted
)

type Emulator struct {
	State EmuState
}

func New() Emulator {
	return Emulator{
		State: Running,
	}
}
