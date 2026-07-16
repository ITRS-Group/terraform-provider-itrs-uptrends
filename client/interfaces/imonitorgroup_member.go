package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IMonitorGroupMember interface {
	AssignMembership(monitorGroupGuid, monitorGuid string) error
	GetGroupMemberships(monitorGroupGuid string) ([]models.MonitorMembershipResponse, error)
	DeleteMembership(monitorGroupGuid, monitorGuid string) error
}
