package main


func main() {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 4

	node1.Next = node2
	node2.Next = node3


	node4 := new(ListNode)
	node1.Val = 1
	node5 := new(ListNode)
	node1.Val = 3
	node6 := new(ListNode)
	node1.Val = 4

	node4.Next = node5
	node5.Next = node6

}

type ListNode struct {
	Val int
	Next *ListNode
}



func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	p := head

	for l1!=nil && l2!=nil {
		if l1.Val < l2.Val{
			p.Next = l1
			l1 = l1.Next
		}else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 == nil {
		p.Next = l2
	}else{
		p.Next = l1
	}
	return head.Next
}


//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	head := new(ListNode)
//	p := new(ListNode)
//	head.Next = p
//	for l1!=nil || l2!=nil {
//		if l1 == nil {
//			p.Next = l2
//			break
//		}
//		if l2 == nil {
//			p.Next = l1
//			break
//		}
//
//		if l1.Val < l2.Val{
//			p.Next = l1
//			l1 = l1.Next
//		}else {
//			p.Next = l2
//			l2 = l2.Next
//		}
//		p = p.Next
//
//	}
//
//	return head.Next.Next
//}



