package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	test := []int{1,2,3,4,5}
	n := len(test)
	test = append(test, 6)
	t := test[n:]
	fmt.Println(t)
}


func levelOrder(root *TreeNode) [][]int {
	treeSet := make([]*TreeNode,0)
	result := make([][]int,0)
	if root != nil {
		treeSet = append(treeSet, root)
		result = append(result, []int{root.Val})
	}

	for len(treeSet) != 0 {
		res := make([]int, 0)
		n := len(treeSet)
		for i:=0;i<n;i++  {
			temp := treeSet[i]
			l := temp.Left
			r := temp.Right
			if l != nil {
				res = append(res, l.Val)
				treeSet = append(treeSet, l)
			}
			if r != nil {
				res = append(res, r.Val)
				treeSet = append(treeSet, r)
			}
		}

		if len(res)>0 {
			result = append(result, res)
		}
		treeSet = treeSet[n:]
	}
	return result
}