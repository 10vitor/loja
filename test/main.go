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

func main() {
    fmt.Println("===== CONTROLE DE ESTOQUE =====")
    fmt.Println("1 - Cadastrar produto")
    fmt.Println("2 - Listar produtos")
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

        product := Product{
            Name:     nome,
            Price:    preco,
            Code:     codigo,
            Quantity: quantidade,
        }

        // Aqui você chamaria seu product_repository pra salvar
        // ex: productRepo.Create(product)
        fmt.Println("Produto cadastrado:", product)

    case 0:
        fmt.Println("Saindo...")
    default:
        fmt.Println("Opção inválida")
    }
}
