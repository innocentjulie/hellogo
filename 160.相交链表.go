/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1 := reverseList(headA)
	l2 := reverseList(headB)
	var list *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return list
		} else {
			curr := l1
			next := l1.Next
			curr.Next = list
			list = curr
			curr = next
		}

	}
	return list

}

// @lc code=end

