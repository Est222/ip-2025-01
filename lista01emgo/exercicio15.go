package main

import (
	"fmt"
)

func main() {
	var i, n int
	i = 2
	fmt.Scanln(&n)
	if i <= n {
		for i <= n {
			fmt.Println(i, "^", i, "=", i*i)
			i = i + 2
		}
	}
}
