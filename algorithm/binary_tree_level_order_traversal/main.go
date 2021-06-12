package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder 二叉树层级遍历
/*
	leetcode: https://leetcode-cn.com/problems/binary-tree-level-order-traversal/submissions/
	讲解: bilibili.com/video/BV12k4y1C7jE?p=93
 */
func LevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{}
	q = append(q, root)

	for len(q) > 0 {
		length := len(q)
		level := []int{}
		for i := 0; i < length; i++ {
			// pop
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)

			// append a
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

