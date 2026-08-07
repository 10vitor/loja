package main

import "fmt"

type Product struct{
    Name     string
    Price    float64
    Code     int
    Quantity int
}

func main() {
    fmt.Println("===== CONTROLE DE ESTOQUE ===== (1 - Cadastrar produto, 2 - Listar produtos, 3 - Buscar produto, 4 - Adicionar estoque, 5 - Remover estoque, 6 - Calcular valor total, 0 - Sair): ")

    product := Product{
        Name:     "Camisas",
        Price:    100.00,
        Code:     12345,
        Quantity: 20,
    }
    fmt.Println(product.Name)
}

