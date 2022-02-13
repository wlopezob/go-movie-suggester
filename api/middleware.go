package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wlopezob/go-movie-suggester/internal/logs"
)

func jwtMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func signToken(tokenKey string, id string) string {
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "wlopezob",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
		"sub":   id,
	}
	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(tokenKey))

	if err != nil {
		return ""
	}
	return t
}

func extractUserIdFromJWT(token, tokenKey string) string {
	token = token[7:]
	logs.Info(token)
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		logs.Error(err.Error())
		return ""
	}
	if !t.Valid {
		logs.Error("Claims is wrong")
		return ""
	}
	claims := t.Claims.(jwt.MapClaims)
	return claims["sub"].(string)
}
