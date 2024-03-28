package Logic

import (
	"API_GatewayMobile/DataStructures"
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
