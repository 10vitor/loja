package inventory

import (
    "encoding/json"
    "loja/internal/product"
    "os"
	"fmt"
)
type Inventory struct {
	Products []product.Product `json:"products"`
}

func (i *Inventory) AddProduct(p product.Product) {
    i.Products = append(i.Products, p)
}

func (i *Inventory) Save() error {

	data, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		return err
	}

    return os.WriteFile("inventory.json", data, 0644)
}

func (i *Inventory) Load() error {
	data, err := os.ReadFile("estoque.json")

	if err != nil {
		return err
	}

	fmt.Println("JSON encontrado!")
	fmt.Println(string(data))

	err = json.Unmarshal(data, i)

	if err != nil {
		return err
	}

	fmt.Println("Produtos carregados:", len(i.Products))

	return nil
}