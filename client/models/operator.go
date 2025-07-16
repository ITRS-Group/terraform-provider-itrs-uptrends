package client

type OperatorGetResponse struct {
	OperatorGuid           string `json:"OperatorGuid"`
	Hash                   string `json:"Hash"`
	FullName               string `json:"FullName"`
	Email                  string `json:"Email"`
	MobilePhone            string `json:"MobilePhone"`
	IsAccountAdministrator bool   `json:"IsAccountAdministrator"`
	BackupEmail            string `json:"BackupEmail"`
	IsOnDuty               bool   `json:"IsOnDuty"`
	SmsProvider            string `json:"SmsProvider"`
	DefaultDashboard       string `json:"DefaultDashboard"`
	OperatorRole           string `json:"OperatorRole"`
}

type OperatorUpdateRequest struct {
	FullName                       string  `json:"FullName"`
	Email                          string  `json:"Email"`
	Password                       *string `json:"Password,omitempty"`
	MobilePhone                    string  `json:"MobilePhone"`
	BackupEmail                    string  `json:"BackupEmail"`
	IsOnDuty                       bool    `json:"IsOnDuty"`
	SmsProvider                    string  `json:"SmsProvider"`
	DefaultDashboard               string  `json:"DefaultDashboard"`
	OutgoingPhoneNumberIdSpecified bool    `json:"OutgoingPhoneNumberIdSpecified"`
	CultureNameSpecified           bool    `json:"CultureNameSpecified"`
	TimeZoneIdSpecified            bool    `json:"TimeZoneIdSpecified"`
	UseNumericSenderSpecified      bool    `json:"UseNumericSenderSpecified"`
	AllowNativeLoginSpecified      bool    `json:"AllowNativeLoginSpecified"`
	AllowSingleSignonSpecified     bool    `json:"AllowSingleSignonSpecified"`
	OperatorRole                   string  `json:"OperatorRole"`
}

type OperatorRequest struct {
	FullName         string  `json:"FullName"`
	Email            string  `json:"Email"`
	Password         string  `json:"Password"`
	MobilePhone      string  `json:"MobilePhone"`
	BackupEmail      *string `json:"BackupEmail,omitempty"`
	IsOnDuty         bool    `json:"IsOnDuty"`
	SmsProvider      string  `json:"SmsProvider"`
	DefaultDashboard string  `json:"DefaultDashboard"`
	OperatorRole     string  `json:"OperatorRole,omitempty"`
}

type OperatorResponse struct {
	OperatorGuid           string `json:"OperatorGuid"`
	Hash                   string `json:"Hash"`
	FullName               string `json:"FullName"`
	Email                  string `json:"Email"`
	MobilePhone            string `json:"MobilePhone"`
	IsAccountAdministrator bool   `json:"IsAccountAdministrator"`
	BackupEmail            string `json:"BackupEmail"`
	IsOnDuty               bool   `json:"IsOnDuty"`
	CultureName            string `json:"CultureName"`
	SmsProvider            string `json:"SmsProvider"`
	DefaultDashboard       string `json:"DefaultDashboard"`
	OperatorRole           string `json:"OperatorRole"`
}
