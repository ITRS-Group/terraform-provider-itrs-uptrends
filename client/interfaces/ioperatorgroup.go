package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IOperatorGroup interface {
	GetOperatorGroups() ([]models.OperatorGroupResponse, int, string, error)
	CreateOperatorGroup(description string) (*models.OperatorGroupResponse, error, string)
	GetOperatorGroup(operatorGroupId string) (*models.OperatorGroupResponse, error, string)
	DeleteOperatorGroup(operatorGroupId string) (error, string)
	UpdateOperatorGroup(description string, operatorGroupID string) (string, error)
}
