package client

import (
	jsonmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IMonitor interface {
	GetMonitor(monitorGuid string) (*jsonmodels.MonitorResponse, error)
	GetMonitors() ([]jsonmodels.MonitorResponse, int, string, error)
	CreateMonitor(payload jsonmodels.MonitorRequest, monitorGroupGuid *string) (*jsonmodels.MonitorResponse, int, string, error)
	UpdateMonitor(monitorGuid string, payload jsonmodels.MonitorRequest) (int, string, error)
	DeleteMonitor(monitorGuid string) (int, string, error)
}
