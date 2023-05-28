package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func makeTransition(states, alphabet int, scanner *bufio.Scanner) [][]int {
  transition := make([][]int, states)
  for i := range transition {
    transition[i] = make([]int, alphabet)
    for j := range transition[i] {
      scanner.Scan()
      transition[i][j], _ = strconv.Atoi(scanner.Text())
    }
  }
  return transition
}

func makeOutput(states, alphabet int, scanner *bufio.Scanner) [][]string {
  output := make([][]string, states)
  for i := range output {
    output[i] = make([]string, alphabet)
    for j := range output[i] {
      scanner.Scan()
      output[i][j] = scanner.Text()
    }
  }
  return output
}

func dfs(q int, time *int, delta [][]int, visited *[]bool, nameStates *[]int) {
  (*nameStates)[q] = *time
  *time++
  (*visited)[q] = true
  for _, v := range delta[q] {
    if !(*visited)[v] {
      dfs(v, time, delta, visited, nameStates)
    }
  }
}

func Canonic(q0 *int, delta [][]int, output [][]string) (canonicDelta [][]int, canonicOutput [][]string) {
  n, m := len(delta), len(delta[0])
  visited, nameStates := make([]bool, n), make([]int, n)
  time := 0
  for i := range delta {
    if !visited[i] {
      dfs(*q0, &time, delta, &visited, &nameStates)
    }
  }
  canonicDelta, canonicOutput = make([][]int, n), make([][]string, n)
  for i := 0; i < n; i++ {
    q := nameStates[i]
    canonicDelta[q], canonicOutput[q] = make([]int, m), make([]string, m)
    for j := 0; j < m; j++ {
      canonicDelta[q][j], canonicOutput[q][j] = nameStates[delta[i][j]], output[i][j]
    }
  }
  *q0 = 0
  return
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  writer := bufio.NewWriter(os.Stdout)

  scanner.Scan()
  n, _ := strconv.Atoi(scanner.Text())
  scanner.Scan()
  m, _ := strconv.Atoi(scanner.Text())
  scanner.Scan()
  q0, _ := strconv.Atoi(scanner.Text())

  delta := makeTransition(n, m, scanner)
  output := makeOutput(n, m, scanner)
  delta, output = Canonic(&q0, delta, output)
  fmt.Fprintf(writer, "%d\n%d\n%d\n", n, m, q0)
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      fmt.Fprintf(writer, "%d ", delta[i][j])
    }
    writer.WriteString("\n")
  }
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      fmt.Fprintf(writer, "%s ", output[i][j])
    }
    writer.WriteString("\n")
  }
  writer.Flush()
}
