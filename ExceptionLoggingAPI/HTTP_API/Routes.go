package HTTP_API

import "github.com/gin-gonic/gin"

func (a *Controller) registerRoutes(engine *gin.Engine) {

	api := engine.Group("/api/v1")
	api.GET("/", a.Ping)

	a.registerExceptionsRoute(api.Group("/exceptions"))
}

func (a *Controller) registerExceptionsRoute(api *gin.RouterGroup) {
	api.GET("/", a.getExceptionLogs)
	api.POST("/", a.createExceptionLog)
}
