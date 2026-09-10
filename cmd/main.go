//Requisitos do estoque da loja:
//Cadastrar produtos com nome, preço, código e quantidade;
//Listar produtos cadastrados;
//Buscar produtos por código;
//Adicionar e remover estoque;
//Validar se o produto existe antes de adicionar ou remover estoque;
//Calcular o valor total de cada item e do estoque.

package main

import (
    "fmt"
    "loja/produtos"
    "loja/estoque"
)

type Product struct{
    Name     string
    Price    float64
    Code     int
    Quantity int
}

func main() {
    fmt.Println("===== CONTROLE DE ESTOQUE =====")
    fmt.Println("1 - Cadastrar produto")
    fmt.Println("2 - Listar produtos")
    fmt.Println("3 - Buscar produto")
    fmt.Println("4 - Adicionar estoque")
    fmt.Println("5 - Remover estoque")
    fmt.Println("6 - Valor total")
    fmt.Println("0 - Sair")

    fmt.Print("Escolha uma opção: ")

    estoque := []Product{
        {Name: "Camiseta", Price: 59.99, Code: 111, Quantity: 10},
        {Name: "Calça Jeans", Price: 99.99, Code: 222, Quantity: 5},
        {Name: "Tênis", Price: 149.99, Code: 333, Quantity: 8},
    }

    fmt.Println("No meu estoque tem", len(estoque), "produtos")

    var opcao int
    fmt.Scanln(&opcao)

    switch opcao {
    case 1:
        var nome       string
        var preco      float64
        var codigo     int
        var quantidade int

        fmt.Println("Nome do produto: ")
        fmt.Scanln(&nome)
        fmt.Println("Preco: ")
        fmt.Scanln(&preco)
        fmt.Println("Codigo: ")
        fmt.Scanln(&codigo)
        fmt.Println("Quantidade: ")
        fmt.Scanln(&quantidade)

        produto := Product{
            Name:     nome,
            Price:    preco,
            Code:     codigo,
            Quantity: quantidade, 
        }

        fmt.Println("Produto cadastrado:", produto)

case 0:
    fmt.Println("Produto nao cadastrado")
default:
    fmt.Println("Opcao invalida: ")

    }
}