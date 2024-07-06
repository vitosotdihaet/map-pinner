package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
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
		sign_lists := api.Group("/sign-lists")
		{
			sign_lists.POST("/", handler.createSignLists)
			sign_lists.GET("/", handler.getSignLists)
			// sign_lists.DELETE("/")

			sign_lists.GET("/:id", handler.getSignListById)
			sign_lists.PUT("/:id", handler.updateSignListById)
			sign_lists.DELETE("/:id", handler.deleteSignListById)

			handler.pointOperations(sign_lists)
		}

		// polygon_lists := api.Group("/polygon-lists")
		// {
		// 	polygon_lists.POST("/")
		// 	polygon_lists.GET("/")
		// 	polygon_lists.DELETE("/")

		// 	polygon_lists.GET("/:id")
		// 	polygon_lists.PUT("/:id")
		// 	polygon_lists.DELETE("/:id")

		// 	handler.pointOperations(sign_lists)
		// }

		// graph_lists := api.Group("/graph-lists")
		// {
		// 	graph_lists.POST("/")
		// 	graph_lists.GET("/")
		// 	graph_lists.DELETE("/")

		// 	graph_lists.GET("/:id")
		// 	graph_lists.PUT("/:id")
		// 	graph_lists.DELETE("/:id")

		// 	handler.pointOperations(sign_lists)
		// }
	}

	return router
}
