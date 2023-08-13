package main

import (
	"fmt"
)

func main() {
	var n,a,b int
	fmt.Scanf("%d %d %d", &n, &a, &b)
	sum := 0
	for i := 1; i <= n; i++ {
		var sum2 int
		for j := i; j > 0; j /= 10 {
			sum2 += j % 10
		}
		if a <= sum2 && sum2 <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}
