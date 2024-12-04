package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
)

// 栈

// // 初始化栈
// var stack []int
// // 元素入栈
// stack = append(stack, 1)
// // 访问栈顶元素
// peek := stack[len(stack)-1]
// // 元素出栈
// pop := stack[len(stack)-1]
// stack = stack[:len(stack)-1]
// // 获取栈长度
// size := len(stack)
// // 判断栈是否为空
// isEmpty := len(stack) == 0

//基于链表实现栈
// 链表的头节点代表栈顶
// 入栈 将元素插入链表头部
// 出栈 将头节点从链表头部删除

type LinkedListStack struct {
	data *list.List
}

// 初始化栈
func newLinkedListStack() *LinkedListStack {
	return &LinkedListStack{data: list.New()}
}

// 入栈
func (s *LinkedListStack) push(value int) {
	s.data.PushBack(value)
}

// 访问栈顶元素
func (s *LinkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	return e.Value
}

// 出栈
func (s *LinkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Remove(s.data.Back())
}

// 获取栈的长度
func (s *LinkedListStack) size() int {
	return s.data.Len()
}

// 判断栈是否为空
func (s *LinkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

func (s *LinkedListStack) toList() *list.List {
	return s.data
}

// 基于数组实现栈
// 将数组的尾部作为栈顶
// 入栈 在数组尾部添加元素
// 出栈 在数组尾部删除元素
type arrayStack struct {
	data []int
}

func newArrayStack() *arrayStack {
	return &arrayStack{data: make([]int, 0, 16)}
}

func (s *arrayStack) size() int {
	return len(s.data)
}

func (s *arrayStack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *arrayStack) push(value int) {
	s.data = append(s.data, value)
}
func (s *arrayStack) pop() any {
	val := s.peek()
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *arrayStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}
func (s *arrayStack) toSlice() []int {
	return s.data
}

// 队列
// 先入先出
// 入队 把元素添加到队尾
// 出队 删除队首元素

// 初始化队列
// queue := list.New()
// 元素入队
// queue.PushBack(1)
// 访问队首元素
// peek := queue.Front()
// 元素出队
// queue.Remove(queue.Front())
// 获取队列的长度
// size := queue.Len()
// 队列是否为空
// isEmpty := queue.Len() == 0

// 基于链表实现队列
// 链表头节点 队首 删除节点
// 链表尾节点 队尾 添加节点

type linkedListQueue struct {
	data *list.List
}

func newLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{data: list.New()}
}

func (s *linkedListQueue) push(value any) {
	s.data.PushBack(value)
}

func (s *linkedListQueue) pop() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Remove(s.data.Front())
}

func (s *linkedListQueue) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Front().Value
}
func (s *linkedListQueue) size() int {
	return s.data.Len()
}

func (s *linkedListQueue) isEmpty() bool {
	return s.data.Len() == 0
}

func (s *linkedListQueue) toList() *list.List {
	return s.data
}

// 基于数组实现队列
// 用两个变量front, rear分别指向队首元素的索引和队尾元素的索引
// 队列长度size=rear-front
// 数组中包含元素的有效区间为[front, rear-1]
// 入队 将元素赋值给rear索引处，并给size加1
// 出队 front加1，并将size减1
// 通过取余操作来实现环形数组

type arrayQueue struct {
	nums        []int
	front, rear int
	queCapacity int
}

func newArrayQueue(cap int) *arrayQueue {
	return &arrayQueue{
		nums:        make([]int, cap),
		front:       0,
		rear:        0,
		queCapacity: cap,
	}
}

func (q *arrayQueue) push(v int) {
	if q.size() == q.queCapacity {
		return
	}
	q.nums[q.rear] = v
	q.rear++
	q.rear = q.rear % q.queCapacity
}

func (q *arrayQueue) peek() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

func (q *arrayQueue) pop() any {
	val := q.peek()
	if val == nil {
		return nil
	}
	q.front++
	q.front = q.front % q.queCapacity
	return val
}

func (q *arrayQueue) size() int {
	if q.rear < q.front {
		return q.rear + q.queCapacity - q.front
	}
	return q.rear - q.front
}
func (q *arrayQueue) isEmpty() bool {
	return q.rear-q.front == 0
}

func (q *arrayQueue) toSlice() []int {
	if q.front <= q.rear {
		return q.nums[q.front:q.rear]
	}
	return append(q.nums[:q.rear], q.nums[q.front:]...)
}

// 双向队列
// 基于双向链表实现双向队列
type linkedListDeque struct {
	data *list.List
}

func newLinkedListDeque() *linkedListDeque {
	return &linkedListDeque{
		data: list.New(),
	}
}
func (s *linkedListDeque) pushFirst(value any) {
	s.data.PushFront(value)
}
func (s *linkedListDeque) pushLast(value any) {
	s.data.PushBack(value)
}
func (s *linkedListDeque) popFirst() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Remove(s.data.Front())
}

func (s *linkedListDeque) popLast() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Remove(s.data.Back())
}
func (s *linkedListDeque) peekFirst() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Front().Value
}
func (s *linkedListDeque) peekLast() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Back().Value
}
func (s *linkedListDeque) size() int {
	return s.data.Len()
}
func (s *linkedListDeque) isEmpty() bool {
	return s.data.Len() == 0
}

// 基于数组实现双向队列
type arrayDeque struct {
	nums                     []int
	front, rear, queCapacity int
}

func newArrayDeque(cap int) *arrayDeque {
	return &arrayDeque{
		nums:        make([]int, 0, cap),
		front:       0,
		rear:        0,
		queCapacity: cap,
	}
}

func (q *arrayDeque) size() int {
	if q.rear < q.front {
		return q.rear + q.queCapacity - q.front
	}
	return q.rear - q.front
}
func (q *arrayDeque) pushFirst(num int) {
	if q.size() == q.queCapacity {
		fmt.Println("deque is full")
		return
	}
	if q.front > 0 {
		q.front--
	} else {
		q.front = q.front + q.queCapacity - 1
	}
	q.nums[q.front] = num
}

func (q *arrayDeque) pushLast(num int) {
	if q.size() == q.queCapacity {
		fmt.Println("deque is full")
		return
	}
	q.nums[q.rear] = num
	q.rear++
	q.rear = q.rear % q.queCapacity
}

// 队首出队
func (q *arrayDeque) popFirst() any {
	num := q.peekFirst()
	if num == nil {
		return nil
	}
	q.front++
	q.front = q.front % q.queCapacity
	return num
}

// 队尾出队
func (q *arrayDeque) popLast() any {
	num := q.peekLast()
	if num == nil {
		return nil
	}
	if q.rear == 0 {
		q.rear = q.rear + q.queCapacity - 1
	} else {
		q.rear--
	}
	return num
}

func (q *arrayDeque) peekFirst() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}
func (q *arrayDeque) peekLast() any {
	if q.isEmpty() {
		return nil
	}
	if q.rear > 0 {
		return q.nums[q.rear-1]
	} else {
		return q.nums[q.rear+q.queCapacity-1]
	}
}
func (q *arrayDeque) isEmpty() bool {
	return q.size() == 0
}

// 栈
// 155 最小栈
// 用两个栈data, min来实现最小栈，
// 每一次在元素入栈data时，始终将当前最小值入栈min栈

// type MinStack struct {
// 	data []int
// 	min  []int
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		data: make([]int, 0),
// 		min:  make([]int, 0),
// 	}
// }

// func (this *MinStack) Push(val int) {
// 	if len(this.min) == 0 || this.min[len(this.min)-1] >= val {
// 		this.min = append(this.min, val)
// 	}
// 	this.data = append(this.data, val)
// }

// func (this *MinStack) Pop() {
// 	top := this.Top()
// 	if top == this.min[len(this.min)-1] {
// 		this.min = this.min[:len(this.min)-1]
// 	}
// 	this.data = this.data[:len(this.data)-1]
// }

// func (this *MinStack) Top() int {
// 	return this.data[len(this.data)-1]
// }

// func (this *MinStack) GetMin() int {
// 	return this.min[len(this.min)-1]
// }

type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	min := this.GetMin()
	if min < val {
		this.min = append(this.min, min)
	} else {
		this.min = append(this.min, val)
	}
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	this.min = this.min[:len(this.min)-1]
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return math.MaxInt
	}
	return this.min[len(this.min)-1]
}

// 150 逆波兰表达式求值
// 适合用栈操作运算：遇到数字则入栈；
// 遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		switch t {
		case "+", "-", "*", "/":
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			var res int
			switch t {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			stack = append(stack, res)
		default:
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}

// 394 字符串解码
func decodeString(s string) string {
	stack := make([]byte, 0)
	// const v = 'c' - '0'
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			temp := make([]byte, 0)
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				temp = append(temp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			j := 1
			for len(stack) >= j && stack[len(stack)-j] >= '0' && stack[len(stack)-j] <= '9' {
				j++
			}
			num := stack[len(stack)-j+1:]
			stack = stack[:len(stack)-j+1]
			k, _ := strconv.Atoi(string(num))

			for m := 0; m < k; m++ {
				for n := len(temp) - 1; n >= 0; n-- {
					stack = append(stack, temp[n])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

// 94 二叉树的中序遍历
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 递归
// func inorderTraversal(root *TreeNode) []int {
// 	var nums []int
// 	inorder(root, &nums)
// 	return nums
// }

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}

// 迭代
// 循环入栈根节点和左子节点，
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nums []int
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, node.Val)
		root = node.Right
	}
	return nums
}

// 133 克隆图
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	res, ok := visited[node]
	if ok {
		return res
	}
	res = &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	visited[node] = res
	// for i := 0; i < len(node.Neighbors); i++ {
	// 	neighbor := clone(node.Neighbors[i], visited)
	// 	res.Neighbors = append(res.Neighbors, neighbor)
	// }
	for _, n := range node.Neighbors {
		neighbor := clone(n, visited)
		res.Neighbors = append(res.Neighbors, neighbor)
	}
	return res
}

// 200 岛屿数量
// 把二维网格看成一个无向图 DFS
func numIslands(grid [][]byte) int {
	var count int
	row := len(grid)
	column := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			// 当顶点为'1'时进入dfs
			if grid[i][j] == '1' && dfs(i, j, grid) > 0 {
				count++
			}
		}
	}
	return count
}

func dfs(i, j int, grid [][]byte) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return 0
	}
	grid[i][j] = '0'
	// 顶点的上下左右都需要DFS
	return dfs(i, j+1, grid) + dfs(i, j-1, grid) + dfs(i-1, j, grid) + dfs(i+1, j, grid) + 1

}

// 84 柱状图中最大的矩形
// 枚举高
// 单调栈 本题中每次入栈的元素都大于栈顶元素，否则出栈栈顶元素，
// 将栈顶元素在heights中对应的元素作为高，
// 将高的左右两边第一个小于高的索引差作为宽
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	max := 0
	stack := make([]int, 0)
	for i := 0; i <= len(heights); i++ {
		var cur int
		if i != len(heights) {
			cur = heights[i]
		} else {
			cur = 0
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= cur {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i
			if len(stack) != 0 {
				peek := stack[len(stack)-1]
				w = i - peek - 1
			}
			a := h * w
			if max < a {
				max = a
			}
		}
		stack = append(stack, i)
	}
	return max
}

// 232 用两个栈实现队列
// type MyQueue struct {
// 	stack []int
// 	temp  []int
// }

// func Constructor() MyQueue {
// 	return MyQueue{
// 		stack: make([]int, 0),
// 		temp:  make([]int, 0),
// 	}
// }

// func (this *MyQueue) Push(x int) {
// 	this.stack = append(this.stack, x)
// }

// func (this *MyQueue) Pop() int {
// 	var pop int
// 	for {
// 		pop = this.stack[len(this.stack)-1]
// 		this.stack = this.stack[:len(this.stack)-1]
// 		if len(this.stack) == 0 {
// 			break
// 		}
// 		this.temp = append(this.temp, pop)
// 	}
// 	for len(this.temp) != 0 {
// 		p := this.temp[len(this.temp)-1]
// 		this.temp = this.temp[:len(this.temp)-1]
// 		this.stack = append(this.stack, p)
// 	}
// 	return pop
// }

// func (this *MyQueue) Peek() int {
// 	var pop int
// 	for len(this.stack) > 0 {
// 		pop = this.stack[len(this.stack)-1]
// 		this.stack = this.stack[:len(this.stack)-1]
// 		this.temp = append(this.temp, pop)
// 	}
// 	for len(this.temp) != 0 {
// 		p := this.temp[len(this.temp)-1]
// 		this.temp = this.temp[:len(this.temp)-1]
// 		this.stack = append(this.stack, p)
// 	}
// 	return pop
// }

// func (this *MyQueue) Empty() bool {
// 	return len(this.stack) == 0
// }

type MyQueue struct {
	stack []int
	temp  []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		temp:  make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	if len(this.stack) == 0 && len(this.temp) != 0 {
		for len(this.temp) != 0 {
			p := this.temp[len(this.temp)-1]
			this.temp = this.temp[:len(this.temp)-1]
			this.stack = append(this.stack, p)
		}
	}
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	// var pop int
	// for {
	// 	pop = this.stack[len(this.stack)-1]
	// 	this.stack = this.stack[:len(this.stack)-1]
	// 	if len(this.stack) == 0 {
	// 		break
	// 	}
	// 	this.temp = append(this.temp, pop)
	// }
	// for len(this.temp) != 0 {
	// 	p := this.temp[len(this.temp)-1]
	// 	this.temp = this.temp[:len(this.temp)-1]
	// 	this.stack = append(this.stack, p)
	// }
	if len(this.stack) == 0 && len(this.temp) == 0 {
		return 0
	}
	pop := this.Peek()
	this.temp = this.temp[:len(this.temp)-1]
	return pop
}

func (this *MyQueue) Peek() int {
	var pop int
	if len(this.stack) == 0 && len(this.temp) == 0 {
		return 0
	}
	if len(this.temp) > 0 {
		return this.temp[len(this.temp)-1]
	}
	for len(this.stack) > 0 {
		pop = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.temp = append(this.temp, pop)
	}
	// for len(this.temp) != 0 {
	// 	p := this.temp[len(this.temp)-1]
	// 	this.temp = this.temp[:len(this.temp)-1]
	// 	this.stack = append(this.stack, p)
	// }
	return pop
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.temp) == 0
}

// 542 01矩阵
// BFS
// func updateMatrix(mat [][]int) [][]int {
// 	m, n := len(mat), len(mat[0])
// 	queue := make([][]int, 0)
// 	visited := make([][]bool, m)
// 	for i := range visited {
// 		visited[i] = make([]bool, n)
// 	}
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if mat[i][j] == 0 {
// 				queue = append(queue, []int{i, j})
// 				visited[i][j] = true
// 			}
// 		}
// 	}
// 	for len(queue) > 0 {
// 		i := queue[0][0]
// 		j := queue[0][1]
// 		queue = queue[1:]

// 		direction := [][]int{[]int{i - 1, j}, []int{i + 1, j}, []int{i, j - 1}, []int{i, j + 1}}
// 		for _, d := range direction {
// 			p, q := d[0], d[1]
// 			if p >= 0 && p < m && q >= 0 && q < n && mat[p][q] == 1 && !visited[p][q] {
// 				mat[p][q] = mat[i][j] + 1
// 				queue = append(queue, []int{p, q})
// 				visited[p][q] = true
// 			}
// 		}

// 	}
// 	return mat
// }

// 动态规划

func updateMatrix(mat [][]int) [][]int {
	row := len(mat)
	column := len(mat[0])
	res := make([][]int, row)
	for i := range res {
		res[i] = make([]int, column)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			res[i][j] = math.MaxInt
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
			}
		}
	}
	// 水平向左 和 竖直向上
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			// 竖直向上
			if i-1 >= 0 {
				res[i][j] = min(res[i][j], res[i-1][j]+1)
			}
			// 水平向左
			if j-1 >= 0 {
				res[i][j] = min(res[i][j], res[i][j-1]+1)
			}
		}
	}
	/*
		// 水平向左 和 竖直向下
		for i := row - 1; i >= 0; i-- {
			for j := 0; j < column; j++ {
				// 竖直向下
				if i+1 < row {
					res[i][j] = min(res[i][j], res[i+1][j]+1)
				}
				// 水平向左
				if j-1 >= 0 {
					res[i][j] = min(res[i][j], res[i][j-1]+1)
				}
			}
		}
		// 水平向右 和 竖直向上
		for i := 0; i < row; i++ {
			for j := column - 1; j >= 0; j-- {
				// 竖直向上
				if i-1 >= 0 {
					res[i][j] = min(res[i][j], res[i-1][j]+1)
				}
				if j+1 < column {
					res[i][j] = min(res[i][j], res[i][j+1]+1)
				}
			}
		}
	*/
	// 水平向右 和 竖直向下
	for i := row - 1; i >= 0; i-- {
		for j := column - 1; j >= 0; j-- {
			if i+1 < row {
				res[i][j] = min(res[i][j], res[i+1][j]+1)
			}
			if j+1 < column {
				res[i][j] = min(res[i][j], res[i][j+1]+1)
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
