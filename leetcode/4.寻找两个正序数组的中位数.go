package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m == 0 && n == 0 {
		return float64(-1)
	}
	// 保证num1的长度<=num2的长度
	if m > n {
		// 交换数组
		return findMedianSortedArrays(nums2, nums1)
	}
	if m == 0 {
		if n%2 == 0 {
			return (float64(nums2[(n-1)/2]) + float64(nums2[n/2])) / 2
		}
		return float64(nums2[n/2])
	}
	// 设i为数组1的切个点，j为数组2的切个点，保证切割后max(left) <= min(right)
	// i + j = m - i + n - j 【m + n为偶数时】
	// i + j = m - i + n - j + 1 【m + n为奇数时，保证左边长度大】
	// 则 j = (m + n + 1) / 2 - i, 其中/表示整除，则该公式完全成立
	// leftIdx表示二分查找时左侧索引， rightIdx表示二分查找时右索引
	leftIdx, rightIdx := 0, m
	var leftMax, rightMin int // 声明左侧最大值与右侧最小值
	var l1, l2, r1, r2 int
	for leftIdx <= rightIdx {
		i := (leftIdx + rightIdx) / 2
		j := (m+n+1)/2 - i

		// 处理边界情况
		if i == 0 {
			l1, l2 = math.MinInt64, nums2[j-1]
			leftMax = l2
		} else if j == 0 {
			l1, l2 = nums1[i-1], math.MinInt64
			leftMax = l1
		} else {
			l1, l2 = nums1[i-1], nums2[j-1]
			leftMax = max(l1, l2)
		}

		// 边界情况
		if j == n {
			r1, r2 = nums1[i], math.MaxInt64
			rightMin = r1
		} else if i == m {
			r1, r2 = math.MaxInt64, nums2[j]
			rightMin = r2
		} else {
			r1, r2 = nums1[i], nums2[j]
			rightMin = min(r1, r2)
		}

		// 符合条件，退出
		if leftMax <= rightMin {
			break
		}

		if l1 >= l2 {
			// 应该继续减小
			rightIdx = i - 1
		} else {
			leftIdx = i + 1
		}
	}

	if (m+n)%2 == 1 {
		// 奇数
		return float64(leftMax)
	}
	return (float64(leftMax) + float64(rightMin)) / 2
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

// @lc code=end
