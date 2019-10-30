package binarySearch

func BinarySearch(key int, array []int) int {

	lo := 0
	hi := len(array)

	for {
		if lo <= hi {
			mid := (lo + hi) / 2
			if key < array[mid] {
				hi = mid - 1
			}else if key > array[mid] {
				lo = mid + 1
			}else {
				return array[mid]
			}
		}else {
			break
		}
	}

	return -1
}

func Rank(key int, a []int) (i int) {
	for _, value := range a {
		if value < key {
			i += 1
		}
	}
	return i
}

func Count(key int, a []int) (j int) {
	for _, value := range a {
		if value == key {
			j += 1
		}
	}
	return j
}