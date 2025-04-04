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
	var r, h, at, ac, al, x float64
	fmt.Scan(&r, &h)
	ac = float64((math.Pow(r, 2))) * 3.14159
	al = 2 * r * h * 3.14159
	at = 2*ac + al
	x = at * 100
	fmt.Println("O valor do custo é=", (roundFloat(x, 2)))
}
