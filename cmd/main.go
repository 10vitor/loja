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

    err := estoque.Load()

    if err != nil {
        fmt.Println("Nenhum estoque encontrado.")
    } else {
        fmt.Println("Estoque carregado com sucesso!")
    }

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
    // CADASTRO
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

        err := estoque.Save()

        if err != nil {
	        fmt.Println("Erro ao salvar estoque:", err)
        } else {
	        fmt.Println("Estoque salvo com sucesso!")
        }

        fmt.Println("Produto cadastrado:", produto)

    case 2:
    // LISTAGEM
    fmt.Println("===== PRODUTOS EM ESTOQUE =====")

    if len(estoque.Products) == 0 {
        fmt.Println("Nenhum produto cadastrado.")
        continue
    }

    for _, produto := range estoque.Products {
        fmt.Println("Nome:", produto.Name)
        fmt.Println("Quantidade:", produto.Quantity)

    }

    case 3: { 
    // BUSCA
        var codigo int
        fmt.Print("Digite o código do produto: ")
        fmt.Scanln(&codigo)

        encontrado := false

        for _, produto := range estoque.Products {

            if produto.Code == codigo {
                fmt.Println("=== PRODUTO ENCONTRADO ===")
                fmt.Println("Nome:", produto.Name)
                fmt.Println("Preço:", produto.Price)
                fmt.Println("Código:", produto.Code)
                fmt.Println("Quantidade:", produto.Quantity)

                encontrado = true
                break
        }
    }
    if !encontrado {
        fmt.Println("Produto não encontrado.")
    }
    }
    case 4: {
    var codigo int
    var quantidade int

    fmt.Print("Digite o código do produto: ")
    fmt.Scanln(&codigo)

    fmt.Print("Quantidade para adicionar: ")
    fmt.Scanln(&quantidade)

    encontrado := false

    for i := range estoque.Products {

        if estoque.Products[i].Code == codigo {

            estoque.Products[i].Quantity += quantidade

            encontrado = true

            err := estoque.Save()

            if err != nil {
                fmt.Println("Erro ao salvar estoque:", err)
            } else {
                fmt.Println("Estoque atualizado com sucesso!")
                fmt.Println("Quantidade atual:", estoque.Products[i].Quantity)
            }

            break
        }
    }

    if !encontrado {
        fmt.Println("Produto não encontrado.")
    }

    if !encontrado {
        fmt.Println("Produto não encontrado.")
    }
    }

    case 5: {
    var codigo int
    var quantidade int

    fmt.Print("Digite o código do produto: ")
    fmt.Scanln(&codigo)

    fmt.Print("Quantidade para remover: ")
    fmt.Scanln(&quantidade)

    encontrado := false

    for i := range estoque.Products {

        if estoque.Products[i].Code == codigo {

            if quantidade > estoque.Products[i].Quantity {
                fmt.Println("Quantidade insuficiente em estoque.")
                encontrado = true
                break
            }

            estoque.Products[i].Quantity -= quantidade

            encontrado = true

            err := estoque.Save()

            if err != nil {
                fmt.Println("Erro ao salvar estoque:", err)
            } else {
                fmt.Println("Estoque atualizado com sucesso!")
                fmt.Println("Quantidade atual:", estoque.Products[i].Quantity)
            }

            break
        }
    }

    if !encontrado {
        fmt.Println("Produto não encontrado.")
    }
    }

    case 6: {
    var valorTotal float64

    for _, produto := range estoque.Products {
        valorTotal += produto.Price * float64(produto.Quantity)
    }

    fmt.Println("===== VALOR TOTAL DO ESTOQUE =====")
    fmt.Printf("Valor total: R$ %.2f\n", valorTotal)
}

    case 0:
        fmt.Println("Saindo...")
        return

    default:
        fmt.Println("Opção inválida")
    }
}
}
