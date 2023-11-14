package util

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ErrorHandler(c *gin.Context, code int, err error) {
	c.JSON(code, ErrorResponse{
		Code:    code,
		Status:  "error",
		Message: err.Error(),
	})
}
