package linked_list

//GetIntersectionNode 160. 相交链表-https://leetcode.cn/problems/intersection-of-two-linked-lists/
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	//0. 判断是否有交点
	for a != b {
		//1. 当链表遍历完更换链表遍历
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
