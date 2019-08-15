package main


type ListNode struct {
	Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p:=l1
	q:=l2

	x,y,carry := 0,0,0
	head := new(ListNode)
	curry:=head

	for p!=nil || q!=nil {
		x,y = 0,0
		if p!=nil {
			x = p.Val
		}
		if q!=nil {
			y = q.Val
		}
		total := x + y + carry
		node := new(ListNode)
		node.Val = total % 10
		carry = total / 10
		curry.Next = node
		curry = curry.Next

		if p!=nil {
			p = p.Next
		}
		if q!=nil {
			q = q.Next
		}
	}
	if carry >0 {
		node := new(ListNode)
		node.Val = carry
		curry.Next = node
	}
	return head.Next
}