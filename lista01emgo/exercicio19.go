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
	var n, i int
	var soma=0
	fmt.Scan(&n)
	if n<=1 {
		fmt.Println("Número inválido!")
	} else {
		for (i==1 && i<=n && i++){ 
			soma=soma+(1/i)
		}
	fmt.Println(roundFloat(soma, 6))
	}
}