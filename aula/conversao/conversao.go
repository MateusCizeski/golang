package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Conversão de Inteiro para Float
	var valor int = 42
	var valorConvertido float64 = float64(valor)

	fmt.Println("Inteiro:", valor)
	fmt.Println("Valor em float:", valorConvertido)

	//Inteiro para uma string
	var valorString string = strconv.Itoa(valor)

	fmt.Println(valorString)

	//Converter string para inteiro
	var valor2 string = "230.50"

	valorInteiro, err := strconv.Atoi(valor2)

	if err != nil {
		fmt.Println("Erro ao converter")
	} else {
		fmt.Println(valorInteiro + 20)
	}

	//Converter de string para float
	var pi string = "3.14159"

	piConvertido, error2 := strconv.ParseFloat(pi, 64)

	if error2 != nil {
		fmt.Println("Erro ao converter para float")
	} else {
		fmt.Println(piConvertido)
	}
}
