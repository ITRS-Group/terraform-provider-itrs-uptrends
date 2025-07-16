package tfsdkmodels

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// MonitorModel represents the Terraform resource model.
type MonitorModel struct {
	MonitorGuid                         types.String                  `tfsdk:"id"`
	Name                                types.String                  `tfsdk:"name"`
	MonitorType                         types.String                  `tfsdk:"monitor_type"`
	GenerateAlert                       types.Bool                    `tfsdk:"generate_alert"`
	IsActive                            types.Bool                    `tfsdk:"is_active"`
	CheckInterval                       types.Int64                   `tfsdk:"check_interval"`
	MonitorMode                         types.String                  `tfsdk:"monitor_mode"`
	Notes                               types.String                  `tfsdk:"notes"`
	CustomMetrics                       *[]CustomMetricModel          `tfsdk:"custom_metrics"`
	CustomFields                        *[]CustomFieldModel           `tfsdk:"custom_fields"`
	SelectedCheckpoints                 *SelectedCheckpointsModel     `tfsdk:"selected_checkpoints"`
	UsePrimaryCheckpointsOnly           types.Bool                    `tfsdk:"use_primary_checkpoints_only"`
	SelfServiceTransactionScript        types.String                  `tfsdk:"self_service_transaction_script"`
	MultiStepApiTransactionScript       types.String                  `tfsdk:"multi_step_api_transaction_script"`
	BlockGoogleAnalytics                types.Bool                    `tfsdk:"block_google_analytics"`
	BlockUptrendsRum                    types.Bool                    `tfsdk:"block_uptrends_rum"`
	BlockUrls                           types.List                    `tfsdk:"block_urls"`
	RequestHeaders                      *[]RequestHeaderModel         `tfsdk:"request_headers"`
	UserAgent                           types.String                  `tfsdk:"user_agent"`
	Username                            types.String                  `tfsdk:"username"`
	Password                            types.String                  `tfsdk:"password_wo"`
	NameForPhoneAlerts                  types.String                  `tfsdk:"name_for_phone_alerts"`
	AuthenticationType                  types.String                  `tfsdk:"authentication_type"`
	ThrottlingOptions                   *ThrottlingOptionsModel       `tfsdk:"throttling_options"`
	DnsBypasses                         *[]DnsBypassModel             `tfsdk:"dns_bypasses"`
	CertificateName                     types.String                  `tfsdk:"certificate_name"`
	CertificateOrganization             types.String                  `tfsdk:"certificate_organization"`
	CertificateOrganizationalUnit       types.String                  `tfsdk:"certificate_organizational_unit"`
	CertificateSerialNumber             types.String                  `tfsdk:"certificate_serial_number"`
	CertificateFingerprint              types.String                  `tfsdk:"certificate_fingerprint"`
	CertificateIssuerName               types.String                  `tfsdk:"certificate_issuer_name"`
	CertificateIssuerCompanyName        types.String                  `tfsdk:"certificate_issuer_company_name"`
	CertificateIssuerOrganizationalUnit types.String                  `tfsdk:"certificate_issuer_organizational_unit"`
	CertificateExpirationWarningDays    types.Int64                   `tfsdk:"certificate_expiration_warning_days"`
	CheckCertificateErrors              types.Bool                    `tfsdk:"check_certificate_errors"`
	IgnoreExternalElements              types.Bool                    `tfsdk:"ignore_external_elements"`
	DomainGroupGuid                     types.String                  `tfsdk:"domain_group_guid"`
	DomainGroupGuidSpecified            types.Bool                    `tfsdk:"domain_group_guid_specified"`
	DnsServer                           types.String                  `tfsdk:"dns_server"`
	DnsQuery                            types.String                  `tfsdk:"dns_query"`
	DnsExpectedResult                   types.String                  `tfsdk:"dns_expected_result"`
	DnsTestValue                        types.String                  `tfsdk:"dns_test_value"`
	Port                                types.Int64                   `tfsdk:"port"`
	IpVersion                           types.String                  `tfsdk:"ip_version"`
	DatabaseName                        types.String                  `tfsdk:"database_name"`
	NetworkAddress                      types.String                  `tfsdk:"network_address"`
	ImapSecureConnection                types.Bool                    `tfsdk:"imap_secure_connection"`
	SftpAction                          types.String                  `tfsdk:"sftp_action"`
	SftpActionPath                      types.String                  `tfsdk:"sftp_action_path"`
	HttpMethod                          types.String                  `tfsdk:"http_method"`
	TlsVersion                          types.String                  `tfsdk:"tls_version"`
	RequestBody                         types.String                  `tfsdk:"request_body"`
	Url                                 types.String                  `tfsdk:"url"`
	BrowserType                         types.String                  `tfsdk:"browser_type"`
	BrowserWindowDimensions             *BrowserWindowDimensionsModel `tfsdk:"browser_window_dimensions"`
	UseConcurrentMonitoring             types.Bool                    `tfsdk:"use_concurrent_monitoring"`
	ConcurrentUnconfirmedErrorThreshold types.Int64                   `tfsdk:"concurrent_unconfirmed_error_threshold"`
	ConcurrentConfirmedErrorThreshold   types.Int64                   `tfsdk:"concurrent_confirmed_error_threshold"`
	ErrorConditions                     *[]ErrorConditionModel        `tfsdk:"error_conditions"`
	CreatedDate                         types.String                  `tfsdk:"created_date"`
}

type CustomMetricModel struct {
	Name         types.String `tfsdk:"name"`
	VariableName types.String `tfsdk:"variable_name"`
}

type CustomFieldModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type SelectedCheckpointsModel struct {
	Checkpoints      types.List `tfsdk:"checkpoints"`
	Regions          types.List `tfsdk:"regions"`
	ExcludeLocations types.List `tfsdk:"exclude_locations"`
}

type ThrottlingOptionsModel struct {
	ThrottlingType      types.String `tfsdk:"throttling_type"`
	ThrottlingValue     types.String `tfsdk:"throttling_value"`
	ThrottlingSpeedUp   types.Int64  `tfsdk:"throttling_speed_up"`
	ThrottlingSpeedDown types.Int64  `tfsdk:"throttling_speed_down"`
	ThrottlingLatency   types.Int64  `tfsdk:"throttling_latency"`
}

type DnsBypassModel struct {
	Source types.String `tfsdk:"source"`
	Target types.String `tfsdk:"target"`
}

type SubStepModel struct {
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	Url      types.String `tfsdk:"url"`
	SetValue types.String `tfsdk:"set_value"`
}

type BrowserWindowDimensionsModel struct {
	IsMobile     types.Bool   `tfsdk:"is_mobile"`
	Width        types.Int64  `tfsdk:"width"`
	Height       types.Int64  `tfsdk:"height"`
	PixelRatio   types.Int64  `tfsdk:"pixel_ratio"`
	MobileDevice types.String `tfsdk:"mobile_device"`
}

type ErrorConditionModel struct {
	ErrorConditionType types.String `tfsdk:"error_condition_type"`
	Value              types.String `tfsdk:"value"`
	Percentage         types.String `tfsdk:"percentage"`
	Level              types.String `tfsdk:"level"`
	MatchType          types.String `tfsdk:"match_type"`
	Effect             types.String `tfsdk:"effect"`
}

type RequestHeaderModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
