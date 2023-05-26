package main

import (
  "fmt"
)

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
  fmt.Println("digraph {")
  fmt.Println("\trankdir = LR")

  for i := range delta {
    for j, v := range delta[i] {
      exit, signal := output[i][j], rune(97+j)
      fmt.Printf("\t%d -> %d [label = \"%c(%s)\"]\n", i, v, signal, exit)
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
