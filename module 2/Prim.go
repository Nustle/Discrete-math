package main

import (
  "fmt"
  "math"
)

type root struct {
  value, dist int
}

type PriorityQueue struct {
  heap    []root
  indices map[int]int
  count   int
}

func InitPriorityQueue(queue *PriorityQueue) {
  queue.heap = make([]root, 0)
  queue.indices = make(map[int]int)
  queue.count = 0
}

func swap(queue *PriorityQueue, i, j int) {
  queue.heap[i], queue.heap[j] = queue.heap[j], queue.heap[i]
  queue.indices[queue.heap[i].value] = i
  queue.indices[queue.heap[j].value] = j
}

func Insert(queue *PriorityQueue, x root) {
  queue.heap = append(queue.heap, x)
  queue.indices[x.value] = queue.count
  queue.count++
  pos := queue.count - 1
  for pos > 0 && queue.heap[pos].dist < queue.heap[(pos-1)/2].dist {
    swap(queue, pos, (pos-1)/2)
    pos = (pos - 1) / 2
  }
}

func Empty(queue *PriorityQueue) bool {
  return queue.count == 0
}

func ExtractMin(queue *PriorityQueue) (ans root) {
  ans = queue.heap[0]
  queue.heap[0] = queue.heap[queue.count-1]
  queue.count--
  queue.heap = queue.heap[:queue.count]
  pos := 0
  for 2*pos+1 < queue.count {
    j := 2*pos + 1
    if 2*pos+2 < queue.count && queue.heap[2*pos+2].dist < queue.heap[j].dist {
      j = 2*pos + 2
    }
    if queue.heap[pos].dist <= queue.heap[j].dist {
      break
    }
    swap(queue, pos, j)
    pos = j
  }
  delete(queue.indices, ans.value)
  return
}

func main() {
  var n, m int
  fmt.Scan(&n)
  fmt.Scan(&m)
  graph := make([][]root, n+1)
  for i := 0; i < m; i++ {
    var x, y, z int
    fmt.Scan(&x)
    fmt.Scan(&y)
    fmt.Scan(&z)
    graph[x] = append(graph[x], root{y, z})
    graph[y] = append(graph[y], root{x, z})
  }
  minEdge, spanning := make([]int, n), make([]bool, n)
  for i := range minEdge {
    minEdge[i] = math.MaxInt32
  }
  minEdge[0] = 0
  var queue PriorityQueue
  InitPriorityQueue(&queue)
  Insert(&queue, root{0, 0})
  way := 0
  for !Empty(&queue) {
    v := ExtractMin(&queue).value
    if spanning[v] {
      continue
    }
    spanning[v] = true
    way += minEdge[v]
    for _, e := range graph[v] {
      to, dist := e.value, e.dist
      if !spanning[to] && dist < minEdge[to] {
        minEdge[to] = dist
        Insert(&queue, root{to, minEdge[to]})
      }
    }
  }
  fmt.Println(way)
}
