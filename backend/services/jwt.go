package services

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	ID    string
	Email string
	jwt.StandardClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func validateToken(signedToken string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, nil
}

func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("BOX_CATALOG_TOKEN")

		log.Println(tokenString)

		if err != nil {
			log.Println(err)
			c.JSON(401, gin.H{"message": "Not authorized"})
			c.Abort()
			return
		}

		if tokenString == "" {
			c.JSON(401, gin.H{"message": "Not authorized"})
			c.Abort()
			return
		}
		claims, err := validateToken(tokenString)
		if err != nil {
			log.Println(err.Error())
			c.JSON(401, gin.H{"message": "Not authorized"})
			c.Abort()
			return
		}
		c.Set("UserID", claims.ID)
		c.Set("Email", claims.Email)
		c.Next()
	}
}

func GenerateJWT(userId string, email string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		ID:    userId,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
