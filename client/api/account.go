package client

import (
	"log"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

// Account represents the account handling structure
type Account struct {
	Client  *resty.Client
	BaseUrl string
}

var _ interfaces.IAccount = (*Account)(nil)

// NewAccount initializes a new Account instance
func NewAccount(baseURL, authenticationHeader, version, platform string) *Account {
	client := resty.New()

	// Set User Properties
	client.SetHeaders(map[string]string{
		"accept":        "application/json",
		"authorization": authenticationHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &Account{
		Client:  client,
		BaseUrl: baseURL,
	}
}

func (a *Account) GetAccountInfo() (*models.AccountResponse, int, error) {
	var account models.AccountResponse
	// Make the GET request
	resp, err := a.Client.R().
		SetResult(&account).
		Get(a.BaseUrl)

	var statusCode = -1
	if resp != nil {
		statusCode = resp.StatusCode()
	}

	log.Printf("Results err: %v and status code %d", err, statusCode)

	return &account, statusCode, err
}
