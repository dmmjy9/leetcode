package c1672_easy

import (
	"fmt"
	"testing"
)

func TestSelf(t *testing.T) {
	var accounts [][]int
	accounts = append(accounts, []int{2, 8, 7})
	accounts = append(accounts, []int{7, 1, 3})
	accounts = append(accounts, []int{1, 9, 5})
	ret := maximumWealth(accounts)

	fmt.Println(ret)
}
