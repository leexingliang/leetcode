package graph

import (
	"fmt"
)

// BFS 广度优先
func (g *Graph) BFS(s, t int) {
	if s == t {
		return
	}

	// 用来标记定点是否被访问过
	isVisited := make([]bool, g.num)
	// 顶点相连的点的队列
	queue := NewQueue(g.num)
	queue.EnQueue(s)
	// 从哪个顶点来的
	prev := make([]int, g.num)
	for i := 0; i < g.num; i++ {
		prev[i] = -1
	}

	isFound := false
	for queue.IsEmpty() && !isFound {
		point := queue.DeQueue()
		isVisited[point] = true

		iter := g.adj[point].root
		for iter != nil && iter.next != nil {
			nextPoint := iter.next.val
			// 访问过....
			if isVisited[nextPoint] {
				continue
			}
			prev[nextPoint] = point
			if nextPoint == t {
				isFound = true
				break
			}
			queue.EnQueue(nextPoint)
			isVisited[nextPoint] = true
		}
	}
	if isFound {
		print(prev, s, t)
	} else {
		fmt.Printf("not found %d -> %d", s, t)
	}
}

func print(path []int, s, t int) {
	if path[t] != -1 && t != s {
		print(path, s, path[t])
	}
	fmt.Printf("%d ", t)
}
