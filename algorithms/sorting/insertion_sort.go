package sorting

// time: O(n^2)
func InsertionSort(n []int) []int {
	for i := 1; i < len(n); i++ {
		if n[i-1] > n[i] {
			for j := i; j >= 1; j-- {
				if n[j] < n[j-1] {
					tmp := n[j]
					n[j] = n[j-1]
					n[j-1] = tmp
				}
			}
		}
	}

	return n
}
