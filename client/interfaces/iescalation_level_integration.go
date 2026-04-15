package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IEscalationLevelIntegration interface {
	GetIntegration(alertDefinitionGuid string, escalationLevelId int, integrationGuid string) (*models.EscalationLevelIntegrationResponse, error)
	AddIntegration(alertDefinitionGuid string, escalationLevelId int, payload models.EscalationLevelIntegrationRequest) (*models.EscalationLevelIntegrationResponse, error)
	UpdateIntegration(alertDefinitionGuid string, escalationLevelId int, integrationGuid string, payload models.EscalationLevelIntegrationRequest) error
	RemoveIntegration(alertDefinitionGuid string, escalationLevelId int, integrationGuid string) error
}
