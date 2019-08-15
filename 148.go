package main

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	p := head
	q := head
	prve := p
	for q!=nil && q.Next!=nil {
		prve = p
		p = p.Next
		q = q.Next.Next
	}
	prve.Next = nil

	l := sortList(head)
	r := sortList(p)
	newHead := mergeTwoLists(l,r)

	return newHead
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	p := head
	for l1!=nil && l2!=nil {
		if l1.Val<l2.Val {
			p.Next = l1
			l1 = l1.Next
		}else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil{
		p.Next = l2
	}else {
		p.Next = l1
	}

	return head.Next

}
