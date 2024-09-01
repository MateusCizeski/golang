package main

import (
	"fmt"
	"time"
)

// Thread 2
func sayHello() {
	fmt.Print("Olá mundo!! \n")
}

// Thread 3
func bemVindo(nome string) {
	fmt.Print("Bem vindo ", nome)
}

// Thread 1
func main() {
	// Thread 2
	go sayHello()

	// Thread 3
	go bemVindo("Mateus")

	fmt.Println("Aprendendo concorrencia")

	time.Sleep(1 * time.Second)
}
