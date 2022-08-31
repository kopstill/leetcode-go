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
