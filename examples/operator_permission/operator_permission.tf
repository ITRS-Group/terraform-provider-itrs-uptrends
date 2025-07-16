resource "operator_permission" "permission123" {
  operator_id    = operator.operator.id
  permission  = "FinancialOperator"
  depends_on     = [operator.operator]
  provider       = itrsuptrends.uptrendsauthenticated
}

variable "operator_permissions_list" {
  description = "List of permissions to be applied to the operator"
  type        = list(string)
  default     = ["FinancialOperator", "TechnicalOperator"]
}

resource "operator_permission" "for_each_example_permissions" {
  for_each       = toset(var.operator_permissions_list) # Loop through the list of permissions
  operator_id    = operator.operator.id
  permission  = each.value
}

# Import example:
# When u create an operator you need to import the operator_permission: AccountAccess.
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = operator_permission.permission_imported
  id = "${operator.operator.id}:AccountAccess"
  provider          = itrsuptrends.uptrendsauthenticated
}