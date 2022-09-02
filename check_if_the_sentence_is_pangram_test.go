package leetcode

import (
	"fmt"
	"testing"
)

func TestCheckIfPangram(t *testing.T) {
	// sentence := "abcdefghijklmnopqrstuvwxyz"
	// sentence := "thequickbrownfoxjumpsoverthelazydog"
	sentence := "leetcode"
	fmt.Println(checkIfPangram(sentence))
}

func Test1CheckIfPangram(t *testing.T) {
	// sentence := "abcdefghijklmnopqrstuvwxyz"
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	// sentence := "leetcode"
	fmt.Println(checkIfPangram1(sentence))
}
