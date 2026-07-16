package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// Membership represents a client for interacting with the Uptrends API.
type Membership struct {
	client  *resty.Client
	baseUrl string
}

// NewMembership creates and returns a new NewMembership instance.
// Parameters:
//   - baseUrl: The base URL of the API (e.g., "https://api.uptrends.com/v4/OperatorGroup").
//   - authHeader: The value for the Authorization header.
func NewMembership(baseUrl, authHeader, version, platform string) *Membership {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"authorization": authHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &Membership{
		client:  client,
		baseUrl: baseUrl,
	}
}

// AssignOperator assigns an operator to an operator group.
// It sends a POST request to the API endpoint.
func (uc *Membership) AssignOperator(operatorGroupGuid, operatorGuid string) error {
	url := fmt.Sprintf("%s/%s/Member/%s", uc.baseUrl, operatorGroupGuid, operatorGuid)
	resp, err := uc.client.R().Post(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to assign operator: %s %s", resp.Status(), resp.Body())
	}
	return nil
}

// GetMemberships retrieves all memberships (operators) of an operator group.
// It sends a GET request to the API endpoint and returns a slice of Membership.
func (uc *Membership) GetMemberships(operatorGroupGuid string) ([]models.MembershipResponse, error) {
	url := fmt.Sprintf("%s/%s/Member", uc.baseUrl, operatorGroupGuid)
	var memberships []models.MembershipResponse

	resp, err := uc.client.R().
		SetResult(&memberships).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to retrieve memberships: %s %s", resp.Status(), resp.Body())
	}

	return memberships, nil
}

// DeleteMembership deletes a membership by removing an operator from an operator group.
// It sends a DELETE request to the API endpoint.
func (uc *Membership) DeleteMembership(operatorGroupGuid, operatorGuid string) error {
	url := fmt.Sprintf("%s/%s/Member/%s", uc.baseUrl, operatorGroupGuid, operatorGuid)
	resp, err := uc.client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("failed to delete membership: %s %s", resp.Status(), resp.Body())
	}
	return nil
}
