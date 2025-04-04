package main

import (
	"fmt"
)

func main() {
	var x, y, i int
	fmt.Scan(&x, &y)
	if x%2 == 0 {
		i = 0
		for i < y {
			fmt.Println(x)
			x = x + 2
			i = i + 1
		}
	} else {
		fmt.Println("O PRIMEIRO NÚMERO NÃO É PAR")
	}
}
