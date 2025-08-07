# To assign an operatorgroup to an alertdefinition, you need first to create both the alertdefinition and the operatorgroup. Then, you can create an `alertdefinition_operatorgroup_membership` resource that links the two.
resource "alertdefinition_operator_membership" "alertdefinition_operator_membership123" {
  provider            = itrs-uptrends.uptrendsauthenticated
  depends_on  = [operatorgroup.operatorgroup123, alertdefinition.alertdefinition123]
  alertdefinition_id    = alertdefinition.alertdefinition123.id
  operator_id = operatorgroup.operatorgroup123.id
  escalationlevel = 3
}

resource "alertdefinition" "alertdefinition123" {
	name = "Alert Definition Resource Test"
	is_active = true
	provider = itrs-uptrends.uptrendsauthenticated
}

resource "operatorgroup" "operatorgroup123" {
 provider    = itrs-uptrends.uptrendsauthenticated
 description = "Operator Group Description"
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = alertdefinition_operator_membership.alertdefinition_operator_membership_imported
  id = "${alertdefinition.alertdefinition123.id}:${operatorgroup.operatorgroup123.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc:046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrs-uptrends.uptrendsauthenticated
}