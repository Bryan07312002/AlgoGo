package sorting

import "testing"

func TestSuiteSelectionSort(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}
		got := SelectionSort(input)
		for i := 0; i < len(expected); i++ {
			if got[i] != expected[i] {
				t.Errorf("expected: " + IntSliceToString(expected) +
					" got: " + IntSliceToString(got))
				break
			}
		}
	})
}
