package main

func main() {
	// 1. 链表反转 1, 2, 3
	// 2. 快排 []int{3, 2, 1, 5}
	// 3. 链表是否有环
	// 4. 层级遍历
	// 5. dfs

}

func quickSort(arr []int, start, end int) {
	if start > end {
		return
	}
	v := arr[start]
	i, j := start, end
	for i < j {
		for i < j && v < arr[j] {
			j--
		}
		if i < j {
			arr[j] = arr[i]
			i++
		}
		for i < j && v > arr[i] {
			i++
		}
		if i < j {
			arr[i] = arr[j]
			j--
		}
	}
	arr[i] = v
	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

type LinkList struct {
	Val  int
	Next *LinkList
}

func reverseLinkList(head *LinkList) *LinkList {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *LinkList
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev

}

func HasRing(head *LinkList) bool {
	// 1. hashmap
	if head == nil || head.Next == nil {
		return false
	}

	cur := head
	lookup := make(map[*LinkList]struct{})
	for cur != nil {
		if _, ok := lookup[cur]; ok {
			return true
		} else {
			lookup[cur] = struct{}{}
		}
		cur = cur.Next
	}
	return false
}

func HasRingV2(head *LinkList) bool {
	// 1. quick slow point
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	quick := head.Next
	cur := head
	for slow != quick {
		if slow == nil || quick == nil {
			return false
		}
		slow = slow.Next
		quick = quick.Next.Next
		cur = cur.Next
	}
	return true

}
