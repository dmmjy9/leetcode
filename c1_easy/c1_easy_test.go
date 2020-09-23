package c1_easy

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	input := []int{3, 2, 4}
	ret := twoSum(input, 6)
	fmt.Println(ret)
}
