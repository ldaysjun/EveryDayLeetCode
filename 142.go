package main

import "fmt"

//func detectCycle(head *ListNode) *ListNode {
//	p := head
//	q := head
//	isLoop := false
//	for q!=nil && q.Next!=nil {
//		p = p.Next
//		q = q.Next.Next
//		if p == q {
//			isLoop = true
//			break
//		}
//	}
//	if !isLoop {
//		fmt.Println("no cycle")
//		return nil
//	}else {
//		i:=0
//		for head != nil {
//			if p == head {
//				fmt.Println("tail connects to node index ",i)
//				break
//			}
//			i++
//			head = head.Next
//			p = p.Next
//		}
//	}
//	return p
//
//}


func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return nil
	}

	m := make(map[*ListNode]int, 0)
	i:=0
	for head != nil{
		if _, ok := m[head]; ok{
			fmt.Println("tail connects to node index ",i)
			return head
		}
		i++
		m[head] = 0
		head = head.Next
	}
	fmt.Println("no cycle")
	return nil
}