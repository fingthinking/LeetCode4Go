package leetcode

import (
	"testing"
)

func TestAddTwoSum(t *testing.T) {
	n1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	n2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	sum := addTwoNumbers(n1, n2)
	t.Logf("====%v", sum)
}
