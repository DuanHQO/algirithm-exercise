package mergeSort

//归并排序
func MergeSort(arr []int) []int  {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr)/2
	left := MergeSort(arr[0:i])
	right := MergeSort(arr[i:])

	result := merge(left, right)
	return result
}

func merge(left []int, right []int) (result []int) {

	m, n := 0, 0
	l, r := len(left), len(right)
	if r > 0 && l > 0 && left[l-1] > right[0] {
		for m < l && n < r {
			if left[m] < right[n] {
				result = append(result, left[m])
				m ++
			} else {
				//这里可以统计逆序对， count += len(left) - m
				result = append(result, right[n])
				n ++
			}
		}
	}
	result = append(result, left[m:]...)
	result = append(result, right[n:]...)

	return result
}
