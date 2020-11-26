package routers

import (
	"github.com/gin-gonic/gin"
	"mtdn.io/Kagerou/services"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/create", services.Response(services.CreateUser))
		user.GET("/info/:telegram_id", services.Response(services.GetUser))
		user.GET("/verify/:telegram_id", services.Response(services.VerifyUser))
		user.PUT("/update", services.Response(services.UpdateUser))
		user.DELETE("/remove", services.Response(services.RemoveUser))
		user.GET("/reset", services.Response(services.ResetUserVerificationCode))
	}

	return router
}
