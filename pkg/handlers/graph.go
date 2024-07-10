package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)


func (handler *Handler) getGraphs(context *gin.Context) {
	graphs, err := handler.service.Graph.GetAll()
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, graphs)
}

func (handler *Handler) createGraphs(context *gin.Context) {
	var inputGraph entities.Graph

	if err := context.BindJSON(&inputGraph); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.service.Graph.Create(inputGraph)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{} {
		"id": id,
	})
}

func (handler *Handler) getGraphById(context *gin.Context) {
	graphIdStr := context.Param("graph_id")

	id, err := strconv.ParseUint(graphIdStr, 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	graph, err := handler.service.Graph.GetById(id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, graph)
}

func (handler *Handler) updateGraphById(context *gin.Context) {
	graphIdStr := context.Param("graph_id")

	id, err := strconv.ParseUint(graphIdStr, 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	var inputGraph entities.GraphUpdate

	if err := context.BindJSON(&inputGraph); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = handler.service.Graph.UpdateById(id, inputGraph)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{} {
		"id": id,
	})
}

func (handler *Handler) deleteGraphById(context *gin.Context) {
	graphIdStr := context.Param("graph_id")

	id, err := strconv.ParseUint(graphIdStr, 10, 64)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = handler.service.Graph.DeleteById(id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{} {
		"id": id,
	})
}
