package aluno

import "fmt"

type Aluno struct {
	Nome string
	N1   float64
	N2   float64
}

func Cadastro(aluno Aluno) {
	fmt.Printf("Aluno: %s foi cadastro com sucesso \n", aluno.Nome)
}

func CalcularMedia(nota1 float64, nota2 float64) float64 {
	media := (nota1 + nota2) / 2

	return media
}

func VerificarStatusAluno(media float64) {
	if media >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Recuperação")
	}
}
