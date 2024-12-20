package datastructures

import (
	"strconv"
	"testing"
)

func TestSuiteStack(t *testing.T) {
	t.Run("should init empty", func(t *testing.T) {
		stack := make(Stack[string], 0)

		if len(stack) != 0 {
			t.Errorf("should be empty but has " +
				strconv.Itoa(len(stack)) + " elements")
		}
	})

	t.Run("should grow size when item is pushed", func(t *testing.T) {
		stack := make(Stack[string], 0)
		stack.push("item")

		if len(stack) != 1 {
			t.Errorf("should have 1 item but has " +
				strconv.Itoa(len(stack)) + " elements")
		}
	})

	t.Run("should reduce in size when item is pop is called", func(t *testing.T) {
		stack := make(Stack[string], 0)
		stack.push("item")
		stack.pop()

		if len(stack) != 0 {
			t.Errorf("should have 0 item but has " +
				strconv.Itoa(len(stack)) + " elements")
		}
	})

	t.Run("pop should return the last inserted item when called", func(t *testing.T) {
		stack := make(Stack[string], 0)
		stack.push("1")
		stack.push("2")
		stack.push("3")

		lastItem, _ := stack.pop()
		if lastItem != "3" {
            t.Errorf("expected: 3 but got:" + lastItem )
		}

		lastItem, _ = stack.pop()
		if lastItem != "2" {
            t.Errorf("expected: 2 but got:" + lastItem )
		}

		lastItem, _ = stack.pop()
		if lastItem != "1" {
            t.Errorf("expected: 1 but got:" + lastItem )
		}
	})

	t.Run("pop should return err if stack is empty", func(t *testing.T) {
		stack := make(Stack[string], 0)
        _, err := stack.pop()

		if err == nil {
			t.Errorf("should returned error but error is nil")
		}
	})

}
