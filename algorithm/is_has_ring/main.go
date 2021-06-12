package main

import "fmt"

func main() {
	root := &ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	root.Next = &node2
	root.Next.Next = &node3
	root.Next.Next.Next = &node4
	root.Next.Next.Next.Next = &node2

	//printLinkList(root)
	ok := hasRingV2(root)
	fmt.Println("has cycle", ok)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	哈希表
	时间复杂度: O(N)
	空间复杂度: O(N)
*/
func hasRing(head *ListNode) bool {
	lookup := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := lookup[head]; ok {
			return true
		}
		lookup[head] = struct{}{}
		head = head.Next
	}
	return false
}

/*
	快慢指针
	时间复杂度: O(N)
	空间复杂度: O(1)
*/
func hasRingV2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	quick := head.Next
	for slow != quick {
		if slow == nil || quick == nil {
			return false
		}
		slow = slow.Next
		quick = quick.Next.Next
	}
	return true
}

// 打印
func printLinkList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}
