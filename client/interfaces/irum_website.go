package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IRumWebsite interface {
	GetRumWebsites() ([]models.RumWebsite, int, string, error)
	CreateRumWebsite(*models.RumWebsite) (*models.RumWebsite, string, error)
	GetRumWebsite(rumWebsiteId string) (*models.RumWebsite, string, error)
	DeleteRumWebsite(rumWebsiteId string) (string, error)
	UpdateRumWebsite(*models.RumWebsite) (string, error)
}
