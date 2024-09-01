package main

import "fmt"

func main() {
	p := 7.500
	i := 0.30
	n := 12
	juros := float64(p) * i * float64(n)

	fmt.Printf("O juros ao mês é: %.2f", juros)
}
