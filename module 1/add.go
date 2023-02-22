package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	return x + y
}

func add(a, b []int32, p int) (sum []int32) {
	mod := int32(p)
	n := max(len(a), len(b)) + 1
	sum = make([]int32, n)
	var rank int32 = 0
	t := min(len(a), len(b))
	for i := 0; i < t; i++ {
		sum[i] = (a[i]+b[i])%mod + rank
		rank = (a[i] + b[i]) / mod
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
	return
}

func scanSlice(sl []int32) {
	for i := 0; i < len(sl); i++ {
		fmt.Scan(&sl[i])
	}
}

func main() {
	var first, second, p int
	fmt.Scanf("%d", &p)
	fmt.Scanf("%d%d", &first, &second)
	a := make([]int32, first)
	b := make([]int32, second)
	scanSlice(a)
	scanSlice(b)
	if len(b) > len(a) {
		a, b = b, a
	}
	fmt.Println(add(a, b, p))
}
