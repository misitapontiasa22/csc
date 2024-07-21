package controller

import (
	"csc/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Service service.ProductService
}

func (pc *ProductController) Insert(ctx *gin.Context) {

	resp := pc.Service.InsertProduct()
	ctx.JSON(http.StatusOK, resp)
}
