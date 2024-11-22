package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

// 二叉树的最近公共祖先
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

// 二叉树的层序遍历

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
			// next = append(next, node.Left, node.Right)
		}
		ans = append(ans, a)
		// ans = append(ans, append([]int{}, a...))
		q = next
	}
	return ans
}

// ans := make([][]int, 0)
// ans1 := make([]int, 0)

// onelevel := make([]*TreeNode, 0)
// onelevel = append(onelevel, root)
// for len(onelevel) > 0 {
// 	node := onelevel[0]
// 	onelevel = onelevel[1:len(onelevel)]
// 	if len(onelevel) == 0 {

// 	}
// 	ans1 = append(ans1, node.Val)
// 	onelevel = append(onelevel, node.Left, node.Right)
// }
// ans := make([][]int, 0)
// queue := list.New()
// queue.PushBack(root)
// for queue.Len() > 0 {
// 	ans1 := make([]int, 0)
// 	nextqueue := list.New()
// 	for queue.Len() > 0 {
// 		node := queue.Remove(queue.Front()).(*TreeNode)
// 		ans1 = append(ans1, node.Val)
// 		if node.Left != nil {
// 			nextqueue.PushBack(node.Left)
// 		}
// 		if node.Right != nil {
// 			nextqueue.PushBack(node.Right)
// 		}
// 	}
// 	ans = append(ans, ans1)
// 	queue = nextqueue
// }
// return ans
