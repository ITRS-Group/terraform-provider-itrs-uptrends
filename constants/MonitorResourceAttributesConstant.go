package constants

import "github.com/itrs-group/terraform-provider-itrs-uptrends/helpers"

var allMonitorAttributes = []string{
	"monitor_type",
	"name",
	"check_interval",
	"use_concurrent_monitoring",
	"concurrent_unconfirmed_error_threshold",
	"concurrent_confirmed_error_threshold",
	"monitor_mode",
	"is_active",
	"generate_alert",
	"name_for_phone_alerts",
	"notes",
	"custom_fields",
	"selected_checkpoints",
	"use_primary_checkpoints_only",
}

var MonitorResourceAttributes = map[string]helpers.ResourceAttributes{
	"Http": {
		RequiredAttributes: []string{},
		OptionalAttributes: []string{},
	},
	"WebserviceHttp": {
		RequiredAttributes: []string{},
		OptionalAttributes: []string{},
	},
	"WebserviceHttps": {
		RequiredAttributes: []string{},
		OptionalAttributes: []string{},
	},
	"Https": {
		RequiredAttributes: []string{
			"url",
			"authentication_type",
		},
		OptionalAttributes: append([]string{
			"user_agent",
			"ip_version",
			"http_method",
			"tls_version",
			"username",
			"password_wo",
			"request_headers",
			"request_body",
			"check_certificate_errors",
			"error_conditions",
		}, allMonitorAttributes...),
	},
	"FullPageCheck": {
		RequiredAttributes: []string{
			"url",
			"browser_type",
			"browser_window_dimensions",
			"authentication_type",
		},
		OptionalAttributes: append([]string{
			"domain_group_guid",
			"domain_group_guid_specified",
			"ignore_external_elements",
			"error_conditions",
			"user_agent",
			"throttling_options",
			"block_google_analytics",
			"block_uptrends_rum",
			"block_urls",
			"dns_bypasses",
			"username",
			"password_wo",
			"request_headers",
		}, allMonitorAttributes...),
	},
	"Transaction": {
		RequiredAttributes: []string{
			"self_service_transaction_script",
			"browser_type",
			"browser_window_dimensions",
			"authentication_type",
		},
		OptionalAttributes: append([]string{
			"user_agent",
			"throttling_options",
			"block_google_analytics",
			"block_uptrends_rum",
			"block_urls",
			"dns_bypasses",
			"error_conditions",
			"username",
			"password_wo",
			"request_headers",
		}, allMonitorAttributes...),
	},
	"MultiStepApi": {
		RequiredAttributes: []string{
			"multi_step_api_transaction_script",
		},
		OptionalAttributes: append([]string{"custom_metrics"}, allMonitorAttributes...),
	},
	"PostmanApi": {
		RequiredAttributes: []string{},
		OptionalAttributes: append([]string{}, allMonitorAttributes...),
	},
	"DNS": {
		RequiredAttributes: []string{
			"dns_query",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"error_conditions",
			"port",
			"dns_server",
			"dns_expected_result",
			"dns_test_value",
		}, allMonitorAttributes...),
	},
	"Certificate": {
		RequiredAttributes: []string{
			"url",
			"authentication_type",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"user_agent",
			"check_certificate_errors",
			"certificate_name",
			"certificate_organization",
			"certificate_organizational_unit",
			"certificate_serial_number",
			"certificate_fingerprint",
			"certificate_issuer_name",
			"certificate_issuer_company_name",
			"certificate_issuer_organizational_unit",
			"certificate_expiration_warning_days",
			"error_conditions",
			"username",
			"password_wo",
		}, allMonitorAttributes...),
	},
	"SFTP": {
		RequiredAttributes: []string{
			"network_address",
			"sftp_action",
			"username",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"error_conditions",
			"sftp_action_path",
			"password_wo",
			"port",
		}, allMonitorAttributes...),
	},
	"FTP": {
		RequiredAttributes: []string{
			"network_address",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"error_conditions",
			"username",
			"password_wo",
			"port",
		}, allMonitorAttributes...),
	},
	"SMTP": {
		RequiredAttributes: []string{
			"network_address",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"error_conditions",
			"username",
			"password_wo",
			"port",
		}, allMonitorAttributes...),
	},
	"POP3": {
		RequiredAttributes: []string{
			"network_address",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"error_conditions",
			"username",
			"password_wo",
			"port",
		}, allMonitorAttributes...),
	},
	"IMAP": {
		RequiredAttributes: []string{
			"network_address",
			"imap_secure_connection",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"error_conditions",
			"username",
			"password_wo",
			"port",
		}, allMonitorAttributes...),
	},
	"MSSQL": {
		RequiredAttributes: []string{
			"network_address",
			"database_name",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"username",
			"password_wo",
			"error_conditions",
			"port",
		}, allMonitorAttributes...),
	},
	"MySQL": {
		RequiredAttributes: []string{
			"network_address",
			"database_name",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"username",
			"password_wo",
			"error_conditions",
			"port",
		}, allMonitorAttributes...),
	},
	"Ping": {
		RequiredAttributes: []string{
			"network_address",
			//"dns_lookup_type",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"error_conditions",
		}, allMonitorAttributes...),
	},
	"Connect": {
		RequiredAttributes: []string{
			"network_address",
		},
		OptionalAttributes: append([]string{
			"ip_version",
			"error_conditions",
			"port",
		}, allMonitorAttributes...),
	},
}
