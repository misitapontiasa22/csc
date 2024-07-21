package repository

import (
	"csc/model"
)

type ProductRepository struct {
}

func (pr *ProductRepository) Insert() model.Product {

	product := model.Product{
		Id:    1,
		Name:  "umbrella v1",
		Price: 10000,
		Qty:   2,
	}
	return product
}
