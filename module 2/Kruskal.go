package main

import (
  "fmt"
  "math"
  "sort"
)

type point struct {
  x, y int
}

type edge struct {
  roots  point
  weight float64
}

func makeSet(nel int, parent, rank *[]int) {
  for v := 0; v < nel; v++ {
    (*parent)[v] = v
    (*rank)[v] = 0
  }
}

func findSet(v int, parent *[]int) int {
  if v == (*parent)[v] {
    return v
  }
  (*parent)[v] = findSet((*parent)[v], parent)
  return (*parent)[v]
}

func unionSets(a, b int, parent, rank *[]int) {
  a = findSet(a, parent)
  b = findSet(b, parent)
  if a != b {
    if (*rank)[a] < (*rank)[b] {
      (*parent)[a] = b
    } else if (*rank)[a] > (*rank)[b] {
      (*parent)[b] = a
    } else {
      (*parent)[b] = a
      (*rank)[a]++
    }
  }
}

func main() {
  var n int
  fmt.Scan(&n)
  attraction := make([]point, n)
  graph := make([]edge, 0)
  for i := 0; i < n; i++ {
    var x, y int
    fmt.Scan(&x)
    fmt.Scan(&y)
    attraction[i] = point{x, y}
  }
  for i := 0; i < n; i++ {
    for j := i + 1; j < n; j++ {
      dist := math.Sqrt(math.Pow(float64(attraction[i].x-attraction[j].x), 2) + math.Pow(float64(attraction[i].y-attraction[j].y), 2))
      graph = append(graph, edge{point{i, j}, dist})
    }
  }
  sort.Slice(graph, func(i, j int) bool {
    return graph[i].weight < graph[j].weight
  })
  nel := n * (n - 1) / 2
  parent, rank := make([]int, 2*nel), make([]int, 2*nel)
  makeSet(2*nel, &parent, &rank)
  ans := 0.0
  for _, e := range graph {
    u, v := e.roots.x, e.roots.y
    if findSet(u, &parent) != findSet(v, &parent) {
      unionSets(u, v, &parent, &rank)
      ans += e.weight
    }
  }
  fmt.Printf("%.2f", ans)
}
