package handler

import (
	"crowdfunding/campaign"
	"crowdfunding/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{
		service,
	}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaignData, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Get List Campaign", http.StatusOK, "success", campaign.CampaignsFormat(campaignData))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaignByID(c *gin.Context) {
	input := campaign.GetCampaignDetailInput{}

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed Get Detail Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignData, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed Get Detail Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Get Campaign Detail", http.StatusOK, "success", campaign.CampaignDetailFormat(campaignData))
	c.JSON(http.StatusOK, response)
}
