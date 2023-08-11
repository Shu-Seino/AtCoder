package main

import (
	"fmt"
)

func main() {
	var a, b int
	// var s string
	fmt.Scanf("%d %d", &a, &b)
	// fmt.Scanf("%s", &s)
	// fmt.Printf("%d %s\n", a+b+c, s)
	var c int	
	c = a * b
	if c % 2 == 0 {
		fmt.Print("Even")
	}else{
		fmt.Print("Odd")
	}
	
	// if c % 2 = 0 {
	// 	print("Even")
	// }else{
	// 	fmt.Print("Odd")
	// }

}
