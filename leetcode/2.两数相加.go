package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	return fmt.Sprintf("ListNode:{Val=%d,Next=%v}", node.Val, node.Next)
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	head := root
	// 构造进位与临时变量
	var tmpVal, carry int
	for l1 != nil && l2 != nil {
		tmpVal = l1.Val + l2.Val + carry
		head.Val = tmpVal % 10
		carry = tmpVal / 10
		// 指针前进
		l1, l2 = l1.Next, l2.Next
		if l1 == nil || l2 == nil {
			break
		}
		head.Next = &ListNode{}
		head = head.Next
	}

	for l1 != nil {
		head.Next = &ListNode{}
		head = head.Next
		tmpVal = l1.Val + carry
		head.Val = tmpVal % 10
		carry = tmpVal / 10
		l1 = l1.Next
	}

	for l2 != nil {
		head.Next = &ListNode{}
		head = head.Next
		tmpVal = l2.Val + carry
		head.Val = tmpVal % 10
		carry = tmpVal / 10
		l2 = l2.Next
	}

	if carry != 0 {
		head.Next = &ListNode{}
		head = head.Next
		head.Val = carry
	}
	return root
}

// @lc code=end
