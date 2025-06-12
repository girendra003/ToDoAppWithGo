package middleware

import (
	"fmt"
	// "log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWT secret key — normally you store this in .env
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// AuthMiddleware verifies JWT tokens

//----------------steps that follow for auth middleware
//get header
//validate header is empty, hasprefix is not then about
//trim prefix
//jwtparse for final validation
//last return  and call c.Next() to continue
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()

			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and verify the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {//----------------most important step tokenstring,scrtKey
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		// Extract claims
		// and store user email in context for later use
		// claims, _ := token.Claims.(jwt.MapClaims)
		// if email, ok := claims["user_email"].(string); ok {
		// 	log.Printf("Authenticated user: %s", email)
		// 	c.Set("user_email", email) // Store email in context for later use
		// } else {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		// 	c.Abort()
		// 	return
		// }
		c.Next() // token is valid, continue
	}
}

// package middleware

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// // //----------------steps that follow for auth middleware
// var jstsecretKey = []byte(os.Getenv("JWT_SECRET"))

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		// //get header
// 		authHeader := c.GetHeader("Authorization")

// 		// //validate header is empty, hasprefix is not then about
// 		if authHeader == " " || !strings.HasPrefix(authHeader, "Bearer") {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
// 			c.Abort()
// 			return
// 		}
// 		// //trim prefix
// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

// 		// //jwtparse for final validation
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("error here in verificaiton ")
// 			}
// 			return jstsecretKey, nil
// 		})
// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
// 			c.Abort()
// 			return
// 		}

// 		// //last return  and call c.Next() to continue

// 		c.Next() // token is valid, continue

// 	}
// }
