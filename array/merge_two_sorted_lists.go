package array

type ListNode struct {
	Val  int
	Next *ListNode
}

//MergeTwoLists 21. 合并两个有序链表-https://leetcode.cn/problems/merge-two-sorted-lists/
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p := dummyHead

	for list1 != nil && list2 != nil {

		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	}

	if list2 != nil {
		p.Next = list2
	}

	return dummyHead.Next
}
