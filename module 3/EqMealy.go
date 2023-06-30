package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

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

func Split1(n int, delta [][]int, output [][]string) (eqClasses []int, count int) {
  parent, rank := make([]int, n), make([]int, n)
  eqClasses = make([]int, n)
  count = n
  makeSet(n, &parent, &rank)
  for i := 0; i < n; i++ {
    for j := i + 1; j < n; j++ {
      if findSet(i, &parent) != findSet(j, &parent) {
        eq := true
        for x := range delta[0] {
          if output[i][x] != output[j][x] {
            eq = false
            break
          }
        }
        if eq {
          unionSets(i, j, &parent, &rank)
          count--
        }
      }
    }
  }
  for i := range eqClasses {
    eqClasses[i] = findSet(i, &parent)
  }
  return
}

func Split(n int, delta [][]int, KeqClasses []int) (eqClasses []int, count int) {
  parent, rank := make([]int, n), make([]int, n)
  makeSet(n, &parent, &rank)
  count = n
  for i := 0; i < n; i++ {
    for j := i + 1; j < n; j++ {
      if KeqClasses[i] == KeqClasses[j] && findSet(i, &parent) != findSet(j, &parent) {
        eq := true
        for x := range delta[0] {
          if KeqClasses[delta[i][x]] != KeqClasses[delta[j][x]] {
            eq = false
            break
          }
        }
        if eq {
          unionSets(i, j, &parent, &rank)
          count--
        }
      }
    }
  }
  for i := range KeqClasses {
    KeqClasses[i] = findSet(i, &parent)
  }
  eqClasses = KeqClasses
  return
}

func MinMealy(q0 *int, n, m int, delta [][]int, output [][]string) ([][]int, [][]string) {
  eqClasses, count := Split1(n, delta, output)
  for true {
    KeqClasses, Kcount := Split(n, delta, eqClasses)
    if count == Kcount {
      break
    }
    eqClasses, count = KeqClasses, Kcount
  }
  minDelta, minOutput := make([][]int, count), make([][]string, count)
  visited := make(map[int]bool)
  index := -1
  helpVisited := make(map[int]bool)
  ind := make([]int, n)
  for q := 0; q < n; q++ {
    minQ := eqClasses[q]
    if !helpVisited[minQ] {
      index++
      helpVisited[minQ] = true
      ind[q] = index
    } else {
      ind[q] = ind[minQ]
    }
  }
  for q := 0; q < n; q++ {
    minQ := eqClasses[q]
    if !visited[minQ] {
      visited[minQ] = true
      index = ind[minQ]
      minDelta[index], minOutput[index] = make([]int, m), make([]string, m)
      for i := 0; i < m; i++ {
        minDelta[index][i] = ind[delta[minQ][i]]
        minOutput[index][i] = output[q][i]
      }
    }
  }
  *q0 = eqClasses[*q0]
  minDelta, minOutput = Canonic(q0, &n, minDelta, minOutput)
  return minDelta, minOutput
}

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

func dfs(q int, time *int, delta [][]int, visited *[]bool, nameStates, numeration map[int]int) {
  nameStates[q] = *time
  numeration[*time] = q
  *time++
  (*visited)[q] = true
  for _, v := range delta[q] {
    if !(*visited)[v] {
      dfs(v, time, delta, visited, nameStates, numeration)
    }
  }
}

func Canonic(q0, n *int, delta [][]int, output [][]string) (canonicDelta [][]int, canonicOutput [][]string) {
  m := len(delta[0])
  visited, nameStates, numeration := make([]bool, *n), make(map[int]int), make(map[int]int)
  time := 0
  dfs(*q0, &time, delta, &visited, nameStates, numeration)
  canonicDelta, canonicOutput = make([][]int, *n), make([][]string, *n)
  for q := 0; q < len(numeration); q++ {
    canonicDelta[q], canonicOutput[q] = make([]int, m), make([]string, m)
    for j := 0; j < m; j++ {
      canonicDelta[q][j], canonicOutput[q][j] = nameStates[delta[numeration[q]][j]], output[numeration[q]][j]
    }
  }
  *n = len(numeration)
  *q0 = 0
  return
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  writer := bufio.NewWriter(os.Stdout)

  scanner.Scan()
  n1, _ := strconv.Atoi(scanner.Text())
  scanner.Scan()
  m1, _ := strconv.Atoi(scanner.Text())
  scanner.Scan()
  q1, _ := strconv.Atoi(scanner.Text())

  delta1 := makeTransition(n1, m1, scanner)
  output1 := makeOutput(n1, m1, scanner)

  scanner.Scan()
  n2, _ := strconv.Atoi(scanner.Text())
  scanner.Scan()
  m2, _ := strconv.Atoi(scanner.Text())
  scanner.Scan()
  q2, _ := strconv.Atoi(scanner.Text())

  delta2 := makeTransition(n2, m2, scanner)
  output2 := makeOutput(n2, m2, scanner)

  delta1, output1 = Canonic(&q1, &n1, delta1, output1)
  delta2, output2 = Canonic(&q2, &n2, delta2, output2)
  delta1, output1 = MinMealy(&q1, n1, m1, delta1, output1)
  delta2, output2 = MinMealy(&q2, n2, m2, delta2, output2)
  delta1, output1 = Canonic(&q1, &n1, delta1, output1)
  delta2, output2 = Canonic(&q2, &n2, delta2, output2)

  if n1 != n2  m1 != m2 {
    fmt.Fprint(writer, "NOT EQUAL")
  } else {
    eq := true
  isEqual:
    for i := 0; i < n1; i++ {
      for j := 0; j < m1; j++ {
        if delta1[i][j] != delta2[i][j] || output1[i][j] != output2[i][j] {
          eq = false
          fmt.Fprint(writer, "NOT EQUAL")
          break isEqual
        }
      }
    }
    if eq {
      fmt.Fprint(writer, "EQUAL")
    }
  }
  writer.Flush()
}
