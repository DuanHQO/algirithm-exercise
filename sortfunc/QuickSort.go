package sortfunc

func Quick3way(arr []int, start, end int) {
	if end <= start {
		return
	}
	lt := start
	i := start + 1
	gt := end
	for i <= gt {
		if arr[i] < arr[lt] {
			arr[i], arr[lt] = arr[lt], arr[i]
			lt++
		} else if arr[i] > arr[lt] {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		} else {
			i++
		}
	}
	//fmt.Printf("%v\n", arr)
	Quick3way(arr, start, lt-1)
	Quick3way(arr, gt+1, end)
}
