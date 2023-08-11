package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scanf("%s", &a)
	var count int
	for i := 0; i < len(a); i++ {
		if a[i] == '1' {
			count++
		}
	}
	fmt.Println(count)
}
