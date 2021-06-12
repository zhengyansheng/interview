package main

import "fmt"

/*
	单链表
		- 基本操作
			- 访问
			- 随机追加
			- 头部添加
			- 尾部添加
			- 删除节点
			- 修改节点
		- 反转
		- 判断链表是否有环
*/

type LinkList struct {
	Val  int
	Next *LinkList
}

// 创建链表
func createLinkList(i int) *LinkList {
	return &LinkList{
		Val:  i,
		Next: nil,
	}
}

// 批量创建链表
func bulkCreateLinkList(arr []int) *LinkList {
	head := createLinkList(arr[0])
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = createLinkList(arr[i])
		cur = cur.Next
	}
	return head
}

// 打印全部链表
func printLinkList(head *LinkList) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
	fmt.Println()
}

// 追加元素到链表头部
func addHeadItem(head *LinkList, i int) *LinkList {
	root := createLinkList(i)
	root.Next = head
	return root
}

// 追加元素到链表尾部
func addTailItem(head *LinkList, i int) {
	for head.Next != nil {
		head = head.Next
	}
	head.Next = createLinkList(i)
}

// InsertHeadItem 随机插入到指定位置 idx
// idx >= 0
func InsertHeadItem(head *LinkList, idx int, i int) {
	upIdx := 0
	node := createLinkList(i)
	for head != nil {
		if idx == upIdx {
			tmp := head.Next
			head.Next = node
			node.Next = tmp
		}
		head = head.Next
		upIdx++
	}
}

func deleteItem(head *LinkList, i int) *LinkList {
	pre := new(LinkList)
	cur := head
	for head != nil {
		if head.Val == i {
			pre.Next = head.Next
			break
		}
		pre = head
		head = head.Next
	}
	return cur
}

func reverseLinkList(head *LinkList) *LinkList {
	// head 为空 或者 head 长度为1时
	if head == nil || head.Next == nil {
		return head
	}
	pre := new(LinkList)
	cur := head
	for cur != nil {
		tmp := cur.Next // last item is nil
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// isLinkRing 判断链表是否有环
// 快慢指针
/*
一个单向链表，判断链表中是否存在环，如果存在环找到环的入口位置并返回，如果不存在环则返回空。

思路：定义快、慢指针，从链表头开始，快指针每次走两步，慢指针每次走一步，如果相遇，说明有环。
	 碰撞之后慢指针回到链表头部，快慢指针每次走一步，第一次相遇就是环入口。
 */
func isLinkRing(head *LinkList) bool  {
	// head.Next
	// head.Next.Next
	slow, fast := head, head
	for head != nil && head.Next != nil && head.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next // panic
		if slow == fast {
			return true
		}
		head = head.Next
	}
	return false
}

func main() {
	
	/*
		head := createLinkList(1)
		head.Next = createLinkList(2)
		head.Next.Next = createLinkList(3)
		head.Next.Next.Next = createLinkList(4)
	
		printLinkList(head)
	*/
	
	/*
	arr := []int{1, 2, 3, 4, 5, 6}
	head := bulkCreateLinkList(arr)
	printLinkList(head)
	
	root := addHeadItem(head, 7)
	root = addHeadItem(root, 8)
	root = addHeadItem(root, 9)
	printLinkList(root)
	
	InsertHeadItem(root, 0, 3)
	printLinkList(root)
	
	addTailItem(root, 8)
	printLinkList(root)
	newHead := deleteItem(root, 1)
	printLinkList(newHead)
	
	
	//newHead := reverseLinkList(root)
	//printLinkList(newHead)
	*/
	
	head1 := createLinkList(1)
	head1.Next = createLinkList(2)
	head1.Next.Next = createLinkList(3)
	head1.Next.Next.Next = head1
	ok := isLinkRing(head1)
	fmt.Println(ok)
	
}
