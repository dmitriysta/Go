//876. Middle of the Linked List

//Given the head of a singly linked list, return the middle node of the linked list.
//
//If there are two middle nodes, return the second middle node.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	slowly := head
	faster := head

	for faster != nil && faster.Next != nil {
		slowly = slowly.Next
		faster = faster.Next.Next
	}
	return slowly

}

