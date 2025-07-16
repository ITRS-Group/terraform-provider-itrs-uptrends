resource "alertdefinition" "alertdefinition123" {
	name = "Alert Definition Resource"
	is_active = true
	escalation_level {
	  threshold_error_count = 2
	  is_active = true
	  number_of_reminders = 10
	  reminder_delay = 20
	  include_trace_route = true
	}
	escalation_level {
	  threshold_error_count = 3
	  is_active = false
	  number_of_reminders = 11
	  reminder_delay = 20
	  include_trace_route = true
	}
	escalation_level {
	  threshold_error_count = 3
	  is_active = false
	  number_of_reminders = 12
	  reminder_delay = 20
	  include_trace_route = true
	}
	provider = itrsuptrends.uptrendsauthenticated
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = alertdefinition.alertdefinition123
  id = "${alertdefinition.alertdefinition123.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrsuptrends.uptrendsauthenticated
}