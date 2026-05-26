package controller

import (
	"go-api/model"
	usecase "go-api/useCase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//UseCase
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		ProductUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID: 1,
			Name: "Batata Frita",
			Price: 20,
		},
	}
	ctx.JSON(http.StatusOK, products)
}