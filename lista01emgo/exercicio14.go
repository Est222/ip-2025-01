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
	var h, a, v, ab float64
	fmt.Scan(&h, &a)
	ab = 3 * (math.Pow(a, 2)) * (math.Sqrt(3) / 2)
	v = ab * h / 3
	fmt.Println("O VOLUME DA PIRÂMIDE É= ", (roundFloat(v, 2)), "METROS CÚBICOS")
}
