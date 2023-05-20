package main

import (
  "fmt"
)

func min(x, y int) (ans int) {
  ans = y
  if x < y {
    ans = x
  }
  return
}

func dfsBridge(now, parent int, time, count *int, graph [][]int, color, timeIn, timeUp *[]int) {
  (*color)[now] = 1
  *time++
  (*timeIn)[now], (*timeUp)[now] = *time, *time
  for _, v := range graph[now] {
    if v != parent {
      if (*color)[v] == 1 {
        (*timeUp)[now] = min((*timeUp)[now], (*timeIn)[v])
      } else if (*color)[v] == 0 {
        dfsBridge(v, now, time, count, graph, color, timeIn, timeUp)
        (*timeUp)[now] = min((*timeUp)[now], (*timeUp)[v])
        if (*timeIn)[now] < (*timeUp)[v] {
          *count++
        }
      }
    }
  }
  (*color)[now] = 2
  return
}

func main() {
  var n, m int
  fmt.Scan(&n)
  fmt.Scan(&m)
  graph := make([][]int, n)
  color := make([]int, n)
  for i := 0; i < m; i++ {
    var x, y int
    fmt.Scan(&x)
    fmt.Scan(&y)
    if len(graph[x]) == 0 {
      graph[x] = make([]int, 0)
    }
    if len(graph[y]) == 0 {
      graph[y] = make([]int, 0)
    }
    graph[x] = append(graph[x], y)
    graph[y] = append(graph[y], x)
  }
  count, time := 0, 0
  timeIn, timeUp := make([]int, n), make([]int, n)
  for i := range graph {
    if color[i] == 0 {
      dfsBridge(i, -1, &time, &count, graph, &color, &timeIn, &timeUp)
    }
  }
  fmt.Print(count)
}
