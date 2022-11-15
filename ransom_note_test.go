package leetcode

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	// ransomNote := "aabc"
	// magazine := "xcabokay"
	ransomNote := "bg"
	magazine := "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func Test1CanConstruct(t *testing.T) {
	// ransomNote := "aabc"
	// magazine := "xcabokay"
	// ransomNote := "bg"
	// magazine := "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"
	ransomNote := "aa"
	magazine := "ab"
	fmt.Println(canConstruct1(ransomNote, magazine))
}

func Test2CanConstruct(t *testing.T) {
	// ransomNote := "aabc"
	// magazine := "xcabokay"
	ransomNote := "bg"
	magazine := "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"
	// ransomNote := "aa"
	// magazine := "ab"
	fmt.Println(canConstruct2(ransomNote, magazine))
}
