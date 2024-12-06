package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) getRoles(context *gin.Context) {
	roles, err := handler.service.Role.GetAll()
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, roles)
}
