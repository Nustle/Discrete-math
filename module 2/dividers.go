package main

import (
  "fmt"
  "sort"
)

func incidental(top int, dividers []int) {
  for _, x := range dividers {
    isIncident := true
    if x <= top || x%top != 0 {
      isIncident = false
    } else {
      for i := 2; i*i < x+1; i++ {
        var del1, del2 int
        if x%i == 0 {
          del1 = i
          del2 = x / i
          if (del1%top == 0 && del1 != top) || (del2%top == 0 && del2 != top) {
            isIncident = false
            break
          }
        }
      }
    }
    if isIncident {
      fmt.Printf("\t%d--%d\n", x, top)
    }
  }
}

func getDividers(dividers []int, n int) []int {
  dividers = append(dividers, 1)
  for i := 2; i*i < n+1; i++ {
    if n%i == 0 {
      dividers = append(dividers, i)
      if i != n/i {
        dividers = append(dividers, n/i)
      }
    }
  }
  if n > 1 {
    dividers = append(dividers, n)
  }
  sort.Slice(dividers, func(i, j int) bool {
    return dividers[i] > dividers[j]
  })

  return dividers
}

func main() {
  var n int
  fmt.Scan(&n)
  dividers := make([]int, 0)
  dividers = getDividers(dividers, n)
  fmt.Println("graph dividers {")
  for _, v := range dividers {
    fmt.Printf("\t%d;\n", v)
  }
  for _, v := range dividers {
    incidental(v, dividers)
  }
  fmt.Println("}")
}
