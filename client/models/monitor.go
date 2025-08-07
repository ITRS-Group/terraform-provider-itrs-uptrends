package client

// MonitorRequest is used for create and update (POST/PUT).
type MonitorRequest struct {
	MonitorGuid                         *string                  `json:"MonitorGuid,omitempty"`
	Name                                string                   `json:"Name"`
	IsActive                            bool                     `json:"IsActive"`
	GenerateAlert                       bool                     `json:"GenerateAlert"`
	CheckInterval                       int                      `json:"CheckInterval"`
	MonitorMode                         string                   `json:"MonitorMode"`
	CustomMetrics                       *[]CustomMetric          `json:"CustomMetrics,omitempty"`
	CustomFields                        []CustomField            `json:"CustomFields,omitempty"`
	SelectedCheckpoints                 SelectedCheckpoints      `json:"SelectedCheckpoints"`
	UsePrimaryCheckpointsOnly           bool                     `json:"UsePrimaryCheckpointsOnly"`
	SelfServiceTransactionScript        *string                  `json:"SelfServiceTransactionScript,omitempty"`
	MonitorType                         string                   `json:"MonitorType"`
	Notes                               *string                  `json:"Notes,omitempty"`
	MultiStepApiTransactionScript       *string                  `json:"MultiStepApiTransactionScript,omitempty"`
	BlockGoogleAnalytics                *bool                    `json:"BlockGoogleAnalytics,omitempty"`
	BlockUptrendsRum                    *bool                    `json:"BlockUptrendsRum,omitempty"`
	BlockUrls                           *[]string                `json:"BlockUrls,omitempty"`
	RequestHeaders                      *[]RequestHeader         `json:"RequestHeaders,omitempty"`
	UserAgent                           *string                  `json:"UserAgent,omitempty"`
	Username                            *string                  `json:"Username,omitempty"`
	Password                            *string                  `json:"Password,omitempty"`
	NameForPhoneAlerts                  *string                  `json:"NameForPhoneAlerts,omitempty"`
	AuthenticationType                  *string                  `json:"AuthenticationType,omitempty"`
	ThrottlingOptions                   *ThrottlingOptions       `json:"ThrottlingOptions,omitempty"`
	DnsBypasses                         *[]DnsBypass             `json:"DnsBypasses,omitempty"`
	CertificateName                     *string                  `json:"CertificateName,omitempty"`
	CertificateOrganization             *string                  `json:"CertificateOrganization,omitempty"`
	CertificateOrganizationalUnit       *string                  `json:"CertificateOrganizationalUnit,omitempty"`
	CertificateSerialNumber             *string                  `json:"CertificateSerialNumber,omitempty"`
	CertificateFingerprint              *string                  `json:"CertificateFingerprint,omitempty"`
	CertificateIssuerName               *string                  `json:"CertificateIssuerName,omitempty"`
	CertificateIssuerCompanyName        *string                  `json:"CertificateIssuerCompanyName,omitempty"`
	CertificateIssuerOrganizationalUnit *string                  `json:"CertificateIssuerOrganizationalUnit,omitempty"`
	CertificateExpirationWarningDays    *int                     `json:"CertificateExpirationWarningDays,omitempty"`
	CheckCertificateErrors              *bool                    `json:"CheckCertificateErrors,omitempty"`
	IgnoreExternalElements              *bool                    `json:"IgnoreExternalElements,omitempty"`
	DomainGroupGuid                     *string                  `json:"DomainGroupGuid,omitempty"`
	DomainGroupGuidSpecified            *bool                    `json:"DomainGroupGuidSpecified,omitempty"`
	DnsServer                           *string                  `json:"DnsServer,omitempty"`
	DnsQuery                            *string                  `json:"DnsQuery,omitempty"`
	DnsExpectedResult                   *string                  `json:"DnsExpectedResult,omitempty"`
	DnsTestValue                        *string                  `json:"DnsTestValue,omitempty"`
	Port                                *int                     `json:"Port,omitempty"`
	IpVersion                           *string                  `json:"IpVersion,omitempty"`
	DatabaseName                        *string                  `json:"DatabaseName,omitempty"`
	NetworkAddress                      *string                  `json:"NetworkAddress,omitempty"`
	ImapSecureConnection                *bool                    `json:"ImapSecureConnection,omitempty"`
	SftpAction                          *string                  `json:"SftpAction,omitempty"`
	SftpActionPath                      *string                  `json:"SftpActionPath,omitempty"`
	HttpMethod                          *string                  `json:"HttpMethod,omitempty"`
	TlsVersion                          *string                  `json:"TlsVersion,omitempty"`
	RequestBody                         *string                  `json:"RequestBody,omitempty"`
	Url                                 *string                  `json:"Url,omitempty"`
	BrowserType                         *string                  `json:"BrowserType,omitempty"`
	BrowserWindowDimensions             *BrowserWindowDimensions `json:"BrowserWindowDimensions,omitempty"`
	UseConcurrentMonitoring             *bool                    `json:"UseConcurrentMonitoring,omitempty"`
	ConcurrentUnconfirmedErrorThreshold *int                     `json:"ConcurrentUnconfirmedErrorThreshold,omitempty"`
	ConcurrentConfirmedErrorThreshold   *int                     `json:"ConcurrentConfirmedErrorThreshold,omitempty"`
	ErrorConditions                     *[]ErrorCondition        `json:"ErrorConditions,omitempty"`
}

// MonitorResponse is returned by GET and POST calls.
type MonitorResponse struct {
	MonitorGuid                         string                   `json:"MonitorGuid"`
	Name                                string                   `json:"Name"`
	IsActive                            bool                     `json:"IsActive"`
	GenerateAlert                       bool                     `json:"GenerateAlert"`
	IsLocked                            bool                     `json:"IsLocked"`
	CheckInterval                       int                      `json:"CheckInterval"`
	MonitorMode                         string                   `json:"MonitorMode"`
	CustomMetrics                       *[]CustomMetric          `json:"CustomMetrics,omitempty"`
	CustomFields                        []CustomField            `json:"CustomFields"`
	SelectedCheckpoints                 SelectedCheckpoints      `json:"SelectedCheckpoints"`
	UsePrimaryCheckpointsOnly           bool                     `json:"UsePrimaryCheckpointsOnly"`
	SelfServiceTransactionScript        *string                  `json:"SelfServiceTransactionScript,omitempty"`
	MonitorType                         string                   `json:"MonitorType"`
	Notes                               string                   `json:"Notes"`
	MultiStepApiTransactionScript       *string                  `json:"MultiStepApiTransactionScript,omitempty"`
	BlockGoogleAnalytics                *bool                    `json:"BlockGoogleAnalytics,omitempty"`
	BlockUptrendsRum                    *bool                    `json:"BlockUptrendsRum,omitempty"`
	BlockUrls                           *[]string                `json:"BlockUrls,omitempty"`
	RequestHeaders                      *[]RequestHeader         `json:"RequestHeaders,omitempty"`
	UserAgent                           *string                  `json:"UserAgent,omitempty"`
	Username                            *string                  `json:"Username,omitempty"`
	Password                            *string                  `json:"Password,omitempty"`
	NameForPhoneAlerts                  *string                  `json:"NameForPhoneAlerts,omitempty"`
	AuthenticationType                  *string                  `json:"AuthenticationType,omitempty"`
	ThrottlingOptions                   *ThrottlingOptions       `json:"ThrottlingOptions,omitempty"`
	DnsBypasses                         *[]DnsBypass             `json:"DnsBypasses,omitempty"`
	CertificateName                     *string                  `json:"CertificateName,omitempty"`
	CertificateOrganization             *string                  `json:"CertificateOrganization,omitempty"`
	CertificateOrganizationalUnit       *string                  `json:"CertificateOrganizationalUnit,omitempty"`
	CertificateSerialNumber             *string                  `json:"CertificateSerialNumber,omitempty"`
	CertificateFingerprint              *string                  `json:"CertificateFingerprint,omitempty"`
	CertificateIssuerName               *string                  `json:"CertificateIssuerName,omitempty"`
	CertificateIssuerCompanyName        *string                  `json:"CertificateIssuerCompanyName,omitempty"`
	CertificateIssuerOrganizationalUnit *string                  `json:"CertificateIssuerOrganizationalUnit,omitempty"`
	CertificateExpirationWarningDays    *int                     `json:"CertificateExpirationWarningDays,omitempty"`
	CheckCertificateErrors              *bool                    `json:"CheckCertificateErrors,omitempty"`
	IgnoreExternalElements              *bool                    `json:"IgnoreExternalElements,omitempty"`
	DomainGroupGuid                     *string                  `json:"DomainGroupGuid,omitempty"`
	DomainGroupGuidSpecified            *bool                    `json:"DomainGroupGuidSpecified,omitempty"`
	DnsServer                           *string                  `json:"DnsServer,omitempty"`
	DnsQuery                            *string                  `json:"DnsQuery,omitempty"`
	DnsExpectedResult                   *string                  `json:"DnsExpectedResult,omitempty"`
	DnsTestValue                        *string                  `json:"DnsTestValue,omitempty"`
	Port                                *int                     `json:"Port,omitempty"`
	IpVersion                           *string                  `json:"IpVersion,omitempty"`
	DatabaseName                        *string                  `json:"DatabaseName,omitempty"`
	NetworkAddress                      *string                  `json:"NetworkAddress,omitempty"`
	ImapSecureConnection                *bool                    `json:"ImapSecureConnection,omitempty"`
	SftpAction                          *string                  `json:"SftpAction,omitempty"`
	SftpActionPath                      *string                  `json:"SftpActionPath,omitempty"`
	HttpMethod                          *string                  `json:"HttpMethod,omitempty"`
	TlsVersion                          *string                  `json:"TlsVersion,omitempty"`
	RequestBody                         *string                  `json:"RequestBody,omitempty"`
	Url                                 *string                  `json:"Url,omitempty"`
	BrowserType                         *string                  `json:"BrowserType,omitempty"`
	BrowserWindowDimensions             *BrowserWindowDimensions `json:"BrowserWindowDimensions,omitempty"`
	UseConcurrentMonitoring             *bool                    `json:"UseConcurrentMonitoring,omitempty"`
	ConcurrentUnconfirmedErrorThreshold *int                     `json:"ConcurrentUnconfirmedErrorThreshold,omitempty"`
	ConcurrentConfirmedErrorThreshold   *int                     `json:"ConcurrentConfirmedErrorThreshold,omitempty"`
	ErrorConditions                     *[]ErrorCondition        `json:"ErrorConditions,omitempty"`
	CreatedDate                         string                   `json:"CreatedDate"`
}

type CustomMetric struct {
	Name         string `json:"Name"`
	VariableName string `json:"VariableName"`
}

type CustomField struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type SelectedCheckpoints struct {
	Checkpoints      *[]int `json:"Checkpoints"`
	Regions          *[]int `json:"Regions"`
	ExcludeLocations *[]int `json:"ExcludeLocations"`
}

type RequestHeader struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type ThrottlingOptions struct {
	ThrottlingType      string  `json:"ThrottlingType"`
	ThrottlingValue     *string `json:"ThrottlingValue"`
	ThrottlingSpeedUp   *int    `json:"ThrottlingSpeedUp"`
	ThrottlingSpeedDown *int    `json:"ThrottlingSpeedDown"`
	ThrottlingLatency   *int    `json:"ThrottlingLatency"`
}

type DnsBypass struct {
	Source string `json:"Source"`
	Target string `json:"Target"`
}

type SubStep struct {
	Name     string `json:"Name"`
	Type     string `json:"Type"`
	Url      string `json:"Url"`
	SetValue string `json:"SetValue"`
}

type BrowserWindowDimensions struct {
	IsMobile     bool   `json:"IsMobile"`
	Width        int    `json:"Width"`
	Height       int    `json:"Height"`
	PixelRatio   int    `json:"PixelRatio"`
	MobileDevice string `json:"MobileDevice"`
}

type ErrorCondition struct {
	ErrorConditionType string  `json:"ErrorConditionType"`
	Value              string  `json:"Value"`
	Percentage         *string `json:"Percentage,omitempty"`
	Level              *string `json:"Level,omitempty"`
	MatchType          *string `json:"MatchType,omitempty"`
	Effect             *string `json:"Effect,omitempty"`
}
