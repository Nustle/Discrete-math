package main

import (
  "fmt"
  "math"
)

type root struct {
  value, dist int
}

type PriorityQueue struct {
  heap  []root
  count int
}

func InitPriorityQueue(queue *PriorityQueue) {
  queue.heap = make([]root, 0)
  queue.count = 0
}

func swap(queue *PriorityQueue, i, j int) {
  queue.heap[i], queue.heap[j] = queue.heap[j], queue.heap[i]
}

func Insert(queue *PriorityQueue, x root) {
  queue.heap = append(queue.heap, x)
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
  return
}

func makeGraph(graph *[][]root, route [][]int) {
  n := len(route)
  (*graph)[0] = append((*graph)[0], root{1, route[0][0] + route[0][1]})
  (*graph)[0] = append((*graph)[0], root{n, route[0][0] + route[1][0]})
  (*graph)[n-1] = append((*graph)[n-1], root{2*n - 1, route[0][n-1] + route[1][n-1]})
  (*graph)[n-1] = append((*graph)[n-1], root{n - 2, route[0][n-1] + route[0][n-2]})
  (*graph)[n*n-n] = append((*graph)[n*n-n], root{n*n - n + 1, route[n-1][0] + route[n-1][1]})
  (*graph)[n*n-n] = append((*graph)[n*n-n], root{n*n - 2*n, route[n-1][0] + route[n-2][0]})
  for i := 1; i < n-1; i++ {
    (*graph)[n*i] = append((*graph)[n*i], root{n * (i + 1), route[i][0] + route[i+1][0]})
    (*graph)[n*i] = append((*graph)[n*i], root{n*i + 1, route[i][0] + route[i][1]})
    (*graph)[n*i] = append((*graph)[n*i], root{n * (i - 1), route[i][0] + route[i-1][0]})
    (*graph)[i] = append((*graph)[i], root{i + 1, route[0][i] + route[0][i+1]})
    (*graph)[i] = append((*graph)[i], root{n + i, route[0][i] + route[1][i]})
    (*graph)[i] = append((*graph)[i], root{i - 1, route[0][i] + route[0][i-1]})
    (*graph)[n*n-n+i] = append((*graph)[n*n-n+i], root{n*n - n + i + 1, route[n-1][i] + route[n-1][i+1]})
    (*graph)[n*n-n+i] = append((*graph)[n*n-n+i], root{n*n - 2*n + i, route[n-1][i] + route[n-2][i]})
    (*graph)[n*n-n+i] = append((*graph)[n*n-n+i], root{n*n - n + i - 1, route[n-1][i] + route[n-1][i-1]})
    (*graph)[n*(i+1)-1] = append((*graph)[n*(i+1)-1], root{n*(i+2) - 1, route[i][n-1] + route[i+1][n-1]})
    (*graph)[n*(i+1)-1] = append((*graph)[n*(i+1)-1], root{n*i - 1, route[i][n-1] + route[i-1][n-1]})
    (*graph)[n*(i+1)-1] = append((*graph)[n*(i+1)-1], root{n*i - 2, route[i][n-1] + route[i][n-2]})
  }
  for i := 1; i < n-1; i++ {
    for j := 1; j < n-1; j++ {
      (*graph)[n*i+j] = append((*graph)[n*i+j], root{n*(i-1) + j, route[i][j] + route[i-1][j]})
      (*graph)[n*i+j] = append((*graph)[n*i+j], root{n*i + j + 1, route[i][j] + route[i][j+1]})
      (*graph)[n*i+j] = append((*graph)[n*i+j], root{n*i + j - 1, route[i][j] + route[i][j-1]})
      (*graph)[n*i+j] = append((*graph)[n*i+j], root{n*(i+1) + j, route[i][j] + route[i+1][j]})
    }
  }
}

func Dijkstra(graph [][]root, route [][]int) int {
  d := make([]int, len(graph))
  for i := range d {
    d[i] = math.MaxInt32
  }
  d[0] = 0
  var queue PriorityQueue
  InitPriorityQueue(&queue)
  Insert(&queue, root{0, d[0]})
  for !Empty(&queue) {
    r := ExtractMin(&queue)
    v := r.value
    for _, u := range graph[v] {
      to, distTo := u.value, u.dist
      dop := route[v/len(route)][v%len(route)]
      if v == 0 {
        dop = 0
      }
      if d[v]+distTo-dop < d[to] {
        d[to] = d[v] + distTo - dop
        Insert(&queue, root{to, d[to]})
      }
    }
  }
  return d[len(graph)-1]
}

func main() {
  var n int
  fmt.Scan(&n)
  nel := n * n
  route := make([][]int, n)
  graph := make([][]root, nel)
  for i := range route {
    route[i] = make([]int, n)
    for j := range route[i] {
      fmt.Scan(&route[i][j])
    }
  }
  ans := 0
  if n == 1 {
    ans = route[0][0]
  } else {
    makeGraph(&graph, route)
    ans = Dijkstra(graph, route)
  }
  fmt.Print(ans)
}
