package client

type OperatorExtendedUpdateRequest struct {
	FullName                       string
	Email                          string
	Password                       *string
	MobilePhone                    string
	BackupEmail                    string
	IsOnDuty                       bool
	SmsProvider                    string
	DefaultDashboard               string
	OutgoingPhoneNumberIdSpecified bool
	CultureNameSpecified           bool
	TimeZoneIdSpecified            bool
	UseNumericSenderSpecified      bool
	AllowNativeLoginSpecified      bool
	AllowSingleSignonSpecified     bool
	OperatorRole                   string
	IsAccountAdministrator         bool
}

type OperatorExtendedGetResponse struct {
	OperatorGuid           string
	Hash                   string
	FullName               string
	Email                  string
	MobilePhone            string
	IsAccountAdministrator bool
	BackupEmail            string
	IsOnDuty               bool
	SmsProvider            string
	DefaultDashboard       string
	OperatorRole           string
}

type OperatorExtendedResponse struct {
	OperatorGuid           string
	Hash                   string
	FullName               string
	Email                  string
	MobilePhone            string
	IsAccountAdministrator bool
	BackupEmail            string
	IsOnDuty               bool
	CultureName            string
	SmsProvider            string
	DefaultDashboard       string
	OperatorRole           string
}

type OperatorExtendedRequest struct {
	FullName         string
	Email            string
	Password         string
	MobilePhone      string
	BackupEmail      *string
	IsOnDuty         bool
	SmsProvider      string
	DefaultDashboard string
	OperatorRole     string
}
