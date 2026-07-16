package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type ICustomIntegration interface {
	GetCustomIntegrations() ([]models.CustomIntegration, error)
	CreateCustomIntegration(*models.CustomIntegration) (*models.CustomIntegration, error)
	GetCustomIntegration(integrationGuid string) (*models.CustomIntegration, error)
	UpdateCustomIntegration(integrationGuid string, integration *models.CustomIntegration) (*models.CustomIntegration, error)
	DeleteCustomIntegration(integrationGuid string) error
}
