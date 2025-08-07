---
page_title: "itrs-uptrends Provider"
subcategory: ""
description: |-
  Interact with ITRS Uptrends monitoring platform.
---

# ITRS Uptrends provider

The ITRS Uptrends provider allows you to manage monitoring resources in the Uptrends platform using Terraform.

## Example usage

```terraform
# Authentication based on the API user credentials
provider "itrs-uptrends" {
  username = "your API user username"
  password = "your API user password"
  alias    = "uptrendsauthenticated"
}
```

## Available resources

### Monitoring resources

- [monitor](resources/monitor.md) - Manage various types of monitors (HTTPS, DNS, Ping, etc.)
- [monitorgroup](resources/monitorgroup.md) - Manage monitor groups
- [monitorgroup_membership](resources/monitorgroup_membership.md) - Manage monitor group memberships

### Alert management

- [alertdefinition](resources/alertdefinition.md) - Manage alert definitions
- [alertdefinition_monitor_membership](resources/alertdefinition_monitor_membership.md) - Manage alert definition monitor memberships
- [alertdefinition_operator_membership](resources/alertdefinition_operator_membership.md) - Manage alert definition operator memberships
- [alertdefinition_operatorgroup_membership](resources/alertdefinition_operatorgroup_membership.md) - Manage alert definition operator group memberships

### User management

- [operator](resources/operator.md) - Manage operators (users)
- [operator_permission](resources/operator_permission.md) - Manage operator permissions
- [operatorgroup](resources/operatorgroup.md) - Manage operator groups
- [operatorgroup_membership](resources/operatorgroup_membership.md) - Manage operator group memberships
- [operatorgroup_permission](resources/operatorgroup_permission.md) - Manage operator group permissions

### Vault management

- [vault_section](resources/vault_section.md) - Manage vault sections
- [vault_item](resources/vault_item.md) - Manage vault items (credentials, certificates, files, etc.)

## Monitor types

The provider supports various monitor types including:

- **HTTPS** - HTTP/HTTPS website monitoring
- **DNS** - DNS resolution monitoring
- **Ping** - Network connectivity monitoring
- **Certificate** - SSL certificate monitoring
- **Full Page Check** - Complete webpage monitoring
- **FTP** - FTP service monitoring
- **IMAP** - Email service monitoring
- **MySQL** - Database monitoring
- **MSSQL** - SQL Server monitoring
- **SFTP** - SFTP service monitoring
- **SMTP** - Email server monitoring
- **POP3** - Email retrieval monitoring
- **Connect** - TCP connection monitoring
- **Transaction** - Multi-step transaction monitoring
- **Multi-step API** - API transaction monitoring
- **Postman API** - Postman collection monitoring

## Vault Item types

The vault system supports various item types:

- **CredentialSet** - Username and password storage
- **Certificate** - SSL certificate storage
- **CertificateArchive** - Password-protected certificate archives
- **File** - File data storage
- **OneTimePassword** - TOTP/HOTP configuration

## Schema

### Required

- `password` (String, Sensitive) Password for Uptrends API authentication.
- `username` (String) Username for Uptrends API authentication.
- `alias` (String) Provider alias for multiple configurations.

## Getting started

1. **Install the provider** by adding it to your Terraform configuration
2. **Configure authentication** using your Uptrends API credentials
3. **Create resources** using the available resource types
4. **Apply your configuration** to provision resources in Uptrends

## Examples

See the [examples](../examples/) directory for complete working examples of each resource type.

## Support

For issues and questions, please contact [our Support team](https://www.uptrends.com/contact).