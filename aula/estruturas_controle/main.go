package main

import "fmt"

//if else | else if
//switch

func main() {
	nota1 := 5.5
	nota2 := 6.5
	media := (nota1 + nota2) / 2

	if media >= 7 {
		fmt.Println("Aprovado!")
		fmt.Println("Sua média foi:", media)
	} else if media >= 5 && media <= 6.99 {
		fmt.Println("Recuperação!")
		fmt.Println("Sua média foi:", media)
	} else {
		fmt.Println("Reprovado!")
		fmt.Println("Sua média foi:", media)
	}

	//===============================================

	dia := 5

	switch dia {
	case 1:
		fmt.Println("Segunda-feira")
	case 2:
		fmt.Println("terça-feira")
	case 3:
		fmt.Println("quarta-feira")
	case 4:
		fmt.Println("quinta-feira")
	case 5:
		fmt.Println("Sexta-feira")
	default:
		fmt.Println("Fim de semana")
	}

	redeSocial := "instagram"

	switch redeSocial {
	case "instagram":
		fmt.Println("teste1")
	case "youtube":
		fmt.Println("teste2")
	case "blog":
		fmt.Println("teste3")
	}
}
