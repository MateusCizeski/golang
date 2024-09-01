package utils

import "fmt"

// >> Exportada (pública): começa com a letra maiúscula e pode ser acessada por outros pacotes
// >> Não exportada (privada): Começa com a letra minúscula e só pode ser acesada dentro do mesmo pacote.
func Mensagem(nome string) {
	fmt.Printf("Olá %s, seja bem vindo a plataforme \n", nome)
}
