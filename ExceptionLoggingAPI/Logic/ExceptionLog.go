package Logic

import (
	"ExceptionLoggingAPI/DataStructures"
	"context"
)

func (c *Controller) CreateExceptionLog(ctx context.Context, exceptionLog DataStructures.ExceptionLog) (err error) {
	return c.db.CreateExceptionLog(ctx, exceptionLog)
}

func (c *Controller) GetExceptionLogs(ctx context.Context) (exceptionLogs []DataStructures.ExceptionLog, err error) {
	return c.db.GetExceptionLogs(ctx)
}
