package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IMonitorGroupClient interface {
	GetMonitorGroups() ([]models.MonitorGroupResponse, int, string, error)
	CreateMonitorGroup(payload models.MonitorGroupRequest) (models.MonitorGroupResponse, int, string, error)
	UpdateMonitorGroup(payload models.MonitorGroupRequest, monitorGroupId string) (int, string, error)
	DeleteMonitorGroup(monitorGroupGuid string) (int, string, error)
	GetMonitorGroup(monitorGroupGuid string) (models.MonitorGroupResponse, string, error)
}
