package c7_easy

import (
	"fmt"
	"testing"
)

func TestSelf(t *testing.T) {
	ret := reverse(-2147483648)
	fmt.Println(ret)
}

func TestSolution(t *testing.T) {
	ret := reverseSolution(1534236469)
	fmt.Println(ret)
}
