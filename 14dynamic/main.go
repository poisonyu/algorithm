package main

import (
	"fmt"
	"math"
	"sort"
)

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

// func minPathSumDFS(grid [][]int)
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

// func main() {
// 	// a := []int{1, 2}
// 	// fmt.Println(subsetSum(a, 3))
// 	// fmt.Println(climbingStairsDFS(5))
// 	// fmt.Println(climbingStairsDFSMem(5))
// 	// fmt.Println(climbingStairsDP(5))
// 	// fmt.Println(climbingStairsDPComp(5))
// 	// cost := []int{0, 1, 10, 1}
// 	// fmt.Println(minCostClimbingStairsDP(cost))

// 	fmt.Println(climbingStairsConstraintDP(4))

// 	n, m := 4, 4
// 	grid := make([][]int, n)
// 	grid[0] = []int{1, 3, 1, 5}
// 	grid[1] = []int{2, 2, 4, 2}
// 	grid[2] = []int{5, 3, 2, 1}
// 	grid[3] = []int{4, 3, 5, 2}
// 	fmt.Println(dynamic(n, m, grid))
// }

// 在限定背包容量下，能放入物品的最大值
// 01背包 回溯 暴力搜索
func knapsackDFS(wgt, val []int, i, c int) int {
	// 当前没用物品或背包容量为0
	if i == 0 || c == 0 {
		return 0
	}
	// 剪枝 当前物品的容量大于当前背包的容量，只能选择不放入物品
	if c < wgt[i-1] {
		return knapsackDFS(wgt, val, i-1, c)
	}
	// 不放入当前物品，返回当前背包中的物品价值
	no := knapsackDFS(wgt, val, i-1, c)
	// 放入当前物品，返回当前背包中的物品价值
	yes := knapsackDFS(wgt, val, i-1, c-wgt[i-1]) + val[i-1]
	return int(math.Max(float64(no), float64(yes)))
}

// 记忆化搜索 重叠子问题，只计算一次
func knapsackDFSMem(wgt, val []int, mem [][]int, i, c int) int {
	// 当前没用物品或背包容量为0
	if i == 0 || c == 0 {
		return 0
	}
	if mem[i][c] != -1 {
		return mem[i][c]
	}
	// 剪枝 当前物品的容量大于当前背包的容量，只能选择不放入物品
	if c < wgt[i-1] {
		return knapsackDFSMem(wgt, val, mem, i-1, c)
	}
	// 不放入当前物品，返回当前背包中的物品价值
	no := knapsackDFSMem(wgt, val, mem, i-1, c)
	// 放入当前物品，返回当前背包中的物品价值
	yes := knapsackDFSMem(wgt, val, mem, i-1, c-wgt[i-1]) + val[i-1]
	mem[i][c] = int(math.Max(float64(no), float64(yes)))
	return mem[i][c]
}

// 动态规划
func knapsackDP(wgt, val []int, cap int) int {
	n := len(wgt)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, cap+1)
	}
	// dp[0][cap-wgt
	for i := 1; i < n+1; i++ {
		for c := 1; c < cap+1; c++ {
			if c < wgt[i-1] {
				dp[i][c] = dp[i-1][c]
			} else {
				dp[i][c] = int(math.Max(float64(dp[i-1][c]), float64(dp[i-1][c-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[n][cap]

}

// 空间优化
func knapsackDPComp(wgt, val []int, cap int) int {
	n := len(wgt)
	dp := make([]int, cap+1)
	for i := 1; i < n+1; i++ {
		for c := cap; c > 0; c-- {
			if wgt[i-1] <= c {
				dp[c] = int(math.Max(float64(dp[c]), float64(dp[c-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[cap]

}

// 完全背包 每个物品可以重复选取
func unboundKnapsackDP(wgt, val []int, cap int) int {
	n := len(wgt)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, cap+1)
	}

	for i := 1; i < n+1; i++ {
		for c := 1; c < cap+1; c++ {
			if c < wgt[i-1] {
				dp[i][c] = dp[i-1][c]
			} else {
				dp[i][c] = int(math.Max(float64(dp[i-1][c]), float64(dp[i][c-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[n][cap]
}

func unboundKnapsackDPComp(wgt, val []int, cap int) int {
	n := len(wgt)
	dp := make([]int, cap+1)
	for i := 1; i < n+1; i++ {
		for c := 1; c < cap+1; c++ {
			if wgt[i-1] <= c {
				dp[c] = int(math.Max(float64(dp[c]), float64(dp[c-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[cap]
}

func greedy(amt int, coins []int) (total int) {
	for i := len(coins) - 1; i >= 0; i-- {
		coin := coins[i]
		total += amt / coin
		amt = amt % coin
	}

	return
}

func moneyDFS(i, amt int, coins []int) int {
	if amt == 0 || i == 0 {
		return amt + 1
	}
	if amt < coins[i-1] {
		return moneyDFS(i-1, amt, coins)
	}
	no := moneyDFS(i-1, amt, coins)
	yes := moneyDFS(i, amt-coins[i-1], coins) + 1
	return int(math.Min(float64(no), float64(yes)))
}
func testknapsack() {
	n := 5
	cap := 50
	wgt := []int{10, 20, 30, 40, 50}
	val := []int{50, 120, 150, 210, 240}
	fmt.Printf("有%d个物品，背包容量为%d，可以得到最大价值为%d\n", n, cap, knapsackDPComp(wgt, val, cap))

}
func coinChangeDP(amt int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amt+1)
	}
	// 当i=0时，没有硬币数量，无法凑出大于1的目标金额，是无效解
	for a := 1; a < len(dp[0]); a++ {
		dp[0][a] = amt + 1
	}
	for i := 1; i < n+1; i++ {
		for a := 1; a < amt+1; a++ {
			if a < coins[i-1] {
				dp[i][a] = dp[i-1][a]
			} else {
				dp[i][a] = int(math.Min(float64(dp[i-1][a]), float64(dp[i][a-coins[i-1]]+1)))
			}
		}
	}
	if dp[n][amt] == amt+1 {
		return -1
	}
	return dp[n][amt]

}
func coinChangeDPComp(amt int, coins []int) int {
	n := len(coins)
	dp := make([]int, amt+1)
	for i := 1; i < amt+1; i++ {
		dp[i] = amt + 1
	}
	for i := 1; i < n+1; i++ {
		for a := 1; a < amt+1; a++ {
			if coins[i-1] <= a {
				dp[a] = int(math.Min(float64(dp[a]), float64(dp[a-coins[i-1]]+1)))
			}
		}
	}
	if dp[amt] != amt+1 {
		return dp[amt]
	}
	return -1
}

func coinChangeIIDP(amt int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amt+1)
	}
	// 当目标金额为0时，也是一种情况，有且只有一种硬币组合数量，就是不选硬币，
	for i := 0; i < n+1; i++ {
		dp[i][0] = 1
		if i != 0 {
			for a := 1; a < amt+1; a++ {
				if a < coins[i-1] {
					dp[i][a] = dp[i-1][a]
				} else {
					dp[i][a] = dp[i-1][a] + dp[i][a-coins[i-1]]
				}
			}
		}
	}
	return dp[n][amt]

}

func coinChangeIIDPComp(amt int, coins []int) int {
	n := len(coins)
	dp := make([]int, amt+1)
	// 当目标金额为0时，也是一种情况，有且只有一种硬币组合数量，就是不选硬币，
	dp[0] = 1
	for i := 1; i < n+1; i++ {
		for a := 1; a < amt+1; a++ {
			if coins[i-1] <= a {
				dp[a] = dp[a] + dp[a-coins[i-1]]
			}
		}
	}
	return dp[amt]

}

func editDistanceDP(s, t string) int {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for j := 1; j < m+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i < n+1; i++ {
		dp[i][0] = i
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			// 当s和t末尾的字符相等时，无需操作
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Min(float64(dp[i][j-1]), math.Min(float64(dp[i-1][j]), float64(dp[i-1][j-1])))) + 1
			}
		}

	}
	return dp[n][m]
}

func editDistanceDPComp(s, t string) int {
	n, m := len(s), len(t)
	dp := make([]int, m+1)
	// 当s为空字符串时，更改为t字符串需要t的长度次数的操作
	for j := 1; j < m+1; j++ {
		dp[j] = j
	}
	for i := 1; i < n+1; i++ {
		leftup := dp[0] // dp[i-1][j-1]
		dp[0] = i       // dp[i, j-1]
		for j := 1; j < m+1; j++ {
			temp := dp[j] // dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[j] = leftup
			} else {
				dp[j] = int(math.Min(float64(dp[j-1]), math.Min(float64(dp[j]), float64(leftup)))) + 1
			}
			leftup = temp
		}
	}
	return dp[m]
}

// func main() {
// 	// coins := []int{1, 2, 5}
// 	// fmt.Println(greedy(11, coins))
// 	// fmt.Println(moneyDFS(3, 11, coins))
// 	// fmt.Println(coinChangeDP(11, coins))
// 	// fmt.Println(coinChangeDPComp(11, coins))
// 	// fmt.Println(coinChangeIIDP(5, coins))
// 	// fmt.Println(coinChangeIIDPComp(5, coins))
// 	// a := []int{1, 2}
// 	// fmt.Println(subsetSum(a, 3))
// 	// fmt.Println(climbingStairsDFS(5))
// 	// fmt.Println(climbingStairsDFSMem(5))
// 	// fmt.Println(climbingStairsDP(5))
// 	// fmt.Println(climbingStairsDPComp(5))
// 	// cost := []int{0, 1, 10, 1}
// 	// fmt.Println(minCostClimbingStairsDP(cost))

// 	// fmt.Println(climbingStairsConstraintDP(4))

// 	// testknapsack()

// 	fmt.Println(editDistanceDPComp("kitten", "sitting"))
// 	fmt.Println(editDistanceDP("hello", "algo"))

// }

// type Item struct {
// 	w int
// 	v int
// }
// func fractionalKnapsack(wgt, val []int,cap int) float64 {
// 	items := make([]Item, len(wgt))
// 	for i:=0 ; i< len(wgt); i++ {
// 		items[i] = Item{w: wgt[i], v: val[i]}
// 	}
// 	sort.Slice(items, func(i, j int) bool {
// 		return float64(items[i].v)/float64(items[i].w) > float64(items[j].v)/float64(items[j].w)
// 	})
// 	for cap > 0 {

// 	}
// 	return
// }

// 分数背包问题 最大化单位重量下的物品价值 每轮贪心的选择单位价值最高的物品
type Item struct {
	w     int
	v     int
	price float64
}

func fractionalKnapsack(wgt, val []int, cap int) float64 {
	items := make([]Item, len(wgt))
	for i := 0; i < len(wgt); i++ {
		items[i] = Item{w: wgt[i], v: val[i], price: float64(val[i]) / float64(wgt[i])}
	}
	sort.Slice(items, func(i, j int) bool {
		// return float64(items[i].v)/float64(items[i].w) > float64(items[j].v)/float64(items[j].w)
		return items[i].price > items[j].price
	})
	var total float64
	for i := 0; i < len(wgt); i++ {
		if cap < items[i].w {
			total += float64(cap) * items[i].price
			break
		} else {
			total += float64(items[i].v)
			cap -= items[i].w
		}
	}
	return total
}

// 最大容量问题
func maxCapacity(ht []int) (cap int) {
	i, j := 0, len(ht)-1
	for i < j {
		l, h := j-i, 0
		if ht[i] < ht[j] {
			h = ht[i]
			i++
		} else {
			h = ht[j]
			j--
		}
		c := l * h
		// c := (j-i) * int(math.Min(float64(ht[i]), float64(ht[j])))
		if c > cap {
			cap = c
		}
	}
	return
}

// 最大切分乘积问题
//
//	func maxProductCutting(n int) (res int) {
//		if n > 3 {
//			a := n / 3
//			res *= a * 3
//			// n = n % 3
//			if n == 1 {
//				// res = (a-1)*3 + 2*2
//				res = int(math.Pow(3, float64(a-1)) * 2 * 2)
//			} else if n == 2 {
//				res = int(math.Pow(3, float64(a)) * 2)
//			} else {
//				res = int(math.Pow(3, float64(a)))
//			}
//		} else {
//			// if n == 3 {
//			// 	res = n - 1
//			// } else if n == 2 {
//			// 	res = 1
//			// }
//			res = 1 * (n - 1)
//		}
//		return
//	}

// 最大切分乘积问题 >=4的整数都应该继续切分，最优切分因子为3，最多只应存在两个2
func maxProductCutting(n int) (res int) {
	if n <= 3 {
		res = 1 * (n - 1)
	}
	a := n / 3
	n = n % 3
	// 余数为1时，1 * 3 < 2 * 2, 应该将最后一个3替换为2
	if n == 1 {
		res = int(math.Pow(3, float64(a-1)) * 2 * 2)
	} else if n == 2 {
		res = int(math.Pow(3, float64(a)) * 2)
	} else {
		res = int(math.Pow(3, float64(a)))
	}
	return
}
func main() {
	// wgt := []int{10, 20, 30, 40, 50}
	// val := []int{50, 120, 150, 210, 240}
	// fmt.Println(fractionalKnapsack(wgt, val, 50))
	a := []int{3, 8, 5, 2, 7, 7, 3, 4}
	fmt.Println(maxCapacity(a))
	fmt.Println(maxProductCutting(11))
}
