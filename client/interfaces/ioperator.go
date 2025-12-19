package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IOperator interface {
	GetOperator(operatorID string) (*models.OperatorResponse, int, string, error)
	GetOperators() ([]models.OperatorResponse, int, string, error)
	UpdateOperator(operatorID string, requestBody models.OperatorRequest) (int, string, error)
	CreateOperator(requestData models.OperatorRequest) (models.OperatorResponse, int, string, error)
	DeleteOperator(operatorID string) (int, string, error)
}
