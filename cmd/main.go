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
)

type Product struct{
    Name     string
    Price    float64
    Code     int
    Quantity int
}

type Pedido struct{
    Cliente Cliente*
    Produto Produto*
    Quantidade int 
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

    var opcao int
    fmt.Scanln(&opcao)

    switch opcao {
    case 1:
        var nome string
        var preco float64
        var codigo int
        var quantidade int

        fmt.Println("Nome do produto: ")
        fmt.Scanln(&nome)
        fmt.Println("Preco: ")
        fmt.Scanln(&preco)
        fmt.Println("Codigo: ")
        fmt.Scanln(&codigo)
        fmt.Println("Quantidade: ")
        fmt.Scanln(&quantidade)

        fmt.Println("Produto cadastrado:", product)

case 0:
    fmt.Println("Produto nao cadastrado")
default:
    fmt.Println("Opcao invalida: ")

    }
}