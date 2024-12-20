package sorting

import "fmt"

func InsertionSort(n []int) []int {
	fmt.Println(IntSliceToString(n))

	fmt.Println("-----")
	for i := 1; i < len(n); i++ {
		pointer := i - 1

		for pointer != 0 && n[pointer] > n[i] {
			fmt.Println(pointer)
			pointer -= 1
		}

		tmp := n[i]
		n[i] = n[pointer]
		n[pointer] = tmp

		fmt.Println(IntSliceToString(n))
	}

	return n
}
