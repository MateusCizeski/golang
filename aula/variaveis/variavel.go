package main

import "fmt"

//Váriaveis do Tipo Var e Const

func main() {
	var nome = "Mateus"
	var idade = 24
	var texto string

	fmt.Println("Variaveis")
	fmt.Println(nome)
	fmt.Println(idade)

	idade = 25
	texto = "Meu primeiro projeto em GO"

	fmt.Println("Agora a idade é", idade)
	fmt.Println(texto)

	//Const
	const pi = 3.14159
	//constante tem que ser inicializada com um valor.
	// const valor int

	fmt.Println(pi)
}
