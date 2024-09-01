package main

import "fmt"

//string
//bool > true false
//int | int8, int 16, int 64, uint8, uint16, uint64 - uint não aceita negativo
//float32 float64
//Bytes, Runes

func main() {
	var nome string = "Mateus"
	var isAdmin bool = false
	idade := 24

	var valor float64 = 900.90
	valor2 := 350.20

	var numero int = 127

	fmt.Println(nome)
	fmt.Println(isAdmin)
	fmt.Println(idade)
	fmt.Println(numero)
	fmt.Println(valor)
	fmt.Println(valor2)

	var teste byte = 'A'
	var runTeste rune = '&'
	fmt.Println(teste)
	fmt.Println(runTeste)
}
