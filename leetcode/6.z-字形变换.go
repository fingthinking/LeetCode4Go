package leetcode

import (
	"unicode/utf8"
)

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	// 算出每个字符应处于的位置
	// 时间复杂度O(N), 空间复杂度O(N)
	rLen := utf8.RuneCountInString(s)
	sRunes := []rune(s)
	targetRunes := make([]rune, rLen)

	r := 0
	idx := 0
	step := 2 * (numRows - 1)
	if step == 0 {
		step = 1
	}
	rowIdx := 0
	for r < numRows && idx < rLen {
		lIdx, rIdx := rowIdx*step+r, (rowIdx+1)*step-r
		println(lIdx, rIdx)
		if lIdx >= rLen { // 进入下一行条件
			r++
			rowIdx = 0
			continue
		} else {
			targetRunes[idx] = sRunes[lIdx]
			idx++
			rowIdx++
			println("===")
		}
		// 最后一行和第一行不需要rIdx赋值
		if lIdx == rIdx || lIdx+step == rIdx {
			continue
		}
		if rIdx < rLen {
			targetRunes[idx] = sRunes[rIdx]
			idx++
		} else {
			r++
			rowIdx = 0
		}
	}
	return string(targetRunes)
}

// @lc code=end
