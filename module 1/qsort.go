package main

import "fmt"

var sl []int

func less(i, j int) (isLess bool) {
  diff := sl[i] - sl[j]
  if diff < 0 {
    isLess = true
  } else {
    isLess = false
  }
  return
}

func swap(i, j int) {
  sl[i], sl[j] = sl[j], sl[i]
}

func partition(low, high int,
  less func(i, j int) bool,
  swap func(i, j int)) int {
  i := low
  for j := low; j < high; j++ {
    if less(j, high) {
      swap(i, j)
      i++
    }
  }
  swap(high, i)
  return i
}

func qsortRec(low, high int,
  less func(i, j int) bool,
  swap func(i, j int)) {
  if low < high {
    q := partition(low, high, less, swap)
    qsortRec(low, q-1, less, swap)
    qsortRec(q+1, high, less, swap)
  }
}

func qsort(n int,
  less func(i, j int) bool,
  swap func(i, j int)) {
  qsortRec(0, n-1, less, swap)
}

func main() {
  var n int
  fmt.Scanf("%d", &n)
  sl = make([]int, n)
  for i := range sl {
    fmt.Scanf("%d", &sl[i])
  }
  qsort(n, less, swap)
  for i := range sl {
    fmt.Printf("%d ", sl[i])
  }
}
