package c121_easy

import (
	"fmt"
	"testing"
)

func TestSelf(t *testing.T) {
	price := []int {7, 1, 5, 3, 6, 4}
	ret := maxProfit(price)
	fmt.Println(ret)
}
