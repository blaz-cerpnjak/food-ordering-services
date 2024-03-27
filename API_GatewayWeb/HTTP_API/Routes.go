package HTTP_API

import "github.com/gin-gonic/gin"

func (a *Controller) registerRoutes(engine *gin.Engine) {

	api := engine.Group("/api/v1")
	api.GET("/", a.Ping)

	a.registerUserRoutes(api.Group("/users"))
	a.registerRestaurantRoutes(api.Group("/restaurants"))
	a.registerProductRoutes(api.Group("/products"))
}

func (a *Controller) registerUserRoutes(api *gin.RouterGroup) {
	api.GET("/", a.getAllUsers)
	api.POST("/", a.createUser)
	api.GET("/:id", a.getUserById)
	api.PUT("/:id", a.updateUser)
	api.DELETE("/:id", a.deleteUser)
}

func (a *Controller) registerRestaurantRoutes(api *gin.RouterGroup) {
	api.GET("/", a.getAllRestaurants)
	api.POST("/", a.createRestaurant)
	api.GET("/:id", a.getRestaurantById)
	api.PUT("/:id", a.updateRestaurant)
	api.DELETE("/:id", a.deleteRestaurant)
}

func (a *Controller) registerProductRoutes(api *gin.RouterGroup) {
	api.GET("/restaurant/:id", a.getAllProductsByRestaurantId)
	api.POST("/", a.createProduct)
	api.GET("/:id", a.getProductById)
	api.PUT("/:id", a.updateProduct)
	api.DELETE("/:id", a.deleteProduct)
}
