package main

import "fmt"

func max(x, y int, a, b []int32) int {
  if x > y {
    return x
  } else {
    c := a
    a = b
    b = c
    return y
  }
}

func min(x, y int, a, b []int32) int {
  return x + y - max(x, y, a, b)
}

func add(a, b []int32, p int) (sum []int32) {
  mod := int32(p)
  n := max(len(a), len(b), a, b) + 1
  sum = make([]int32, n)
  var rank int32 = 0
  t := min(len(a), len(b), a, b)
  for i := 0; i < t; i++ {
    sum[n-i-1] = (a[i]+b[i])%mod + rank
    rank = (a[i] + b[i]) / mod
  }
  for j := t; j < n-1; j++ {
    sum[n-j-1] = (a[j] + rank) % mod
    rank = (a[j] + rank) / mod
  }
  if sum[0] == 0 {
    if rank > 0 {
      sum[0] = rank
    } else {
      sum = sum[1:]
    }
  }
  return
}

func scanSlice(sl []int32) {
  for i := 0; i < len(sl); i++ {
    fmt.Scanf("%d", &sl[i])
  }
}

func main() {
  var first, second, p int
  fmt.Scanf("%d", &p)
  fmt.Scanf("%d %d", &first, &second)
  a := make([]int32, first)
  b := make([]int32, second)
  scanSlice(a)
  scanSlice(b)
  fmt.Println(add(a, b, p))
}
