package client

type OperatorRequest struct {
	FullName         string  `json:"FullName"`
	Email            string  `json:"Email"`
	Password         *string `json:"Password,omitempty"`
	MobilePhone      string  `json:"MobilePhone"`
	BackupEmail      string  `json:"BackupEmail"`
	IsOnDuty         bool    `json:"IsOnDuty"`
	SmsProvider      string  `json:"SmsProvider"`
	DefaultDashboard string  `json:"DefaultDashboard"`
	OperatorRole     string  `json:"OperatorRole"`
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
