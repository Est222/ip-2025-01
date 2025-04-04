package main

import "fmt"

func main() {
	var f, n int
	var c float64
	var x = 1
	fmt.Println("Quantos valores quer transformar?")
	fmt.Scan(&n)
	for x <= n {
		fmt.Println("Informe a temperatura em Fahrenheit para ser transformado em Celsius")
		fmt.Scan(&f)
		c = float64(5*(f-32)) / 9
		fmt.Println(f, "\nFahrenheit equivale a", c, "Celsius")
		x++
	}
}
