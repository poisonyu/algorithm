package main

import (
	"math"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(num int) *TreeNode {
	return &TreeNode{Left: nil, Right: nil, Val: num}
}

// 二叉搜索数
// 根节点的值大于左子节点值，小于右子节点的值
// 左右子节点也满足上面条件

func BinarySearchTree(root *TreeNode, num int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > num {
		return BinarySearchTree(root.Left, num)
	} else if root.Val < num {
		return BinarySearchTree(root.Right, num)
	} else {
		return root
	}
}

type binarySearchTree struct {
	root *TreeNode
}

// 查找节点
func (bst *binarySearchTree) search(num int) *TreeNode {
	root := bst.root
	for root != nil {
		if root.Val < num {
			root = root.Right
		} else if root.Val > num {
			root = root.Left
		} else {
			break
		}
	}
	return root
}

// 插入节点
func (bst *binarySearchTree) insert(num int) {
	node := bst.root
	// 根节点为空
	if node == nil {
		bst.root = NewTreeNode(num)
		return
	}
	var pre *TreeNode
	for node != nil {
		if node.Val == num {
			// 节点已经存在，无需插入节点
			return
		}
		pre = node
		if node.Val < num {
			node = node.Right
		} else {
			// pre = node
			node = node.Right
		}
	}
	newNode := NewTreeNode(num)
	if pre.Val < num {
		pre.Right = newNode
	} else {
		pre.Left = newNode
	}
}

// 删除节点
func (bst *binarySearchTree) remove(num int) {
	cur := bst.root
	if cur == nil {
		return
	}
	var pre *TreeNode
	for cur != nil {
		if cur.Val == num {
			break
		}
		pre = cur
		if cur.Val < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	// 无待删除节点
	if cur == nil {
		return
	}
	// if cur.Left == nil && cur.Right == nil {
	// 	if pre.Left.Val == num {
	// 		pre.Left = nil
	// 	} else {
	// 		pre.Right = nil
	// 	}
	// }
	// 子节点数为0或1
	if cur.Left == nil || cur.Right == nil {
		var node *TreeNode
		if cur.Left != nil {
			node = cur.Left
		} else {
			node = cur.Right
		}
		if cur != bst.root {
			if pre.Left == cur {
				pre.Left = node
			} else {
				pre.Right = node
			}
		} else {
			bst.root = node
		}
		// 子节点数为2
	} else {
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		// 递归删除节点tmp
		bst.remove(tmp.Val)
		// 通过覆盖cur的值来删除cur
		cur.Val = tmp.Val
	}
}

// func inOrder(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	inOrder(root.Left)
// 	root.Val
// 	inOrder(root.Right)
// }

// 543 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree?envType=problem-list-v2&envId=f2qQnjWL
func diameterOfBinaryTree(root *TreeNode) int {
	var res int = 1
	maxDepth(root, &res)
	return res - 1
}

// 返回当前节点为根节点的最大深度
// 根节点到最远叶节点的最大路径的节点数
func maxDepth(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	// 左子的最大深度
	l := maxDepth(root.Left, res)
	// 右子树的最大深度
	r := maxDepth(root.Right, res)
	// 当前根节点左右子树组成的路径
	*res = int(math.Max(float64(*res), float64(l+r+1)))
	return int(math.Max(float64(l), float64(r))) + 1
}

// 124 二叉树的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum?envType=problem-list-v2&envId=f2qQnjWL
// 路径中至少包含一个节点
func maxPathSum(root *TreeNode) int {
	var s int = math.MinInt
	pathSum(root, &s)
	return s
}

// 当前根节点的最大
func pathSum(root *TreeNode, s *int) int {
	if root == nil {
		return 0
	}
	l := pathSum(root.Left, s)
	r := pathSum(root.Right, s)
	// 左子节点最大路径和为负数时，将和记为0，即不选这条路径
	if l < 0 {
		l = 0
	}
	// 右子节点最大路径和为负数时，将和记为0，即不选这条路径
	if r < 0 {
		r = 0
	}
	*s = int(math.Max(float64(*s), float64(l+r+root.Val)))
	return int(math.Max(float64(l), float64(r))) + root.Val
}

// 236 二叉树的最近公共祖先
// p、q均存在与给定二叉树中
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// 102 二叉树的层序遍历
// 在将节点放入队列中时，需要先判断节点是否为nil,需要注意下面三处nil的判断
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		a := make([]int, 0)
		next := make([]*TreeNode, 0)
		for len(q) > 0 {
			node := q[0]
			q = q[1:len(q)]
			a = append(a, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		ans = append(ans, a)
		// ans = append(ans, append([]int{}, a...))
		q = next
	}
	return ans
}

// 107 二叉树的层序遍历II
// 将102的结果反转即可
func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		ans1 := make([]int, 0)
		q := make([]*TreeNode, 0)
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			ans1 = append(ans1, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, ans1)
		queue = q
	}
	slices.Reverse(ans)
	return ans
}

// 103 二叉树的锯齿形层序遍历
// 正常层序遍历，把偶数层的结果反转放入ａｎｓ中

// 下面是自己的思路与实现
// 将每一层的节点放进一个栈中，层数分为奇偶层，
// 区别是放入左右子节点的先后顺序不同
// 第1层（奇数层）先放入左子节点
// 第2层（偶数层）先放入右子节点
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	level := 1
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		// 开始一层的遍历
		ans1 := make([]int, 0)
		next := make([]*TreeNode, 0)
		// 每一次循环从当前层中取出一个节点
		for len(queue) > 0 {
			node := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if level%2 == 0 {
				if node.Right != nil {
					next = append(next, node.Right)
				}
				if node.Left != nil {
					next = append(next, node.Left)
				}
			} else {
				if node.Left != nil {
					next = append(next, node.Left)
				}
				if node.Right != nil {
					next = append(next, node.Right)
				}
			}
			ans1 = append(ans1, node.Val)
		}
		// 进入下一层前给level+1
		level++
		ans = append(ans, ans1)
		queue = next
		// 结束一层的遍历
	}
	return ans
}
