resource "operatorgroup_permission" "permission123" {
  operator_id    = operatorgroup.operatorgroup.id
  permission  = "FinancialOperator"
  depends_on     = [operatorgroup.operatorgroup]
  provider       = itrs-uptrends.uptrendsauthenticated
}

variable "operatorgroup_permissions_list" {
  description = "List of permissions to be applied to the operatorgroup"
  type        = list(string)
  default     = ["FinancialOperator", "TechnicalOperator"]
}

resource "operatorgroup_permission" "for_each_example_permissions" {
  for_each       = toset(var.operatorgroup_permissions_list) # Loop through the list of permissions
  operator_id    = operatorgroup.operatorgroup.id
  permission  = each.value
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = operatorgroup_permission.permission_imported
  id = "${operatorgroup.operatorgroup.id}:CreateAlertDefinition" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc:CreateAlertDefinition")
  provider          = itrs-uptrends.uptrendsauthenticated
}