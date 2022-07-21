package leetcode

import (
	"fmt"
	"testing"
)

func Test0LengthOfLongestSubstring(t *testing.T) {
	// s := "你你好，,世界！"
	s := "Hello, 你好呀花花，希望你你健康成长！！！"
	fmt.Println(lengthOfLongestSubstring(s))
}

func Test1LengthOfLongestSubstring1(t *testing.T) {
	// s := "abcabcbb"
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring1(s))
}

func Test2LengthOfLongestSubstring2(t *testing.T) {
	// s := "bbbbb"
	// s := "dvdf"
	// s := " "
	s := ""
	fmt.Println(lengthOfLongestSubstring2(s))
}
