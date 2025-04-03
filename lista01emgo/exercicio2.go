package main

import "fmt"

func main() {
	var total float64
	var popular, geral, arquibancada, cadeiras, n int
	var x = 1
	fmt.Println("Quantos jogos foram jogados neste dia? ")
	fmt.Scan(&n)
	for x < n {
		fmt.Println("\nInforme os valores do jogo")
		fmt.Scan(&popular, &geral, &arquibancada, &cadeiras)
		total = float64(popular + (geral * 2) + (arquibancada * 10) + (cadeiras * 20))
		fmt.Println("\nA renda total do jogo foi de:", total)
		x++
	}

}
