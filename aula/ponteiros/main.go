package main

import "fmt"

//Um ponteiro é uma variável que armazena  o endereço de memória de outra variável
func Somar(num *int) {
	*num = *num + 1
}

func Subtrair(num *int) {
	*num = *num - 1
}

func main() {
	numero := 10

	fmt.Println("Valor inicial da variavel numero:", numero)
	fmt.Println("Valor de memória da variavel numero:", &numero)

	Somar(&numero)
	Somar(&numero)
	Somar(&numero)
	Somar(&numero)

	fmt.Println("Valor atual da variavel numero após somar:", numero)

	Subtrair(&numero)
	Subtrair(&numero)
	Subtrair(&numero)

	fmt.Println("Valor atual da variavel numero após subtrair:", numero)
}
