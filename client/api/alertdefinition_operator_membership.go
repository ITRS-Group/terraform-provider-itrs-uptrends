package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// AlertDefinitionOperatorMembership represents a client for interacting with the Uptrends API.
type AlertDefinitionOperatorMembership struct {
	client  *resty.Client
	baseUrl string
}

// NewAlertDefinitionOperatorMembership creates and returns a new AlertDefinitionOperatorMembership instance.
// Parameters:
//   - baseUrl: The base URL of the API (e.g., "https://api.uptrends.com/v4/AlertDefinition").
//   - authHeader: The value for the Authorization header.
func NewAlertDefinitionOperatorMembership(baseUrl, authHeader, version, platform string) *AlertDefinitionOperatorMembership {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &AlertDefinitionOperatorMembership{
		client:  client,
		baseUrl: baseUrl,
	}
}

// CreateMembership sends a POST call to create a relation between an alert definition and an operator.
func (adm *AlertDefinitionOperatorMembership) CreateMembership(alertDefinitionGuid string, escalationLevelNumber int, operatorGuid string) (*models.AlertDefinitionOperatorMembershipResponse, error) {
	url := fmt.Sprintf("%s/%s/EscalationLevel/%d/Member/Operator/%s", adm.baseUrl, alertDefinitionGuid, escalationLevelNumber, operatorGuid)

	resp, err := adm.client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&models.AlertDefinitionOperatorMembershipResponse{}).
		Post(url)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*models.AlertDefinitionOperatorMembershipResponse), nil
}

// GetMembership sends a GET call to retrieve the membership details for the specified alert definition and escalation level.
func (adm *AlertDefinitionOperatorMembership) GetMembership(alertDefinitionGuid string, escalationLevelNumber int) ([]models.GetMembershipResponse, error) {
	url := fmt.Sprintf("%s/%s/EscalationLevel/%d/Member", adm.baseUrl, alertDefinitionGuid, escalationLevelNumber)

	resp, err := adm.client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&[]models.GetMembershipResponse{}).
		Get(url)
	if err != nil {
		return nil, err
	}

	// Since SetResult returns a pointer to a slice, we dereference it.
	return *resp.Result().(*[]models.GetMembershipResponse), nil
}

// DeleteMembership sends a DELETE call to remove the specified relation.
func (adm *AlertDefinitionOperatorMembership) DeleteMembership(alertDefinitionGuid string, escalationLevelNumber int, operatorGuid string) error {
	url := fmt.Sprintf("%s/%s/EscalationLevel/%d/Member/Operator/%s", adm.baseUrl, alertDefinitionGuid, escalationLevelNumber, operatorGuid)

	resp, err := adm.client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)
	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("delete failed with status code: %d", resp.StatusCode())
	}
	return nil
}
