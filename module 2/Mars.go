package main

import (
  "fmt"
  "sort"
)

type biGraph struct {
  left  []int
  right []int
}

func compareCrew(first, second biGraph) int {
  for i := 0; i < len(first.left); i++ {
    if first.left[i] < second.left[i] {
      return 1
    } else if first.left[i] > second.left[i] {
      return -1
    }
  }
  return 0
}

func sortMask(binaryMask *[]biGraph) {
  for c := 0; c < len(*binaryMask); c++ {
    sort.Slice((*binaryMask)[c].left, func(i, j int) bool {
      return (*binaryMask)[c].left[i] < (*binaryMask)[c].left[j]
    })
    sort.Slice((*binaryMask)[c].right, func(i, j int) bool {
      return (*binaryMask)[c].right[i] < (*binaryMask)[c].right[j]
    })
    if len((*binaryMask)[c].right) < len((*binaryMask)[c].left) {
      (*binaryMask)[c].left, (*binaryMask)[c].right = (*binaryMask)[c].right, (*binaryMask)[c].left
    }
  }
}

func InitBiGraph(crew *biGraph, index int, binaryMask []biGraph) {
  if index == -1 {
    crew.left = make([]int, 0)
    crew.right = make([]int, 0)
  } else {
    crew.left = make([]int, len(binaryMask[index].left))
    crew.right = make([]int, len(binaryMask[index].right))
    copy(crew.left, binaryMask[index].left)
    copy(crew.right, binaryMask[index].right)
  }
}

func makeCrew(now, share, index int, color, parent *[]int, antiGraph [][]int, binaryMask []biGraph) biGraph {
  var newCrew biGraph
  InitBiGraph(&newCrew, index, binaryMask)
  ok := dfs(now, share, color, parent, antiGraph, &newCrew)
  if !ok {
    var errCrew biGraph
    InitBiGraph(&errCrew, -1, binaryMask)
    errCrew.left = append(errCrew.left, -1)
    return errCrew
  }
  return newCrew
}

func updateMask(binaryMask *[]biGraph, subMask []biGraph) {
  binaryLen, subLen := len(*binaryMask), len(subMask)
  for i := 0; i < subLen; i++ {
    if i < binaryLen {
      (*binaryMask)[i] = subMask[i]
    } else {
      *binaryMask = append(*binaryMask, subMask[i])
    }
  }
}

func getCycleLen(start, end int, parent []int) (cycleLen int) {
  for v := end; v != start; v = parent[v] {
    cycleLen++
  }
  return
}

func dfs(now, share int, color, parent *[]int, antiGraph [][]int, newCrew *biGraph) bool {
  (*color)[now] = 1
  if share == 0 {
    newCrew.left = append(newCrew.left, now)
  } else {
    newCrew.right = append(newCrew.right, now)
  }
  for _, v := range antiGraph[now] {
    if (*color)[v] == 0 {
      (*parent)[v] = now
      ok := dfs(v, (share+1)%2, color, parent, antiGraph, newCrew)
      if !ok {
        return ok
      }
    } else if (*color)[v] == 1 && (*parent)[now] != v {
      if (getCycleLen(v, now, *parent)+1)%2 == 1 {
        return false
      }
    }
  }
  (*color)[now] = 2
  return true
}

func main() {
  var n int
  fmt.Scan(&n)
  antiGraph := make([][]int, n+1)
  var sign string
  for i := 1; i < n+1; i++ {
    if len(antiGraph[i]) == 0 {
      antiGraph[i] = make([]int, 0)
    }
    for j := 1; j < n+1; j++ {
      fmt.Scan(&sign)
      if sign == "+" {
        antiGraph[i] = append(antiGraph[i], j)
      }
    }
  }
  binaryMask := make([]biGraph, 0)

  color, lastColor := make([]int, n+1), make([]int, n+1)
  parent, lastParent := make([]int, n+1), make([]int, n+1)
  start := makeCrew(1, 0, -1, &color, &parent, antiGraph, binaryMask)
  if start.left[0] == -1 {
    fmt.Println("No solution")
  } else {
    copy(lastColor, color)
    copy(lastParent, parent)
    binaryMask = append(binaryMask, start)
    haveSolution := true
  Mask:
    for i := 2; i < n+1; i++ {
      if color[i] == 0 {
        subMask := make([]biGraph, 0)
        newColor, newParent := make([]int, n+1), make([]int, n+1)
        for j := 0; j < 2*len(binaryMask); j++ {
          crew := makeCrew(i, j%2, j/2, &color, &parent, antiGraph, binaryMask)
          if crew.left[0] == -1 {
            fmt.Println("No solution")
            haveSolution = false
            break Mask
          }
          copy(newColor, color)
          copy(color, lastColor)
          copy(newParent, parent)
          copy(parent, lastParent)
          subMask = append(subMask, crew)
        }
        updateMask(&binaryMask, subMask)
        copy(color, newColor)
        copy(lastColor, newColor)
        copy(parent, newParent)
        copy(lastParent, newParent)
      }
    }
    if haveSolution {
      sortMask(&binaryMask)
      minCrew := binaryMask[0]
      minDiff := len(binaryMask[0].right) - len(binaryMask[0].left)
      for i := 1; i < len(binaryMask); i++ {
        diff := len(binaryMask[i].right) - len(binaryMask[i].left)
        if (diff < minDiff) || (diff == minDiff && len(minCrew.left) == len(binaryMask[i].left) && compareCrew(binaryMask[i], minCrew) > 0) {
          minCrew = binaryMask[i]
          minDiff = diff
        }
      }
      for _, v := range minCrew.left {
        fmt.Printf("%d ", v)
      }
    }
  }
}
