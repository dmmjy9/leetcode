package c14_easy

import (
	"fmt"
	"testing"
)

func TestSelf(t *testing.T) {
	input := []string {
		"flower",
		"flow",
		"flight",
	}
	output := longestCommonPrefix(input)
	fmt.Println(output)
}
