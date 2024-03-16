package main

import "fmt"

//https://leetcode.com/problems/binary-root-preorder-traversal/
// Решение через рекурсивный подход

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var root *TreeNode

	fmt.Println(preorderTraversal(root))
}

func preorderTraversal(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	handle(root, &result)

	return result
}

func handle(root *TreeNode, result *[]int) {
	*result = append(*result, root.Val)
	if root.Left != nil {
		handle(root.Left, result)
	}
	if root.Right != nil {
		handle(root.Right, result)
	}
}
