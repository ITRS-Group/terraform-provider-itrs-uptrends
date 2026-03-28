package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IAlertDefinitionMonitorGroupMember interface {
	AssignMonitorGroup(alertDefinitionGuid, monitorGroupGuid string) (*models.AssignResponse, error)
	RemoveAssignment(alertDefinitionGuid, monitorGroupGuid string) error
	GetAssignments(alertDefinitionGuid string) ([]models.Assignment, error)
}
