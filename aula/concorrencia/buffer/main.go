package main

import (
	"fmt"
	"math/rand"
	"time"
)

func prepareOrder(orderId int, completedOrders chan string) {
	fmt.Printf("preparando o pedido #%d... \n", orderId)

	//simulação de tempo de preparação
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	//enviando o pedido para o channel
	completedOrders <- fmt.Sprintf("Pedido #%d concluido!!", orderId)
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	orderCount := 4
	completedOrders := make(chan string, orderCount)

	// Lança uma goroutine para cada pedido
	for i := 1; i <= orderCount; i++ {
		go prepareOrder(i, completedOrders)
	}

	for i := 1; i <= orderCount; i++ {
		fmt.Println(<-completedOrders)
	}

	//fechando o canal após o uso
	close(completedOrders)
}
