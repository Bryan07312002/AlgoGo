package sorting

// time: O(n^2)
func BubbleSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n)-1; j++ {
			if n[j] > n[j+1] {
				tmp := n[j]
				n[j] = n[j+1]
				n[j+1] = tmp
			}
		}
	}

    return n
}
