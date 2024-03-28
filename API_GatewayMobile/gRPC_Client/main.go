package gRPC_Client

import (
	"API_GatewayMobile/DataStructures/com.blazc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

type Controller struct {
	Client com_blazc.OrderServiceClient
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) EstablishConnection() (err error) {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(getEnv("ORDERS_GRPC_API", "localhost:9000"), opts...)
	if err != nil {
		return err
	}

	c.Client = com_blazc.NewOrderServiceClient(conn)
	return
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
