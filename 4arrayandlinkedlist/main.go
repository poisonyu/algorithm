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

// 148 排序链表
// func sortList(head *ListNode) *ListNode {
// 	return sort(head, nil)
// }

// func sort(head, tail *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	if head.Next == tail {
// 		head.Next = nil
// 		return head
// 	}
// 	slow, fast := head, head
// 	for fast != tail {
// 		slow = slow.Next
// 		fast = fast.Next
// 		if fast != nil {
// 			fast = fast.Next
// 		}
// 	}
// 	mid := slow
// 	return merge(sort(head, mid), sort(mid, tail))
// }
// func merge(head1, head2 *ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	temp, temp1, temp2 := dummy, head1, head2
// 	for temp1 != nil && temp2 != nil {
// 		if temp1.Val < temp2.Val {
// 			temp.Next = temp1
// 			temp1 = temp1.Next
// 		} else {
// 			temp.Next = temp2
// 			temp2 = temp2.Next
// 		}
// 		temp = temp.Next
// 	}
// 	if temp1 != nil {
// 		temp.Next = temp1
// 	} else {
// 		temp.Next = temp2
// 	}
// 	return dummy.Next
// }

// 链表排序
// 归并排序，每一次都将链表一分为二（通过快慢指针来找到链表中点），
// 需要传入链表的头节点和尾节点
// 直到链表元素只有一个时，合并链表

func sortList(head *ListNode) *ListNode {
	return merge(head, nil)
}

func merge(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}
	//快慢指针找中点
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return sort(merge(head, mid), merge(mid, tail))
}

func sort(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp, left, right := dummy, head1, head2
	for left != nil && right != nil {
		if left.Val <= right.Val {
			temp.Next = left
			left = left.Next
		} else {
			temp.Next = right
			right = right.Next
		}
		temp = temp.Next
	}
	if left != nil {
		temp.Next = left
	} else {
		temp.Next = right
	}
	return dummy.Next
}

// 143重排链表
// 方法一线性表 遍历链表，放进切片中，再原地组合链表
// func reorderList(head *ListNode) {

// 	var count, i int
// 	node := head
// 	for node != nil {
// 		count++
// 		node = node.Next
// 	}
// 	nodeTable := make([]*ListNode, count)
// 	for head != nil {
// 		nodeTable[i] = head
// 		i++
// 		head = head.Next
// 	}
// 	for i, j := 0, count-i-1; i < j; i++ {
// 		nodeTable[i].Next = nodeTable[j]
// 		if i+1 < j {
// 			nodeTable[j].Next = nodeTable[i+1]
// 		} else {
// 			nodeTable[j].Next = nil
// 		}
// 	}

// }

// 方法二 找链表中点，分成两个链表，把后面的链表反转，再合并

func reorderList(head *ListNode) {
	// 快慢指针找中点
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow
	r := reverse(mid.Next)
	mid.Next = nil
	merge2(head, r)

}

// func reverse(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	node := reverse(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return node
// }

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	for head != nil {
		n := head.Next
		head.Next = prev
		prev = head
		if n == nil {
			break
		}
		head = n
	}
	return head
}
func merge2(head1, head2 *ListNode) {
	// dummy := &ListNode{}
	// temp, left, right := dummy , head1, head2
	var n1, n2 *ListNode
	for head1 != nil && head2 != nil {
		n1, n2 = head1.Next, head2.Next
		head1.Next = head2
		head1 = n1
		head2.Next = head1
		head2 = n2
		// head1.Next = head2
	}
}

// 141 环形链表
// 快慢指针
// 快指针会追上慢指针 有环
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 142 环形链表ii
func detectCycle(head *ListNode) *ListNode {

}
