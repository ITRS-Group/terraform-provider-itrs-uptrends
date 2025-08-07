package client

import (
	jsonmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IMonitor interface {
	GetMonitor(monitorGuid string) (*jsonmodels.MonitorResponse, error)
	CreateMonitor(payload jsonmodels.MonitorRequest) (*jsonmodels.MonitorResponse, int, string, error)
	UpdateMonitor(monitorGuid string, payload jsonmodels.MonitorRequest) (int, string, error)
	DeleteMonitor(monitorGuid string) (int, string, error)
}
