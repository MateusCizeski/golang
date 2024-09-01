package main

import "fmt"

func main() {
	anoAtual := 2024
	anoNascimento := 2000

	idadeAtual := anoAtual - anoNascimento
	fmt.Println("Idade atual é:", idadeAtual)

	horas := 1
	totalMinutos := horas * 60

	fmt.Println("Total em minutos:", totalMinutos)
	fmt.Printf("%d hora tem %d minutos", horas, totalMinutos)
}
