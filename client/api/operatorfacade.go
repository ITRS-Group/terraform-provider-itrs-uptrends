package client

import (
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// OperatorFacade implements IOperatorFacade.
type OperatorFacade struct {
	operator      interfaces.IOperator
	operatorGroup interfaces.IOperatorGroup
	membership    interfaces.IMembership
}

// NewOperatorFacade constructs a new OperatorFacade instance.
func NewOperatorFacade(op interfaces.IOperator, opGroup interfaces.IOperatorGroup, mem interfaces.IMembership) *OperatorFacade {
	return &OperatorFacade{
		operator:      op,
		operatorGroup: opGroup,
		membership:    mem,
	}
}

// UpdateOperator updates an operator's details and synchronizes admin membership.
func (of *OperatorFacade) UpdateOperator(operatorID string, requestBody models.OperatorExtendedUpdateRequest) (int, string, error) {
	// Convert the extended update request to the basic update request.
	basicUpdate := models.OperatorUpdateRequest{
		FullName:                       requestBody.FullName,
		Email:                          requestBody.Email,
		MobilePhone:                    requestBody.MobilePhone,
		BackupEmail:                    requestBody.BackupEmail,
		IsOnDuty:                       requestBody.IsOnDuty,
		SmsProvider:                    requestBody.SmsProvider,
		DefaultDashboard:               requestBody.DefaultDashboard,
		OutgoingPhoneNumberIdSpecified: requestBody.OutgoingPhoneNumberIdSpecified,
		CultureNameSpecified:           requestBody.CultureNameSpecified,
		TimeZoneIdSpecified:            requestBody.TimeZoneIdSpecified,
		UseNumericSenderSpecified:      requestBody.UseNumericSenderSpecified,
		AllowNativeLoginSpecified:      requestBody.AllowNativeLoginSpecified,
		AllowSingleSignonSpecified:     requestBody.AllowSingleSignonSpecified,
		OperatorRole:                   requestBody.OperatorRole,
	}

	// Handle the optional password field
	if requestBody.Password != nil {
		basicUpdate.Password = requestBody.Password // Set the password if provided
	}

	status, msg, err := of.operator.UpdateOperator(operatorID, basicUpdate)
	if err != nil {
		return status, msg, err
	}

	return status, msg, nil
}

// CreateOperator creates a new operator and assigns admin membership as indicated.
func (of *OperatorFacade) CreateOperator(requestData models.OperatorExtendedRequest) (models.OperatorExtendedResponse, int, error, string) {
	// Convert the extended request to the basic operator request.
	basicRequest := models.OperatorRequest{
		FullName:         requestData.FullName,
		Email:            requestData.Email,
		Password:         requestData.Password,
		MobilePhone:      requestData.MobilePhone,
		BackupEmail:      requestData.BackupEmail,
		IsOnDuty:         requestData.IsOnDuty,
		SmsProvider:      requestData.SmsProvider,
		DefaultDashboard: requestData.DefaultDashboard,
		OperatorRole:     requestData.OperatorRole,
	}
	opResp, statusCode, err, msg := of.operator.CreateOperator(basicRequest)
	if err != nil {
		return models.OperatorExtendedResponse{}, statusCode, err, msg
	}

	// Convert opResp to OperatorExtendedResponse directly.
	extendedResp := models.OperatorExtendedResponse(opResp)

	return extendedResp, statusCode, nil, msg
}

// DeleteOperator delegates the deletion to the underlying IOperator.
func (of *OperatorFacade) DeleteOperator(operatorID string) (int, string, error) {
	return of.operator.DeleteOperator(operatorID)
}

// GetOperator retrieves a single operator by ID and enriches with IsAdmin flag if needed.
func (of *OperatorFacade) GetOperator(operatorID string) (models.OperatorExtendedGetResponse, int, error, string) {
	op, statusCode, err, msg := of.operator.GetOperator(operatorID)
	if err != nil {
		return models.OperatorExtendedGetResponse{}, statusCode, err, msg
	}
	return models.OperatorExtendedGetResponse(*op), statusCode, nil, msg
}
