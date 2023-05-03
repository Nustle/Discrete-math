package main

import (
	"fmt"
	"math"
)

type root struct {
	v, color int
}

type pair struct {
	x, y int
}

func dfs(graph [][]root, visited map[int]bool, now int, count *int, min *int, color int, edges map[pair]int) {
	visited[now] = true
	for i, v := range graph[now] {
		var p pair
		p.x, p.y = int(math.Min(float64(now), float64(v.v))), int(math.Max(float64(now), float64(v.v)))
		edges[p] = 1
		if v.color == 0 {
			*count++
			graph[now][i].color = color
			*min = int(math.Min(float64(*min), float64(v.v)))
			dfs(graph, visited, v.v, count, min, color, edges)
		}
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	graph := make([][]root, n)
	for i := 0; i < m; i++ {
		var x, y int
		var r1, r2 root
		fmt.Scan(&x)
		fmt.Scan(&y)
		if len(graph[x]) == 0 {
			graph[x] = make([]root, 0)
		}
		if len(graph[y]) == 0 {
			graph[y] = make([]root, 0)
		}
		r1.v, r2.v = y, x
		r1.color, r2.color = 0, 0
		graph[x] = append(graph[x], r1)
		if r1 != r2 {
			graph[y] = append(graph[y], r2)
		}
	}
	visited := make(map[int]bool)
	color := 1
	resultColor := 1
	maxCount := 0
	countEdges := -1
	min := n
	edges := make(map[pair]int)
	first, second := 0, len(edges)
	for i := range graph {
		count := 0
		currentMin := n
		if !visited[i] {
			dfs(graph, visited, i, &count, &currentMin, color, edges)
			first = second
			second = len(edges)
		}
		if count > maxCount || (count == maxCount && second-first > countEdges) || (count == maxCount && second-first == countEdges && currentMin < min) {
			maxCount = count
			countEdges = second - first
			min = currentMin
			resultColor = color
		}
		color++
	}
	fmt.Println("graph G {")
	for i := range graph {
		fmt.Printf("\t%d", i)
		if len(graph[i]) > 0 && graph[i][0].color == resultColor || (m == 0 && i == 0) {
			fmt.Print(" [color = red]")
		}
		fmt.Println()
	}
	for i := range graph {
		for _, v := range graph[i] {
			var p pair
			p.x, p.y = i, v.v
			_, ok := edges[p]
			if ok {
				fmt.Printf("\t%d--%d", i, v.v)
				if v.color == resultColor {
					fmt.Print(" [color = red]")
				}
				fmt.Println()
			}
		}
	}
	fmt.Println("}")
}
