# 图论

## 存储图

在图论中，存储图的方法主要有矩阵存图法和邻接表存图法，两者各有特点，适用于不同的场景：

### 1. 矩阵存图法（邻接矩阵）
- **定义**：使用一个二维数组（矩阵）表示图中顶点之间的连接关系。

- **实现方式**：对于一个有 `n` 个顶点的图，构建一个 `n x n` 的二维数组 `adj`，其中 `adj[i][j]` 表示顶点 `i` 到顶点 `j` 之间是否有边。
    - **无向图**：`adj[i][j] = 1` 表示顶点 `i` 和 `j` 之间有边，`adj[i][j] = 0` 表示没有边。由于是无向图，`adj[i][j] == adj[j][i]`。
    - **有向图**：`adj[i][j] = 1` 表示存在一条从顶点 `i` 指向顶点 `j` 的边。
    - **权重图**：如果是带权图，`adj[i][j]` 可以存储权值，若无边则用无穷大或其他特殊值表示。
  
- **优点**：
    - 判断两点之间是否有边的时间复杂度为 `O(1)`，非常高效。
    - 表示稠密图时非常适合，因为边的数量接近顶点的平方。

- **缺点**：
    - 占用空间较大，尤其是当图非常稀疏（边的数量远小于顶点的平方）时，存储大量无用的空信息。
    
    ```go
    package main
    
    // 邻接矩阵存图法
    type GraphMatrix struct {
    	n   int     //顶点个数
    	adj [][]int //邻接矩阵
    }
    func creat(n int) *GraphMatrix {
    	adj := make([][]int, n)
    	for i := range adj {
    		adj[i] = make([]int, n)
    	}
    	return &GraphMatrix{n: n, adj: adj}
    }
    
    // 无向图添加边
    func (g *GraphMatrix) addEdge(u, v int) {
    	g.adj[u][v] = 1
    	g.adj[v][u] = 1
    }
    func creatGraph() {
    	g := creat(4)
    	g.addEdge(0, 1)
    	g.addEdge(1, 3)
    	g.addEdge(1, 2)
    }
    ```

### 2. 邻接表存图法
- **定义**：使用多个链表（或数组）存储每个顶点的邻接顶点。
- **实现方式**：对于每个顶点 `i`，使用一个链表或数组来存储与顶点 `i` 相邻的所有顶点。对于无向图，每一条边 `i-j` 需要在 `i` 和 `j` 的链表中分别记录一次；对于有向图，只需在出发顶点的链表中记录一次。
    - **无向图**：顶点 `i` 的链表中存储与其直接相连的所有顶点 `j`。
    - **有向图**：顶点 `i` 的链表中存储所有从 `i` 出发的顶点 `j`。
    - **权重图**：链表中的元素可以存储与边相关的权值。

- **优点**：
    - 节省空间，特别适用于稀疏图，因为只需存储实际存在的边。
    - 遍历某一顶点的邻接点的时间复杂度为 `O(k)`，其中 `k` 是该顶点的度数。

- **缺点**：
    - 判断两点是否有边的时间复杂度为 `O(k)`，相对于邻接矩阵的 `O(1)` 要慢一些。

```go
//数组作为邻接表
package main
type GraphMatrix struct {
	n   int     //顶点个数
	adj [][]int //邻接表，也可用map[int][]int替代
}
func creat2(n int) *GraphMatrix {
	adj := make([][]int, n)
	return &GraphMatrix{n: n, adj: adj}
}
func (g *GraphMatrix) addEdge2(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}
func creatGraph2() {
	g := creat2(4)
	g.addEdge2(0, 1) // 添加边 (1-2)
    g.addEdge2(0, 3) // 添加边 (1-4)
	g.addEdge2(1, 2) // 添加边 (2-3)
	g.addEdge2(2, 3) // 添加边 (3-4)
}
```

```go
//链表作为邻接表
package main

type Node struct {
	Val       int
	Neighbors []*Node
}
func main() {
	node1 := &Node{
		Val: 1, Neighbors: []*Node{},
	}
	node2 := &Node{
		Val: 2, Neighbors: []*Node{},
	}
	node3 := &Node{
		Val: 3, Neighbors: []*Node{},
	}
	node4 := &Node{
		Val: 4, Neighbors: []*Node{},
	}
	node1.Neighbors = append(node1.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node1, node3)
	node3.Neighbors = append(node3.Neighbors, node2, node4)
	node4.Neighbors = append(node4.Neighbors, node1, node3)
}
```



```go
邻接表:
1: [2 4]      1--------2
2: [1 3]	   \	 /
3: [2 4]		 \ /
4: [1 3]		 / \
				/	 \
			  3--------4
```

**选择依据：**

- **稠密图**（边多，接近于顶点平方数）：适合使用邻接矩阵存储。
- **稀疏图**（边少，远少于顶点平方数）：适合使用邻接表存储。

你可以根据图的特点选择合适的存储方式。

### 3、带权图

```go
// 定义边的结构
type Edge struct {
	to     int // 目标顶点
	weight int // 边的权重
}

// 带权图的结构
type GraphWeighted struct {
	n   int           // 顶点个数
	adj map[int][]Edge // 邻接表
}

// 创建图
func createWeightedGraph(n int) *GraphWeighted {
	adj := make(map[int][]Edge, n)
	return &GraphWeighted{n: n, adj: adj}
}

// 添加带权边
func (g *GraphWeighted) addWeightedEdge(u, v, weight int) {
	g.adj[u] = append(g.adj[u], Edge{to: v, weight: weight})
	g.adj[v] = append(g.adj[v], Edge{to: u, weight: weight}) // 如果是有向图，去掉这行
}
```

## 图的搜索

### 1、BFS（广度优先）

可以使用list库来构建queue，从而提升性能，这里为了简单直接使用切片替代了。

```go
func (bf *GraphMatrix2) BFS(n int) {
    // visitd表示节点是否访问了，queue是队列
	visitd := make(map[int]bool, n)
	var queue []int
    //将n入队列并访问
	visitd[n] = true
	fmt.Println(n)
	queue = append(queue, n)
    //如果队列不为空就访问队列头部元素，访问与当前顶点相邻的所有顶点
	for len(queue) != 0 {
		for _, i := range bf.adj[queue[0]] {
			if !visitd[i] {
				visitd[i] = true
				fmt.Println(i)
				queue = append(queue, i)
			}
		}
		queue = queue[1:]
	}
}
```

### 2、DFS（带权深度优先）

首先访问出发点v，将其标记为已访问，然后依次从v出发搜索v每个邻接点w，若w没被访问，则以w为新的出发点继续深度优先遍历，直到所有邻接点都被访问

```go
package main

import "fmt"

type Edge struct {
	to     int //目标顶点
	weight int //权重
}
type GraphWeighted struct {
	n   int            //顶点个数
	adj map[int][]Edge //邻接表
}

// 创建图
func creatWeightedGrapth(n int) *GraphWeighted {
	adj := make(map[int][]Edge, n)
	return &GraphWeighted{
		n:   n,
		adj: adj,
	}
}

func (g *GraphWeighted) addWeightEdge(u, v, weight int) {
	g.adj[u] = append(g.adj[u], Edge{to: v, weight: weight})
	g.adj[v] = append(g.adj[v], Edge{to: u, weight: weight})
}

func (g *GraphWeighted) DFS(n int) {
	visited := make(map[int]bool, g.n)
	g.dfs(n, visited)
}

func (g *GraphWeighted) dfs(n int, visited map[int]bool) {
	visited[n] = true
	for _, i := range g.adj[n] {
		if !visited[i.to] {
			fmt.Printf("From %d to %d, weight: %d\n", n, i.to, i.weight)
			g.dfs(i.to, visited)
		}
	}
}
func DFSMain() {
	// 创建带权图
	g := creatWeightedGrapth(4)
	g.addWeightEdge(0, 1, 10) // 边 (0-1) 权重 10
	g.addWeightEdge(1, 2, 5)  // 边 (1-2) 集合当中，并且v∈V（如果存在权重 5
	g.addWeightEdge(2, 3, 8)  // 边 (2-3) 权重 8
	g.addWeightEdge(2, 0, 3)  // 边 (2-0) 权重 3

	// 从顶点 0 开始的带权图 DFS
	g.DFS(0)
}
```

## 并查集

### 初始化并构建相应函数

### 结构体（集合）

并查集结构体，数组parent[i]记录第i个位置元素的parent，rank表示秩(树高)。有些情况下不用数组而用map存储

```go
type UnionFind struct{
  parent []int
  size []int
}
```

### INIT函数

对并查集进行初始化，长度为n，第i个位置的parent是自身，所有rank都为1

```go
func INIT(n int) *UnionFind{
  uf:=&UnionFind{
    parent:make([]int,n),
    rank:make([]int,n),
  }
  for i:=0;i<n;i++{
    uf.parent[i]=i
    uf.size[i]=1
  }
  return uf
}
```

### Find函数（查）

如果第x位的parent不是自身，就递归的寻找祖先并赋值给parent[x]，这个过程是并查集的**路径压缩**压缩操作——在查找操作时，将沿途的节点都直接指向根节点，从而优化树的深度。

```go
func(uf *UnionFind)Find(x int){
  if uf.parent[x]!=x{
    uf.parent[x]=uf.Find(uf.parent[x])
  }
  return uf.parent[x]
}
```

### Union函数（并）

Union函数将两个不同祖先的点合并，使其拥有相同的祖先。

```go
func(uf *UnionFind)Union(x,y int){
  rootx:=uf.Find(uf.parent[x])
  rooty:=uf.Find(uf.parent[y])
  
 if rootx != rooty {
		if uf.size[rootx] >= uf.size[rooty] {
			uf.parent[rooty] = uf.parent[rootx]
      uf.size[rootx]+=uf.size[rooty]
		} else{
			uf.parent[rootx] = uf.parent[rooty]
      uf.size[rooty]+=uf.size[rootx]
		} 
	}
}
```

### 例题

[200. 岛屿数量](https://leetcode.cn/problems/number-of-islands/)

给你一个由 `'1'`（陆地）和 `'0'`（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

**解答：**

- 初始化一个并查集，大小为 `rows * cols`，对应网格中每个格子。
- 遍历网格，对于每个为 `'1'` 的位置，将其与其四周的相邻陆地（上下左右）合并在一起。
- 最后，通过遍历网格，统计每个格子的根节点来确定有多少个独立的岛屿。

```go
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	uf := NewUnionFind(rows * cols)

	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	// 将所有陆地相邻的部分合并
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				for _, d := range directions {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '1' {
						uf.Union(r*cols+c, nr*cols+nc)
					}
				}
			}
		}
	}

	// 统计有多少个岛屿
	islands := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' && uf.Find(r*cols+c) == r*cols+c {
				islands++
			}
		}
	}
	return islands
}
```

## 最小生成树

在一个带权无向图中选择若干条边，使得这些边构成一棵树，并且使得所有边的权值之和最小。最小生成树适用于连通图的场景，且必须包含所有节点但不能形成回路。

### prim算法（贪心）

核心思路：任选一个节点标记，遍历所有标记了的节点，在与之相连且未标记节点之间选出其中最小的边，然后将该未标记节点标记，重复以上过程，最终所有节点都被标记，得到的就是最小生成树。(每次加入一个节点)

代码举例：

先定义了一个稠密图，V表示有多少个节点，调用prim算法得到最小长度以及节点之间相连情况parent数组第i位存储与i节点相连的节点值。本例的数组为[0,1,0,1]

```go
// 定义图的邻接矩阵表示法
graph := [][]int{
    {0, 2, 0, 6, 0},
    {2, 0, 3, 8, 5},
    {0, 3, 0, 0, 7},
    {6, 8, 0, 0, 9},
    {0, 5, 7, 9, 0},
}
V := len(graph)		
totalWeight, parent := prim(graph, V)  
```



prim算法:

```go
// Prim算法生成最小生成树
func prim(graph [][]int, V int) (int, []int) {
	// 初始化数据结构
	selected := make([]bool, V) // 标记每个节点是否被选中加入生成树
	key := make([]int, V)       // 每个节点与生成树之间的最小连接权值
	parent := make([]int, V)    // 记录生成树结构的数组
	for i := 0; i < V; i++ {
		key[i] = INF  // 将所有节点的key值初始化为无穷大
		parent[i] = -1 // 初始化parent数组，-1表示无父节点
	}

	// 从节点0开始，设定为已选中
	selected[0] = true
	var totalWeight int // 存储生成树的总权重

	// 外层循环：每次找到一个新的节点加入生成树，共V-1次循环
	for h := 1; h < V; h++ {
		m := INF // 重置最小边权重为无穷大
		var x, y int // 用于记录当前找到的最小边的两个节点

		// 内层循环：遍历已加入生成树的节点，寻找连接到未选中节点的最小边
		for i := 0; i < V; i++ {
			if selected[i] { // 如果节点i已经在生成树中
				for j := 0; j < V; j++ {
					// 找到一个未加入生成树的节点j，且边(i, j)存在
					if !selected[j] && graph[i][j] > 0 {
						// 更新最小边权重和节点位置
						if graph[i][j] < m {
							x, y = i, j
							m = graph[i][j]
						}
					}
				}
			}
		}

		// 将最小边(y)加入生成树
		selected[y] = true     // 标记节点y为已选中
		key[y] = graph[x][y]   // 更新节点y的最小连接权值
		parent[y] = x          // 记录生成树结构，节点y的父节点是x
		totalWeight += key[y]  // 累加总权重
	}

	return totalWeight, parent // 返回生成树的总权重和结构
}
```

### kruskal算法（贪心）

核心思路：先将边权值按照从小到大排序，然后将其依次添加到生成树中，只要这条边不会与之前的边形成环路。通常使用并查集来检测环路。

并查集模板在这就省略了

```go
type Edge struct {
	u, v, weight int
}
// Kruskal算法，返回最小生成树的总权重和边列表
func kruskal(edges []Edge, v int) (int, []Edge) {
	uf := Init(v) // 初始化并查集
	// 按边权重从小到大排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})
	var mst []Edge   // 存储最小生成树的边
	totalWeight := 0 // 存储最小生成树的总权重

	// 遍历排序后的边，依次加入生成树
	for _, edge := range edges {
		uroot := uf.Find(edge.u) // 查找起点的根节点
		vroot := uf.Find(edge.v) // 查找终点的根节点

		// 如果起点和终点不在同一集合，加入该边
		if uroot != vroot {
			mst = append(mst, edge)
			totalWeight += edge.weight
			uf.union(uroot, vroot) // 合并起点和终点的集合

			// 当生成树包含v-1条边时，生成树构建完成
			if len(mst) == v-1 {
				break
			}
		}
	}
	return totalWeight, mst
}
```

## 最短路径

