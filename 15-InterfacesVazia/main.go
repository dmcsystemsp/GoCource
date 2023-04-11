package main

import "fmt"

func main() {

	var minhaVar interface{} = "Hello World"
	println(minhaVar)          // vai imprimir o endereçamento de memória
	println(minhaVar.(string)) // vai imprimir o conteúdo da variável

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é: %v e resultado de OK é: %v\n", res, ok)

	res2 := minhaVar.(int)
	fmt.Printf("O valor de res é: %v\n", res2)

	/*
		var x interface{} = 10
		var y interface{} = "Hello World"

		showType(x)
		showType(y)
	*/

}

func showType(t interface{}) {
	fmt.Printf("O Tipo da variável é: %T e o valor é: %v\n", t, t)
}
