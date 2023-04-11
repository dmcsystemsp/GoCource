package main

import "fmt"

func main() {

	salarios := map[string]int{"Daniel": 5000, "Paulo": 3000, "João": 2000, "Maria": 1000}
	fmt.Println(salarios["Paulo"])
	delete(salarios, "Paulo")
	salarios["Paul"] = 500
	fmt.Println(salarios["Paul"])

	/*
		sal := make(map[string]int)
		sal1 = map[string]int{}

		sal["Daniel"] = 1000
		sal1["Pedro"] = 500

	*/

	for nome, salario := range salarios {
		fmt.Printf("O Salário de %s é de %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O Salário é %d\n", salario)
	}

}
