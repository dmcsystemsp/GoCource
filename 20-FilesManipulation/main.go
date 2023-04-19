package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//tamanho1, err := f.WriteString("Daniel Carneiro")
	tamanho, err := f.Write([]byte("Escrevendo dados no Arquivo"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo criado com sucesso! Tamanho: %d bytes", tamanho)

	f.Close()

	//Lendo um arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//Lendo um arquivo linha a linha
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	arquivo2.Close()

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
