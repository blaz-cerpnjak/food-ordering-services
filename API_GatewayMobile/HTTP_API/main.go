package HTTP_API

import (
	"API_GatewayMobile/Logic"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type Controller struct {
	logic *Logic.Controller
	done  chan bool
}

func New(logic *Logic.Controller) *Controller {
	return &Controller{
		logic: logic,
	}
}

func (a *Controller) Start() {

	engine := gin.Default()
	a.registerRoutes(engine)

	a.done = make(chan bool, 0)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      engine,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {

		<-a.done

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println(err.Error())
		}

	}()

	go func() {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err.Error())
		}
		os.Exit(1)

	}()

}

func (a *Controller) Stop() {
	a.done <- true
}

func (a *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
