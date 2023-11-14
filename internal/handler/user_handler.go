package handler

import (
	"go-with-kafka/internal/model"
	"go-with-kafka/internal/repository"
	"go-with-kafka/internal/service"
	"go-with-kafka/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type userHandler struct {
	service  service.UserService
	validate *validator.Validate
}

func InitUserHandler(db *gorm.DB, e *gin.Engine, validator *validator.Validate) {

	userRepo := repository.NewUserRepositoryDB(db)
	service := service.NewUserService(userRepo)
	handler := userHandler{service, validator}

	v1 := e.Group("/v1")
	v1.POST("/users", handler.insert)
}

func (h userHandler) insert(c *gin.Context) {

	request := model.DefaultPayload[model.UserRequest]{}
	if err := c.Bind(&request); err != nil {
		util.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}

	if err := h.validate.Struct(request); err != nil {
		util.ErrorHandler(c, http.StatusUnprocessableEntity, err)
		return
	}

	response, err := h.service.Insert(request)
	if err != nil {
		util.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
