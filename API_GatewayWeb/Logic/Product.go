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

func (c *Controller) CreateProduct(ctx context.Context, product DataStructures.Product) (createdProduct DataStructures.Product, err error) {
	url := fmt.Sprintf("%s/product", getEnv("RESTAURANTS_API", "http://localhost:8082"))

	body, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
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

	err = json.NewDecoder(response.Body).Decode(&createdProduct)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (c *Controller) UpdateProduct(ctx context.Context, product DataStructures.Product) (updatedProduct DataStructures.Product, err error) {
	url := fmt.Sprintf("%s/product", getEnv("RESTAURANTS_API", "http://localhost:8082"))

	body, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
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

	err = json.NewDecoder(response.Body).Decode(&updatedProduct)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (c *Controller) DeleteProduct(ctx context.Context, id primitive.ObjectID) (err error) {
	url := fmt.Sprintf("%s/product/%s", getEnv("RESTAURANTS_API", "http://localhost:8082"), id.Hex())

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
