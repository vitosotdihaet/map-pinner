package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) getRegions(context *gin.Context) {
	regions, err := handler.service.Region.GetAll()
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, regions)
}
