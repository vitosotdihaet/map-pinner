package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) getPolygons(context *gin.Context) {
	polygons, err := handler.service.Polygon.GetAll()
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, polygons)
}

func (handler *Handler) createPolygons(context *gin.Context) {
	// var inputPolygon entities.Polygon

	// if err := context.BindJSON(&inputPolygon); err != nil {
	// 	newErrorResponse(context, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// id, err := handler.service.Polygon.Create(inputPolygon)
	// if err != nil {
	// 	newErrorResponse(context, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// context.JSON(http.StatusOK, map[string]interface{} {
	// 	"id": id,
	// })
}

func (handler *Handler) getPolygonById(context *gin.Context) {
	// polygonIdStr := context.Param("polygon_id")

	// id, err := strconv.ParseUint(polygonIdStr, 10, 64)
	// if err != nil {
	// 	newErrorResponse(context, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// polygon, err := handler.service.Polygon.GetById(id)
	// if err != nil {
	// 	newErrorResponse(context, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// context.JSON(http.StatusOK, polygon)
}

func (handler *Handler) updatePolygonById(context *gin.Context) {
	// var inputPolygon entities.Polygon

	// if err := context.BindJSON(&inputPolygon); err != nil {
	// 	newErrorResponse(context, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// err := handler.service.Polygon.UpdateById(inputPolygon)
	// if err != nil {
	// 	newErrorResponse(context, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// context.JSON(http.StatusOK, map[string]interface{} {
	// 	"id": inputPolygon.ID,
	// })
}

func (handler *Handler) deletePolygonById(context *gin.Context) {
	// polygonIdStr := context.Param("polygon_id")

	// id, err := strconv.ParseUint(polygonIdStr, 10, 64)
	// if err != nil {
	// 	newErrorResponse(context, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// err = handler.service.Polygon.DeleteById(id)
	// if err != nil {
	// 	newErrorResponse(context, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// context.JSON(http.StatusOK, id)
}
