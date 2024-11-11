package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

func (handler *Handler) getLines(context *gin.Context) {
	lines, err := handler.service.Line.GetAll()
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, lines)
}

func (handler *Handler) createLines(context *gin.Context) {
	var inputLine entities.Line

	if err := context.BindJSON(&inputLine); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.service.Line.Create(inputLine)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) getLineById(context *gin.Context) {
	lineIdStr := context.Param("line_id")

	id, err := strconv.ParseUint(lineIdStr, 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	line, err := handler.service.Line.GetById(id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, line)
}

func (handler *Handler) updateLineById(context *gin.Context) {
	lineIdStr := context.Param("line_id")

	id, err := strconv.ParseUint(lineIdStr, 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	var inputLine entities.LineUpdate

	if err := context.BindJSON(&inputLine); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = handler.service.Line.UpdateById(id, inputLine)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) deleteLineById(context *gin.Context) {
	lineIdStr := context.Param("line_id")

	id, err := strconv.ParseUint(lineIdStr, 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = handler.service.Line.DeleteById(id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
