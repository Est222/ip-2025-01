package main

import (
	"fmt"
	"math"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
	var n, m float64
	fmt.Scan(&n)
	if n <= 300 {
		m = n + 0.5*n
	} else {
		m = n + 0.3*n
	}
	fmt.Println("Salário com reajuste= ", (roundFloat(m, 2)))
}
