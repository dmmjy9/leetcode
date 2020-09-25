package c14_easy

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		for ; strings.Index(strs[i], prefix) != 0; {
			prefix = string([]byte(prefix)[:len(prefix)-2])
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}
