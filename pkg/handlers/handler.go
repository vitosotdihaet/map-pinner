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

func (handler *Handler) InitEndpoints() *gin.Engine {
	router := gin.New()

	router.Use(requestLogger())
	router.Use(responseLogger())

	router.Static("/static", "./web")
	router.StaticFile("/", "./web/auth.html")

	users := router.Group("/users")
	users.GET("/", handler.getUsers)
	users.POST("/", handler.createUser)
	users.GET("/bynamepassword", handler.getUserByNamePassword)
	users.GET("/current-user", handler.JWTAuthMiddleware, handler.getAuthenticatedUser)

	api := router.Group("/api", handler.JWTAuthMiddleware)
	markers := api.Group("/markers")
	markers.POST("/:type", handler.createMarker)
	markers.GET("/", handler.getMarkers)
	markers.GET("/:type/:id", handler.getMarkerById)
	markers.PUT("/:type/:id", handler.updateMarkerById)
	markers.DELETE("/:type/:id", handler.deleteMarkerById)

	groups := api.Group("/groups")
	groups.POST("/", handler.createGroup)
	groups.GET("/:id", handler.getGroupById)
	groups.GET("/", handler.getGroups)
	groups.POST("/:id/:username/:role_id", handler.addUserToGroup)

	regions := api.Group("/regions")
	regions.GET("/", handler.getRegions)
	regions.POST("/", handler.createRegion)

	roles := api.Group("/roles")
	roles.GET("/", handler.getRoles)
	roles.GET("/is-owner", handler.isOwner)

	return router
}
