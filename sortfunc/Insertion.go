package sortfunc

func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		cur := a[i]
		key := i
		for j := i ; 1 <= j && cur < a[j-1]; j-- {
			key = j - 1
			a[j] = a[j-1]
		}
		a[key] = cur
	}
}
