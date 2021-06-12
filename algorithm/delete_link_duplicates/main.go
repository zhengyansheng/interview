package main

/*
https://www.nowcoder.com/practice/c087914fae584da886a0091e877f2c79?tpId=117&tab=answerKey

题目描述
	删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次

	例如：
	给出的链表为1\to1\to21→1→2,返回1 \to 21→2.
	给出的链表为1\to1\to 2 \to 3 \to 31→1→2→3→3,返回1\to 2 \to 31→2→3.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {

}
