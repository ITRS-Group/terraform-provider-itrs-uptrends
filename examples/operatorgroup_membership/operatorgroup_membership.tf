resource "itrs-uptrends_operatorgroup_membership" "operatorgroup_membership123" {
  provider            = itrs-uptrends.uptrendsauthenticated
  depends_on  = [itrs-uptrends_operator.operator123, itrs-uptrends_operatorgroup.operatorgroup123]
  operator_id = itrs-uptrends_operator.operator123.id
  operatorgroup_id    = itrs-uptrends_operatorgroup.operatorgroup123.id
}

resource "itrs-uptrends_operatorgroup" "operatorgroup123" {
  provider    = itrs-uptrends.uptrendsauthenticated
  description = "Operator Group Description"
}

resource "itrs-uptrends_operator" "operator123" {
  provider                      = itrs-uptrends.uptrendsauthenticated
  backup_email                  = ""
  default_dashboard             = "UseAccountSpecifiedDashboard"
  email                         = "operatoremail@email.com"
  full_name                     = "Firstname Lastname"
  is_on_duty                    = true
  mobile_phone                  = ""
  operator_role                 = "Unspecified"
  password_wo                      = "password123!!"
  sms_provider                  = "UseAccountSetting"
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = itrs-uptrends_operatorgroup_membership.operatorgroup_membership_imported
  id = "${itrs-uptrends_operator.operator123.id}:${itrs-uptrends_operatorgroup.operatorgroup123.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc:046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrs-uptrends.uptrendsauthenticated
}