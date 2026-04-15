package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IAlertDefinitionMonitorGroupMember interface {
	AssignMonitorGroup(alertDefinitionGuid, monitorGroupGuid string) (*models.AlertDefinitionMonitorGroupMembershipResponse, error)
	RemoveAssignment(alertDefinitionGuid, monitorGroupGuid string) error
	GetMonitorGroupAssignments(alertDefinitionGuid string) ([]models.GetMonitorGroupMembershipResponse, error)
}
