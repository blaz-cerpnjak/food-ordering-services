package HTTP_API

import (
	"ExceptionLoggingAPI/DataStructures"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (a *Controller) createExceptionLog(c *gin.Context) {

	exceptionLog := DataStructures.ExceptionLog{}
	err := c.BindJSON(&exceptionLog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exceptionLog.Timestamp = time.Now()

	err = a.logic.CreateExceptionLog(c.Request.Context(), exceptionLog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exceptionLog)
}

func (a *Controller) getExceptionLogs(c *gin.Context) {

	exceptionLogs, err := a.logic.GetExceptionLogs(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": exceptionLogs})
}
