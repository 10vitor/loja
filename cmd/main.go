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
    fmt.Println("3 - Buscar produto")
    fmt.Println("4 - Adicionar estoque")
    fmt.Println("5 - Remover estoque")
    fmt.Println("6 - Valor total")
    fmt.Println("0 - Sair")

    fmt.Print("Escolha uma opção: ")

    var opcao int
    fmt.Scanln(&opcao)

    var input string
    fmt.Scanln(&input)

    product := Product{
        Name:     "Camisas",
        Price:    100.00,
        Code:     12345,
        Quantity: 2,
    }
    fmt.Println(product)
}

