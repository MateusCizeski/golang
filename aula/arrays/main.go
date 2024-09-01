package main

import "fmt"

//Array > Um array em GO é uma sequência de elementos de um único tipo,
//com um tamanho fixo definido na sua declaração

func main() {

	var numeros [4]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 15
	numeros[3] = 5

	fmt.Println(numeros)

	var valores = [3]int{3, 10, 90}

	fmt.Println(valores)

	permissoes := [3]string{"usuario", "admin", "editor"}

	fmt.Println(permissoes)
}
