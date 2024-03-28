package Logic

import (
	"API_GatewayWeb/DataStructures"
	"bytes"
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

func (c *Controller) CreateUser(ctx context.Context, user DataStructures.User) (createdUser DataStructures.User, err error) {
	url := fmt.Sprintf("%s/users", getEnv("USERS_API", "http://localhost:8081/api/v1"))

	body, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&createdUser)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (c *Controller) UpdateUser(ctx context.Context, user DataStructures.User) (updatedUser DataStructures.User, err error) {
	url := fmt.Sprintf("%s/users/%s", getEnv("USERS_API", "http://localhost:8081/api/v1"), user.Id.Hex())

	body, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&updatedUser)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (c *Controller) DeleteUser(ctx context.Context, id primitive.ObjectID) (err error) {
	url := fmt.Sprintf("%s/users/%s", getEnv("USERS_API", "http://localhost:8081/api/v1"), id.Hex())

	request, err := http.NewRequest("DELETE", url, nil)
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

	return
}
