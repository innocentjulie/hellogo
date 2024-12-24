/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var list *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = list
		list = curr
		curr = next
	}
	return list
	// return reverseList(head.Next)
	// return list
}

// @lc code=end

