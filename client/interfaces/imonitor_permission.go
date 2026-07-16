package client

import models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"

type IMonitorPermission interface {
	GetMonitorAuthorizations(monitorGuid string) ([]models.MonitorAuthorization, error)
	CreateMonitorAuthorization(monitorGuid string, auth models.MonitorAuthorization) (*models.MonitorAuthorization, error)
	DeleteMonitorAuthorization(monitorGuid, authorizationGuid string) error
}
