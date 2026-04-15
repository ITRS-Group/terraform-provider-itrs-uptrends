package client

type OperatorRequest struct {
	FullName                       string  `json:"FullName"`
	Email                          string  `json:"Email"`
	Password                       *string `json:"Password,omitempty"`
	MobilePhone                    string  `json:"MobilePhone"`
	OutgoingPhoneNumberId          *int    `json:"OutgoingPhoneNumberId,omitempty"`
	OutgoingPhoneNumberIdSpecified bool    `json:"OutgoingPhoneNumberIdSpecified"`
	BackupEmail                    string  `json:"BackupEmail"`
	IsOnDuty                       bool    `json:"IsOnDuty"`
	CultureName                    *string `json:"CultureName,omitempty"`
	CultureNameSpecified           bool    `json:"CultureNameSpecified"`
	TimeZoneId                     *int    `json:"TimeZoneId,omitempty"`
	TimeZoneIdSpecified            bool    `json:"TimeZoneIdSpecified"`
	SmsProvider                    string  `json:"SmsProvider"`
	UseNumericSender               *bool   `json:"UseNumericSender,omitempty"`
	UseNumericSenderSpecified      bool    `json:"UseNumericSenderSpecified"`
	AllowNativeLogin               *bool   `json:"AllowNativeLogin,omitempty"`
	AllowNativeLoginSpecified      bool    `json:"AllowNativeLoginSpecified"`
	AllowSingleSignon              *bool   `json:"AllowSingleSignon,omitempty"`
	AllowSingleSignonSpecified     bool    `json:"AllowSingleSignonSpecified"`
	DefaultDashboard               string  `json:"DefaultDashboard"`
	SetupMode                      string  `json:"SetupMode"`
	OperatorRole                   string  `json:"OperatorRole"`
}

type OperatorResponse struct {
	OperatorGuid           string `json:"OperatorGuid"`
	Hash                   string `json:"Hash"`
	FullName               string `json:"FullName"`
	Email                  string `json:"Email"`
	MobilePhone            string `json:"MobilePhone"`
	OutgoingPhoneNumberId  int    `json:"OutgoingPhoneNumberId"`
	IsAccountAdministrator bool   `json:"IsAccountAdministrator"`
	BackupEmail            string `json:"BackupEmail"`
	IsOnDuty               bool   `json:"IsOnDuty"`
	CultureName            string `json:"CultureName"`
	TimeZoneId             int    `json:"TimeZoneId"`
	SmsProvider            string `json:"SmsProvider"`
	UseNumericSender       bool   `json:"UseNumericSender"`
	AllowNativeLogin       bool   `json:"AllowNativeLogin"`
	AllowSingleSignon      bool   `json:"AllowSingleSignon"`
	DefaultDashboard       string `json:"DefaultDashboard"`
	SetupMode              string `json:"SetupMode"`
	OperatorRole           string `json:"OperatorRole"`
}
