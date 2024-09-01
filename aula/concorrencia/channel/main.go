package main

import (
	"fmt"
)

func sayHello(chName chan string) {
	chName <- "Bem vindo ao sistema"
}

func main() {
	ch := make(chan int)
	chName := make(chan string)

	go sayHello(chName)

	go func() {
		ch <- 123
	}()

	fmt.Print("utilizando channels \n")

	valor := <-ch
	fmt.Print("Valor do numero: ", valor)
	fmt.Print("\n")

	msg := <-chName
	fmt.Print(msg)
}
