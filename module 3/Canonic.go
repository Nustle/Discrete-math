package main

import "fmt"

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

func dfs(q int, time *int, delta [][]int, visited *[]bool, numeration *[]int, nameStates *[]int) {
  (*numeration)[*time] = q
  (*nameStates)[q] = *time
  *time++
  (*visited)[q] = true
  for _, v := range delta[q] {
    if !(*visited)[v] {
      dfs(v, time, delta, visited, numeration, nameStates)
    }
  }
}

func Canonic(n, m int, q0 *int, delta [][]int, output [][]string) {
  visited, numeration := make([]bool, n), make([]int, n)
  nameStates := make([]int, n)
  time := 0
  dfs(*q0, &time, delta, &visited, &numeration, &nameStates)
  *q0 = 0
  fmt.Println(n)
  fmt.Println(m)
  fmt.Println(*q0)
  for i := 0; i < n; i++ {
    q := delta[numeration[i]]
    for j := 0; j < m; j++ {
      fmt.Printf("%d ", nameStates[q[j]])
    }
    fmt.Println()
  }
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      fmt.Printf("%s ", output[numeration[i]][j])
    }
    fmt.Println()
  }
}

func main() {
  var n, m, q0 int
  fmt.Scan(&n)
  fmt.Scan(&m)
  fmt.Scan(&q0)
  delta := makeTransition(n, m)
  output := makeOutput(n, m)
  Canonic(n, m, &q0, delta, output)
}
