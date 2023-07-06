package Services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	dto "proyecto.com/Domain/Dto"
	"proyecto.com/Domain/Entities"
	"strconv"
	"time"
)

// Función para encriptar una contraseña
func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

// Función para verificar una contraseña encriptada
func verifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func uintToString(number uint) string {
	str := strconv.FormatUint(uint64(number), 10)
	return str
}

func RegisterUser(user *Entities.User) bool { //Hugo
	var _user Entities.IUserRepository = Entities.User{}
	hashPassword := hashPassword(user.Password)

	if len(hashPassword) == 0 {
		return false
	}

	user.Password = hashPassword

	var userCreated = _user.SignUp(user)

	return userCreated
}

func LoginUser(user *dto.UserDto) (dto.UserLoginDto, bool) { //Hugo
	var _user Entities.IUserRepository = Entities.User{}

	userStored, err := _user.GetUser(user)
	var userLoginDto = dto.UserLoginDto{}

	if err != nil {
		return userLoginDto, false
	}

	if !verifyPassword(userStored.Password, user.Password) {
		return userLoginDto, false
	}

	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	claims := jwt.MapClaims{
		"userId": userStored.ID,
		"expt":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //crear objeto token los los claims y el metodo de firma

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Println("Error al firmar el token:", err)
		return userLoginDto, false
	}
	userLoginDto.Email = userStored.Email
	userLoginDto.Name = userStored.Name
	userLoginDto.Token = tokenString

	return userLoginDto, true
}
