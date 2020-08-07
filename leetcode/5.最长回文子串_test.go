package leetcode

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("abaccabacca"))
	t.Log(longestPalindrome("abbc"))
	t.Log(longestPalindrome("abcbc"))
	t.Log(longestPalindrome("acde"))
}
