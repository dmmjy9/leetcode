package c5_medium

func longestPalindrome(s string) string {
	strLength := len(s)
	maxLength := 0
	maxStr := ""

	if strLength > 1000 {
		return ""
	}

	if strLength == 1 {
		return s
	}
	if strLength == 2 {
		if s[0] != s[1] {
			return string(s[0])
		} else {
			return s
		}
	}

	for i := 0; i < strLength; i++ {
		for j := i + 1; j < strLength; j++ {
			subStr := s[i:j+1]
			if isHw(subStr) && len(subStr) > maxLength {
				maxLength = len(subStr)
				maxStr = subStr
			}
		}
	}

	if maxLength == 0 {
		return string(s[0])
	}

	return maxStr
}

func isHw(s string) bool {
	strLength := len(s)
	halfLength := len(s) / 2

	j := strLength - 1
	for i := 0; i < halfLength; i++ {
		if s[i] != s[j] {
			return false
		}
		j -= 1
	}
	return true
}
