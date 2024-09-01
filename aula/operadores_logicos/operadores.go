package main

import "fmt"

func main() {
	//and &&
	//or ||
	//negação !

	estoque := false
	vendaLiberada := false
	freteGratis := false

	nome1 := "mateus"
	nome2 := "mateus"

	fmt.Println("AND:", estoque && vendaLiberada)

	fmt.Println("OR:", estoque || vendaLiberada)

	fmt.Println("Negação:", !freteGratis)

	fmt.Println(nome1 != nome2)
}
