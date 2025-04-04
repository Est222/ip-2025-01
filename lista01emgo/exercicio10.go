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
	var a, b, c, d, det float64
	fmt.Scan(&a, &b, &c, &d)
	det = a*d - b*c
	fmt.Println("O valor do determinante é= ", roundFloat(det, 2))
}
