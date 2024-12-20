package sorting

// time: O(n^2)
func SelectionSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		smallest := i
		for j := i; j < len(n); j++ {
			if n[j] < n[smallest] {
				smallest = j
			}
		}

		tmp := n[i]
		n[i] = n[smallest]
		n[smallest] = tmp
	}

	return n
}
