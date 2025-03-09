package memory

import (
	"errors"
	"fmt"
)

const LastProtectedAddress = 0x1FF
const FinalAddress = 0xFFF

type Memory struct {
	space [4096]byte
}

func New() Memory {
	return Memory{}
}

func (m *Memory) Load(b byte, addr uint16) error {
	if addr <= LastProtectedAddress {
		return errors.New("memory: trying to insert into protected space")
	}

	if addr > FinalAddress {
		return errors.New("memory: addr is too large for memory")
	}

	m.space[addr] = b

	return nil
}

func (m *Memory) Copy(source []byte, startAddr uint16) error {
	if len(source) > (FinalAddress - LastProtectedAddress) {
		return errors.New("memory: memory overflow, trying to copy too big source")
	}

	for i, byteValue := range source {
		err := m.Load(byteValue, startAddr+uint16(i))
		if err != nil {
			return fmt.Errorf("memory/copy: error loading byte / %v", err)
		}
	}

	return nil
}
