package main

import (
	"API_GatewayMobile/HTTP_API"
	"API_GatewayMobile/Logic"
	"API_GatewayMobile/gRPC_Client"
	"fmt"
	"os"
	"os/signal"
)

func main() {

	grpcClient := gRPC_Client.NewController()
	err := grpcClient.EstablishConnection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("gRPC connection established")

	logicController, err := Logic.New(grpcClient)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	httpAPI := HTTP_API.New(logicController)
	httpAPI.Start()

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, os.Interrupt)

	<-quit

	httpAPI.Stop()
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
