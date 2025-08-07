resource "alertdefinition" "alertdefinition123" {
	name = "Alert Definition Resource"
	is_active = true
	escalation_levels = [
		  {
		    escalation_mode        = "AlertOnErrorCount"
		    include_trace_route    = true
		    is_active              = true
		    message                = "message"
		    number_of_reminders    = 5
		    reminder_delay         = 10
		    threshold_error_count  = 1
		    threshold_minutes      = 5
		    id = 1
		  },
		  {
		    escalation_mode        = "AlertOnErrorCount"
		    include_trace_route    = true
		    is_active              = false
		    message                = ""
		    number_of_reminders    = 5
		    reminder_delay         = 10
		    threshold_error_count  = 1
		    threshold_minutes      = 5
		    id = 2  
		  },
		  {
		    escalation_mode        = "AlertOnErrorCount"
		    include_trace_route    = true
		    is_active              = false
		    message                = ""
		    number_of_reminders    = 5
		    reminder_delay         = 10
		    threshold_error_count  = 1
		    threshold_minutes      = 5
		    id = 3
		  }
		]
	provider = itrs-uptrends.uptrendsauthenticated
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = alertdefinition.alertdefinition123
  id = "${alertdefinition.alertdefinition123.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrs-uptrends.uptrendsauthenticated
}