package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	// s := "()[]{}"
	// s := "([])"
	// s := "()()()(){}{}[][][]{}"
	// s := "()()[][]{}{}({{[[([()])]]}}){}"
	s := "["
	fmt.Println(isValid(s))
}

func Test1IsValid(t *testing.T) {
	// s := "()[]{}"
	// s := "()()[][]{}{}({{[[([()])]]}}){}"
	// s := "(([[]]))"
	// s := "}{}"
	s := "){"
	fmt.Println(isValid1(s))
}
