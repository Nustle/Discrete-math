package main

import (
  "fmt"
)

type Queue struct {
  queue []int
}

func Push(q *Queue, v int) {
  q.queue = append(q.queue, v)
}

func Front(q *Queue) int {
  return q.queue[0]
}

func Pop(q *Queue) {
  q.queue = q.queue[1:]
}

func Empty(q *Queue) bool {
  return len(q.queue) == 0
}

func makeTransition(states, alphabet int) [][]int {
  transition := make([][]int, states)
  for i := range transition {
    transition[i] = make([]int, alphabet)
    for j := range transition[i] {
      fmt.Scan(&transition[i][j])
    }
  }
  return transition
}

func makeOutput(states, alphabet int) [][]string {
  output := make([][]string, states)
  for i := range output {
    output[i] = make([]string, alphabet)
    for j := range output[i] {
      fmt.Scan(&output[i][j])
    }
  }
  return output
}

func visMealy(q0 int, delta [][]int, output [][]string) {
  var queue Queue
  fmt.Println("digraph {")
  fmt.Println("\trankdir = LR")
  used := make([]bool, len(delta))
  used[q0] = true
  Push(&queue, q0)
  for !Empty(&queue) {
    v := Front(&queue)
    Pop(&queue)
    for i := range delta[v] {
      to := delta[v][i]
      exit, signal := output[v][i], rune(97+i)
      fmt.Printf("\t%d -> %d [label = \"%c(%s)\"]\n", v, to, signal, exit)
      if !used[to] {
        used[to] = true
        Push(&queue, to)
      }
    }
  }
  fmt.Print("}")
}

func main() {
  var n, m, q0 int
  fmt.Scan(&n)
  fmt.Scan(&m)
  fmt.Scan(&q0)
  delta := makeTransition(n, m)
  output := makeOutput(n, m)
  visMealy(q0, delta, output)
}
