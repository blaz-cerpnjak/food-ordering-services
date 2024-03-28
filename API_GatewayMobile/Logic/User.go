package Logic

import (
	"API_GatewayMobile/DataStructures"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (c *Controller) GetAllUsers(ctx context.Context) (users []DataStructures.User, err error) {
	url := fmt.Sprintf("%s/users", getEnv("USERS_API", "http://localhost:8081/api/v1"))

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&users)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return
}

func (c *Controller) GetUserById(ctx context.Context, id primitive.ObjectID) (user DataStructures.User, err error) {
	url := fmt.Sprintf("%s/users/%s", getEnv("USERS_API", "http://localhost:8081/api/v1"), id.Hex())

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
