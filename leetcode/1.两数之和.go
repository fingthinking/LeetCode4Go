package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	cache := make(map[int]int, 0)
	for idx, num := range nums {
		if preIdx, ok := cache[target-num]; ok {
			return []int{preIdx, idx}
		}
		cache[num] = idx
	}
	return nil
}

// @lc code=end
