package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func point_operations(group *gin.RouterGroup) {
	points := group.Group(":id/points")
	{
		points.GET("/")
		points.POST("/")
		points.DELETE("/")

		points.GET("/:point_id")
		points.PUT("/:point_id")
		points.DELETE("/:point_id")
	}
}

func (h *Handler) InitEndpoints() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		sign_lists := api.Group("/sign-lists")
		{
			sign_lists.GET("/")
			sign_lists.POST("/")
			sign_lists.DELETE("/")

			sign_lists.GET("/:id")
			sign_lists.PUT("/:id")
			sign_lists.DELETE("/:id")

			point_operations(sign_lists)
		}

		polygon_lists := api.Group("/polygon-lists")
		{
			polygon_lists.GET("/")
			polygon_lists.POST("/")
			polygon_lists.DELETE("/")

			polygon_lists.GET("/:id")
			polygon_lists.PUT("/:id")
			polygon_lists.DELETE("/:id")

			point_operations(sign_lists)
		}

		graph_lists := api.Group("/graph-lists")
		{
			graph_lists.GET("/")
			graph_lists.POST("/")
			graph_lists.DELETE("/")

			graph_lists.GET("/:id")
			graph_lists.PUT("/:id")
			graph_lists.DELETE("/:id")

			point_operations(sign_lists)
		}
	}

	return router
}
