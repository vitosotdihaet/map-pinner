package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vitosotdihaet/map-pinner/pkg/services"
)

func requestLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()

		context.Next()

		latency := time.Since(now)

		logrus.Infof("%s %s %s %s\n",
			context.Request.Method,
			context.Request.RequestURI,
			context.Request.Proto,
			latency,
		)
	}
}

func responseLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		context.Next()

		logrus.Infof("%d %s %s\n",
			context.Writer.Status(),
			context.Request.Method,
			context.Request.RequestURI,
		)
	}
}

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

	router.Use(requestLogger())
	router.Use(responseLogger())

	router.Static("/static", "./web")
	router.StaticFile("/", "./web/index.html")

	api := router.Group("/api")

	{
		markers := api.Group("/markers")
		{
			points := markers.Group("/points")
			{
				points.POST("/", handler.createPoints)
				points.GET("/", handler.getPoints)
				// points.DELETE("/")

				points.GET("/:point_id", handler.getPointById)
				points.PUT("/:point_id", handler.updatePointById)
				points.DELETE("/:point_id", handler.deletePointById)
			}

			polygons := markers.Group("/polygons")
			{
				polygons.POST("/", handler.createPolygons)
				polygons.GET("/", handler.getPolygons)
				// polygons.DELETE("/")

				polygons.GET("/:polygon_id", handler.getPolygonById)
				polygons.PUT("/:polygon_id", handler.updatePolygonById)
				polygons.DELETE("/:polygon_id", handler.deletePolygonById)
			}

			lines := markers.Group("/lines")
			{
				lines.POST("/", handler.createLines)
				lines.GET("/", handler.getLines)
				// lines.DELETE("/")

				lines.GET("/:line_id", handler.getLineById)
				lines.PUT("/:line_id", handler.updateLineById)
				lines.DELETE("/:line_id", handler.deleteLineById)

				handler.pointOperations(lines)
			}
		}
	}

	return router
}
