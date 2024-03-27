package HTTP_API

import (
	"API_GatewayWeb/DataStructures"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Controller) getAllRestaurants(ctx *gin.Context) {
	restaurants, err := a.logic.GetAllRestaurants(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, restaurants)
}

func (a *Controller) getRestaurantById(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	restaurant, err := a.logic.GetRestaurantById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, restaurant)
}

func (a *Controller) createRestaurant(ctx *gin.Context) {
	var restaurant DataStructures.Restaurant
	err := ctx.BindJSON(&restaurant)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdRestaurant, err := a.logic.CreateRestaurant(ctx.Request.Context(), restaurant)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, createdRestaurant)
}

func (a *Controller) updateRestaurant(ctx *gin.Context) {
	var restaurant DataStructures.Restaurant
	err := ctx.BindJSON(&restaurant)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updated, err := a.logic.UpdateRestaurant(ctx.Request.Context(), restaurant)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, updated)
}

func (a *Controller) deleteRestaurant(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = a.logic.DeleteRestaurant(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{})
}
