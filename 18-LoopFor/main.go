package main

func main() {

	for i := 1; i <= 10; i++ {
		println(i)
	}

	numeros := []string{"Um", "Dois", "Três", "Quatro"}

	for k, v := range numeros {
		println(k, v)
	}

	i := 1
	for i <= 10 {
		println(i)
		i++
	}

	for {
		println("Loop Infinito!")
	}
}
