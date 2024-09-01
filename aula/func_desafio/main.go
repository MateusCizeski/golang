package main

import "fmt"

type Aluno struct {
	Nome  string
	Nota1 float64
	Nota2 float64
}

func Media(aluno Aluno) (string, float64) {
	media := (aluno.Nota1 + aluno.Nota2) / 2

	return aluno.Nome, media
}

func VerificarAprovacao(media float64) (string, bool) {
	if media >= 7 {
		return "Aprovado", true
	} else {
		return "Recuperação", false
	}
}

func main() {
	aluno1 := Aluno{Nome: "Mateus 50 de braço", Nota1: 8.5, Nota2: 7.5}
	nome, media := Media(aluno1)
	aprovado, status := VerificarAprovacao(media)

	fmt.Printf("A média do %s é de %f \n", nome, media)
	fmt.Printf("O aluno está %s, %t \n", aprovado, status)
}
