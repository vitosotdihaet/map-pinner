package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

func (handler *Handler) getGroups(context *gin.Context) {
	groups, err := handler.service.Group.GetAll()
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, groups)
}

func (handler *Handler) createGroup(context *gin.Context) {
	var inputGroup entities.Group

	if err := context.BindJSON(&inputGroup); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.service.Group.Create(inputGroup)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
