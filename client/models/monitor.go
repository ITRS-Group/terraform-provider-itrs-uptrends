package client

// MonitorRequest is used for create and update (POST/PUT).
type MonitorRequest struct {
	MonitorGuid                         *string                  `json:"MonitorGuid,omitempty"`
	Name                                string                   `json:"Name"`
	IsActive                            bool                     `json:"IsActive"`
	GenerateAlert                       bool                     `json:"GenerateAlert"`
	CheckInterval                       *int                     `json:"CheckInterval,omitempty"`
	CheckIntervalSeconds                *int                     `json:"CheckIntervalSeconds,omitempty"`
	MonitorMode                         string                   `json:"MonitorMode"`
	CustomMetrics                       *[]CustomMetric          `json:"CustomMetrics,omitempty"`
	CustomFields                        []CustomField            `json:"CustomFields,omitempty"`
	SelectedCheckpoints                 SelectedCheckpoints      `json:"SelectedCheckpoints"`
	UsePrimaryCheckpointsOnly           *bool                    `json:"UsePrimaryCheckpointsOnly"`
	SelfServiceTransactionScript        *string                  `json:"SelfServiceTransactionScript,omitempty"`
	MonitorType                         string                   `json:"MonitorType"`
	Notes                               *string                  `json:"Notes,omitempty"`
	MultiStepApiTransactionScript       *string                  `json:"MultiStepApiTransactionScript,omitempty"`
	MsaSteps                            *[]MsaStep               `json:"MsaSteps,omitempty"`
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
	PostmanCollectionJson               *string                  `json:"PostmanCollectionJson,omitempty"`
	PredefinedVariables                 *[]PredefinedVariables   `json:"PredefinedVariables,omitempty"`
	HttpVersion                         *string                  `json:"HttpVersion,omitempty"`
	UseW3CTotalTime                     *bool                    `json:"UseW3CTotalTime,omitempty"`
}

// MonitorResponse is returned by GET and POST calls.
type MonitorResponse struct {
	MonitorGuid                         string                   `json:"MonitorGuid"`
	Name                                string                   `json:"Name"`
	IsActive                            bool                     `json:"IsActive"`
	GenerateAlert                       bool                     `json:"GenerateAlert"`
	IsLocked                            bool                     `json:"IsLocked"`
	CheckInterval                       *int                     `json:"CheckInterval,omitempty"`
	CheckIntervalSeconds                *int                     `json:"CheckIntervalSeconds,omitempty"`
	MonitorMode                         string                   `json:"MonitorMode"`
	CustomMetrics                       *[]CustomMetric          `json:"CustomMetrics,omitempty"`
	CustomFields                        []CustomField            `json:"CustomFields"`
	SelectedCheckpoints                 SelectedCheckpoints      `json:"SelectedCheckpoints"`
	UsePrimaryCheckpointsOnly           bool                     `json:"UsePrimaryCheckpointsOnly"`
	SelfServiceTransactionScript        *string                  `json:"SelfServiceTransactionScript,omitempty"`
	MonitorType                         string                   `json:"MonitorType"`
	Notes                               string                   `json:"Notes"`
	MultiStepApiTransactionScript       *string                  `json:"MultiStepApiTransactionScript,omitempty"`
	MsaSteps                            *[]MsaStep               `json:"MsaSteps,omitempty"`
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
	PostmanCollectionJson               *string                  `json:"PostmanCollectionJson,omitempty"`
	PredefinedVariables                 *[]PredefinedVariables   `json:"PredefinedVariables,omitempty"`
	HttpVersion                         *string                  `json:"HttpVersion,omitempty"`
	UseW3CTotalTime                     *bool                    `json:"UseW3CTotalTime,omitempty"`
}

type CustomMetric struct {
	Name         string `json:"Name"`
	VariableName string `json:"VariableName"`
}

type CustomField struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type PredefinedVariables struct {
	Key   string `json:"Key"`
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

// MsaStep represents one step of a MultiStepApi monitor transaction.
type MsaStep struct {
	Name                       *string                  `json:"Name,omitempty"`
	StepType                   *string                  `json:"StepType,omitempty"`
	Url                        *string                  `json:"Url,omitempty"`
	Method                     *string                  `json:"Method,omitempty"`
	Body                       *string                  `json:"Body,omitempty"`
	BodyType                   *string                  `json:"BodyType,omitempty"`
	VaultFileId                *string                  `json:"VaultFileId,omitempty"`
	MultiPartForm              *[]MsaFormPart           `json:"MultiPartForm,omitempty"`
	RequestHeaders             *[]MsaHttpHeaderValue    `json:"RequestHeaders,omitempty"`
	Variables                  *[]MsaVariableDefinition `json:"Variables,omitempty"`
	Assertions                 *[]MsaAssertion          `json:"Assertions,omitempty"`
	Authentication             *MsaAuthenticationInfo   `json:"Authentication,omitempty"`
	UseFixedClientCertificate  *bool                    `json:"UseFixedClientCertificate,omitempty"`
	ClientCertificateVaultItem *string                  `json:"ClientCertificateVaultItem,omitempty"`
	IgnoreCertificateErrors    *bool                    `json:"IgnoreCertificateErrors,omitempty"`
	Delay                      *int                     `json:"Delay,omitempty"`
	RetryUntilSuccessful       *bool                    `json:"RetryUntilSuccessful,omitempty"`
	MaxAttempts                *int                     `json:"MaxAttempts,omitempty"`
	RetryWaitMilliseconds      *int                     `json:"RetryWaitMilliseconds,omitempty"`
	PreRequestScript           *string                  `json:"PreRequestScript,omitempty"`
	PostResponseScript         *string                  `json:"PostResponseScript,omitempty"`
	CalculatedContentType      *string                  `json:"CalculatedContentType,omitempty"`
	AllowedTlsVersions         *[]string                `json:"AllowedTlsVersions,omitempty"`
	MsaStepHttpVersion         *string                  `json:"MsaStepHttpVersion,omitempty"`
}

type MsaFormPart struct {
	Type  string `json:"Type"`
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type MsaHttpHeaderValue struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type MsaAssertion struct {
	Source      string `json:"Source"`
	Property    string `json:"Property"`
	Comparison  string `json:"Comparison"`
	TargetValue string `json:"TargetValue"`
}

// MsaVariableDefinition is recursive: Arguments holds nested variable definitions
// for composite sources (e.g. Sum, Concatenation, ExecuteFunction).
type MsaVariableDefinition struct {
	Source    string                  `json:"Source"`
	Property  string                  `json:"Property"`
	Name      string                  `json:"Name"`
	Arguments []MsaVariableDefinition `json:"Arguments,omitempty"`
}

type MsaAuthenticationInfo struct {
	Id                 *string `json:"Id,omitempty"`
	AuthenticationType string  `json:"AuthenticationType"`
	UserName           *string `json:"UserName,omitempty"`
	Password           *string `json:"Password,omitempty"`
	PasswordSpecified  *bool   `json:"PasswordSpecified,omitempty"`
}
