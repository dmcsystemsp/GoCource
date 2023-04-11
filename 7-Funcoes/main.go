package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(sum(5, 5))

	valor, err := soma(8, 8)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

}

func sum(a int, b int) int {
	return (a + b)
}

func soma(a, b int) (int, error) {

	if a+b >= 20 {
		return 0, errors.New("A Soma é maior que 20")
	}
	return a + b, nil

}
