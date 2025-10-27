package SetupRoutes

import (
	"authentication/controllers"
	"authentication/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/signup", controllers.Signup())
	r.POST("/signin", controllers.Login())

	protected := r.Group("/")
	protected.Use(middlewares.Authenticate())
	{
		protected.GET("/users", controllers.GetUsers())
		protected.GET("/users/:id", controllers.GetUser())
	}
}