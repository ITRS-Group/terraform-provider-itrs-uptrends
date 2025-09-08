package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IAccount interface {
	// GetAccountInfo retrieves account information
	GetAccountInfo() (*models.AccountResponse, int, error)
}
