---
page_title: "operatorgroup_permission Resource - itrs-uptrends"
subcategory: ""
description: |-
  Manages permissions for operator groups in the Uptrends monitoring platform.  
  A list of relevant fields and their meaning can be found in the [API documentation for operator group permissions](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/OperatorGroup) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/account/users/operators/operator-permissions).
---

# operatorgroup_permission (Resource)

## Example usage

```terraform
# Create an operator
resource "operatorgroup" "example_group" {
  ...
}

# Assign a permission to the operator group
resource "operatorgroup_permission" "technical_contact" {
  operatorgroup_id = operatorgroup.example_group.id
  permission       = "TechnicalContact"
  depends_on       = [operatorgroup.example_group]
  provider         = itrs-uptrends.uptrendsauthenticated
}

# Assign multiple permissions using for_each
variable "operator_permissions_list" {
  description = "List of permissions to be applied to the operator group"
  type        = list(string)
  default     = ["FinancialOperator", "TechnicalContact", "ShareDashboards"]
}

resource "operatorgroup_permission" "multiple_permissions" {
  for_each         = toset(var.operator_permissions_list)
  operatorgroup_id = operatorgroup.example_group.id
  permission       = each.value
  depends_on       = [operatorgroup.example_group]
  provider         = itrs-uptrends.uptrendsauthenticated
}
```

## Use cases

Operator group permissions enable you to control access and assign appropriate roles to different operator groups within your organization.

## Available permissions

The following permissions can be assigned to operator groups:

- `ShareDashboards` - Ability to share dashboards.
- `AllowInfra` - Access to Infra features.
- `TechnicalContact` - Operator(s) that will be contacted in case of technical changes or issues related to the account.
- `FinancialOperator` - Operator(s) that can place orders and view invoices. These operators will get notified about account expiration, reaching the SMS credits limit, and will also receive payment reminders.
- `BasicOperator` - Basic operator permissions.
- `CreateAlertDefinition` - Allow operators to create new alert definitions.
- `CreateIntegration` - Allow operators to create new integrations.
- `CreatePrivateLocations` - Allow operators to create private locations.
- `ManageMonitorTemplates` - Allow operators to manage and apply monitor templates.

**Note:** The `Administrator` permission is a special case and cannot be assigned.

## Related resources

- [operatorgroup](operatorgroup.md) - Create and manage operator groups
- [operatorgroup_membership](operatorgroup_membership.md) - Add operators to operator groups
- [operator_permission](operator_permission.md) - Manage permissions for individual operators

## Schema

### Required

- `operatorgroup_id` (String) The unique identifier of the operator group to assign the permission to.
- `permission` (String) The permission to assign. Must be one of the available permission values listed above.

### Read-only

- `id` (String) The unique identifier of the permission assignment (composite key in format `operatorgroup_id:permission`).

## Import

Import is supported using the following syntax:

```shell
# Operator group permission can be imported by specifying the composite identifier.
terraform import operatorgroup_permission.example "046a727c-7a90-4776-9e41-ab050bdda5dc:TechnicalContact"
```

## Notes

- The `operatorgroup_id` and `permission` fields are immutable and require resource replacement when changed.
- Each permission assignment is a separate resource instance.
- The resource automatically handles the composite ID format (`operatorgroup_id:permission`).
- Make sure the operator group exists before creating permission assignments.
- Use `depends_on` to ensure proper resource creation order.
