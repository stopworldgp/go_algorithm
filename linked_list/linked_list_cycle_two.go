package linked_list

//DetectCycle 142. 环形链表 II-https://leetcode.cn/problems/linked-list-cycle-ii/
func DetectCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	//0. 快慢指针第一次相遇，注意判断 fast 和fast.next 都不能为空
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//1. 当相遇时跳出，不能将该判断放到for里，因为第一次 fast= slow = head
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	//1. 将slow指针放到head处，两个指针开始从相同步数开始走
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
