package main

import "fmt"

func main() {
	var numeros []int

	numeros = append(numeros, 50, 10, 99, 30)

	fmt.Println("Conteudo do slice numeros:", numeros)
	fmt.Println("Pegando a posição 0:", numeros[0])

	frutas := []string{"banana", "maça", "mamão"}

	fmt.Println("Solice de futas:", frutas)

	frutas = append(frutas, "pêra")

	fmt.Println("Solice de futas após append:", frutas)

	nomes := make([]string, 0)

	fmt.Println("Slice nomes:", nomes)

	nomes = append(nomes, "mateus", "meteus", "mateus")

	fmt.Println("Slice de nomes:", nomes)
}
