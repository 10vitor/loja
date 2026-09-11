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
    "loja/internal/product"
    "loja/internal/product/inventory"
)

func main() {

    estoque := inventory.Inventory{}

    for {

    fmt.Println("===== CONTROLE DE ESTOQUE =====")

    fmt.Println("1 - Cadastrar produtos")
    fmt.Println("2 - Listar produtos")
    fmt.Println("3 - Buscar produto")
    fmt.Println("4 - Adicionar estoque")
    fmt.Println("5 - Remover estoque")
    fmt.Println("6 - Calcular valor total")
    fmt.Println("0 - Sair")

    fmt.Print("Escolha uma opção: ")

    var opcao int
    fmt.Scanln(&opcao)

    switch opcao {

    case 1:

        var nome string
        var preco float64
        var codigo int
        var quantidade int

        fmt.Print("Nome do produto: ")
        fmt.Scanln(&nome)

        fmt.Print("Preço: ")
        fmt.Scanln(&preco)

        fmt.Print("Código: ")
        fmt.Scanln(&codigo)

        fmt.Print("Quantidade: ")
        fmt.Scanln(&quantidade)

        produto := product.Product{
            Name:     nome,
            Price:    preco,
            Code:     codigo,
            Quantity: quantidade,
        }

        estoque.AddProduct(produto)

        fmt.Println("Produto cadastrado:", produto)

    case 2:
    fmt.Println("===== PRODUTOS EM ESTOQUE =====")

    for _, produto := range estoque.Products {
        fmt.Println("Nome:", produto.Name)
        fmt.Println("Preço:", produto.Price)
        fmt.Println("Código:", produto.Code)
        fmt.Println("Quantidade:", produto.Quantity)

    }

    case 0:
        fmt.Println("Saindo...")

    default:
        fmt.Println("Opção inválida")
    }
}
}