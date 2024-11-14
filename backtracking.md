

## 回溯算法
通过穷举来解决问题，通常采用深度优先搜索
>例1给定一个二叉树，搜索并记录所有值为7的节点，返回节点列表
```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preOrderI(node *TreeNode, target int, res *[]*TreeNode) {
	if node == nil {
		return
	}
	if node.Val == target {
		*res = append(*res, node)
	}
	preOrderI(node.Left, target, res)
	preOrderI(node.Right, target, res)
}
```
1. 尝试与回退
>例2在二叉树中搜索所有值为7的节点，返回根节点到这些节点的路径
```
func preOrderII(root *TreeNode, target int, res *[][]*TreeNode, path *[]*TreeNode) {
	if root == nil {
		return
	}
	// 尝试 path用于记录访问过的节点
	*path = append(*path, root)
	if root.Val == target {
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderII(root.Left, target, res, path)
	preOrderII(root.Right, target, res, path)
	// 回退 删除path中的当前节点
	*path = (*path)[:len(*path)-1]
}
```
2. 剪枝
>例3在二叉树中搜索所有值为7的节点，返回根节点到这些节点的路径，**并要求路径中不包含值为3的节点**
```
func preOrderII(root *TreeNode, target int, res *[][]*TreeNode, path *[]*TreeNode) {
	// 剪枝
	if root == nil || root.Val == 3 {
		return
	}
	// 尝试 path用于记录访问过的节点
	*path = append(*path, root)
	if root.Val == target {
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderII(root.Left, target, res, path)
	preOrderII(root.Right, target, res, path)
	// 回退 删除path中的当前节点
	*path = (*path)[:len(*path)-1]
}
```
3. 框架代码
```
// 回溯通用代码 剪枝 尝试 回退
// state表示问题的当前状态
// choices表示当前状态下的选择 
func backtrack(state *State, choices []Choice, res *[]State) {
	// 判断是否是解
	if isSolution(state) {
		// 记录解
		recordSolution(state, res)
		// return
	}
	for _, choice := range choices {
		// 剪枝
		if isValid(state, choice) {
			// 尝试 做出选择，更新状态
			makeChoice(state, choice)
			backtrack(state, choices, res)
			// 回退 撤销选择，恢复之前的状态
			undoChoice(state, choice)
		}
	}
}
```

```
// 回溯框架处理例3
func backtrack(state *[]*TreeNode, choices *[]*TreeNode, res *[][]*TreeNode) {
	if isSolution(state) {
		recordSolution(state, res)
	}
	for _, choice := range *choices {
		if isValid(state, choice) {
			makeChoice(state, choice)
			temp := make([]*TreeNode, 0)
			temp = append(temp, choice.Left, choice.Right)
			backtrack(state, &temp, res)
			undoChoice(state)
		}
	}
}
func undoChoice(state *[]*TreeNode) {
	*state = (*state)[:len(*state)-1]
}
func makeChoice(state *[]*TreeNode, choice *TreeNode) {
	*state = append(*state, choice)
}
func isValid(state *[]*TreeNode, choice *TreeNode) bool {
	return choice != nil && choice.Val != 3
}
func recordSolution(state *[]*TreeNode, res *[][]*TreeNode) {
	*res = append(*res, append([]*TreeNode{}, (*state)...))
}
func isSolution(state *[]*TreeNode) bool {
	return len(*state) != 0 && (*state)[len(*state)-1].Val == 7
}
```

6. 回溯典型例题
搜索问题
* 全排列问题
* 子集和问题
* 汉诺塔问题
约束满足问题
* n皇后
* 数独
* 图着色问题
组合优化问题