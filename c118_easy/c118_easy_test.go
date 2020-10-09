package c118_easy

import (
	"fmt"
	"testing"
)

func TestSelf(t *testing.T) {
	ret := generate(6)
	for _, val := range ret {
		fmt.Println(val)
	}
}
