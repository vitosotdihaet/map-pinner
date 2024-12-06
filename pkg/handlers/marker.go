package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)

func (handler *Handler) getMarkers(context *gin.Context) {
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

	regionId, err := strconv.ParseUint(context.Query("region_id"), 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	markers, err := handler.service.Marker.GetAll(regionId, user.ID)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, markers)
}

func (handler *Handler) createMarker(context *gin.Context) {
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

	entityType, err := entities.TypeFromString(context.Param("type"))
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	regionId, err := strconv.ParseUint(context.Query("region_id"), 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	var inputMarker entities.Marker
	switch entityType {
	case entities.PointType:
		inputMarker = &entities.Point{}
	case entities.PolygonType:
		inputMarker = &entities.Polygon{}
	case entities.LineType:
		inputMarker = &entities.Line{}
	}

	if err := context.BindJSON(&inputMarker); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.service.Marker.Create(regionId, user.ID, inputMarker)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) getMarkerById(context *gin.Context) {
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

	entityType, err := entities.TypeFromString(context.Param("type"))
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	marker, err := handler.service.Marker.GetById(entityType, id, user.ID)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, marker)
}

func (handler *Handler) updateMarkerById(context *gin.Context) {
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

	entityType, err := entities.TypeFromString(context.Param("type"))
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	var inputMarkerUpdate entities.Marker
	switch entityType {
	case entities.PointType:
		inputMarkerUpdate = &entities.PointUpdate{}
	case entities.PolygonType:
		inputMarkerUpdate = &entities.PolygonUpdate{}
	case entities.LineType:
		inputMarkerUpdate = &entities.LineUpdate{}
	}

	if err := context.BindJSON(&inputMarkerUpdate); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = handler.service.Marker.UpdateById(id, inputMarkerUpdate, user.ID)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) deleteMarkerById(context *gin.Context) {
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

	entityType, err := entities.TypeFromString(context.Param("type"))
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = handler.service.Marker.DeleteById(entityType, id, user.ID)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
