package leetcode

import (
	"unicode/utf8"
)

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
// 动态规划实现
func longestPalindrome(s string) string {
	// 时间复杂度 O(N^2) 空间复杂度O(N^2)
	// 长度小于2，直接返回
	// 如果字符串收尾不相同，则不是回文
	// 如果字符串收尾相同，则继续判断
	runeLen := utf8.RuneCountInString(s)
	if runeLen < 2 {
		return s
	}
	// 状态定义
	// p[i][j] : 表示字符串i...j是否是回文，其中i,j为左闭右闭区间
	// p[i][j] = p[i+1][j-1] & s[i]==s[j]

	p := make([][]bool, runeLen)

	for i := 0; i < runeLen; i++ {
		p[i] = make([]bool, runeLen)
	}

	maxLen := 1
	begin := 0
	runes := []rune(s)
	for j := 1; j < runeLen; j++ {
		for i := 0; i < j; i++ {
			if runes[i] != runes[j] {
				// 表示不为回文数
				p[i][j] = false
			} else {
				if j-i < 3 { // 表示相等，且只有3个数，则一定为true
					p[i][j] = true
				} else {
					p[i][j] = p[i+1][j-1]
				}
			}

			if p[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return string(runes[begin : begin+maxLen])

}

// @lc code=end
