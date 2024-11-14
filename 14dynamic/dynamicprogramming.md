

## 14.2 动态规划问题特性
动态规划对问题进行递归分解，并且子问题是相互依赖的，可能会出现许多子问题重叠  
特性
1. 重叠子问题
2. 最优子结构
>给定一个楼梯，每步可以上1阶或者2阶，每一阶楼梯上都贴有一个非负整数，表示在该台阶需要付出的代价。给定非负整数数组cost，其中cost[i]表示在第i个台阶需要付出的代价，cost[0]为地面（起始点）。计数**最少**需要付出多少代价才能到达顶部？  
**原问题的最优解是从子问题的最优解构建得来的**
```
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
```
3. 无后效性
给定一个确定的状态， 它的未来发展只与当前状态有关，与过去经历的所有状态无关  
>带约束的爬楼梯  
给定一个共有n阶的楼梯，每步可以上1阶或者2阶，**但不能连续两轮跳1阶**，有多少种方案可以爬到楼顶？  
**此问题不满足无后效性**  
扩展状态定义：状态[i,j]表示处在第i阶并且上一轮跳了j阶，其中j属于{1,2}  
dp[i,1] = dp[i-1,2]  
dp[i,2] = dp[i-2,1] + dp[i-2,2]  
爬到第n阶的方案数：dp[n,1] + dp[n,2]
```
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
```

## 14.3 动态规划解题思路
1. 问题判断
**重叠子问题** **最优子结构** **无后效性**  
可以先观察问题是否适合使用回溯（穷举）解决（决策树模型）  
加分项  
找出最优解 最多（少）最大（小）  
问题的状态使用列表、多维矩阵、树表示，状态与周围状态存在递推关系  
减分项  
找出所有可能的解决方案  
明显的排列组合，需要返回具体的多个方案  
2. 问题求解步骤

