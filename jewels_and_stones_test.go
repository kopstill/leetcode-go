package leetcode

import (
	"fmt"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	// jewels := "aA"
	// stones := "aAAbbbb"
	jewels := "z"
	stones := "ZZ"
	fmt.Println(numJewelsInStones(jewels, stones))
}

func Test1NumJewelsInStones(t *testing.T) {
	jewels := "abc"
	stones := "asdfbsdfcsdfabc"
	fmt.Println(numJewelsInStones1(jewels, stones))
}

func Test2NumJewelsInStones(t *testing.T) {
	jewels := "leetcode"
	stones := "leetcodeleetcode"
	fmt.Println(numJewelsInStones2(jewels, stones))
}
