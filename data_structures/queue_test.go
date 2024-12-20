package datastructures

import (
	"strconv"
	"testing"
)

func TestSuiteQueue(t *testing.T) {
	t.Run("Should start empty", func(t *testing.T) {
		queue := new(Queue[int])

		if len(*queue) != 0 {
			t.Errorf("Should start empty but has: " + strconv.Itoa(len(*queue)))
		}
	})

	t.Run("Should grow in size when push is called", func(t *testing.T) {
		queue := new(Queue[int])
		queue.push(0)

		if len(*queue) != 1 {
			t.Errorf("expected: 1 but got: " + strconv.Itoa(len(*queue)))
		}
	})

	t.Run("Should reduce in size when pop is called", func(t *testing.T) {
		queue := make(Queue[int], 0)
		queue.push(0)
		queue.pop()

		if len(queue) != 0 {
			t.Errorf("expected: 0 but got: " + strconv.Itoa(len(queue)))
		}
	})

	t.Run("Should pop in FIFO order", func(t *testing.T) {
		queue := make(Queue[int], 0)
		queue.push(1)
		queue.push(2)
		queue.push(3)

		item, _ := queue.pop()
		if item != 1 {
			t.Errorf("expected: 1 but got: " + strconv.Itoa(item))
		}

		item, _ = queue.pop()
		if item != 2 {
			t.Errorf("expected: 2 but got: " + strconv.Itoa(item))
		}

        item, _ = queue.pop()
		if item != 3 {
			t.Errorf("expected: 3 but got: " + strconv.Itoa(item))
		}
	})

	t.Run("Should return error if pop is called when queue is empty", func(t *testing.T) {
		queue := make(Queue[int], 0)
        _,err := queue.pop()

		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
