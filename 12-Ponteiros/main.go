package main

func main() {

	// Variável -> Endereço  de Memória que tem um -> Valor
	a := 10
	println(a)
	println(&a)

	// A variável aponta para o ponteiro que tem um endereço a memória ->valormain
	var ponteiro *int = &a
	*ponteiro = 20

	println(a)
	b := &a
	println(b)
	println(*b)

	*b = 30
	println(a)

}
