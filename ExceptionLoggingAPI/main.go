package main

import (
	"ExceptionLoggingAPI/DB/MongoDB"
	"ExceptionLoggingAPI/HTTP_API"
	"ExceptionLoggingAPI/Logic"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	mongoDB := MongoDB.New(
		getEnv("MONGO_DB_NAME", "exception_logs"),
	)

	err := mongoDB.Init(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		if err = mongoDB.Close(context.Background()); err != nil {
			panic(err)
		}
	}()

	logicController, err := Logic.New(mongoDB)
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
