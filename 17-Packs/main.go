package main

import (
	"fmt"

	"aprendizadoGo-Course/matematica"
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Printf("O resultado é %d\n", s)
	fmt.Printf("Valor é %d\n", matematica.A)

	carro := matematica.Carro{Marca: "Fiat"}

	println("Carro é: ", carro.Marca)

	println(carro.Andar())

}
