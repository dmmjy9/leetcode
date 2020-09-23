package c1_easy

func twoSum(nums []int, target int) []int {
	var n int
	var ret []int
	hash := make(map[int]int)

	for idx, val := range nums {
		n = target - val
		if _, ok := hash[n]; !ok {
			hash[val] = idx
		} else {
			ret = append(ret, hash[n])
			ret = append(ret, idx)
			return ret
		}
	}
	return ret
}
