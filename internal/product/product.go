package product

import (
	"errors"
	"fmt"
)


type Product struct {
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Code     int      `json:"code"`
	Quantity int  	  `json:"quantity"`
}

func RegisterProduct(products []Product) []Product {
	var code, quantity int
	var name string
	var price float64

	fmt.Print("Código: ")
	fmt.Scan(&code)


	if _, found := FindProductByCode(products, code); found {
		fmt.Println("Erro: Código já existe!")
		return products
	}

	fmt.Print("Nome: ")
	fmt.Scan(&name)
	fmt.Print("Preço: ")
	fmt.Scan(&price)
	fmt.Print("Quantidade: ")
	fmt.Scan(&quantity)

	newProduct := Product{
		Code: code, Name: name, Price: price, Quantity: quantity,
	}

	fmt.Println("Produto cadastrado!")
	return append(products, newProduct)
}

func ListProducts(products []Product) {
	if len(products) == 0 {
		fmt.Println("Estoque vazio!")
		return
	}
	fmt.Println("\n--- ESTOQUE ---")
	for _, p := range products {
		totalItem := float64(p.Quantity) * p.Price
		fmt.Printf("Cód: %d | %s | Qtd: %d | R$ %.2f | Total Item: R$ %.2f\n",
			p.Code, p.Name, p.Quantity, p.Price, totalItem)
	}
	fmt.Printf("VALOR TOTAL GERAL: R$ %.2f\n", CalculateTotalStock(products))
}

func FindProductByCode(products []Product, code int) (int, bool) {
	for i, p := range products {
		if p.Code == code {
			return i, true
		}
	}
	return -1, false
}

func AddStock(products []Product, code int, quantity int) error {
	index, found := FindProductByCode(products, code)
	if!found {
		return errors.New("produto não existe, cadastre primeiro")
	}
	products[index].Quantity += quantity
	return nil
}

func RemoveStock(products []Product, code int, quantity int) error {
	index, found := FindProductByCode(products, code)
	if!found {
		return errors.New("produto não existe")
	}
	if quantity > products[index].Quantity {
		return fmt.Errorf("estoque insuficiente, você só tem %d", products[index].Quantity)
	}
	products[index].Quantity -= quantity
	return nil
}

func CalculateTotalStock(products []Product) float64 {
	var total float64
	for _, p := range products {
		total += float64(p.Quantity) * p.Price
	}
	return total
}