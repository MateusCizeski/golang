package main

import "fmt"

func main() {
	//adição
	valor1 := 10
	valor2 := 7

	soma := valor1 + valor2

	fmt.Println("Resultado da soma:", soma)

	//subtração ( - )
	subtracao := valor1 - valor2
	fmt.Println("Resultado da subtração:", subtracao)

	//multiplicação ( * )
	multiplicacao := valor1 * valor2
	fmt.Println("Resultado da multiplicação:", multiplicacao)

	//divisão
	nota1 := 10.00
	nota2 := 2.00

	divisao := (nota1 + nota2) / 2
	fmt.Println("Resultado da nota:", divisao)

	x := 10
	y := 3

	resto := x % y

	fmt.Println("Resto da divisão:", resto)

}
