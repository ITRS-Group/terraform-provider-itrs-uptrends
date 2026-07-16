package client

import models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"

type IMonitorGroupPermission interface {
	GetMonitorGroupAuthorizations(monitorGroupGuid string) ([]models.MonitorGroupAuthorization, error)
	CreateMonitorGroupAuthorization(monitorGroupGuid string, auth models.MonitorGroupAuthorization) (*models.MonitorGroupAuthorization, error)
	DeleteMonitorGroupAuthorization(monitorGroupGuid, authorizationGuid string) error
}
