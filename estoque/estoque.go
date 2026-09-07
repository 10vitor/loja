package estoque

import (
	"encoding/json"
	"errors"
	"loja/produtos"
	"os"
)

const arquivoProdutos = "estoque.json"
var Estoque = make(map[int]*produtos.Product)

func carregarEstoque() {
	dados, err := os.ReadFile(arquivoProdutos)
	if err != nil {
		Estoque[100] = &produtos.Product{
		Code: 100,
		Name: "Camisa Polo",
		Price: 79.90,
		Quantity: 10}
		SalvarEstoque()
		return
	}
	json.Unmarshal(dados, &Estoque)

}

func SalvarEstoque() {
	dados, _ := json.MarshalIndent(Estoque, "", "  ")
	os.WriteFile(arquivoProdutos, dados, 0644)
}

func CadastrarProduto(p *produtos.Product) error {
	if _, existe := Estoque[p.Code]; existe {
		return errors.New("Código já existe")
	}
	Estoque[p.Code] = p
	SalvarEstoque()
	return nil
}

func adicionarEstoque(codigo int, quantidade int) error {
	p , existe := Estoque[codigo]
	if !existe {
		return errors.New("Produto não encontrado")
	}
	p.Quantity += quantidade
	SalvarEstoque()
	return nil
}

func removerEstoque(codigo int, quantidade int) error {
	p, existe := Estoque[codigo]
	if !existe {
		return errors.New("Produto não encontrado")
	}
	if quantidade > p.Quantity {
		return errors.New("Quantidade insuficiente em estoque")
	}
	p.Quantity -= quantidade
	SalvarEstoque()
	return nil
}

func buscarProduto(codigo int) (*produtos.Product, bool) {
	p, ok := Estoque[codigo]
	return p, ok
}

func calcularValorTotalEstoque() float64 {
	var total float64
	for _, p := range Estoque {
		total += p.TotalValue()
	}
	return total
}