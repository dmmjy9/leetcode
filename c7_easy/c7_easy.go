package c7_easy

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	if x == 0 {
		return x
	}

	if float64(x) - math.MaxInt32 > 0 ||
		float64(x) - math.MinInt32 < 0 {
		return 0
	}

	var retSlice []string
	var retStr string
	strX := strconv.Itoa(int(math.Abs(float64(x))))

	for i := len(strX)-1; i != -1; i-- {
		retSlice = append(retSlice, string(strX[i]))
	}
	for _, v := range retSlice {
		retStr = retStr + v
	}

	retInt, _ := strconv.Atoi(retStr)

	if float64(retInt) - math.MaxInt32 > 0 ||
		float64(retInt) - math.MinInt32 < 0 {
		return 0
	}

	if x < 0 {
		return 0 - retInt
	}
	return retInt
}

func reverseSolution(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	var pop int
	var rev = 0
	for ; x != 0; x /= 10 {
		pop = x % 10
		if rev > math.MaxInt32 / 10 ||
			(rev == math.MaxInt32 / 10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32 / 10 ||
			(rev == math.MinInt32 / 10 && pop < -8) {
			return 0
		}
		rev = rev * 10 + pop
	}
	return rev
}
