package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitosotdihaet/map-pinner/package/services"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) pointOperations(group *gin.RouterGroup) {
	points := group.Group(":id/points")
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

	api := router.Group("/api")
	{
		handler.pointOperations(api)

		// polygon_lists := api.Group("/polygon-lists")
		// {
		// 	polygon_lists.POST("/")
		// 	polygon_lists.GET("/")
		// 	polygon_lists.DELETE("/")

		// 	polygon_lists.GET("/:id")
		// 	polygon_lists.PUT("/:id")
		// 	polygon_lists.DELETE("/:id")

		// 	handler.pointOperations(points_lists)
		// }

		// graph_lists := api.Group("/graph-lists")
		// {
		// 	graph_lists.POST("/")
		// 	graph_lists.GET("/")
		// 	graph_lists.DELETE("/")

		// 	graph_lists.GET("/:id")
		// 	graph_lists.PUT("/:id")
		// 	graph_lists.DELETE("/:id")

		// 	handler.pointOperations(points_lists)
		// }
	}

	return router
}
