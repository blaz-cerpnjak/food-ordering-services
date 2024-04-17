package Logic

import (
	"API_GatewayWeb/DataStructures"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (c *Controller) GenerateApiToken(userId primitive.ObjectID, expirationTime time.Time) (tokenString string, err error) {

	// First should also check if the user exists in the database
	// and has paid for the service or has the right to access the service, etc.
	// We should also save api tokens in key-value store like Redis so that we can
	// invalidate the token if needed.

	tk := &DataStructures.ApiKey{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err = token.SignedString(c.jwtSecret)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
