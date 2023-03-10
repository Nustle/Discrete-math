package main

import "fmt"

func max(x, y int) (ans int) {
  ans = x
  if y > x {
    ans = y
  }
  return
}

func incidental(top int, dividers []int, incidentMatrix map[int][]int) {
  for _, x := range dividers {
    isIncident := true
    if x == top || (x < top && top % x != 0  x > top && x % top != 0) {
      isIncident = false
    } else {
      first := max(x, top)
      second := x + top - first
      for i := 2; i*i < first+1; i++ {
        var del1, del2 int
        if first % i == 0 {
          del1 = i
          del2 = first / i
          if (del1 % second == 0 && del1 != second) || (del2 % second == 0 && del2 != second) {
            isIncident = false
            break
          }
        }
      }
    }
    if isIncident {
      incidentMatrix[top] = append(incidentMatrix[top], x)
    }
  }
}

func getDividers(dividers []int, n int) []int {
  dividers = append(dividers, 1)
  for i := 2; i*i < n+1; i++ {
    if n % i == 0 {
      dividers = append(dividers, i)
      if i != n/i {
        dividers = append(dividers, n/i)
      }
    }
  }
  if n > 1 {
    dividers = append(dividers, n)
  }
  return dividers
}

func makeDot(incidentMatrix map[int][]int) {
  fmt.Println("graph dividers {")
  for i := range incidentMatrix {
    fmt.Printf("\t%d;\n", i)
  }
  for i := range incidentMatrix {
    for _, v := range incidentMatrix[i] {
      if i < v {
        fmt.Printf("\t%d -- %d\n", i, v)
      }
    }
  }
  fmt.Println("}")
}

func main() {
  var n int
  fmt.Scan(&n)
  dividers := make([]int, 0)
  dividers = getDividers(dividers, n)
  var incidentMatrix map[int][]int
  incidentMatrix = make(map[int][]int)
  for _, v := range dividers {
    incidental(v, dividers, incidentMatrix)
  }
  makeDot(incidentMatrix)
}
