package cpu

import (
	"fmt"

	"github.com/erick-tmr/chip-8-go/internal/data_structures"
)

var program_counter uint16
var i_register uint16
var call_stack data_structures.Stack32[uint16]

func Start() {
	fmt.Println("CPU initialized")

}
