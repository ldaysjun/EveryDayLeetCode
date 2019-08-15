package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headB == nil || headA == nil {
		return nil
	}
	pA := headA
	pB := headB
	for pA != pB {
		if pA!=nil{
			pA = pA.Next
		}else {
			pA = headB
		}
		if pB!=nil {
			pB = pB.Next
		}else {
			pB = headA
		}
	}
	return pA
}