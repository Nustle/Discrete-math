package main

import (
  "fmt"
  "sort"
)

func topSortedDfs(now int, graph [][]int, color []int, topSorted []int, i *int) {
  color[now] = 1
  for _, v := range graph[now] {
    if color[v] == 0 {
      topSortedDfs(v, graph, color, topSorted, i)
    }
  }
  topSorted[*i] = now
  *i--
  color[now] = 2
}

func dfsT(now int, graphT [][]int, visited map[int]bool, component *[]int, color *[]int, currentColor int) {
  if len(*component) == 0 {
    *component = append(*component, now)
    (*color)[now] = currentColor
  } else if len(*component) > 0 && (*component)[0] < now {
    *component = append(*component, now)
    (*color)[now] = currentColor
  } else {
    t := (*component)[0]
    (*component)[0] = now
    *component = append(*component, t)
    (*color)[now] = currentColor
  }
  visited[now] = true
  for _, v := range graphT[now] {
    if !visited[v] {
      dfsT(v, graphT, visited, component, color, currentColor)
    }
  }
}

func main() {
  var n, m int
  fmt.Scan(&n)
  fmt.Scan(&m)
  graph := make([][]int, n)
  graphT := make([][]int, n)
  topSorted := make([]int, n)
  color := make([]int, n)
  for i := 0; i < m; i++ {
    var x, y int
    fmt.Scan(&x)
    fmt.Scan(&y)
    if len(graph[x]) == 0 {
      graph[x] = make([]int, 0)
    }
    if len(graphT[y]) == 0 {
      graphT[y] = make([]int, 0)
    }
    graph[x] = append(graph[x], y)
    graphT[y] = append(graphT[y], x)
  }
  q := n - 1
  for i := range graph {
    if color[i] == 0 {
      topSortedDfs(i, graph, color, topSorted, &q)
    }
  }
  visited := make(map[int]bool)
  components := make([][]int, 0)
  currentColor := 1
  color = make([]int, n)
  for _, v := range topSorted {
    if !visited[v] {
      component := make([]int, 0)
      dfsT(v, graphT, visited, &component, &color, currentColor)
      components = append(components, component)
      currentColor++
    }
  }
  answer := make([]int, 0)

  for i := range components {
    base := true
    pivot := components[i][0]
  isBased:
    for _, v := range components[i] {
      for _, z := range graphT[v] {
        if color[pivot] != color[z] {
          base = false
          break isBased
        }
      }
    }
    if base {
      answer = append(answer, pivot)
    }
  }
  sort.Slice(answer, func(i, j int) bool {
    return answer[i] < answer[j]
  })
  for _, v := range answer {
    fmt.Printf("%d ", v)
  }
}
