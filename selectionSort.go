package sorting

// Selection Sort
// 1. Find the smallest element
// 2. Place it at the leftmost side of array/subarray by swapping
// Time: O(n2)
// Space: O(n)
// Best time: O(n2); Avg. time: O(n2)

// finding the smallest element in the subarray and placing at the beginning
func SelectionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
