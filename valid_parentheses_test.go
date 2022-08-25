package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	// s := "()[]{}"
	// s := "([])"
	// s := "()()()(){}{}[][][]{}"
	s := "()()[][]{}{}({{[[([()])]]}}){}"
	fmt.Println(isValid(s))
}
