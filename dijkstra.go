package main

import (
	"fmt"
	"math"
)

type Edge struct {
	to     int
	weight int
}
type Graph struct {
	nodes map[int][]Edge
}

func Newgraph() *Graph {
	return &Graph{nodes: make(map[int][]Edge)}
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.nodes[from] = append(g.nodes[from], Edge{to: to, weight: weight})
	g.nodes[to] = append(g.nodes[to], Edge{to: from, weight: weight})
}
func (g *Graph) Dijkstra(start int) map[int]int {
	//建立dist用于存储start点到其他节点的距离
	dist := make(map[int]int)
	for node := range g.nodes {
		dist[node] = math.MaxInt64
	}
	//起始节点已经到自身距离为0
	dist[start] = 0
	// visited 记录已经确定最短路径的节点
	visited := make(map[int]bool)
	visited[start] = true

	//now表示当前遍历的节点
	now := start

	//被访问的节点数大于等于总节点数循环结束
	for len(visited) < len(g.nodes) {
		//遍历当前节点所有相连接的点
		for _, node := range g.nodes[now] {
			if !visited[node.to] {
				//如果初始节点到该节点相邻节点的距离>初始节点到该节点+该节点到相邻节点的距离则更新
				if dist[now]+node.weight < dist[node.to] {
					dist[node.to] = dist[now] + node.weight
				}
			}
		}
		//找到初始节点到所有节点中的最小值
		minKey := -1
		minValue := math.MaxInt64

		for keys, value := range dist {
			//该节点不能被visitd，因为visitd的节点已经被访问过了
			if value < minValue && !visited[keys] {
				minValue = value
				minKey = keys
			}
		}
		visited[minKey] = true
		now = minKey

		// 如果 minKey 未更新，说明所有节点都访问完毕
		if minKey == -1 {
			break
		}
	}
	return dist
}

func main() {
	g := Newgraph()
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 1)
	g.AddEdge(2, 1, 2)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 5)
	g.AddEdge(3, 4, 3)

	start := 0
	dist := g.Dijkstra(start)

	// 输出每个节点的最短路径
	for node, d := range dist {
		fmt.Printf("Shortest distance from node %d to node %d is %d\n", start, node, d)
	}
}
