package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IOperatorPermission interface {
	AssignOperatorPermission(operatorGuid, permission string) error
	GetOperatorPermission(operatorGuid string) (models.OperatorPermissionResponse, error)
	DeleteOperatorPermission(operatorGuid, permission string) error
}
