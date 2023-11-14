package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type defaultHandler struct {
}

func InitDefaultHandler(e *gin.Engine) {
	handler := defaultHandler{}

	e.GET("/health-check", handler.healthCheck)
}

func (h defaultHandler) healthCheck(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
