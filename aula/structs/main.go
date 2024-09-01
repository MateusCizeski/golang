package main

import "fmt"

type Usuario struct {
	Nome      string
	Idade     int
	Status    bool
	Permissao string
}

func main() {
	usuarios := map[int]Usuario{
		1: {Nome: "Mateus", Idade: 24, Status: true, Permissao: "Admin"},
		2: {Nome: "Mateus2", Idade: 24, Status: false, Permissao: "User"},
	}

	usuarios[3] = Usuario{Nome: "Mateus3", Idade: 25, Status: false, Permissao: "user"}

	usuarios[len(usuarios)+1] = Usuario{Nome: "Mateus4", Idade: 25, Status: true, Permissao: "admin"}

	i := 1

	for i <= len(usuarios) {
		// fmt.Println(usuarios[i])
		fmt.Printf("Nome: %s, Idade: %d, permissao: %s \n", usuarios[i].Nome, usuarios[i].Idade, usuarios[i].Permissao)
		i++
	}

	fmt.Println("=================================================")
	usuarios[len(usuarios)+1] = Usuario{Nome: "Mateus5", Idade: 25, Status: false, Permissao: "admin"}

	usuarios[1] = Usuario{Nome: "Mateus Programador", Idade: 24, Status: true, Permissao: "Admin"}

	fmt.Println(usuarios)
}
