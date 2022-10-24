//206. Reverse Linked List

//Given the head of a singly linked list, reverse the list, and return the reversed list.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var reverse *ListNode
	for head != nil {
		temporarily := head.Next
		head.Next = reverse
		reverse = head
		head = temporarily
	}
	return reverse

}
