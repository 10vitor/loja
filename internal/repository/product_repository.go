package repository

import (
	"errors"
	"loja-go/internal/model"
)

type ProductRepository struct {
	products []model.Produvt
	ultimoID uint
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: []model.Produvt{},
		ultimoID: 0,
	}
}

func (r *ProductRepository) Create(product model.Produvt) model.Produvt {
	r.ultimoID++
	product.ID = r.ultimoID
	r.products = append(r.products, product)
	return product
}

func (r *ProductRepository) FindAll() []model.Produto {
	return r.produtos
}

func (r *ProductRepository) FindByID(id uint) (*model.Produto, error) {
	for _, p := range r.produtos {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("produto não encontrado")
}