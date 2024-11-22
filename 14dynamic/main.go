package main

import "fmt"

// 给定一个共有n阶的楼梯，每步可以上1阶或者2阶，请问有多少种方案可以爬到楼顶

func backtrack(target int, state *[]int, nums *[]int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, *state...))
		return
	}

	for _, num := range *nums {
		if target-num < 0 {
			// nums排序，用break,没排序，用continue
			continue
		}
		*state = append(*state, num)
		backtrack(target-num, state, nums, res)
		*state = (*state)[:len(*state)-1]
	}
}

func subsetSum(nums []int, target int) [][]int {
	state := make([]int, 0)
	res := make([][]int, 0)
	backtrack(target, &state, &nums, &res)
	return res
}

// 暴力搜索
func dfs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return dfs(n-1) + dfs(n-2)
}

func climbingStairsDFS(n int) int {
	return dfs(n)
}

// 记忆化搜索
// 重叠子问题都只被计算一次

func dfsMem(i int, mem []int) int {
	if i == 1 || i == 2 {
		return i
	}
	// fmt.Printf("dp[%d]: %v\n", i, mem)
	// dp[i]存在，直接返回
	if mem[i] != -1 {
		return mem[i]
	}
	// dp[i]不存在，递归计算dp[i]
	count := dfsMem(i-1, mem) + dfsMem(i-2, mem)
	// 记录dp[i]
	mem[i] = count
	return count
}

func climbingStairsDFSMem(n int) int {
	mem := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		mem[i] = -1
	}
	return dfsMem(n, mem)
}

// 动态规划
func climbingStairsDP(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbingStairsDPComp(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i < n+1; i++ {
		a, b = b, a+b
	}
	return b
}

func minCostClimbingStairsDP(cost []int) int {
	c := len(cost)
	n := c - 1
	if n == 1 || n == 2 {
		return cost[n]
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 初始化dp，存储子问题的解
	dp := make([]int, c)
	dp[1] = cost[1]
	dp[2] = cost[2]
	// 状态转移 从较小的子问题逐步求解较大子问题
	for i := 3; i < c; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return dp[n]
}

// 空间优化
func minCostClimbingStairsDPComp(cost []int) int {
	c := len(cost)
	n := c - 1
	if n == 1 || n == 2 {
		return cost[n]
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 初始状态：预设最小子问题的解
	a, b := cost[1], cost[2]
	// 状态转移 从较小的子问题逐步求解较大子问题
	for i := 3; i < c; i++ {
		a, b = b, min(a, b)+cost[i]
	}
	return b
}

// 带约束爬楼梯
func climbingStairsConstraintDP(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[1][1] = 1
	dp[2][2] = 1
	for i := 3; i < n+1; i++ {
		dp[i][1] = dp[i-1][2]
		dp[i][2] = dp[i-2][1] + dp[i-2][2]
	}
	return dp[n][1] + dp[n][2]
}

func minPathSumDFS(grid [][]int)
func dynamic(n, m int, grid [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}
			if i > 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[n-1][m-1]
}

func bag(start, n int, wgt, val []int) int {
	// return

	for i := start; i < n; i++ {
		// 放入
		bag()
		// 不放入

	}
}

func main() {
	// a := []int{1, 2}
	// fmt.Println(subsetSum(a, 3))
	// fmt.Println(climbingStairsDFS(5))
	// fmt.Println(climbingStairsDFSMem(5))
	// fmt.Println(climbingStairsDP(5))
	// fmt.Println(climbingStairsDPComp(5))
	// cost := []int{0, 1, 10, 1}
	// fmt.Println(minCostClimbingStairsDP(cost))

	fmt.Println(climbingStairsConstraintDP(4))

	n, m := 4, 4
	grid := make([][]int, n)
	grid[0] = []int{1, 3, 1, 5}
	grid[1] = []int{2, 2, 4, 2}
	grid[2] = []int{5, 3, 2, 1}
	grid[3] = []int{4, 3, 5, 2}
	fmt.Println(dynamic(n, m, grid))
}
