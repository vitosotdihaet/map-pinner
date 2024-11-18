package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
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

	context.JSON(http.StatusOK, user)
}
