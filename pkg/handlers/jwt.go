package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
	"github.com/vitosotdihaet/map-pinner/pkg/middleware"
)

func (handler *Handler) validateToken(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		newErrorResponse(context, http.StatusUnauthorized, "Missing token")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims := &entities.UserClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return middleware.JWTKey, nil
	})

	if err != nil || !token.Valid {
		newErrorResponse(context, http.StatusUnauthorized, "Invalid token")
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
}
