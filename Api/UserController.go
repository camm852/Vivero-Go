package Api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proyecto.com/Application/Services"
	dto "proyecto.com/Domain/Dto"
	"proyecto.com/Domain/Entities"
)

func SignUp(context *gin.Context) {
	var user Entities.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var isCreated = Services.RegisterUser(&user)

	if !isCreated {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Cannot created user"})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"msg": "User created succesfully"})
}

func SignIn(context *gin.Context) {
	var user dto.UserDto
	if err := context.ShouldBindJSON(&user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userLogin, isLogin := Services.LoginUser(&user)

	if !isLogin {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Cannot login user"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"msg": userLogin})
}
