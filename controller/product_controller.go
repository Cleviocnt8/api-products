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
	
	products, err := p.ProductUseCase.GetProducts()
	if err != nill{
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}