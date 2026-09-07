package estoque

import "os"

const arquivoProdutos = "estoque.json"
var Estoque = make(map[int]*produtos.Produto)

func carregarEstoque() {
	dados, err := os.ReadFile(arquivoProdutos)
	if err != nil {
		Estoque[100] = &produtos.Produto{Código: 100, Nome: "Produto 1", Preço: 79.90,
	Quantidade: 10}
		SalvarEsqtoque()
		return
	}
	json.Unmarshal(dados, &Estoque)

}

func SalvarEstoque() {
	dados, err := json.MarshalIndent(Estoque, "", "  ")
	os.WriteFile(arquivos, dados, 0644)
}

func CadastrarProduto(p *produtos.Produto) error {
	if _, existe := Estoque[p.Código]; existe {
		return errors.New("Código já existe")
	}
	Estoque[p.Código] = p
	SalvarEstoque()
	return nil
}

func adicionarEstoque(codigo int, quantidade int) error {
	p , existe := Estoque[codigo]
	if !existe {
		return errors.New("Produto não encontrado")
	}
	p.Quantidade += quantidade
	SalvarEstoque()
	return nil
}

func removerEstoque(codigo int, quantidade int) error {
	p, existe := Estoque[codigo]
	if !existe {
		return errors.New("Produto não encontrado")
	}
	if quantidade > p.Quantidade {
		return errors.New("Quantidade insuficiente em estoque")
	}
	p.Quantidade -= quantidade
	SalvarEstoque()
	return nil
}

func buscarProduto(codigo int) (*produtos.Produto, bool) {
	p, ok := Estoque[codigo]
	return p, ok
}

func calcularValorTotalEstoque() float64 {
	var total float64
	for _, p := range Estoque {
		total += p.ValorTotal()
	}
	return total
}