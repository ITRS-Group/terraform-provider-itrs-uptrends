package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IOperatorGroupPermission interface {
	AssignOperatorGroupPermission(operatorGroupGuid, permission string) error
	GetOperatorGroupPermission(operatorGroupGuid string) (models.OperatorGroupPermissionResponse, error)
	DeleteOperatorGroupPermission(operatorGroupGuid, permission string) error
}
