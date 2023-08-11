package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	//nの数だけ配列を作成し、入力値を格納する
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	var count int = 0
	//配列の3つの値が何回で２で割り切れるかmain関数に返す
	for {
		for i := 0; i < n; i++ {
			if a[i]%2 == 0 {
				a[i] = a[i] / 2
			} else {
				fmt.Printf("%d\n", count)
				return
			}
		}
		count++
	}
	
}
