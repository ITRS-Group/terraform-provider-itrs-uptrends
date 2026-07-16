package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type ICustomIntegrationPermission interface {
	GetIntegrationAuthorizations(integrationGuid string) ([]models.IntegrationAuthorization, error)
	CreateIntegrationAuthorization(integrationGuid string, auth models.IntegrationAuthorization) (*models.IntegrationAuthorization, error)
	DeleteIntegrationAuthorization(integrationGuid, authorizationGuid string) error
}
