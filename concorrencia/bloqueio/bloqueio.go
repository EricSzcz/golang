package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // Operação bloqueante
	fmt.Println("Só depois que for lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // Operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c)
	fmt.Println("Fim")
}
