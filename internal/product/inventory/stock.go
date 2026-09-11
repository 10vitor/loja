package inventory

import "loja/internal/product"

type Inventory struct {
	Products []product.Product
}

func (i *Inventory) AddProduct(p product.Product) {
    i.Products = append(i.Products, p)
}

var stockInventory = Inventory{
    Products: []product.Product{
        {
            Name:     "Camisa Preta ",
            Price:    79.90,
            Code:     101,
            Quantity: 20,
        },
        {
            Name:     "Shorts Jeans",
            Price:    119.90,
            Code:     102,
            Quantity: 15,
        },
    },
}