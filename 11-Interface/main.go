package main

import "fmt"

type Endereco struct {
	Logadouro string
	Numero    int
	Cidade    string
	Estado    string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	nome string
}

func (e Empresa) Desativar() {

}

type Client struct {
	Id      int
	Name    string
	Address Endereco
	Phone   string
	Ativo   bool
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O Cliente %s foi desativado!", c.Name)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	cli := Client{
		Id:    1,
		Name:  "Daniel Carneiro",
		Phone: "11 3645-2247",
		Ativo: true,
	}

	Desativacao(cli)
	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)
}
