package leetcode

import (
	"fmt"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	n := 84
	fmt.Println(guessNumber(n))
}

func Test1GuessNumber(t *testing.T) {
	n := 1000
	fmt.Println(guessNumber1(n))
}
