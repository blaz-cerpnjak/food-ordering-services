package HTTP_API

import (
	"API_GatewayWeb/DataStructures"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Controller) createOrder(ctx *gin.Context) {
	var order DataStructures.Order
	err := ctx.BindJSON(&order)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdOrder, err := a.logic.CreateOrder(ctx.Request.Context(), order)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, createdOrder)
}

func (a *Controller) updateOrder(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	confirmation, err := a.logic.UpdateOrder(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, confirmation)
}

func (a *Controller) deleteOrder(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	confirmation, err := a.logic.DeleteOrder(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, confirmation)
}

func (a *Controller) getOrdersBySellerId(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	orders, err := a.logic.GetOrdersBySellerId(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, orders)
}
