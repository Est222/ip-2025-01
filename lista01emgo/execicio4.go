import (
	"fmt"
	"math"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
	var sal, qkw, custo, consumo, custod float64
	fmt.Scan(&sal, &qkw)
	custo = (sal * 7) / 1000
	fmt.Println(custo)
	consumo = (custo * qkw)
	custod = consumo - (consumo * 0.1)
	fmt.Println("Custo por kW: R$", (roundFloat(consumo, 2)))
	fmt.Println("\nCusto do consumo: R$", (roundFloat(custo, 2)))
	fmt.Println("\nCusto com desconto: R$", (roundFloat(custod, 2)))
}