package main

import (
	"fmt"
)

func main() {

	var n float64
	fmt.Scan(&n)
	if n < 6 {
		fmt.Println("NOTA=", n, "CONCEITO=D")
	} else if n >= 6 && n < 7.5 {
		fmt.Println("NOTA=", n, "CONCEITO=C")
	} else if n >= 9 {
		fmt.Println("NOTA=", n, "CONCEITO=A")
	}
}
