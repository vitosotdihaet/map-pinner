package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/pkg/services"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) pointOperations(group *gin.RouterGroup) {
	points := group.Group("/points")
	{
		points.POST("/", handler.createPoints)
		points.GET("/", handler.getPoints)
		// points.DELETE("/")

		points.GET("/:point_id", handler.getPointById)
		points.PUT("/:point_id", handler.updatePointById)
		points.DELETE("/:point_id", handler.deletePointById)
	}
}

func (handler *Handler) InitEndpoints() *gin.Engine {
	router := gin.New()

	router.Static("/static", "./web")

	api := router.Group("/api")
	{
		handler.pointOperations(api)

		polygons := api.Group("/polygons")
		{
			polygons.POST("/", handler.createPolygons)
			polygons.GET("/", handler.getPolygons)
			// polygons.DELETE("/")

			polygons.GET("/:polygon_id", handler.getPolygonById)
			polygons.PUT("/:polygon_id", handler.updatePolygonById)
			polygons.DELETE("/:polygon_id", handler.deletePolygonById)

			handler.pointOperations(polygons)
		}

		// graphs := api.Group("/graphs")
		// {
		// 	graphs.POST("/")
		// 	graphs.GET("/")
		// 	// graphs.DELETE("/")

		// 	graphs.GET("/:id")
		// 	graphs.PUT("/:id")
		// 	graphs.DELETE("/:id")

		// 	handler.pointOperations(graphs)
		// }
	}

	return router
}
