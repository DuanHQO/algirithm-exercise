package isCoprime

func IsCoprime(x, y int) (result bool) {
	if x == 1 && y == 1 {
		result = true
	}else if x <= 0 || y <= 0 || x == y {
		result = false
	}else if x == 1 || y == 1 {
		result = true
	}else {
		tmp := 0
		for {
			tmp = x % y
			if tmp == 0 {
				break
			}else {
				x = y
				y = tmp
			}
		}
		if y == 1 {
			result = true
		}else {
			result = false
		}
	}
	return result
}
