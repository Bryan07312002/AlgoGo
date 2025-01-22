package datastructures

// return 1 to the item go down on the list
// return 0 if i == j
// return -1 if the item go up on the list
type Comparator[T any] func(i, j T) int

// MinHeap represents a min-heap data structure.
type PriorityQueue[T any] struct {
	Data    []T
	Compare Comparator[T]
}

// buildMinHeap builds a min-heap from an array.
func buildMinHeap[T any](data *[]T, comparator Comparator[T]) PriorityQueue[T] {
	heap := PriorityQueue[T]{
		Data:    *data,
		Compare: comparator,
	}

	for i := heap.Len()/2 - 1; i >= 0; i-- {
		downHeap(&heap, i)
	}

	return heap
}

func (h PriorityQueue[T]) Len() int {
    return len(h.Data)
}

func (h PriorityQueue[T]) Less(i, j int) bool {
    return h.Compare(h.Data[i], h.Data[j]) == -1
}

func (h PriorityQueue[T]) Swap(i, j int) {
    h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

func (h *PriorityQueue[T]) Push(x T) {
	*&h.Data = append(*&h.Data, x)

	upHeap(*h, h.Len()-1)
}

func (h *PriorityQueue[T]) Pop() interface{} {
	old := *h
	n := old.Len()
	x := old.Data[0]
	*&h.Data = old.Data[1:n]

	if n > 1 {
		downHeap(h, 0)
	}

	return x
}

// upHeap restores the heap property after an insertion.
func upHeap[T any](h PriorityQueue[T], i int) {
	for {
		parent := (i - 1) / 2
		isGreaterOrEqual := h.Compare(h.Data[i], h.Data[parent])
		if parent < 0 || isGreaterOrEqual == 1 || isGreaterOrEqual == 0 {
			break
		}
		h.Swap(i, parent)
		i = parent
	}
}

// downHeap restores the heap property after a removal.
func downHeap[T any](h *PriorityQueue[T], i int) {
	n := h.Len()

	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.Compare(h.Data[left], h.Data[smallest]) == -1 {
			smallest = left
		}
		if right < n && h.Compare(h.Data[right], h.Data[smallest]) == -1 {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.Swap(i, smallest)
		i = smallest
	}
}
