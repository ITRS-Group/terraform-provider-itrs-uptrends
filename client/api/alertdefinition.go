package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// AlertDefinition represents a client for interacting with the Alert Definition API.
type AlertDefinition struct {
	authHeader string
	baseURL    string
	client     *resty.Client
}

var _ interfaces.IAlertDefinition = (*AlertDefinition)(nil)

// NewAlertDefinition creates a new instance of AlertDefinition.
// baseURL should be the full URL ending with "AlertDefinition" (e.g., "https://domain/v4/AlertDefinition").
func NewAlertDefinition(baseURL, authHeader, version, platform string) *AlertDefinition {
	client := resty.New()
	// Set common headers for all requests
	client.SetHeaders(map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &AlertDefinition{
		authHeader: authHeader,
		baseURL:    baseURL,
		client:     client,
	}
}

// GetAlertDefinition retrieves a single alert definition by its GUID.
func (a *AlertDefinition) GetAlertDefinition(alertDefinitionGuid string) (*models.AlertDefinitionItem, error) {
	var definition models.AlertDefinitionItem
	url := fmt.Sprintf("%s/%s", a.baseURL, alertDefinitionGuid)

	resp, err := a.client.R().
		SetResult(&definition).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error fetching alert definition: %s", resp.Status())
	}

	return &definition, nil
}

// CreateAlertDefinition creates a new alert definition.
func (a *AlertDefinition) CreateAlertDefinition(reqData models.CreateAlertDefinitionRequest) (*models.AlertDefinitionItem, error) {
	var definition models.AlertDefinitionItem

	resp, err := a.client.R().
		SetBody(reqData).
		SetResult(&definition).
		Post(a.baseURL)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error creating alert definition: %s", resp.Status())
	}

	return &definition, nil
}

// UpdateAlertDefinition updates an existing alert definition by its GUID.
func (a *AlertDefinition) UpdateAlertDefinition(alertDefinitionGuid string, reqData models.UpdateAlertDefinitionRequest) error {
	url := fmt.Sprintf("%s/%s", a.baseURL, alertDefinitionGuid)

	resp, err := a.client.R().
		SetBody(reqData).
		Patch(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("error updating alert definition: %s", resp.Status())
	}

	return nil
}

// GetEscalationLevels retrieves all three escalation levels for a given alert definition.
func (a *AlertDefinition) GetEscalationLevels(alertDefinitionGuid string) ([]models.EscalationLevel, error) {
	var levels []models.EscalationLevel
	url := fmt.Sprintf("%s/%s/EscalationLevel", a.baseURL, alertDefinitionGuid)

	resp, err := a.client.R().
		SetResult(&levels).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error fetching escalation levels: %s", resp.Status())
	}

	return levels, nil
}

// PatchEscalationLevel updates a specific escalation level for a given alert definition.
func (a *AlertDefinition) PatchEscalationLevel(alertDefinitionGuid string, levelId int, reqData models.PatchEscalationLevelRequest) error {
	url := fmt.Sprintf("%s/%s/EscalationLevel/%d", a.baseURL, alertDefinitionGuid, levelId)

	resp, err := a.client.R().
		SetBody(reqData).
		Patch(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("error patching escalation level %d: %s", levelId, resp.Status())
	}

	return nil
}

func (a *AlertDefinition) DeleteAlertDefinition(alertDefinitionGuid string) error {
	url := fmt.Sprintf("%s/%s", a.baseURL, alertDefinitionGuid)

	resp, err := a.client.R().
		Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("error updating alert definition: %s", resp.Status())
	}

	return nil

}
