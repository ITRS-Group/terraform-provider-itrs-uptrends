# To assign a monitor group to an alert definition, you need first to create both the alert definition and the itrs-uptrends_monitorgroup. Then, you can create a `itrs-uptrends_alertdefinition_monitorgroup_membership` resource that links the two.
resource "itrs-uptrends_alertdefinition_monitorgroup_membership" "alertdefinition_monitorgroup_membership_example" {
  provider            = itrs-uptrends.uptrendsauthenticated
  alertdefinition_id    = itrs-uptrends_alertdefinition.alertdefinition_example.id
  monitorgroup_id = itrs-uptrends_monitorgroup.monitorgroup_example.id
  depends_on = [itrs-uptrends_alertdefinition.alertdefinition_example, itrs-uptrends_monitorgroup.monitorgroup_example]
}

# Import example:
import {
  to = itrs-uptrends_alertdefinition_monitorgroup_membership.alertdefinition_monitorgroup_membership_imported
  id = "${itrs-uptrends_alertdefinition.alertdefinition_example.id}:${itrs-uptrends_monitorgroup.monitorgroup_example.id}"
  provider          = itrs-uptrends.uptrendsauthenticated
}

resource "itrs-uptrends_alertdefinition" "alertdefinition_example" {
	name = "Alert Definition Resource Test"
	is_active = true
	provider = itrs-uptrends.uptrendsauthenticated
}

resource "itrs-uptrends_monitorgroup" "monitorgroup_example" {
  description = "Monitor Group Resource Test"
  provider = itrs-uptrends.uptrendsauthenticated
}
