package Router

import (
	"github.com/gin-gonic/gin"
	"proyecto.com/Api"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/api/user")
	{
		userRouter.POST("/sign-in", Api.SignIn)
		userRouter.POST("/sign-up", Api.SignUp)
	}
}
