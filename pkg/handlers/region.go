package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

func (handler *Handler) getRegions(context *gin.Context) {
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

	groupId, err := strconv.ParseUint(context.Query("group_id"), 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	regions, err := handler.service.Region.GetAll(groupId, user.ID)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, regions)
}

func (handler *Handler) createRegion(context *gin.Context) {
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

	groupId, err := strconv.ParseUint(context.Query("group_id"), 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	var inputRegion entities.Region
	if err := context.BindJSON(&inputRegion); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.service.Region.Create(inputRegion, groupId, user.ID)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
