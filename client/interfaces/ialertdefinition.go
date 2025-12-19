package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IAlertDefinition interface {
	GetAlertDefinition(alertDefinitionGuid string) (*models.AlertDefinitionResponse, error)
	GetAlertDefinitions() ([]models.AlertDefinitionResponse, int, string, error)
	CreateAlertDefinition(payload models.AlertDefinitionRequest) (*models.AlertDefinitionResponse, error)
	UpdateAlertDefinition(alertDefinitionGuid string, payload models.AlertDefinitionRequest) error
	DeleteAlertDefinition(alertDefinitionGuid string) error
	UpdateEscalationLevel(payload models.EscalationLevel) error
	GetEscalationLevels(alertDefinitionGuid string) ([]models.EscalationLevel, error)
}
