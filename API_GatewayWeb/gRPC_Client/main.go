package gRPC_Client

import (
	"API_GatewayWeb/DataStructures/com.blazc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	conn, err := grpc.Dial("localhost:9000", opts...)
	if err != nil {
		return err
	}

	c.Client = com_blazc.NewOrderServiceClient(conn)
	return
}
