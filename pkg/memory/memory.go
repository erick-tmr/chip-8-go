package memory

import "fmt"

type Memory struct {
	Space   [4096]byte
	private string
}

func New() Memory {
	fmt.Println("Memory initialized")
	mem_instance := Memory{
		[4096]byte{},
		"Im private",
	}

	return mem_instance
}
