package sorting

// Insertion Sort
// Insert the next element at the appropriate place in the left side
// At i-th iteration the array upto size i is sorted
// Time: O(n2)
// Avg. time: O(n2); Best time: O(n)
// Space: O(n)+1
// Best case occurs when array is already sorted

// uses a key; copying consecutive values
func InsertionSort(a []int) {
	var j, key int
	for i := 1; i < len(a); i++ {
		key = a[i]
		j = i - 1
		for j >= 0 && key < a[j] {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}

// without using key; swapping consecutive values
// Space: O(n)
func InsertionSort2(a []int) {
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 && a[j] < a[j-1] {
			a[j], a[j-1] = a[j-1], a[j]
			j--
		}
	}
}
