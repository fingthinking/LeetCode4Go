package leetcode

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb"))
	t.Log(lengthOfLongestSubstring("bbbb"))
	t.Log(lengthOfLongestSubstring("pwwkewk"))
	t.Log(lengthOfLongestSubstring(""))
	t.Log(lengthOfLongestSubstring(" "))
	t.Log(lengthOfLongestSubstring("abcdefg"))
	t.Log(lengthOfLongestSubstring("aab"))
}
