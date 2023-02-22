package main

import "fmt"

func add(a, b []int32, p int) []int32 {
  mod := int32(p)
  n := len(a) + 1
  sum := make([]int32, n)
  var rank int32 = 0
  t := len(b)
  for i := 0; i < t; i++ {
    sum[i] = (a[i] + b[i] + rank) % mod
    rank = (a[i] + b[i] + rank) / mod
  }
  for j := t; j < n-1; j++ {
    sum[j] = (a[j] + rank) % mod
    rank = (a[j] + rank) / mod
  }
  if sum[n-1] == 0 {
    if rank > 0 {
      sum[n-1] = rank
    } else {
      sum = sum[:(n - 1)]
    }
  }
  return sum
}

func scanSlice(sl []int32) {
  for i := 0; i < len(sl); i++ {
    fmt.Scan(&sl[i])
  }
}

func main() {
  var first, second, p int
  fmt.Scan(&p)
  fmt.Scan(&first, &second)
  a := make([]int32, first)
  b := make([]int32, second)
  scanSlice(a)
  scanSlice(b)
  if len(b) > len(a) {
    a, b = b, a
  }
  fmt.Println(add(a, b, p))
}
