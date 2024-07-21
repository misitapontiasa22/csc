package service

import (
	"csc/model"
	"csc/repository"
	"net/http"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func (ps *ProductService) InsertProduct() model.ProductResponse {

	product := ps.Repo.Insert()
	return model.ProductResponse{
		Code:    http.StatusOK,
		Message: "success insert product",
		Product: product,
	}
}
