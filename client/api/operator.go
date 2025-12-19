package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	httpclient "github.com/itrs-group/terraform-provider-itrs-uptrends/client/httpclient"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type Operator struct {
	Client  *resty.Client
	BaseUrl string
}

// Ensure MyStruct implements MyInterface
var _ interfaces.IOperator = (*Operator)(nil)

func NewOperator(baseURL, authenticationHeader, version, platform string) *Operator {
	client := resty.New()

	// Set User Properties
	client.SetHeaders(map[string]string{
		"accept":        "application/json",
		"Content-Type":  "application/json",
		"authorization": authenticationHeader,
	})
	customHTTPClient := httpclient.NewHTTPClient(version, platform)
	client.SetTransport(customHTTPClient.Transport)
	return &Operator{
		Client:  client,
		BaseUrl: baseURL,
	}
}

// UpdateOperator sends a PUT request to update an operator and returns the error, status code, and response body.
func (a *Operator) UpdateOperator(operatorID string, requestBody models.OperatorRequest) (int, string, error) {

	var updateUrl = a.BaseUrl + "/" + operatorID
	resp, err := a.Client.R().
		SetBody(requestBody). // Use the passed struct as the body
		Patch(updateUrl)

	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode(), resp.String(), nil
}

func (a *Operator) CreateOperator(requestData models.OperatorRequest) (models.OperatorResponse, int, string, error) {
	var operatorResponse models.OperatorResponse
	resp, err := a.Client.R().
		SetBody(requestData).
		SetResult(&operatorResponse).
		Post(a.BaseUrl)

	if err != nil {
		return models.OperatorResponse{}, 0, resp.String(), fmt.Errorf("failed to execute HTTP request: %v", err)
	}

	// Return the response details
	return operatorResponse, resp.StatusCode(), resp.String(), nil
}

func (a *Operator) DeleteOperator(operatorID string) (int, string, error) {
	var updateUrl = a.BaseUrl + "/" + operatorID
	resp, err := a.Client.R().
		Delete(updateUrl)

	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode(), resp.String(), nil
}

func (a *Operator) GetOperator(operatorID string) (*models.OperatorResponse, int, string, error) {
	var operator models.OperatorResponse
	url := a.BaseUrl + "/" + operatorID
	resp, err := a.Client.R().
		SetResult(&operator).
		Get(url)

	statusCode := -1
	responseBody := ""
	if resp != nil {
		statusCode = resp.StatusCode()
		responseBody = resp.String()
	}

	if err != nil {
		return nil, statusCode, responseBody, err
	}
	if !resp.IsSuccess() {
		return nil, statusCode, responseBody, fmt.Errorf("failed to get operator: %s", resp.Status())
	}
	return &operator, statusCode, responseBody, nil
}

func (a *Operator) GetOperators() ([]models.OperatorResponse, int, string, error) {
	var operators []models.OperatorResponse
	url := a.BaseUrl
	resp, err := a.Client.R().
		SetResult(&operators).
		Get(url)

	statusCode := -1
	responseBody := ""
	if resp != nil {
		statusCode = resp.StatusCode()
		responseBody = resp.String()
	}

	if err != nil {
		return nil, statusCode, responseBody, err
	}
	if !resp.IsSuccess() {
		return nil, statusCode, responseBody, fmt.Errorf("failed to get operators: %s", resp.Status())
	}
	return operators, statusCode, responseBody, nil
}
