package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IAlertDefinition interface {
	GetAlertDefinition(alertDefinitionGuid string) (*models.AlertDefinitionItem, error)
	CreateAlertDefinition(reqData models.CreateAlertDefinitionRequest) (*models.AlertDefinitionItem, error)
	UpdateAlertDefinition(alertDefinitionGuid string, reqData models.UpdateAlertDefinitionRequest) error
	GetEscalationLevels(alertDefinitionGuid string) ([]models.EscalationLevel, error)
	PatchEscalationLevel(alertDefinitionGuid string, levelId int, reqData models.PatchEscalationLevelRequest) error
	DeleteAlertDefinition(alertDefinitionGuid string) error
}
