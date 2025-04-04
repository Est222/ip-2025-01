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
	var a, b, c float64
	var delta float64
	fmt.Scan(&a, &b, &c)
	delta = (math.Pow(b, 2)) - 4*a*c
	fmt.Println("O valor de delta é= ", roundFloat(delta, 2))

}
