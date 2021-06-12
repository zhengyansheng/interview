package main

import "fmt"

/*
	二叉树
		- 遍历
			- 前序遍历 PreOrder   根-左-右
			- 中序遍历 InOrder    左-根-右
			- 后序遍历 PostOrder  左-右-根

			- 深度优先遍历
			- 广度优先遍历
			- 搜索
 */

// BinaryTree 定义二叉树节点
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// CreateTree 新建节点
func CreateTree(value int) *BinaryTree {
	return &BinaryTree{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func BulkCreateTree(arr []int) *BinaryTree {
	root := CreateTree(arr[0])

	node := CreateTree(arr[1])
	if root.Value > arr[1] {
		root.Left = node
	} else {
		root.Right = node
	}
	return root

}

func (bt *BinaryTree) GetTreeHeight(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}
	lHeight := bt.GetTreeHeight(tree.Left)
	rHeight := bt.GetTreeHeight(tree.Right)
	if lHeight > rHeight {
		return lHeight + 1
	} else {
		return rHeight + 1
	}
}

// PreOrder 递归前序遍历
func (bt *BinaryTree) PreOrder(tree *BinaryTree) {
	if tree != nil {
		fmt.Printf("%v, ", tree.Value)
		bt.PreOrder(tree.Left)
		bt.PreOrder(tree.Right)
	}
}

// InOrder 递归中序遍历
func (bt *BinaryTree) InOrder(tree *BinaryTree) {
	if tree != nil {
		bt.InOrder(tree.Left)
		fmt.Printf("%v, ", tree.Value)
		bt.InOrder(tree.Right)
	}
}

// PostOrder 递归后序遍历
func (bt *BinaryTree) PostOrder(tree *BinaryTree) {
	if tree != nil {
		bt.PostOrder(tree.Left)
		bt.PostOrder(tree.Right)
		fmt.Printf("%v, ", tree.Value)
	}
}

func main() {
	root := CreateTree(5)
	root.Left = CreateTree(2)
	root.Right = CreateTree(4)
	root.Left.Right = CreateTree(7)
	root.Left.Right.Left = CreateTree(6)
	root.Right.Left = CreateTree(8)
	root.Right.Right = CreateTree(9)

	treeHeight := root.GetTreeHeight(root)
	fmt.Printf("binary tree heigh: %d\n", treeHeight)

	root.PreOrder(root)
	fmt.Println()
	root.InOrder(root)
	fmt.Println()
	root.PostOrder(root)

	fmt.Println("")

	arr := []int{4, 2, 5, 1, 6, 3, 7}
	bt := BulkCreateTree(arr)
	bt.InOrder(bt)
	fmt.Println(bt)
	fmt.Println(bt.Left, bt.Right)

}
