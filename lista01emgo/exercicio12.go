package main

import (
	"fmt"
)

func main() {
	var n int
	var x float64
	fmt.Scan(&n)
	if n%3 == 0 {
		x = float64(n*10) / 3
	} else {
		x = float64(((n-n%3)/3)*10 + (n%3)*5)
	}
	fmt.Println("\nO valor a pagar é= ", x)
}
