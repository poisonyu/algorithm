package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 基于邻接矩阵实现无向图
type graphAdjMat struct {
	vertices []int
	adjMat   [][]int
}

func newGraphAdjMat(vertices []int, edges [][]int) *graphAdjMat {
	n := len(vertices)
	adjMat := make([][]int, n)
	for i := range adjMat {
		adjMat[i] = make([]int, n)
	}
	g := &graphAdjMat{
		vertices: vertices,
		adjMat:   adjMat,
	}
	for i := 0; i < len(edges); i++ {
		g.addEdge(edges[i][0], edges[i][1])
	}
	return g
}

func (g *graphAdjMat) size() int {
	return len(g.vertices)
}

func (g *graphAdjMat) addVertex(val int) {
	n := g.size()
	g.vertices = append(g.vertices, val)
	// 添加行
	newRow := make([]int, n)
	g.adjMat = append(g.adjMat, newRow)
	// 添加列
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
}

func (g *graphAdjMat) removeVertex(index int) {
	if index >= g.size() {
		return
	}
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	g.adjMat = append(g.adjMat[:index], g.adjMat[index+1:]...)
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i][:index], g.adjMat[i][index+1:]...)
	}

}
func (g *graphAdjMat) addEdge(i, j int) {
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Errorf("%s", "Index Out Of Bounds Exception")
	}
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

func (g *graphAdjMat) removeEdge(i, j int) {
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Errorf("%s", "Index Out Of Bounds Expection")
	}
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}

func (g *graphAdjMat) print() {
	fmt.Printf("\t顶点列表 = %v\n", g.vertices)
	fmt.Printf("\t邻接矩阵 = \n")
	for i := range g.adjMat {
		fmt.Printf("\t\t\t%v\n", g.adjMat[i])
	}
}

// 基于邻接表的实现无向图
// 使用动态数组代替链表
// 使用哈希表存储邻接表

type Vertex struct {
	Val int
}
type graphAdjList struct {
	adjList map[Vertex][]Vertex
}

func DeleteSliceElms(vertexs []Vertex, vet Vertex) []Vertex {
	i := 0
	for i < len(vertexs) {
		if vertexs[i].Val != vet.Val {
			i++
		} else {
			break
		}
	}
	if i == len(vertexs) {
		return vertexs
	}
	return append(vertexs[:i], vertexs[i+1:]...)
}

// 初始化邻接表
func newGraphAdjList(edges [][]Vertex) *graphAdjList {
	g := &graphAdjList{
		adjList: make(map[Vertex][]Vertex),
	}
	for _, edge := range edges {
		g.addVertex(edge[0])
		g.addVertex(edge[1])
		g.addEdge(edge[0], edge[1])
	}
	return g
}

// 获取顶点数量
func (g *graphAdjList) size() int {
	return len(g.adjList)
}

// 添加边 vet1 - vet2
func (g *graphAdjList) addEdge(vet1, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("error")
	}
	g.adjList[vet1] = append(g.adjList[vet1], vet2)
	g.adjList[vet2] = append(g.adjList[vet2], vet1)
}

// 删除边 vet1 - vet2
func (g *graphAdjList) removeEdge(vet1, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("error")
	}
	g.adjList[vet1] = DeleteSliceElms(g.adjList[vet1], vet2)
	g.adjList[vet2] = DeleteSliceElms(g.adjList[vet2], vet1)
}

// 添加顶点
func (g *graphAdjList) addVertex(vet Vertex) {
	if _, ok := g.adjList[vet]; !ok {
		g.adjList[vet] = make([]Vertex, 0)
	}
}

// 删除顶点
func (g *graphAdjList) removeVertex(vet Vertex) {
	if _, ok := g.adjList[vet]; !ok {
		panic("error")
	}
	delete(g.adjList, vet)
	for i := range g.adjList {
		g.adjList[i] = DeleteSliceElms(g.adjList[i], vet)
	}
}

func (g *graphAdjList) print() {
	var builder strings.Builder
	fmt.Printf("邻接表 = \n")
	for k, v := range g.adjList {
		builder.WriteString("\t\t" + strconv.Itoa(k.Val) + ": ")
		for _, vet := range v {
			builder.WriteString(strconv.Itoa(vet.Val) + " ")
		}
		fmt.Println(builder.String())
		builder.Reset()
	}
}
