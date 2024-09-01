package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// while (enquanto)

	numero := 0

	for numero < 5 {
		numero++
		fmt.Println(numero)
	}

	for {
		fmt.Println("Loop infinito")
		break
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("Valor do i é igual a 3")
			continue
		}

		if i == 4 {
			break
		}

		fmt.Println("Valor do i", i)
	}
}
