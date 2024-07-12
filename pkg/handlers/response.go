package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


type handleError struct {
	Message string `json:"message"`
}

func newErrorResponse(context *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	context.AbortWithStatusJSON(statusCode, handleError{message})
}