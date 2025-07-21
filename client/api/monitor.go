package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	jsonmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type Monitor struct {
	client  *resty.Client
	baseURL string
}

var _ interfaces.IMonitor = (*Monitor)(nil)

// NewMonitorClient creates a new Monitor client with default headers.
// The authHeader parameter must include the full value for the "Authorization" header,
// and baseURL must be the full endpoint ending with "Monitor".
func NewMonitorClient(authHeader, baseURL, version, platform string) *Monitor {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &Monitor{
		client:  client,
		baseURL: baseURL,
	}
}

func (m *Monitor) GetMonitor(monitorGuid string) (*jsonmodels.MonitorResponse, error) {
	var monitorResponse jsonmodels.MonitorResponse
	url := fmt.Sprintf("%s/%s", m.baseURL, monitorGuid)

	resp, err := m.client.R().
		SetResult(&monitorResponse).
		Get(url)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to get monitor (HTTP %d): %s", resp.StatusCode(), resp.String())
	}
	return &monitorResponse, nil
}

func (m *Monitor) CreateMonitor(payload jsonmodels.MonitorRequest) (*jsonmodels.MonitorResponse, int, string, error) {
	var monitorResponse jsonmodels.MonitorResponse

	marshalRequestData, err := json.Marshal(payload)

	if err != nil {
		return nil, 0, "", fmt.Errorf("failed to marshal request data: %v", err)
	}
	resp, err := m.client.R().
		SetBody(marshalRequestData).
		SetResult(&monitorResponse).
		Post(m.baseURL)

	if err != nil {
		return nil, 0, resp.String(), fmt.Errorf("failed to execute HTTP request: %v", err)
	}

	return &monitorResponse, resp.StatusCode(), resp.String(), nil
}

func (m *Monitor) UpdateMonitor(monitorGuid string, payload jsonmodels.MonitorRequest) (int, string, error) {
	url := fmt.Sprintf("%s/%s", m.baseURL, monitorGuid)

	marshalRequestData, err := json.Marshal(payload)

	if err != nil {
		return 0, "", fmt.Errorf("failed to marshal request data: %v", err)
	}

	resp, err := m.client.R().
		SetBody(marshalRequestData).
		Patch(url)

	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode(), resp.String(), nil
}

func (m *Monitor) DeleteMonitor(monitorGuid string) (int, string, error) {
	url := fmt.Sprintf("%s/%s", m.baseURL, monitorGuid)

	resp, err := m.client.R().
		Delete(url)

	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode(), resp.String(), nil
}
