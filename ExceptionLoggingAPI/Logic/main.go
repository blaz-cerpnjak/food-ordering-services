package Logic

import (
	"ExceptionLoggingAPI/DB"
	"os"
)

type Controller struct {
	db DB.DB
}

func New(db DB.DB) (*Controller, error) {
	return &Controller{
		db: db,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
