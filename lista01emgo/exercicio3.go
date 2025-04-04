package main

import "fmt"

func main() {

	var n1, n2, n3 float64
	fmt.Println("Escreva 3 números")
	fmt.Scan(&n1, &n2, &n3)
	fmt.Println(&n1, &n2, &n3)
	var fff = n1*100 + n2*10 + n3
	var quadrada = fff * fff
	fmt.Println("\n", quadrada)

}
