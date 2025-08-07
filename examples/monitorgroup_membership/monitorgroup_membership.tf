# To assign a monitor to a monitor group, you need first to create both the monitor group and the monitor. Then, you can create a `monitorgroup_membership` resource that links the two.
resource "monitorgroup_membership" "monitorgroup_membership_example" {
  provider    = itrs-uptrends.uptrendsauthenticated
  monitorgroup_id = monitorgroup.monitorgroup_example.id
  monitor_id = monitor.certificate_monitor.id
  depends_on = [monitorgroup.monitorgroup_example, monitor.certificate_monitor]
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = monitorgroup_membership.monitorgroup_membership_imported
  id = "${monitor.certificate_monitor.id}:${monitorgroup.monitorgroup_example.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc:046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrs-uptrends.uptrendsauthenticated
}

resource "monitorgroup" "monitorgroup_example" {
  provider    = itrs-uptrends.uptrendsauthenticated
  description = "Monitor Group"
  is_quota_unlimited = true
}

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

