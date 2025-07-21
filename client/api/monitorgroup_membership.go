package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// MonitorGroupMember encapsulates methods to interact with the Membership API.
type MonitorGroupMember struct {
	baseURL    string
	authHeader string
	client     *resty.Client
}

// NewMonitorGroupMember creates a new instance of MonitorGroupMember based on a base URL and an auth header.
func NewMonitorGroupMember(baseURL, authHeader, version, platform string) *MonitorGroupMember {
	client := resty.New().
		SetBaseURL(baseURL).
		SetHeaders(map[string]string{
			"Authorization": authHeader,
		})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &MonitorGroupMember{
		baseURL:    baseURL,
		authHeader: authHeader,
		client:     client,
	}
}

// AssignMembership sends a POST request to assign a membership to a monitor group.
func (mc *MonitorGroupMember) AssignMembership(monitorGroupGuid, monitorGuid string) error {
	url := fmt.Sprintf("/%s/Member/%s", monitorGroupGuid, monitorGuid)
	resp, err := mc.client.R().Post(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("error assigning membership: %s", resp.Status())
	}
	return nil
}

// GetGroupMemberships retrieves all memberships for a given monitor group.
func (mc *MonitorGroupMember) GetGroupMemberships(monitorGroupGuid string) ([]models.MonitorMembershipResponse, error) {
	url := fmt.Sprintf("/%s/Member", monitorGroupGuid)
	var memberships []models.MonitorMembershipResponse
	resp, err := mc.client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&memberships).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error retrieving memberships: %s", resp.Status())
	}
	return memberships, nil
}

// DeleteMembership sends a DELETE request to remove a membership from a monitor group.
func (mc *MonitorGroupMember) DeleteMembership(monitorGroupGuid, monitorGuid string) error {
	url := fmt.Sprintf("/%s/Member/%s", monitorGroupGuid, monitorGuid)
	resp, err := mc.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("error deleting membership: %s", resp.Status())
	}
	return nil
}
