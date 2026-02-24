resource "itrs-uptrends_operator_permission" "permission123" {
  operator_id    = itrs-uptrends_operator.operator.id
  permission  = "FinancialOperator"
  depends_on     = [itrs-uptrends_operator.operator]
  provider       = itrs-uptrends.uptrendsauthenticated
}

variable "operator_permissions_list" {
  description = "List of permissions to be applied to the operator"
  type        = list(string)
  default     = ["FinancialOperator", "TechnicalOperator"]
}

resource "itrs-uptrends_operator_permission" "for_each_example_permissions" {
  for_each       = toset(var.operator_permissions_list) # Loop through the list of permissions
  operator_id    = itrs-uptrends_operator.operator.id
  permission  = each.value
}

# Import example:
# When u create an operator you need to import the operator_permission: AccountAccess.
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = itrs-uptrends_operator_permission.permission_imported
  id = "${itrs-uptrends_operator.operator.id}:AccountAccess"
  provider          = itrs-uptrends.uptrendsauthenticated
}