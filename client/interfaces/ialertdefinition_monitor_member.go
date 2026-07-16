package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IAlertDefinitionMonitorMember interface {
	AssignMonitor(alertDefinitionGuid, monitorGuid string) (*models.AssignResponse, error)
	RemoveAssignment(alertDefinitionGuid, monitorGuid string) error
	GetAssignments(alertDefinitionGuid string) ([]models.Assignment, error)
}
