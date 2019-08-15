package main

import (
	"fmt"
)


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	test := []int{1,2,3,4,5}

	temp := make([]int,0,len(test))
	for j:=len(test)-1;j>=0;j--  {
		temp = append(temp, test[j])
	}
	test = temp
	fmt.Println(test)
	//levelOrder(nil)
}


func zigzagLevelOrder(root *TreeNode) [][]int {
	treeSet := make([]*TreeNode,0)
	result := make([][]int,0)
	flage := 1
	if root != nil {
		treeSet = append(treeSet, root)
		result = append(result, []int{root.Val})
	}

	for len(treeSet) != 0 {
		res := make([]int, 0)
		n := len(treeSet)
		flage++
		for i:=0;i<n;i++  {
			temp := treeSet[i]
			l := temp.Left
			r := temp.Right

			lval := -1
			rval := -1

			if l != nil {
				treeSet = append(treeSet, l)
				lval = l.Val
				res = append(res, lval)
			}
			if r != nil {
				treeSet = append(treeSet, r)
				rval = r.Val
				res = append(res,rval)
			}

		}

		if flage % 2 == 0 {
			temp := make([]int,0,len(res))
			for j:=len(res)-1;j>=0;j--  {
				temp = append(temp, res[j])
			}
			res = temp
			fmt.Println(res,flage)
		}


		if len(res)>0 {
			result = append(result, res)
		}
		treeSet = treeSet[n:]
	}
	return result
}

func everse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

