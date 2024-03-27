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

func (c *Controller) GetAllRestaurants(ctx context.Context) (restaurants []DataStructures.Restaurant, err error) {
	url := fmt.Sprintf("%s/restaurant", getEnv("RESTAURANTS_API", "http://localhost:8082"))

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

	err = json.NewDecoder(response.Body).Decode(&restaurants)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return
}

func (c *Controller) GetRestaurantById(ctx context.Context, id primitive.ObjectID) (restaurant DataStructures.Restaurant, err error) {
	url := fmt.Sprintf("%s/restaurant/%s", getEnv("RESTAURANTS_API", "http://localhost:8082"), id.Hex())

	fmt.Println(url)

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

	err = json.NewDecoder(response.Body).Decode(&restaurant)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (c *Controller) CreateRestaurant(ctx context.Context, restaurant DataStructures.Restaurant) (createdRestaurant DataStructures.Restaurant, err error) {
	url := fmt.Sprintf("%s/restaurant", getEnv("RESTAURANTS_API", "http://localhost:8082"))

	body, err := json.Marshal(restaurant)
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

	err = json.NewDecoder(response.Body).Decode(&createdRestaurant)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (c *Controller) UpdateRestaurant(ctx context.Context, restaurant DataStructures.Restaurant) (updatedRestaurant DataStructures.Restaurant, err error) {
	url := fmt.Sprintf("%s/restaurant/%s", getEnv("RESTAURANTS_API", "http://localhost:8082"), restaurant.Id.Hex())

	body, err := json.Marshal(restaurant)
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

	err = json.NewDecoder(response.Body).Decode(&updatedRestaurant)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (c *Controller) DeleteRestaurant(ctx context.Context, id primitive.ObjectID) (err error) {
	url := fmt.Sprintf("%s/restaurant/%s", getEnv("RESTAURANTS_API", "http://localhost:8082"), id.Hex())

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
