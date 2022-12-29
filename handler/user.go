package handler

import (
	"crowdfunding/auth"
	"crowdfunding/helper"
	"crowdfunding/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{
		userService,
		authService,
	}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	input := user.RegisterUserInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorFormatValidation(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register User Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userData, err := h.userService.RegisterUser(input)
	if err != nil {
		errors := helper.ErrorFormatValidation(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register User Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(userData.ID)
	if err != nil {
		response := helper.APIResponse("Register User Failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.UserFormat(userData, token)
	response := helper.APIResponse("Success Register User", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	input := user.LoginInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorFormatValidation(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userData, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(userData.ID)
	if err != nil {
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.UserFormat(userData, token)
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

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Upload Avatar Failed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 4

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Upload Avatar Failed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Upload Avatar Failed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}

	response := helper.APIResponse("Success Upload Avatar", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
