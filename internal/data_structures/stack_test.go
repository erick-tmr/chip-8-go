package data_structures

import (
	"testing"
)

func Test_stack_Push(t *testing.T) {
	var test_stack Stack32[uint]

	err := test_stack.Push(42)

	if err != nil {
		t.Fatal("Cannot push on empty stack, ", err)
	}

	for i := 0; i < 31; i++ {
		err := test_stack.Push(uint(i))

		if err != nil {
			t.Fatal("Stack overflow before expected, ", err)
		}
	}

	expected_err := test_stack.Push(43)

	if expected_err != ErrStackOverflow {
		t.Fatal("Stack overflow is expected here but got another error, ", expected_err)
	}
}

func Test_stack_Pop(t *testing.T) {
	var test_stack Stack32[uint]
	for i := 1; i < 11; i++ {
		err := test_stack.Push(uint(i))

		if err != nil {
			t.Fatal("Stack overflow before expected, ", err)
		}
	}

	t.Logf("Stack before first pop: %v", test_stack)
	elem, success := test_stack.Pop()

	if !success {
		t.Fatal("Cannot remove any element from stack")
	}
	if elem != 10 {
		t.Fatalf("Popped wrong element from stack, popped: %v", elem)
	}

	for i := 9; i > 0; i-- {
		elem, success := test_stack.Pop()

		if !success {
			t.Fatal("Cannot remove when expected from stack")
		}
		if elem != uint(i) {
			t.Fatalf("Popped wrong element from stack, expected: %v, popped %v", i, elem)
		}
	}

	unexpected, popped := test_stack.Pop()

	if popped {
		t.Fatalf("Was able to pop without any element on the stack, popped: %v", unexpected)
	}
}
