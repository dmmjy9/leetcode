package c1678_easy

import (
	"fmt"
	"testing"
)

func TestSelf(t *testing.T) {
	ret := interpret("G()()()()(al)")
	fmt.Println(ret)
}
