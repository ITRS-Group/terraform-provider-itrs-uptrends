package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IAlertDefinitionOperatorGroupMembership interface {
	CreateMembership(alertDefinitionGuid string, escalationLevelNumber int, operatorGuid string) (*models.AlertDefinitionOperatorGroupMembershipResponse, error)
	GetMembership(alertDefinitionGuid string, escalationLevelNumber int) ([]models.GetMembershipResponse, error)
	DeleteMembership(alertDefinitionGuid string, escalationLevelNumber int, operatorGuid string) error
}
