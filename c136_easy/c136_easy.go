package c136_easy

func singleNumber(nums []int) int {
	L:
	for idxout, valout := range nums {
		for idxin, valin := range nums {
			if idxin == idxout {
				continue
			}
			if valin == valout {
				continue L
			}
		}
		return valout
	}
	return 0
}

func singleNumberBitM(nums []int) int {
	a := 0
	for _, val := range nums {
		a ^= val
	}
	return a
}
