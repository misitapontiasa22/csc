package main

import (
	"csc/controller"
	"csc/repository"
	"csc/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := repository.ProductRepository{}
	service := service.ProductService{
		Repo: repo,
	}
	controller := controller.ProductController{
		Service: service,
	}

	router.POST("/product/insert", controller.Insert)
	router.Run("localhost:8080")
}
