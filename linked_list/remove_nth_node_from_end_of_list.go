package linked_list

//RemoveNthFromEnd 19. 删除链表的倒数第 N 个结点-https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	//1. slow = 虚拟节点，因为要知道倒数第n的前一个节点用于next
	dummy := &ListNode{0, head}
	slow, fast := dummy, head
	//2. fast先走n步
	for i := 0; i < n; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
