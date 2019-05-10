package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
}

func main() {

	node1 := new(ListNode)
	node1.Val = 1

	node2 := new(ListNode)
	node2.Val = 2
	//
	node3 := new(ListNode)
	node3.Val = 3

	node4 := new(ListNode)
	node4.Val = 4

	node5 := new(ListNode)
	node5.Val = 5
	//
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	test := removeNthFromEnd(node1,5)
	for test != nil{
		fmt.Println("结果",test.Val)
		test = test.Next
	}
}



//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	if head == nil{
//		return nil
//	}
//	if head.Next == nil {
//		return nil
//	}
//	p := head
//	total := 1
//	for p.Next != nil{
//		p = p.Next
//		total ++
//	}
//	target := total - n
//	if target == 0 {
//		return head.Next
//	}
//	p = head
//	q := p.Next
//	for i:=1;i<target;i++  {
//		if p.Next != nil {
//			p = p.Next
//		}
//		if q.Next != nil {
//			q = q.Next
//		}
//	}
//
//	p.Next = q.Next
//	return head
//}



func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil{
		return nil
	}
	if head.Next == nil {
		return nil
	}
	p := head
	q := head
	i := 1

	for p.Next != nil {
		if i >= n+1 {
			q = q.Next
		}
		p = p.Next
		i ++
	}

	if q == head && i == n{
		return head.Next
	}
	q.Next = q.Next.Next

	return head
}