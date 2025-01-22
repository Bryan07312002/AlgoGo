package datastructures

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	t.Run("Push and Pop", func(t *testing.T) {
		// Initialize a MinHeap
		h := buildMinHeap(
			&[]int{12, 11, 13, 5, 6, 7},
			func(i, j int) int {
				if i > j {
					return 1
				}

				if i < j {
					return -1
				}

				return 0
			},
		)

		// Test the minimum element after building the heap
		expected := 5
		if (h.Data)[0] != expected {
			t.Errorf("Expected minimum element to be %d, but got %d", expected, (h.Data)[0])
		}

		// Test popping the minimum element
		min := h.Pop().(int)
		if min != 5 {
			t.Errorf("Expected to pop 5, but got %d", min)
		}

		// Test the new minimum after popping the root
		expected = 6
		if (h.Data)[0] != expected {
			t.Errorf("Expected new minimum element to be %d, but got %d",
				expected, (h.Data)[0])
		}
	})

	t.Run("Push", func(t *testing.T) {
		// Initialize an empty MinHeap
		h := buildMinHeap(
			&[]int{},
			func(i, j int) int {
				if i > j {
					return 1
				}

				if i < j {
					return -1
				}

				return 0
			},
		)

		// Push elements into the heap
		h.Push(10)
		h.Push(20)
		h.Push(5)
		h.Push(3)
		h.Push(8)

		// Test the minimum element after pushing
		expected := 3
		if (h.Data)[0] != expected {
			t.Errorf("Expected minimum element to be %d, but got %d", expected, (h.Data)[0])
		}

		// Pop all elements and check if they are in ascending order
		popped := []int{}
		for h.Len() > 0 {
			popped = append(popped, h.Pop().(int))
		}

		// Check if elements are in ascending order
		for i := 1; i < len(popped); i++ {
			if popped[i] < popped[i-1] {
				t.Errorf("Heap is not in ascending order: %v", popped)
			}
		}
	})

	t.Run("Empty Pop", func(t *testing.T) {
		// Initialize an empty MinHeap
		h := buildMinHeap(&[]int{}, func(i, j int) int { return 0 })

		// Try popping from an empty heap
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when popping from an empty heap, but got none")
			}
		}()

		h.Pop() // This should panic
	})

	t.Run("check if can min heap", func(t *testing.T) {
		// Test building a MinHeap from an unsorted array
		h := buildMinHeap(
			&[]int{12, 11, 13, 5, 6, 7},
			func(i, j int) int {
				if i > j {
					return 1
				}

				if i < j {
					return -1
				}

				return 0
			},
		)

		// After building the heap, the minimum element should be at the root
		expected := 5
		if (h.Data)[0] != expected {
			t.Errorf("Expected minimum element to be %d, but got %d", expected, (h.Data)[0])
		}

		// Check if the heap is a valid MinHeap by ensuring the heap property
		for i := 0; i < h.Len(); i++ {
			left := 2*i + 1
			right := 2*i + 2
			if left < h.Len() && (h.Data)[i] > (h.Data)[left] {
				t.Errorf("Heap property violated at index %d: %v", i, h.Data)
			}
			if right < h.Len() && (h.Data)[i] > (h.Data)[right] {
				t.Errorf("Heap property violated at index %d: %v", i, h.Data)
			}
		}
	})
}
