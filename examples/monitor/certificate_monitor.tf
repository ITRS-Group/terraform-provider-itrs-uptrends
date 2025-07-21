resource "monitor" "certificate_monitor" {
	name           = "Certificate monitor"
	monitor_type   = "Certificate"
	generate_alert = true
	is_active      = true
	monitor_mode   = "Production"
	check_interval = 7
	notes          = "This is a sample monitor"
	url            = "https://example.com"

	error_conditions = [
		{
		error_condition_type = "LoadTimeLimit1"
		value                = "3500"
		effect               = "Error"
		},
		{
		error_condition_type = "LoadTimeLimit2"
		value                = "5000"
		effect               = "Error"
		}
	]

	custom_fields = []

	selected_checkpoints = {}

	check_certificate_errors    = false
	ip_version                  = "IpV6"
	use_primary_checkpoints_only = false
	provider = itrs-uptrends.uptrendsauthenticated
	username = "1234"
	password_wo = "abc"
	authentication_type="Basic"
	certificate_name = ""
	certificate_organization =""
	certificate_organizational_unit = ""
	certificate_serial_number = ""
	certificate_fingerprint =""
	certificate_issuer_name = ""
	certificate_issuer_company_name=""
	certificate_issuer_organizational_unit=""
	certificate_expiration_warning_days=0
}