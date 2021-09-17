package sorting

// Quick Sort
// 1. Pick an element as pivot and place it at the right place such that all element to its left are smaller
// 2. Partition around pivot creating subarrays and go through recursively
// Time: O(n)
// Avg. time: O(nlogn); Best time: O(nlogn)
// Space: O(n)
// Best case occurs when middle element is chosen as pivot

// last element as pivot
func QuickSort(a []int, lh ...int) {
	var low, high int
	if len(lh) == 2 {
		low = lh[0]
		high = lh[1]
	} else {
		low = 0
		high = len(a) - 1
	}
	if low < high {
		// func partition()
		p := func(a []int, low, high int) int {
			i := low - 1
			pivot := a[high]
			for j := low; j < high; j++ {
				if a[j] <= pivot {
					i++
					a[i], a[j] = a[j], a[i]
				}
			}
			a[i+1], a[high] = a[high], a[i+1]
			return i + 1
		}(a, low, high)
		QuickSort(a, low, p-1)
		QuickSort(a, p+1, high)
	}
}

// first element as pivot
func QuickSortFirst(a []int, lh ...int) {
	var low, high int
	if len(lh) == 2 {
		low = lh[0]
		high = lh[1]
	} else {
		low = 0
		high = len(a) - 1
	}
	if low < high {
		p := func(a []int, low, high int) int {
			i := low
			pivot := a[low]
			for j := low + 1; j <= high; j++ {
				if a[j] <= pivot {
					i++
					a[j], a[i] = a[i], a[j]
				}
			}
			a[i], a[low] = a[low], a[i]
			return i
		}(a, low, high)
		QuickSortFirst(a, low, p-1)
		QuickSortFirst(a, p+1, high)
	}
}

// not currently working for array containing same elements
// middle element as pivot (Best Case)
func QuickSortBest(a []int, lh ...int) {
	var low, high int
	if len(lh) > 0 {
		low = lh[0]
		high = lh[1]
	} else {
		low = 0
		high = len(a) - 1
	}
	if low < high {
		p := func(a []int, low, high int) int {
			i := low - 1
			pivot := a[low+(high-low)/2]
			for j := low; j <= high; j++ {
				if a[j] < pivot {
					i++
					a[j], a[i] = a[i], a[j]
				}
			}
			a[i+1], a[low+(high-low)/2] = a[low+(high-low)/2], a[i+1]
			return i + 1
		}(a, low, high)
		QuickSortBest(a, low, p-1)
		QuickSortBest(a, p+1, high)
	}
}
