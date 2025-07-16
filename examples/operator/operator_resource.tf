resource "operator" "operator123" {
  provider          = itrsuptrends.uptrendsauthenticated
  backup_email      = ""
  default_dashboard = "UseAccountSpecifiedDashboard"
  email             = "operatoremail@email.com"
  full_name         = "Operator Name"
  is_on_duty        = true
  mobile_phone      = ""
  password_wo          = var.operator_password # Mandatory for create, optional for update
  operator_role     = "Unspecified"
  sms_provider      = "UseAccountSetting"
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = operator.operator123
  id = "${operator.operatorr123.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrsuptrends.uptrendsauthenticated
}