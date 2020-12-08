package c1672_easy

func maximumWealth(accounts [][]int) int {
	var maxWealth int
	for _, account := range accounts {
		wealth := 0
		for _, val := range account {
			wealth += val
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}
