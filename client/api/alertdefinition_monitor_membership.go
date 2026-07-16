package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// AlertDefinitionMonitorMember manages monitor assignments for alert definitions.
type AlertDefinitionMonitorMember struct {
	baseUrl    string
	authHeader string
	client     *resty.Client
}

// NewAlertDefinitionMonitorMember creates a new instance of AlertDefinitionMonitorMember.
// baseUrl should be the full URL ending with "AlertDefinition" (e.g., "https://api.example.com/v4/AlertDefinition").
// authHeader should be the complete authorization header (e.g., "Basic <token>").
func NewAlertDefinitionMonitorMember(baseUrl, authHeader, version, platform string) *AlertDefinitionMonitorMember {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &AlertDefinitionMonitorMember{
		baseUrl:    baseUrl,
		authHeader: authHeader,
		client:     client,
	}
}

// AssignMonitor assigns a monitor to an alert definition.
func (adm *AlertDefinitionMonitorMember) AssignMonitor(alertDefinitionGuid, monitorGuid string) (*models.AssignResponse, error) {
	url := fmt.Sprintf("%s/%s/Member/Monitor/%s", adm.baseUrl, alertDefinitionGuid, monitorGuid)
	resp, err := adm.client.R().
		SetResult(&models.AssignResponse{}).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error assigning monitor: %s %s", resp.Status(), resp.Body())
	}
	return resp.Result().(*models.AssignResponse), nil
}

// RemoveAssignment removes the assignment of a monitor from an alert definition.
func (adm *AlertDefinitionMonitorMember) RemoveAssignment(alertDefinitionGuid, monitorGuid string) error {
	url := fmt.Sprintf("%s/%s/Member/Monitor/%s", adm.baseUrl, alertDefinitionGuid, monitorGuid)
	resp, err := adm.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("error removing assignment: %s %s", resp.Status(), resp.Body())
	}
	return nil
}

// GetAssignments retrieves all monitor assignments for a given alert definition.
func (adm *AlertDefinitionMonitorMember) GetAssignments(alertDefinitionGuid string) ([]models.Assignment, error) {
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
