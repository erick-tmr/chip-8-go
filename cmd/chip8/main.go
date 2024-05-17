package main

import (
	"fmt"

	"github.com/erick-tmr/chip-8-go/pkg/memory"
)

func main() {
	fmt.Println("CHIP-8 EMU MAIN EXECUTABLE")
	memory_instance := memory.New()
	memory_instance.Space[0] = 1

	fmt.Println(memory_instance.Space)
}
