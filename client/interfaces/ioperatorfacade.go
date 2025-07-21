package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IOperatorFacade interface {
	GetOperator(operatorID string) (models.OperatorExtendedGetResponse, int, error, string)
	UpdateOperator(operatorID string, requestBody models.OperatorExtendedUpdateRequest) (int, string, error)
	CreateOperator(requestData models.OperatorExtendedRequest) (models.OperatorExtendedResponse, int, error, string)
	DeleteOperator(operatorID string) (int, string, error)
}
