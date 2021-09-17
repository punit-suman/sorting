package sorting

// Merge Sort
// Split into single units and then merge in order
// Time: O(nlogn)
// Space O(n) + O(n)
// Best time: O(nlogn); Avg. time: O(nlogn)

func MergeSort(a []int, se ...int) {
	var start, end, mid int
	if len(se) == 0 {
		start = 0
		end = len(a) - 1
	} else {
		start = se[0]
		end = se[1]
	}
	if start < end {
		mid = (start + end) / 2
		MergeSort(a, start, mid)
		MergeSort(a, mid+1, end)
		// func Merge
		func(a []int, start, mid, end int) {
			temp := make([]int, end-start+1)
			i, j, k := start, mid+1, 0
			for i <= mid && j <= end {
				if a[i] <= a[j] {
					temp[k] = a[i]
					i++
				} else {
					temp[k] = a[j]
					j++
				}
				k++
			}
			for i <= mid {
				temp[k] = a[i]
				k++
				i++
			}
			for j <= end {
				temp[k] = a[j]
				k++
				j++
			}
			for m := start; m <= end; m++ {
				a[m] = temp[m-start]
			}
		}(a, start, mid, end)
	}
}
