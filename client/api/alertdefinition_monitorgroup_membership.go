package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// AlertDefinitionMonitorGroupMember manages monitor group assignments for alert definitions.
type AlertDefinitionMonitorGroupMember struct {
	baseUrl    string
	authHeader string
	client     *resty.Client
}

// NewAlertDefinitionMonitorGroupMember creates a new instance of AlertDefinitionMonitorGroupMember.
// baseUrl should be the full URL ending with "AlertDefinition" (e.g., "https://api.example.com/v4/AlertDefinition").
// authHeader should be the complete authorization header (e.g., "Basic <token>").
func NewAlertDefinitionMonitorGroupMember(baseUrl, authHeader, version, platform string) *AlertDefinitionMonitorGroupMember {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &AlertDefinitionMonitorGroupMember{
		baseUrl:    baseUrl,
		authHeader: authHeader,
		client:     client,
	}
}

// AssignMonitorGroup assigns a monitor group to an alert definition.
func (adm *AlertDefinitionMonitorGroupMember) AssignMonitorGroup(alertDefinitionGuid, monitorGroupGuid string) (*models.AssignResponse, error) {
	url := fmt.Sprintf("%s/%s/Member/MonitorGroup/%s", adm.baseUrl, alertDefinitionGuid, monitorGroupGuid)
	resp, err := adm.client.R().
		SetResult(&models.AssignResponse{}).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error assigning monitor group: %s %s", resp.Status(), resp.Body())
	}
	return resp.Result().(*models.AssignResponse), nil
}

// RemoveAssignment removes the assignment of a monitor group from an alert definition.
func (adm *AlertDefinitionMonitorGroupMember) RemoveAssignment(alertDefinitionGuid, monitorGroupGuid string) error {
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

// GetAssignments retrieves all monitor group assignments for a given alert definition.
func (adm *AlertDefinitionMonitorGroupMember) GetAssignments(alertDefinitionGuid string) ([]models.Assignment, error) {
	url := fmt.Sprintf("%s/%s/Member", adm.baseUrl, alertDefinitionGuid)
	resp, err := adm.client.R().
		SetResult(&[]models.Assignment{}).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error retrieving assignments: %s %s", resp.Status(), resp.Body())
	}
	assignments := *resp.Result().(*[]models.Assignment)
	return assignments, nil
}
