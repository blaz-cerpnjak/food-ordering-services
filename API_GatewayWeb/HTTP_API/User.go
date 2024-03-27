package HTTP_API

import (
	"API_GatewayWeb/DataStructures"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Controller) getAllUsers(ctx *gin.Context) {
	users, err := a.logic.GetAllUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, users)
}

func (a *Controller) getUserById(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.logic.GetUserById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}

func (a *Controller) createUser(ctx *gin.Context) {
	var user DataStructures.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(user)

	createdUser, err := a.logic.CreateUser(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, createdUser)
}

func (a *Controller) updateUser(ctx *gin.Context) {
	var user DataStructures.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := a.logic.UpdateUser(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, updatedUser)
}

func (a *Controller) deleteUser(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = a.logic.DeleteUser(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "User deleted"})
}
