package main

import "fmt"

func main() {

	fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	defer fmt.Println("Terceira Linha")
	fmt.Println("Quarta Linha")
}
