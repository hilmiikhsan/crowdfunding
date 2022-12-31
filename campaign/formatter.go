package campaign

import (
	"strings"
	"time"
)

type CampaignFormatter struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	BackerCount      int       `json:"backer_count"`
	GoalAmount       int       `json:"goal_amount"`
	CurrentAmount    int       `json:"current_amount"`
	Slug             string    `json:"slug"`
	ImageURL         string    `json:"image_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func CampaignFormat(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.Description = campaign.Description
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageURL = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func CampaignsFormat(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := CampaignFormat(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	UserID           int                      `json:"user_id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	BackerCount      int                      `json:"backer_count"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	Slug             string                   `json:"slug"`
	ImageURL         string                   `json:"image_url"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormtter     `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormtter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func CampaignDetailFormat(campaign Campaign) CampaignDetailFormatter {
	camaignDetailFormatter := CampaignDetailFormatter{}
	camaignDetailFormatter.ID = campaign.ID
	camaignDetailFormatter.UserID = campaign.UserID
	camaignDetailFormatter.Name = campaign.Name
	camaignDetailFormatter.ShortDescription = campaign.ShortDescription
	camaignDetailFormatter.Description = campaign.Description
	camaignDetailFormatter.GoalAmount = campaign.GoalAmount
	camaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	camaignDetailFormatter.ImageURL = ""
	camaignDetailFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		camaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	perks := []string{}

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	camaignDetailFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormtter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName
	camaignDetailFormatter.User = campaignUserFormatter

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	camaignDetailFormatter.Images = images

	return camaignDetailFormatter
}
