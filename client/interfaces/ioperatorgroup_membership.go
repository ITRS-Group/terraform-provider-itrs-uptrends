package client

import (
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

type IMembership interface {
	AssignOperator(operatorGroupGuid, operatorGuid string) error
	GetMemberships(operatorGroupGuid string) ([]models.MembershipResponse, error)
	DeleteMembership(operatorGroupGuid, operatorGuid string) error
}
