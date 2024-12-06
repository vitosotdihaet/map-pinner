package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

var JWTKey []byte

func (handler *Handler) JWTAuthMiddleware(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		context.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims := &entities.UserClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})

	if err != nil || !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		context.Abort()
		return
	}

	_, err = handler.service.User.GetById(claims.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "No such user"})
		context.Abort()
		return
	}

	context.Set("user", entities.User{
		ID:   claims.ID,
		Name: claims.Name,
	})

	context.Set("token", tokenString)

	context.Next()
}
