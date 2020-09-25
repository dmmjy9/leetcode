package c9_easy

func isPalinddome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	var rev = 0
	for ; x > rev; x /= 10 {
		rev = rev * 10 + x % 10
	}
	return x == rev || x == rev/10
}
