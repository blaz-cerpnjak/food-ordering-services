package Logic

import (
	"API_GatewayWeb/gRPC_Client"
	"net/http"
	"os"
)

type Controller struct {
	httpClient *http.Client
	grpc       *gRPC_Client.Controller
	jwtSecret  []byte
}

func New(grpc *gRPC_Client.Controller, jwtSecret string) (*Controller, error) {
	httpClient := &http.Client{}

	return &Controller{
		httpClient: httpClient,
		grpc:       grpc,
		jwtSecret:  []byte(jwtSecret),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
