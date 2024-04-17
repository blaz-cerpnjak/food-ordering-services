package MongoDB

import (
	"ExceptionLoggingAPI/DataStructures"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *MongoDB) CreateExceptionLog(ctx context.Context, exceptionLog DataStructures.ExceptionLog) (err error) {

	_, err = db.client.Database(db.database).Collection("ExceptionLogs").InsertOne(ctx, exceptionLog)
	if err != nil {
		return
	}

	return
}

func (db *MongoDB) GetExceptionLogs(ctx context.Context) (exceptionLogs []DataStructures.ExceptionLog, err error) {

	cur, err := db.client.Database(db.database).Collection("ExceptionLogs").Find(ctx, bson.M{})
	if err != nil {
		return
	}

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			return
		}
	}(cur, ctx)

	exceptionLogs = make([]DataStructures.ExceptionLog, 0)

	for cur.Next(ctx) {
		var exceptionLog DataStructures.ExceptionLog
		err = cur.Decode(&exceptionLog)
		if err != nil {
			return
		}

		exceptionLogs = append(exceptionLogs, exceptionLog)
	}

	return
}
