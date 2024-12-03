package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

func (handler *Handler) getGroups(context *gin.Context) {
	userany, exists := context.Get("user")
	if !exists {
		newErrorResponse(context, http.StatusUnauthorized, "User not found")
		return
	}

	user, ok := userany.(entities.User)
	if !ok {
		newErrorResponse(context, http.StatusInternalServerError, "Could not unpack user")
		return
	}

	groups, err := handler.service.Group.GetAll(user.ID)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, groups)
}

func (handler *Handler) createGroup(context *gin.Context) {
	userany, exists := context.Get("user")
	if !exists {
		newErrorResponse(context, http.StatusUnauthorized, "User not found")
		return
	}

	user, ok := userany.(entities.User)
	if !ok {
		newErrorResponse(context, http.StatusInternalServerError, "Could not unpack user")
		return
	}

	var inputGroup entities.Group
	if err := context.BindJSON(&inputGroup); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.service.Group.Create(inputGroup, user.ID)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) getGroupById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	group, err := handler.service.Group.GetById(id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, group)
}
