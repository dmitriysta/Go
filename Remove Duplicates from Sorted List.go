//83. Remove Duplicates from Sorted List

//Given the head of a sorted linked list,
//delete all duplicates such that each element appears only once.
//Return the linked list sorted as well.


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	h := head
	for h.Next != nil{
		if h.Val == h.Next.Val {
			h.Next = h.Next.Next
		}else{
			h = h.Next
		}
	}
	return head
}