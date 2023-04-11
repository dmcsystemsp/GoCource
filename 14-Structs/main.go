package main

/*
type Client struct {
	nome string
}

func (c Client) andou() {
	c.nome = "Daniel Carneiro"
	fmt.Printf("O Cliente %v andou.\n", c.nome)

}
*/

type Conta struct {
	saldo int
}

func newConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {

	conta := Conta{
		saldo: 100,
	}

	conta.simular(200)
	println(conta.saldo)

	/*
		cli := Client{
			nome: "Daniel",
		}
		cli.andou()
		fmt.Printf("O valor da Struct com nome %v\n", cli.nome)

	*/
}
