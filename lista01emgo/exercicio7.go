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

	var p, m, f, c float64
	fmt.Scan(&f, &p)
	c = 5 * (f - 32) / 9
	m = 25.4 * p
	fmt.Println("O valor em celsius= ", roundFloat(c, 2))
	fmt.Println("\nA quantidade de chuva é= ", roundFloat(m, 2))

}
