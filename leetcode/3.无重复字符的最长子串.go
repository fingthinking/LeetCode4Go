package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// 长度小于2可直接返回结果
	sLen := len(s)
	if sLen < 2 {
		return sLen
	}
	// 基于滑动窗口
	runeCache := make(map[rune]int, sLen) // rune: idx
	tail := 0                             // 定义快慢指针
	subLen := 0
	for idx, r := range s {
		if rIdx, ok := runeCache[r]; ok {
			tmpLen := idx - tail
			if tmpLen > subLen {
				subLen = tmpLen
			}
			if rIdx >= tail {
				// tail移动
				tail = rIdx + 1
			}
		}
		runeCache[r] = idx
	}
	// 最长子串在最后的情况
	if sLen-tail > subLen {
		subLen = sLen - tail
	}
	return subLen
}

// @lc code=end
