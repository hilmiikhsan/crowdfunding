package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{
		userService,
	}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	input := user.RegisterUserInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorFormatValidation(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed Register User", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userData, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Failed Register User", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.UserFormat(userData, "tokentokentokentokentoken")

	response := helper.APIResponse("Success Register User", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
