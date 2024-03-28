package Logic

import (
	"API_GatewayMobile/gRPC_Client"
	"net/http"
	"os"
)

type Controller struct {
	httpClient *http.Client
	grpc       *gRPC_Client.Controller
}

func New(grpc *gRPC_Client.Controller) (*Controller, error) {
	httpClient := &http.Client{}

	return &Controller{
		httpClient: httpClient,
		grpc:       grpc,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
