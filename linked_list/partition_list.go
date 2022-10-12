package linked_list

//Partition 86. 分隔链表-https://leetcode.cn/problems/partition-list/
func Partition(head *ListNode, x int) *ListNode {

	smallHead := &ListNode{}
	small := smallHead
	bigHead := &ListNode{}
	big := bigHead
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			big.Next = head
			big = big.Next
		}
		head = head.Next
	}
	big.Next = nil
	small.Next = bigHead.Next

	return smallHead.Next

}
