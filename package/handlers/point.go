package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/package/entities"
)

func (handler *Handler) getPoints(context *gin.Context) {
	points, err := handler.service.Point.GetAll()
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, points)
}

func (handler *Handler) createPoints(context *gin.Context) {
	var inputPoint entities.Point

	if err := context.BindJSON(&inputPoint); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.service.Point.Create(inputPoint)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{} {
		"id": id,
	})
}

func (handler *Handler) getPointById(context *gin.Context) {
	pointIdStr := context.Param("point_id")

	id, err := strconv.ParseUint(pointIdStr, 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	point, err := handler.service.Point.GetById(id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, point)
}

func (handler *Handler) updatePointById(context *gin.Context) {
	var inputPoint entities.Point

	if err := context.BindJSON(&inputPoint); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err := handler.service.Point.UpdateById(inputPoint)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{} {
		"id": inputPoint.ID,
	})
}

func (handler *Handler) deletePointById(context *gin.Context) {
	pointIdStr := context.Param("point_id")

	id, err := strconv.ParseUint(pointIdStr, 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = handler.service.Point.DeleteById(id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, id)
}
