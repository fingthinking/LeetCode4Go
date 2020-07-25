package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	idxs := twoSum([]int{2, 7, 11, 15}, 9)
	if reflect.DeepEqual(idxs, []int{0, 1}) {
		t.Log("Assert True")
	}
}
