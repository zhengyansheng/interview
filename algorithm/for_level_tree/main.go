package main

import "fmt"

/*

给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
https://blog.csdn.net/s15738841819/article/details/84259014
*/

//var result [][]int

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func createNode(value int) *TreeNode {
	return &TreeNode{
		Left:  nil,
		Right: nil,
		Value: value,
	}
}

// 层级遍历
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	pool := make([]*TreeNode, 0)
	pool = append(pool, root)
	for len(pool) > 0 {
		count := len(pool)
		children := make([]int, 0)
		for count > 0 {
			node := pool[0]
			count--
			if node.Left != nil {
				pool = append(pool, node.Left)
			}
			if node.Right != nil {
				pool = append(pool, node.Right)
			}
			fmt.Println(pool)
			children = append(children, pool[0].Value)
			pool = pool[1:]
		}
		res = append(res, children)
	}
	return res
}

func main() {
	root := createNode(3)
	root.Left = createNode(9)
	root.Right = createNode(20)
	root.Right.Left = createNode(15)
	root.Right.Right = createNode(7)
	ret := levelOrder(root)
	fmt.Println(ret)
}
