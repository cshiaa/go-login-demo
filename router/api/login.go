package api

import (
	"github.com/gin-gonic/gin"
	"github.com/cshiaa/go-login-demo/controller"

)

func Routers(e *gin.Engine) {

	public := e.Group("/api")
	{
		public.POST("/register", controller.Register)
		public.POST("/login", controller.Login)
	}
}