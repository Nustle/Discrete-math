package main

import "fmt"

type frac struct {
  num, denom int
}

func max(x, y int) (ans int) {
  ans = x
  if y > x {
    ans = y
  }
  return
}

func abs(x int) int {
  if x < 0 {
    x = -x
  }
  return x
}

func nod(x, y int) int {
  x = abs(x)
  y = abs(y)
  x, y = max(x, y), x+y-max(x, y)
  for y > 0 {
    x, y = y, x%y
  }
  return x
}

func appendSolution(ans []frac, i int, matrix [][]frac) []frac {
  var x frac
  x.num = matrix[i-1][i-1].denom * matrix[i-1][len(matrix)].num
  x.denom = matrix[i-1][i-1].num * matrix[i-1][len(matrix)].denom
  div := nod(x.num, x.denom)
  x.num /= div
  x.denom /= div
  if (x.num < 0 && x.denom < 0) || (x.num >= 0 && x.denom < 0) {
    x.num, x.denom = -x.num, -x.denom
  }
  ans = append(ans, x)
  return ans
}

func forwardGaussRec(step int, matrix [][]frac) {
  for i := step + 1; i < len(matrix); i++ {
    k := matrix[i][step]
    for j := step; j < len(matrix[i]); j++ {
      matrix[i][j].num = matrix[i][j].num*matrix[step][j].denom*matrix[step][step].num*k.denom - 
                         matrix[i][j].denom*matrix[step][j].num*matrix[step][step].denom*k.num
      matrix[i][j].denom = matrix[i][j].denom * matrix[step][j].denom * matrix[step][step].num * k.denom
    }
  }
}

func backwardGaussRec(step int, matrix [][]frac) {
  n := len(matrix)
  for i := step - 1; i >= 0; i-- {
    matrix[i][n].num = matrix[step][n].denom*matrix[step][step].num*matrix[i][step].denom*matrix[i][n].num - 
                       matrix[step][n].num*matrix[step][step].denom*matrix[i][step].num*matrix[i][n].denom
    matrix[i][n].denom = matrix[step][n].denom * matrix[step][step].num * matrix[i][step].denom * matrix[i][n].denom
    matrix[i][step].num, matrix[i][step].denom = 0, 1
    div := nod(matrix[i][n].num, matrix[i][n].denom)
    matrix[i][n].num /= div
    matrix[i][n].denom /= div
  }
}

func forwardGauss(matrix [][]frac) {
  for i := range matrix {
    forwardGaussRec(i, matrix)
  }
}

func backwardGauss(matrix [][]frac) (ans []frac) {
  ans = appendSolution(ans, len(matrix), matrix)
  for i := len(matrix) - 1; i > 0; i-- {
    backwardGaussRec(i, matrix)
    ans = appendSolution(ans, i, matrix)
  }
  return
}

func isJoint(matrix [][]frac) (ok bool) {
  ok = true
  for i := range matrix {
    haveNulls := true
    for j := 0; j < len(matrix[i])-1; j++ {
      if matrix[i][j].num != 0 {
        haveNulls = false
        break
      }
    }
    if haveNulls {
      if matrix[i][len(matrix[i])-1].num != 0 {
        ok = false
      }
    }
  }
  return
}

func main() {
  var n int
  fmt.Scan(&n)
  matrix := make([][]frac, n)
  for i := range matrix {
    matrix[i] = make([]frac, n+1)
    for j := range matrix[i] {
      fmt.Scan(&matrix[i][j].num)
      matrix[i][j].denom = 1
    }
  }
  forwardGauss(matrix)
  if !isJoint(matrix) {
    fmt.Println("No solution")
  } else {
    ans := backwardGauss(matrix)
    for i := len(ans) - 1; i >= 0; i-- {
      if ans[i].denom == 1 {
        fmt.Printf("%d", ans[i].num)
      } else {
        fmt.Printf("%d/%d", ans[i].num, ans[i].denom)
      }
      if i > 0 {
        fmt.Println()
      }
    }
  }
}
