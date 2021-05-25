package sorting

func QuickSort(a []int, lh ...int) {
	var low, high int
	if len(lh) == 0 {
		low = 0
		high = len(a) - 1
	} else {
		low = lh[0]
		high = lh[1]
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

func BubbleSort(a []int) {
	for j := 0; j < len(a)-1; j++ {
		for i := 0; i < len(a)-j-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
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

func SelectionSort(a []int) {
	var mI int
	m := func(a []int) int {
		var minI int
		min := a[0]
		for i := 1; i < len(a); i++ {
			if a[i] < min {
				min = a[i]
				minI = i
			}
		}
		return minI
	}
	for i := 0; i < len(a)-1; i++ {
		mI = m(a[i+1:]) + i + 1
		if a[i] > a[mI] {
			a[i], a[mI] = a[mI], a[i]
		}
	}
}
