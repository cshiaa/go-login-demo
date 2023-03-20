package main


import (
	"github.com/gin-gonic/gin"

	"github.com/cshiaa/go-login-demo/controller"
	"github.com/cshiaa/go-login-demo/models"
	"github.com/cshiaa/go-login-demo/middlewares"
)

func main() {

	models.ConnectDataBase()

	router := gin.Default()
	router.Use(middlewares.Cors())
	public := router.Group("/api")
	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controller.CurrentUser)
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8089")
}