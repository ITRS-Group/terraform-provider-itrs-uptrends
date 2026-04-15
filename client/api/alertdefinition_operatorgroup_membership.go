package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// AlertDefinitionOperatorGroupMembership represents a client for interacting with the Uptrends API.
type AlertDefinitionOperatorGroupMembership struct {
	client  *resty.Client
	baseUrl string
}

// NewAlertDefinitionOperatorGroupMembership creates and returns a new AlertDefinitionOperatorGroupMembership instance.
// Parameters:
//   - baseUrl: The base URL of the API (e.g., "https://api.uptrends.com/v4/AlertDefinition").
//   - authHeader: The value for the Authorization header.
func NewAlertDefinitionOperatorGroupMembership(baseUrl, authHeader, version, platform string) *AlertDefinitionOperatorGroupMembership {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &AlertDefinitionOperatorGroupMembership{
		client:  client,
		baseUrl: baseUrl,
	}
}

// CreateMembership sends a POST call to create a relation between an alert definition and an operator.
func (adm *AlertDefinitionOperatorGroupMembership) CreateMembership(alertDefinitionGuid string, escalationLevelNumber int, operatorGuid string) (*models.AlertDefinitionOperatorGroupMembershipResponse, error) {
	url := fmt.Sprintf("%s/%s/EscalationLevel/%d/Member/OperatorGroup/%s", adm.baseUrl, alertDefinitionGuid, escalationLevelNumber, operatorGuid)

	resp, err := adm.client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&models.AlertDefinitionOperatorGroupMembershipResponse{}).
		Post(url)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*models.AlertDefinitionOperatorGroupMembershipResponse), nil
}

// GetMembership sends a GET call to retrieve the membership details for the specified alert definition and escalation level.
func (adm *AlertDefinitionOperatorGroupMembership) GetMembership(alertDefinitionGuid string, escalationLevelNumber int) ([]models.GetMembershipResponse, error) {
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
func (adm *AlertDefinitionOperatorGroupMembership) DeleteMembership(alertDefinitionGuid string, escalationLevelNumber int, operatorGuid string) error {
	url := fmt.Sprintf("%s/%s/EscalationLevel/%d/Member/OperatorGroup/%s", adm.baseUrl, alertDefinitionGuid, escalationLevelNumber, operatorGuid)

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
