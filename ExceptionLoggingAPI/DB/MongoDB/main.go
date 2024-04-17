package MongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type MongoDB struct {
	client   *mongo.Client
	database string
}

func New(database string) *MongoDB {
	return &MongoDB{
		database: database,
	}
}

func (db *MongoDB) Init(ctx context.Context) (err error) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	connectionURI := getEnv("MONGO_DB_URI", "mongodb://localhost:27017")

	db.client, err = mongo.NewClient(options.Client().ApplyURI(connectionURI).SetServerAPIOptions(serverAPI))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = db.client.Connect(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (db *MongoDB) Close(ctx context.Context) (err error) {

	ctx, cancelFunc := context.WithTimeout(ctx, 3*time.Second)
	defer cancelFunc()

	err = db.client.Disconnect(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
