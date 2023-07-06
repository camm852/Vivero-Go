package Middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func AuthRequired(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader != "" {
		bearerToken := strings.Split(authHeader, "Bearer ")

		if len(bearerToken) == 2 {
			token := bearerToken[1]

			var secretKey = []byte(os.Getenv("SECRET_KEY"))

			// Decodificar y verificar el token
			decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				// Verificar el algoritmo de firma
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Algoritmo de firma inválido")
				}

				// Devolver la clave secreta utilizada para firmar el token
				return secretKey, nil
			})
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// Verificar si el token es válido
			if decodedToken.Valid {
				// Obtener los claims (datos) del token decodificado
				claims, ok := decodedToken.Claims.(jwt.MapClaims)
				if !ok {
					fmt.Println("Error al obtener los claims del token")
					c.AbortWithStatus(http.StatusUnauthorized)
				}
				c.Set("userId", claims["userId"]) // pasarle el id a los endpoints
				c.Next()
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
