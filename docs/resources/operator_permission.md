---
page_title: "operator_permission Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages permissions for individual operators in the Uptrends monitoring platform.  
  A list of relevant fields and their meaning can be found in the [API documentation for operator permissions](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/Operator) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/account/users/operators/operator-permissions).
---

# operator_permission (Resource)

## Example usage

```terraform
# Create an operator resource
resource "operator" "operator123"{
  ...
}

# Assign permission to an operator
resource "operator_permission" "permission123" {
  operator_id = operator.operator123.id
  permission  = "FinancialOperator"
  depends_on  = [operator.operator123]
  provider    = itrs-uptrends.uptrendsauthenticated
}

# Assign multiple permissions using for_each
variable "operator_permissions_list" {
  description = "List of permissions to be applied to the operator"
  type        = list(string)
  default     = ["FinancialOperator", "TechnicalOperator"]
}

resource "operator_permission" "for_each_example_permissions" {
  for_each    = toset(var.operator_permissions_list)
  operator_id = operator.operator123.id
  permission  = each.value
  provider    = itrs-uptrends.uptrendsauthenticated
}
```

## Use cases

Operator permissions enable fine-grained, role-based access control to ensure operators have only the permissions necessary for their responsibilities and compliance requirements.

## Available permissions

The following permissions can be assigned to operators:

- `ShareDashboards` - Ability to share dashboards.
- `AllowInfra` - Access to Infra features.
- `TechnicalContact` - Operator(s) that will be contacted in case of technical changes or issues related to the account.
- `FinancialOperator` - Operator(s) that can place orders and view invoices. These operators will get notified about account expiration, reaching the SMS credits limit, and will also receive payment reminders.
- `AccountAccess` - The default permissions for an operator to be able to view monitors and dashboards.
- `CreateAlertDefinition` - Allow operators to create new alert definitions.
- `CreateIntegration` - Allow operators to create new integrations.
- `CreatePrivateLocations` - Allow operators to create private locations.
- `ManageMonitorTemplates` - Allow operators to manage and apply monitor templates.

**Note:** The `AccountAdministrator` permission is a special case and cannot be assigned.

## Related resources

- [operator](operator.md) - Create and manage individual operators
- [operatorgroup](operatorgroup.md) - Create and manage operator groups
- [operatorgroup_permission](operatorgroup_permission.md) - Manage permissions for operator groups

## Schema

### Required

- `operator_id` (String) The unique identifier of the operator to assign the permission to.
- `permission` (String) The permission to assign. Must be one of the available permission values listed above.

### Read-only

- `id` (String) The unique identifier of the permission assignment (composite key in format `operator_id:permission`).

## Import

Import is supported using the following syntax:

```shell
# Operator permission can be imported by specifying the composite identifier.
terraform import operator_permission.example "046a727c-7a90-4776-9e41-ab050bdda5dc:FinancialOperator"

# When you create an operator, you need to import the default AccountAccess permission:
terraform import operator_permission.account_access "046a727c-7a90-4776-9e41-ab050bdda5dc:AccountAccess"
```

## Notes

- The `operator_id` and `permission` fields are immutable and require resource replacement when changed.
- Each permission assignment is a separate resource instance.
- The resource automatically handles the composite ID format (`operator_id:permission`).
- Make sure the operator exists before creating permission assignments.
- Use `depends_on` to ensure proper resource creation order.
- When you create an operator, it has a default `AccountAccess` permission that should be imported.