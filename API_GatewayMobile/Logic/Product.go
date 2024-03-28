package Logic

import (
	"API_GatewayMobile/DataStructures"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (c *Controller) GetAllProductsByRestaurantId(ctx context.Context, id primitive.ObjectID) (products []DataStructures.Product, err error) {
	url := fmt.Sprintf("%s/product/restaurant/%s", getEnv("RESTAURANTS_API", "http://localhost:8082"), id.Hex())

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

	err = json.NewDecoder(response.Body).Decode(&products)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return
}

func (c *Controller) GetProductById(ctx context.Context, id primitive.ObjectID) (product DataStructures.Product, err error) {
	url := fmt.Sprintf("%s/product/%s", getEnv("RESTAURANTS_API", "http://localhost:8082"), id.Hex())

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

	err = json.NewDecoder(response.Body).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
