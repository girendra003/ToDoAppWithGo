package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: Hardcoded login check (in real life, query DB)
		var creds struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		if creds.Email != "g@gmail.com" || creds.Password != "password123" {  //----setep for backend check api hit
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate JWT token
		claims := jwt.MapClaims{
			"user_email": creds.Email,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

//first take credentials like loging or pas for generate token
//generate claims using     jwt.MapClaims
//create a new token with   jwt.NewWithClaims
//sign the token with a secret key and return it in the response token.SignedString
// Note: In a real application, you would validate the user's credentials against a database