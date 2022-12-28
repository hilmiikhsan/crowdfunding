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
		errors := helper.ErrorFormatValidation(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed Register User", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.UserFormat(userData, "tokentokentokentokentoken")
	response := helper.APIResponse("Success Register User", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	input := user.LoginInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorFormatValidation(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed Login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userData, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed Login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.UserFormat(userData, "tokentokentokentokentoken")
	response := helper.APIResponse("Success Login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	input := user.CheckEmailInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorFormatValidation(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvaliable, err := h.userService.CheckEmailAvailability(input)
	if err != nil {
		errorMessage := gin.H{"errors": http.StatusInternalServerError}
		response := helper.APIResponse("Email Checking Failed", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := gin.H{
		"is_avaliable": isEmailAvaliable,
	}

	message := "Email has been registered"

	if isEmailAvaliable {
		message = "Email is available"
	}

	response := helper.APIResponse(message, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
