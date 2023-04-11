package main

import "fmt"

const a = "Hello World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Daniel"
	e float32 = 4.33
	f ID      = 1
)

func main() {
	fmt.Printf("O valor de E é: %T", e)
	fmt.Printf("O valor de E é: %V", e)
	fmt.Printf("O valor de E é: %T", f)

}
