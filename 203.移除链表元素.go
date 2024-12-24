/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var dfs func(head *ListNode) *ListNode
	dfs = func(head *ListNode) *ListNode {
		if head == nil {
			return head
		}
		head.Next = dfs(head.Next)
		if head.Val == val {
			return head.Next
		} else {
			return head
		}

	}
	return dfs(head)
}

// @lc code=end

