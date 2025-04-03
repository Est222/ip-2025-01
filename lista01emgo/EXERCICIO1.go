package main

import "fmt"
"math"

func main() {

	var n1, n2, n3, media float64
	fmt.Println("Informe as notas do aluno:")
	fmt.Scan(&n1, &n1 &n3)
	media=(n1+n2+n3)/3
	fmt.Println("A média das notas do aluno é:", media)

	if(media<6) {	
		fmt.Println("\nO aluno foi APROVADO")
	} else {
		fmt.Println("\nO aluno foi REPROVADO")
	}
}