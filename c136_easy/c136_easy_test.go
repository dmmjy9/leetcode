package c136_easy

import (
	"fmt"
	"testing"
)

func TestSelf(t *testing.T) {
	nums := []int {2, 2, 1}
	fmt.Println(singleNumber(nums))
}

func TestBitM(t *testing.T) {
	nums := []int {2, 1, 2, 4, 1}
	fmt.Println(singleNumberBitM(nums))
}
