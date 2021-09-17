package sorting

import "fmt"

// Bubble Sort
// checks each element with each next elements
// Place largest element at the rightmost side of the array/subarray by continuous swapping
// Time: O(n2)
// Space: O(n)
// Best time: O(n); Avg. time: O(n2)

// placing largest element at the rightmost side by successive swapping
func BubbleSort(a []int) {
	for j := 0; j < len(a)-1; j++ {
		for i := 0; i < len(a)-j-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
			fmt.Println(a)
		}
	}
}

func BubbleSortRecursion(a []int) {
	if len(a) == 1 {
		return
	}
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			a[i], a[i+1] = a[i+1], a[i]
		}
	}
	BubbleSortRecursion(a[:len(a)-1])
}
