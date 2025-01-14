package sorting

import "testing"

func TestSuiteInsertionSort(t *testing.T) {
	t.Run("should sort", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}
		got := InsertionSort(input)
		for i := 0; i < len(expected); i++ {
			if got[i] != expected[i] {
				t.Errorf(
					"expected: %s got: %s",
					IntSliceToString(expected),
					IntSliceToString(got),
				)
				break
			}
		}
	})

	t.Run("should sort with repeted numbers", func(t *testing.T) {
		input := []int{5, 4, 5, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5, 5}
		got := InsertionSort(input)
		for i := 0; i < len(expected); i++ {
			if got[i] != expected[i] {
				t.Errorf(
					"expected: %s got: %s",
					IntSliceToString(expected),
					IntSliceToString(got),
				)
				break
			}
		}
	})
}
