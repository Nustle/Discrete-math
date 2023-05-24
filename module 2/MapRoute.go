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
}
