package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}
func main() {

	n1 := 10
	n2 := 20

	println(soma(&n1, &n2))
	println(n1)

}
