package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var values [][]int
var tmpValue []int

func threeOrders(root *TreeNode) [][]int {
	// write code here
	values = make([][]int, 0)
	tmpValue = make([]int, 0)
	preOrder(root)
	values = append(values, tmpValue)
	tmpValue = []int{}
	InOrder(root)
	values = append(values, tmpValue)
	tmpValue = []int{}
	PostOrder(root)
	values = append(values, tmpValue)

	return values
}

func preOrder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Val, ",")
		tmpValue = append(tmpValue, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}

}

func InOrder(root *TreeNode) {
	if root != nil {
		InOrder(root.Left)
		fmt.Print(root.Val, ",")
		tmpValue = append(tmpValue, root.Val)
		InOrder(root.Right)
	}
}

func PostOrder(root *TreeNode) {
	if root != nil {
		PostOrder(root.Left)
		PostOrder(root.Right)
		fmt.Print(root.Val, ",")
		tmpValue = append(tmpValue, root.Val)
	}
}

func createLink(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	root := createLink(1)
	root.Left = createLink(2)
	root.Right = createLink(3)
	fmt.Print(threeOrders(root))
}
