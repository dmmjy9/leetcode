package c8_medium

import (
	"math"
)

func myAtoi(s string) int {
	i, sign, result := 0, 1, 0
	if len(s) == 0 {
		return 0
	}

	for i < len(s) && s[i] == ' ' {
		i += 1
	}

	if i < len(s) && s[i] == '+' || s[i] == '-' {
		if s[i+1] == '-' {
			sign = -1
		} else {
			sign = 1
		}
	}

	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		if result > math.MaxInt32 / 10 || (result == math.MinInt32 / 10 && int(s[i] - '0') > math.MinInt32 % 10) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		result = result * 10 + int(s[i+1] - '0')
	}
	return result * sign
}
