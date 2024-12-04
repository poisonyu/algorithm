package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 83 删除排序链表中的重复元素
// 关键信息是链表是排序的
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	root := head
	cur := head.Val
	for head != nil && head.Next != nil {
		if head.Next.Val != cur {
			cur = head.Next.Val
			head = head.Next
		} else {
			// 下一个节点和head相等时，head指向下下个节点
			// 下一个循环再次判断head与下下个节点是否相同
			head.Next = head.Next.Next
		}
	}
	return root
}

// 82 删除排序链表中的重复元素II
// 所有重复数字的节点都要删除
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	root := &ListNode{Val: 101, Next: head}
	cur := root
	for cur.Next != nil && cur.Next.Next != nil {
		// 在一个节点和下下个节点相等时，出现了重复节点，
		// 先保存当前节点，
		if cur.Next.Val == cur.Next.Next.Val {
			// 保存当前节点
			node := cur
			// 再循环比较删除重复节点
			v := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == v {
				cur = cur.Next
			}
			node.Next = cur.Next
			cur = node
		} else {
			cur = cur.Next
		}
	}
	return root.Next
}

// 206 反转链表
// 遍历
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	var prev *ListNode
// 	cur := head
// 	for cur != nil && cur.Next != nil {
// 		next := cur.Next // 2
// 		cur.Next = prev
// 		prev = cur
// 		cur = next
// 	}
// 	cur.Next = prev
// 	return cur
// }

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

// 92 反转链表II
// 本题的目标是将第left个链表到第right个链表反转
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	cur := dummy
	for i := 0; i < left-1; i++ {
		cur = cur.Next
	}
	node := cur.Next
	leftnode := node
	var prev *ListNode

	for j := 0; j < right-left+1; j++ { // 2  3
		next := node.Next // 3 // 4
		node.Next = prev
		prev = node
		node = next
	}
	leftnode.Next = node
	cur.Next = prev
	return dummy.Next
}

// 21 合并两个有序链表
// 遍历
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil {
// 		return list2
// 	}
// 	if list2 == nil {
// 		return list1
// 	}
// 	dummy := &ListNode{Next: nil}
// 	cur := dummy
// 	for list1 != nil && list2 != nil {
// 		if list1.Val < list2.Val {
// 			cur.Next = list1

// 			list1 = list1.Next
// 		} else {
// 			cur.Next = list2
// 			list2 = list2.Next
// 		}
// 		cur = cur.Next
// 	}
// 	if list1 == nil {
// 		cur.Next = list2
// 	} else {
// 		cur.Next = list1
// 	}
// 	return dummy.Next
// }

// 递归
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

// 86 分隔链表
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	tail := &ListNode{Next: nil}
	large := tail
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val < x {
			cur = cur.Next
		} else {
			large.Next = cur.Next
			cur.Next = cur.Next.Next
			large = large.Next
		}
	}
	large.Next = nil
	cur.Next = tail.Next
	return dummy.Next
}

// 排序链表
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}
func merge(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp, temp1, temp2 := dummy, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val < temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else {
		temp.Next = temp2
	}
	return dummy.Next
}

// func mergeSort(nums []int, i, j int) {

// 	if i >= j {
// 		return
// 	}
// 	mid := i + (j-i)/2
// 	mergeSort(nums, i, mid)
// 	mergeSort(nums, mid+1, j)
// 	merge(nums, i, mid, j)
// }

// func merge(nums []int, left, mid, right int) {
// 	temp := make([]int, right-left+1)
// 	i, j, k := left, mid+1, 0
// 	for i < j && j <= right {
// 		if nums[i] <= nums[j] {
// 			temp[k] = nums[i]
// 			i++
// 		} else {
// 			temp[k] = nums[j]
// 			j++
// 		}
// 		k++
// 	}
// 	for i < j {
// 		temp[k] = nums[i]
// 		i++
// 		k++
// 	}
// 	for j <= right {
// 		temp[k] = nums[j]
// 		j++
// 		k++
// 	}
// 	for i, num := range temp {
// 		nums[left+i] = num
// 	}
// }
