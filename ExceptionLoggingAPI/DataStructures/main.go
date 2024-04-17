package DataStructures

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ExceptionLog struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Message   string             `json:"message" bson:"message"`
	Severity  string             `json:"severity" bson:"severity"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
