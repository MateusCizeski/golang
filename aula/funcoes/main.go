package main

import (
	"fmt"
)

func novoUsuario(userName string) (string, bool) {
	if userName != "" {
		return userName, true
	}

	return "Sem username", false
}

func main() {
	usuario1 := "Mateus"

	resultado, status := novoUsuario(usuario1)
	fmt.Printf("usuário atual: %s, status atual %t \n", resultado, status)

	resultado2, status2 := novoUsuario("")
	fmt.Printf("usuário atual: %s, status atual %t \n", resultado2, status2)

	func() {
		fmt.Println("Função anonima foi chamada!")
	}()
}
