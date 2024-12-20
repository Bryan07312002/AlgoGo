/*
Example 1:
Input: prices = [8,4,6,2,3]
Output: [4,2,4,2,3]
Explanation:
For item 0 with price[0]=8 you will receive a discount equivalent to prices[1]=4, therefore, the final price you will pay is 8 - 4 = 4.
For item 1 with price[1]=4 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 4 - 2 = 2.
For item 2 with price[2]=6 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 6 - 2 = 4.
For items 3 and 4 you will not receive any discount at all.

Example 2:
Input: prices = [1,2,3,4,5]
Output: [1,2,3,4,5]
Explanation: In this case, for all items, you will not receive any discount at all.

Example 3:
Input: prices = [10,1,1,6]
Output: [9,0,1,6]
*/

package leetgo

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func IntSliceToStringSlice(slice []int) []string {
	strSlice := make([]string, len(slice))

	for i, num := range slice {
		strSlice[i] = strconv.Itoa(num)
	}

	return strSlice
}

func Test1475(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		prices := []int{8, 4, 6, 2, 3}
		expected := []int{4, 2, 4, 2, 3}
		got := FinalPrices(prices)

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected: [" + strings.Join(IntSliceToStringSlice(expected), ",") + "] but got: [" + strings.Join(IntSliceToStringSlice(got), ",") + "]")
		}
	})

	t.Run("example 2", func(t *testing.T) {
		prices := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		got := FinalPrices(prices)

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected: [" + strings.Join(IntSliceToStringSlice(expected), ",") + "] but got: [" + strings.Join(IntSliceToStringSlice(got), ",") + "]")
		}
	})

	t.Run("example 3", func(t *testing.T) {
		prices := []int{10, 1, 1, 6}
		expected := []int{9, 0, 1, 6}
		got := FinalPrices(prices)

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected: [" + strings.Join(IntSliceToStringSlice(expected), ",") + "] but got: [" + strings.Join(IntSliceToStringSlice(got), ",") + "]")
		}
	})
}
