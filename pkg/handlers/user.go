package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
	"github.com/vitosotdihaet/map-pinner/pkg/misc"
)

func (handler *Handler) getUsers(context *gin.Context) {
	users, err := handler.service.User.GetAll()
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, users)
}

func (handler *Handler) createUser(context *gin.Context) {
	var inputPassword entities.Password
	var inputUser entities.User

	inputPassword.Value = context.Query("password")
	inputUser.Name = context.Query("username")

	passwordLength := len(inputPassword.Value)
	nameLength := len(inputUser.Name)
	if passwordLength < 8 || nameLength < 8 {
		newErrorResponse(context, http.StatusBadRequest, "Input data is too small")
		return
	} else if passwordLength > 72 || nameLength > 32 {
		newErrorResponse(context, http.StatusBadRequest, "Input data is too big")
		return
	}

	passwordHash, err := misc.Hash(inputPassword.Value)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	var hashedPassword = entities.HashedPassword{Value: passwordHash}

	id, err := handler.service.User.Create(inputUser, hashedPassword)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) getUserByNamePassword(context *gin.Context) {
	var inputPassword entities.Password
	var inputUser entities.User

	inputPassword.Value = context.Query("password")
	inputUser.Name = context.Query("username")

	passwordLength := len(inputPassword.Value)
	nameLength := len(inputUser.Name)
	if passwordLength < 8 || nameLength < 8 {
		newErrorResponse(context, http.StatusBadRequest, "Input data is too small")
		return
	} else if passwordLength > 72 || nameLength > 32 {
		newErrorResponse(context, http.StatusBadRequest, "Input data is too big")
		return
	}

	user, hashedPassword, err := handler.service.User.GetByName(inputUser)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	if user == nil {
		newErrorResponse(context, http.StatusForbidden, "User not found")
		return
	}
	if !misc.VerifyPassword(inputPassword.Value, hashedPassword.Value) {
		newErrorResponse(context, http.StatusForbidden, "Wrong password")
		return
	}

	claims := &entities.UserClaim{
		ID:   user.ID,
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, "Failed to create a JWT token")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": tokenString,
	})
}

func (handler *Handler) getAuthenticatedUser(context *gin.Context) {
	user, exists := context.Get("user")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	tokenString, exists := context.Get("token")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": tokenString,
	})
}
