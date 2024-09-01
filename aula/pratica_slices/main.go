package main

import "fmt"

func main() {
	var tarefas []string

	//Adicionar itens no slice
	tarefas = append(tarefas, "Estudar GO")
	tarefas = append(tarefas, "jogar Elden ring")
	tarefas = append(tarefas, "jogar mais elden ring")

	fmt.Println("Tarefas:", tarefas)
	fmt.Println("Tamanho do slice atual:", len(tarefas))

	//Slicing / cortar
	// tarefas = tarefas[1:]
	// fmt.Println("Remover a primeira tarefa:", tarefas)

	//ultimo item do slice
	// tarefas = tarefas[:len(tarefas)-1]
	// fmt.Println(tarefas)

	tarefas = append(tarefas[:1], tarefas[2:]...)
	fmt.Println("Removendo o segundo item do slice:", tarefas)
}
