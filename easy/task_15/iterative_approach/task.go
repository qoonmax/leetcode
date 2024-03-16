package main

import "fmt"

//https://leetcode.com/problems/binary-root-preorder-traversal/
// Решение через итеративный подход

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(preorderTraversal(root))
}

func preorderTraversal(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		result = append(result, curr.Val)
		stack = stack[:len(stack)-1]

		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}

	return result
}
