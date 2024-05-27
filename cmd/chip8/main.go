package main

import (
	"fmt"

	"github.com/erick-tmr/chip-8-go/internal/data_structures"
)

func main() {
	fmt.Println("CHIP-8 EMU MAIN EXECUTABLE")

	var random_stack data_structures.Stack32[uint16]
	fmt.Println(random_stack)
	fmt.Println("Adding element to stack")

	for i := 0; i < 33; i++ {
		error := random_stack.Push(uint16(i))

		if error == data_structures.ErrStackOverflow {
			fmt.Println(error)
			break
		}
	}

	fmt.Println("Pop element")
	elem, _ := random_stack.Pop()
	fmt.Println(random_stack)
	fmt.Println(elem)
}
