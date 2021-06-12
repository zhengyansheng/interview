package main

import (
	"fmt"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 1 2 1
	// 1 2 2 1
	root := new(ListNode)
	root.Next = &ListNode{Val: 1}
	root.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	ok := isPalindrome1(root)
	fmt.Println(ok)
}

// 方法1: 时间复杂度O(N), 空间复杂度 O(N)
func isPalindrome1(head *ListNode) bool {
	arr := []int{}

	cur := head.Next
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true

}
