package DB

import (
	"ExceptionLoggingAPI/DataStructures"
	"context"
)

type DB interface {
	Init(ctx context.Context) (err error)
	Close(ctx context.Context) (err error)

	// Exception Log
	CreateExceptionLog(ctx context.Context, exceptionLog DataStructures.ExceptionLog) (err error)
	GetExceptionLogs(ctx context.Context) (exceptionLogs []DataStructures.ExceptionLog, err error)
}
