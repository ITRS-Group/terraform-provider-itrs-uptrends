# To assign an operatorgroup to an alertdefinition, you need first to create both the alertdefinition and the itrs-uptrends_operatorgroup. Then, you can create an `itrs-uptrends_alertdefinition_operatorgroup_membership` resource that links the two.
resource "itrs-uptrends_alertdefinition_operator_membership" "alertdefinition_operator_membership123" {
  provider            = itrs-uptrends.uptrendsauthenticated
  depends_on  = [itrs-uptrends_operatorgroup.operatorgroup123, itrs-uptrends_alertdefinition.alertdefinition123]
  alertdefinition_id    = itrs-uptrends_alertdefinition.alertdefinition123.id
  operator_id = itrs-uptrends_operatorgroup.operatorgroup123.id
  escalationlevel = 3
}

resource "itrs-uptrends_alertdefinition" "alertdefinition123" {
	name = "Alert Definition Resource Test"
	is_active = true
	provider = itrs-uptrends.uptrendsauthenticated
}

resource "itrs-uptrends_operatorgroup" "operatorgroup123" {
 provider    = itrs-uptrends.uptrendsauthenticated
 description = "Operator Group Description"
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = itrs-uptrends_alertdefinition_operator_membership.alertdefinition_operator_membership_imported
  id = "${itrs-uptrends_alertdefinition.alertdefinition123.id}:${itrs-uptrends_operatorgroup.operatorgroup123.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc:046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrs-uptrends.uptrendsauthenticated
}