resource "monitor" "dns_monitor" {
	name           = "DNS monitor"
	monitor_type   = "DNS"
	generate_alert = true
	is_active      = true
	monitor_mode   = "Production"
	check_interval = 5
	notes          = "This is a sample monitor"
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
	ip_version                  = "IpV6"
	use_primary_checkpoints_only = false
	dns_query = "ARecord"
	dns_expected_result = ""
	dns_server = ""
	dns_test_value =""
	port = 1
	provider = itrsuptrends.uptrendsauthenticated
}