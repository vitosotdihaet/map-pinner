package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

func (handler *Handler) getRoles(context *gin.Context) {
	roles, err := handler.service.Role.GetAll()
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, roles)
}

func (handler *Handler) isOwner(context *gin.Context) {
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

	isOwner, err := handler.service.HasAtLeastSystemRole(user.ID, "owner")
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	if isOwner {
		context.Status(http.StatusOK)
	} else {
		newErrorResponse(context, http.StatusUnauthorized, "User is not an owner")
	}
}

func (handler *Handler) getCurretRole(context *gin.Context) {
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

	groupId, err := strconv.ParseUint(context.Param("group_id"), 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	roleId, err := handler.service.GetRoleID(user.ID, groupId)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"role_id": roleId,
	})
}
