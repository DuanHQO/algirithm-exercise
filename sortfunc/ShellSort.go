package sortfunc

func ShellSort(a []byte) {
	N := len(a)
	h := 1
	for  {
		if h < N/3 {
			h = h*3 + 1
		} else {
			break
		}
	}
	for {
		if h >= 1 {
			for i := h; i < N; i++ {
				for j := i; j >= h && a[j] < a[j-h]; j -= h {
					a[j], a[j-h] = a[j-h], a[j]
				}
			}
			h = h/3
		} else {
			break
		}
	}
}
