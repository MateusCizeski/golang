package main

//go mod init projeto

import (
	"fmt"
	"projeto/aluno"
	"projeto/utils"
)

func main() {
	utils.Mensagem("Mateus")

	aluno1 := aluno.Aluno{Nome: "Mateus", N1: 10, N2: 7.5}

	aluno.Cadastro(aluno1)
	media := aluno.CalcularMedia(aluno1.N1, aluno1.N2)

	fmt.Println("media do aluno é:", media)

	aluno.VerificarStatusAluno(media)
}
