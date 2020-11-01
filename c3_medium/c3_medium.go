package c3_medium

import "math"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	j := 0
	strMap := make(map[uint8]int)

	for i := 0; i < n; i++ {
		if _, ok := strMap[s[i]]; ok {
			if strMap[s[i]] > j {
				j = strMap[s[i]]
			}
		}
		ans = int(math.Max(float64(ans), float64(i-j+1)))
		strMap[s[i]] = i + 1
	}
	return ans
}
