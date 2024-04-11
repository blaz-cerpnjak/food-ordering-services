package HTTP_API

import (
	"API_GatewayWeb/DataStructures"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Controller) getAllProducts(ctx *gin.Context) {
	products, err := a.logic.GetAllProducts(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, products)
}

func (a *Controller) getAllProductsByRestaurantId(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	restaurant, err := a.logic.GetAllProductsByRestaurantId(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, restaurant)
}

func (a *Controller) getProductById(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product, err := a.logic.GetProductById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, product)
}

func (a *Controller) createProduct(ctx *gin.Context) {
	var product DataStructures.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := a.logic.CreateProduct(ctx.Request.Context(), product)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, createdProduct)
}

func (a *Controller) updateProduct(ctx *gin.Context) {
	var product DataStructures.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updated, err := a.logic.UpdateProduct(ctx.Request.Context(), product)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, updated)
}

func (a *Controller) deleteProduct(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = a.logic.DeleteProduct(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Product deleted"})
}
