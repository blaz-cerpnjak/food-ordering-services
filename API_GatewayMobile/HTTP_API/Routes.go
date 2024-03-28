package HTTP_API

import "github.com/gin-gonic/gin"

func (a *Controller) registerRoutes(engine *gin.Engine) {

	api := engine.Group("/api/v1")
	api.GET("/", a.Ping)

	a.registerUserRoutes(api.Group("/users"))
	a.registerRestaurantRoutes(api.Group("/restaurants"))
	a.registerProductRoutes(api.Group("/products"))
	a.registerOrderRoutes(api.Group("/orders"))
}

func (a *Controller) registerUserRoutes(api *gin.RouterGroup) {
	api.GET("/", a.getAllUsers)
	api.GET("/:id", a.getUserById)
}

func (a *Controller) registerRestaurantRoutes(api *gin.RouterGroup) {
	api.GET("/", a.getAllRestaurants)
	api.GET("/:id", a.getRestaurantById)
}

func (a *Controller) registerProductRoutes(api *gin.RouterGroup) {
	api.GET("/restaurant/:id", a.getAllProductsByRestaurantId)
	api.GET("/:id", a.getProductById)
}

func (a *Controller) registerOrderRoutes(api *gin.RouterGroup) {
	api.POST("/", a.createOrder)
	api.GET("/restaurant/:id", a.getOrdersBySellerId)
	api.PUT("/:id", a.updateOrder)
	api.DELETE("/:id", a.deleteOrder)
}
