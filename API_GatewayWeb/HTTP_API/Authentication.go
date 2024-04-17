package HTTP_API

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (a *Controller) generateApiToken(c *gin.Context) {

	id, err := primitive.ObjectIDFromHex(c.Param("userId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	expirationTime := time.Now().AddDate(0, 0, 30)

	tokenString, err := a.logic.GenerateApiToken(id, expirationTime)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"api_token": tokenString})
}
