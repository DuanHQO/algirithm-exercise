package quickSort


//快速排序
func QuickSort(arr []int, start, end int) {
	if end - start > 1 {
		mid := partition(arr, start, end)
		QuickSort(arr, start, mid)
		QuickSort(arr, mid + 1, end)
	}
}

func partition(list []int, begin, end int) (i int) {
	cValue := list[begin]
	i = begin
	for j := i + 1; j < end; j++ {
		if list[j] < cValue {
			i++ // 这里一定要先加1后在交换值，因为支点现在不必交换，现在i 和 j（小于支点和大于支点）正在划分边界
			list[j], list[i] = list[i], list[j]
			//fmt.Println("list:", list)
		}
	}
	/* 到此，i和j的边界已经划分完成，把i对应的值和支点对应的值交换后，就符合了快分的要求：i左边对应的值都小于等于且右边的都大于支点对应的值
	   此时i的值就是新的支点, 对应的值就是新的主元
	*/
	list[i], list[begin] = list[begin], list[i]
	//fmt.Println("partition:", list)
	return i
}

func Quick3way(arr []int, start, end int) {
	if start <= end {
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
	Quick3way(arr, start, lt-1)
	Quick3way(arr, gt+1, end)
}
