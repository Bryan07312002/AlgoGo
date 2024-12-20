package sorting

import (
	"strconv"
	"strings"
	"testing"
)

func TestSuiteBubbleSort(t *testing.T) {
	t.Run("Should sort array correctly", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}
		got := BubbleSort(input)
		for i := 0; i < len(expected); i++ {
			if got[i] != expected[i] {
				t.Errorf("expected: " + IntSliceToString(expected) +
					" got: " + IntSliceToString(got))
				break
			}
		}
	})
}

func IntSliceToString(slice []int) string {
	var out []string

	for _, n := range slice {
		out = append(out, strconv.Itoa(n))
	}


	return "[ " + strings.Join(out, ", ") + " ]"
}


