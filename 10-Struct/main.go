package main

type Endereco struct {
	Logadouro string
	Cidade    string
}

type Client struct {
	Id      int
	Name    string
	Address Endereco
	Phone   string
	Ativo   bool
}

func main() {
	cliente := Client{
		Id:    1,
		Name:  "Daniel Carneiro",
		Phone: "11 3645-2247",
		Ativo: true,
	}
	cliente.Address.Logadouro = "Rua Frei Caneca"

	println(cliente.Name)
	println(cliente.Address.Logadouro)
}
