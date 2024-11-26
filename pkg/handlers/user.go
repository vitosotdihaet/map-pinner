package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
	"github.com/vitosotdihaet/map-pinner/pkg/middleware"
	"github.com/vitosotdihaet/map-pinner/pkg/misc"
)

func (handler *Handler) GetUsers(context *gin.Context) {
	users, err := handler.service.User.GetAll()
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, users)
}

func (handler *Handler) CreateUser(context *gin.Context) {
	var inputPassword entities.Password
	var inputUser entities.User

	inputPassword.Value = context.Query("password")
	inputUser.Name = context.Query("username")

	// TODO: min length is 8
	if len(inputPassword.Value) < 1 || len(inputUser.Name) < 1 {
		newErrorResponse(context, http.StatusBadRequest, "Input data is too small")
		return
	}

	var hashedPassword = entities.HashedPassword{Value: misc.Hash(inputPassword.Value)}

	id, err := handler.service.User.Create(inputUser, hashedPassword)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) GetUserByNamePassword(context *gin.Context) {
	var inputPassword entities.Password
	var inputUser entities.User

	inputPassword.Value = context.Query("password")
	inputUser.Name = context.Query("username")

	// TODO: min length is 8
	if len(inputPassword.Value) < 1 || len(inputUser.Name) < 1 {
		newErrorResponse(context, http.StatusBadRequest, "Input data is too small")
		return
	}

	var hashedPassword = entities.HashedPassword{Value: misc.Hash(inputPassword.Value)}

	user, err := handler.service.User.GetByNamePassword(inputUser, hashedPassword)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	if user == nil {
		newErrorResponse(context, http.StatusForbidden, "User not found")
		return
	}

	claims := &entities.UserClaim{
		ID:   user.ID,
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	logrus.Tracef("CLAIM: %v", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JWTKey)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, "Failed to create a JWT token")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
