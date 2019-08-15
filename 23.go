package main


type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	mergeKLists(nil)
}


func mergeKLists(lists []*ListNode) *ListNode {
	var newListHead *ListNode
	for i:=0;i<len(lists);i++{
		l1 := newListHead
		l2 := lists[i]

		newListHead = mergeTwoLists(l1,l2)
	}
	return newListHead
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	p := head
	for l1!=nil && l2!=nil {
		if l1.Val < l2.Val {
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
	}else {
		p.Next = l1
	}
	return head.Next
}
