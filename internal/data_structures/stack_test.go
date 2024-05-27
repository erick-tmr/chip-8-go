package data_structures

import "testing"

func Test_stack_Push(t *testing.T) {
	var test_stack Stack32[uint]

	err := test_stack.Push(42)

	if err != nil {
		t.Error("Cannot push on empty stack, ", err)
	}

	for i := 0; i < 31; i++ {
		err := test_stack.Push(uint(i))

		if err != nil {
			t.Error("Stack overflow before expected, ", err)
		}
	}

	expected_err := test_stack.Push(43)

	if expected_err != ErrStackOverflow {
		t.Error("Stack overflow is expected here but got another error, ", expected_err)
	}
}
