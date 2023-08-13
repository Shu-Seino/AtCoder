package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	A, B := 0, 0
	for i, v := range a {
		if i%2 == 0 {
			A += v
		} else {
			B += v
		}
	}
	fmt.Println(A - B)
}
