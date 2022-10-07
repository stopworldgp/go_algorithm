package linked_list

//MergeTwoLists 21. 合并两个有序链表-https://leetcode.cn/problems/merge-two-sorted-lists/
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//0.建立虚拟头结点
	dummyHead := &ListNode{}
	p := dummyHead

	//1. 场景1 双链表不为空比较
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
	//2. 场景2 其中一个链表为空，包含了 两种，一种开始一个链表为空和第二种 两个链表不一边长
	if list1 != nil {
		p.Next = list1
	}

	if list2 != nil {
		p.Next = list2
	}

	return dummyHead.Next
}
