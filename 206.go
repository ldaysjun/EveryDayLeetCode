package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	prev := head
	cur := head.Next

	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		if temp == nil {
			break
		}
		cur = temp
	}
	head.Next = nil
	return cur
}
