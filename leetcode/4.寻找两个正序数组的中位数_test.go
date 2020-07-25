package leetcode

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log(findMedianSortedArrays([]int{1, 2, 5, 7}, []int{2, 4, 5}))
	t.Log(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	t.Log(findMedianSortedArrays([]int{1, 2, 5}, []int{2, 4, 5}))
	t.Log(findMedianSortedArrays([]int{1, 3}, []int{2}))
	t.Log(findMedianSortedArrays([]int{1}, []int{2}))
	t.Log(findMedianSortedArrays([]int{1}, []int{1}))
	t.Log(findMedianSortedArrays([]int{}, []int{2}))
}
