package leetcode

import (
	"fmt"
	"testing"
)

func TestAddDigits(t *testing.T) {
	num := 12138
	fmt.Println(addDigits(num))
}

func Test1AddDigits(t *testing.T) {
	num := 9527
	fmt.Println(addDigits1(num))
}

func Test2AddDigits(t *testing.T) {
	num := 4242
	fmt.Println(addDigits2(num))
}
