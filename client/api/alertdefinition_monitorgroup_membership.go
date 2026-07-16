package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// AlertDefinitionMonitorMember manages monitor assignments for alert definitions.
type AlertDefinitionMonitorGroupMembership struct {
	baseUrl    string
	authHeader string
	client     *resty.Client
}

// NewAlertDefinitionMonitorMember creates a new instance of AlertDefinitionMonitorMember.
// baseUrl should be the full URL ending with "AlertDefinition" (e.g., "https://api.example.com/v4/AlertDefinition").
// authHeader should be the complete authorization header (e.g., "Basic <token>").
func NewAlertDefinitionMonitorGroupMembership(baseUrl, authHeader, version, platform string) *AlertDefinitionMonitorGroupMembership {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &AlertDefinitionMonitorGroupMembership{
		baseUrl:    baseUrl,
		authHeader: authHeader,
		client:     client,
	}
}

// AssignMonitor assigns a monitor to an alert definition.
func (adm *AlertDefinitionMonitorGroupMembership) AssignMonitorGroup(alertDefinitionGuid, monitorGroupGuid string) (*models.AlertDefinitionMonitorGroupMembershipResponse, error) {
	url := fmt.Sprintf("%s/%s/Member/MonitorGroup/%s", adm.baseUrl, alertDefinitionGuid, monitorGroupGuid)
	resp, err := adm.client.R().
		SetResult(&models.AlertDefinitionMonitorGroupMembershipResponse{}).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error assigning monitor: %s %s", resp.Status(), resp.Body())
	}
	return resp.Result().(*models.AlertDefinitionMonitorGroupMembershipResponse), nil
}

// RemoveAssignment removes the assignment of a monitor from an alert definition.
func (adm *AlertDefinitionMonitorGroupMembership) RemoveAssignment(alertDefinitionGuid, monitorGroupGuid string) error {
	url := fmt.Sprintf("%s/%s/Member/MonitorGroup/%s", adm.baseUrl, alertDefinitionGuid, monitorGroupGuid)
	resp, err := adm.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("error removing assignment: %s %s", resp.Status(), resp.Body())
	}
	return nil
}

// GetAssignments retrieves all monitor group assignments for a given alert definition
func (adm *AlertDefinitionMonitorGroupMembership) GetMonitorGroupAssignments(alertDefinitionGuid string) ([]models.GetMonitorGroupMembershipResponse, error) {
	url := fmt.Sprintf("%s/%s/Member", adm.baseUrl, alertDefinitionGuid)
	resp, err := adm.client.R().
		SetResult(&[]models.GetMonitorGroupMembershipResponse{}).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error retrieving assignments: %s %s", resp.Status(), resp.Body())
	}
	return *resp.Result().(*[]models.GetMonitorGroupMembershipResponse), nil
}
