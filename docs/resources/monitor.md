---
page_title: "monitor Resource - itrs-uptrends"
subcategory: ""
description: |-

---

# itrs-uptrends_monitor (Resource)
  Manages monitors in the Uptrends monitoring platform.  
  A list of relevant fields and their meaning can be found in the [API documentation for monitors](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/Monitor) and the [Uptrends support knowledge base](https://www.uptrends.com/support/kb/api/monitor-api).

## Monitor types

The following monitor types are supported:

- `Http` - HTTP monitoring
- `Https` - HTTPS monitoring with SSL/TLS support
- `WebserviceHttp` - Web service HTTP monitoring
- `WebserviceHttps` - Web service HTTPS monitoring
- `FullPageCheck` - Full page browser-based monitoring
- `Transaction` - Transaction monitoring with browser automation
- `MultiStepApi` - Multi-step API monitoring
- `PostmanApi` - Postman API monitoring
- `DNS` - DNS monitoring
- `Certificate` - SSL/TLS certificate monitoring
- `SFTP` - SFTP monitoring
- `FTP` - FTP monitoring
- `SMTP` - SMTP monitoring
- `POP3` - POP3 monitoring
- `IMAP` - IMAP monitoring
- `MSSQL` - Microsoft SQL Server monitoring
- `MySQL` - MySQL monitoring
- `Ping` - Ping monitoring
- `Connect` - TCP connection monitoring

## Example usage - HTTPS monitor

```terraform
resource "itrs-uptrends_monitor" "https_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "HTTPS monitor"
  monitor_type      = "Https"
  generate_alert    = false
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Staging"
  notes             = "Monitoring https://example.com"
  url               = "https://example.com"
  user_agent        = "Chrome"
  username          = "1234"
  password_wo       = "abcdef"
  authentication_type = "Basic"
  http_method         = "Get"
  tls_version         = "Tls12"
  http_version        = "Negotiate"
  ip_version          = "IpV4"
  error_conditions = [
    {
      error_condition_type = "LoadTimeLimit1"
      value                = "3500"
      effect               = "Error"
    },
    {
      error_condition_type = "LoadTimeLimit2"
      value                = "5000"
      effect               = "Error"
    },
    {
      error_condition_type = "TotalMinBytes"
      value                = "12"
    },
    {
      error_condition_type = "ContentMatch"
      value                = "This is the content check option"
    }
  ]
  selected_checkpoints = {}
}
```

## Example usage - Certificate monitor

```terraform
resource "itrs-uptrends_monitor" "certificate_monitor" {
  name           = "Certificate monitor"
  monitor_type   = "Certificate"
  generate_alert = true
  is_active      = true
  monitor_mode   = "Production"
  check_interval = 10
  notes          = "This is a sample monitor"
  url            = "https://example.com"
  error_conditions = [
    {
      error_condition_type = "LoadTimeLimit1"
      value                = "3500"
      effect               = "Error"
    },
    {
      error_condition_type = "LoadTimeLimit2"
      value                = "5000"
      effect               = "Error"
    }
  ]
  custom_fields = []
  selected_checkpoints = {}
  check_certificate_errors    = false
  ip_version                  = "IpV6"
  use_primary_checkpoints_only = false
  provider = itrs-uptrends.uptrendsauthenticated
  username = "1234"
  password_wo = "abc"
  authentication_type = "Basic"
  certificate_name = ""
  certificate_organization = ""
  certificate_organizational_unit = ""
  certificate_serial_number = ""
  certificate_fingerprint = ""
  certificate_issuer_name = ""
  certificate_issuer_company_name = ""
  certificate_issuer_organizational_unit = ""
  certificate_expiration_warning_days = 0
}
```

## Example usage - DNS monitor

```terraform
resource "itrs-uptrends_monitor" "dns_monitor" {
  name     = "DNS monitor"
  monitor_type      = "DNS"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"
  ip_version        = "IpV4"
  dns_query         = "ARecord"
  dns_expected_result = "127.0.0.1"
  dns_test_value    = "test"
  dns_server        = "8.8.8.8"
  port              = 53
  provider = itrs-uptrends.uptrendsauthenticated
}
```

## Example usage - Full Page Check monitor

```terraform
resource "itrs-uptrends_monitor" "fullpagecheck_monitor" {
  name                         = "Full Page Check monitor"
  monitor_type                 = "FullPageCheck"
  generate_alert               = true
  is_active                    = true
  monitor_mode                 = "Production"
  check_interval               = 10
  notes                        = "This is a sample monitor"
  url                          = "https://example.com"
  user_agent                   = "Chrome"
  block_google_analytics       = false
  block_uptrends_rum           = false
  username                     = ""
  authentication_type          = "None"
  throttling_options = {
    throttling_type = "Inactive"
  }
  domain_group_guid_specified = true
  browser_type             = "Chrome"
  error_conditions = [
    {
      error_condition_type = "LoadTimeLimit1"
      value                = "2500"
      effect               = "Error"
    },
    {
      error_condition_type = "LoadTimeLimit2"
      value                = "5000"
      effect               = "Error"
    },
    {
      error_condition_type = "TotalMinBytes"
      value                = "12"
    },
    {
      error_condition_type = "ContentMatch"
      value                = "This is the content check option"
    }
  ]
  browser_window_dimensions = {
    is_mobile = false
    width = 1280
    height = 800
    pixel_ratio = 1
    mobile_device = "iPhone SE"
  }
  use_w3c_total_time = false
  provider = itrs-uptrends.uptrendsauthenticated
}
```

## Example usage - Transaction monitor

```terraform
resource "itrs-uptrends_monitor" "transaction_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "Transaction monitor"
  monitor_type      = "Transaction"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"
  self_service_transaction_script = jsonencode({
    steps = [{
      actions = [{
        navigate = {
          url = "https://www.example.com"
        }
        }, {
        testDocumentContent = {
          testType = "Contains"
          value    = "example text on the page"
        }
      }]
      collectPageSource = false
      name              = "Example step"
    }]
  })
  browser_type      = "Chrome"
  browser_window_dimensions = {
    is_mobile = false
    width = 1920
    height = 1080
    pixel_ratio = 1
    mobile_device = ""
  }
  authentication_type = "Basic"
  username = "test"
  password_wo = "test123"
  use_w3c_total_time = false
}
```

## Example usage - Multi-Step API monitor

```terraform
resource "itrs-uptrends_monitor" "multistepapi_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "MultiStepApi monitor"
  monitor_type      = "MultiStepApi"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"  
  multi_step_api_transaction_script = jsonencode({
    MsaSteps = [{
      AllowedTlsVersions = []
      Assertions         = []
      Authentication = {
        AuthenticationType = "None"
        Id                 = "9f9818f7-d804-48a5-a601-cb03a79515aa"
        PasswordSpecified  = false
        UserName           = ""
      }
      Body                      = ""
      BodyType                  = "Raw"
      CalculatedContentType     = ""
      Encoding                  = "Utf8"
      IgnoreCertificateErrors   = false
      IgnoreErrors              = false
      MaxAttempts               = 2
      Method                    = "GET"
      MultiPartForm             = []
      Name                      = ""
      PostResponseScript        = ""
      PreRequestScript          = ""
      RequestHeaders            = []
      RetryUntilSuccessful      = false
      RetryWaitMilliseconds     = 1000
      StepType                  = "HttpRequest"
      Url                       = "https://go.dev/doc/tutorial/"
      UseFixedClientCertificate = false
      Variables                 = []
    }]
    PredefinedVariables  = []
    UserDefinedFunctions = []
    Version              = 2
  })
}
```

## Example usage - SFTP monitor

```terraform
resource "itrs-uptrends_monitor" "sftp_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "SFTP monitor"
  monitor_type      = "SFTP"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"
  ip_version        = "IpV6"
  network_address   = "127.0.0.1"
  port              = 22
  sftp_action       = "ConnectOnly"
  username          = "user"
}
```

## Example usage - FTP monitor

```terraform
resource "itrs-uptrends_monitor" "ftp_monitor" {
  name     = "FTP monitor"
  monitor_type      = "FTP"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"
  ip_version        = "IpV4"
  network_address   = "127.0.0.1"
  port              = 21
  provider = itrs-uptrends.uptrendsauthenticated
}
```

## Example usage - SMTP monitor

```terraform
resource "itrs-uptrends_monitor" "smtp_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "SMTP monitor"
  monitor_type      = "SMTP"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"
  ip_version        = "IpV6"
  network_address   = "127.0.0.1"
  port              = 25
}
```

## Example usage - POP3 monitor

```terraform
resource "itrs-uptrends_monitor" "pop3_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "POP3 monitor"
  monitor_type      = "POP3"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"
  network_address   = "127.0.0.1"
  port              = 110
}
```

## Example usage - IMAP monitor

```terraform
resource "itrs-uptrends_monitor" "imap_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "IMAP monitor"
  monitor_type      = "IMAP"
  generate_alert    = true
  is_active         = true
  imap_secure_connection = true
  check_interval    = 10
  monitor_mode      = "Production"
  ip_version        = "IpV4"
  network_address   = "127.0.0.1"
}
```

## Example usage - MSSQL monitor

```terraform
resource "itrs-uptrends_monitor" "mssql_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "MSSQL monitor"
  monitor_type      = "MSSQL"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  database_name     = "master"
  monitor_mode      = "Production"
  network_address   = "127.0.0.1"
}
```

## Example usage - MySQL monitor

```terraform
resource "itrs-uptrends_monitor" "mysql_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "MySQL monitor"
  monitor_type      = "MySQL"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"
  network_address   = "127.0.0.1"
  database_name     = "master"
}
```

## Example usage - Ping monitor

```terraform
resource "itrs-uptrends_monitor" "ping_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "Ping monitor"
  monitor_type      = "Ping"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"
  network_address   = "127.0.0.1"
}
```

## Example usage - Connect monitor

```terraform
resource "itrs-uptrends_monitor" "connect_monitor" {
  name     = "Connect monitor"
  monitor_type      = "Connect"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"
  network_address   = "127.0.0.1"
  port              = 50
  provider = itrs-uptrends.uptrendsauthenticated
}
```

## Example usage - Connect monitor

```terraform
resource "itrs-uptrends_monitor" "postman_monitor" {
  name     = "Postman monitor"
  monitor_type      = "PostmanApi"
  generate_alert    = true
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Production"		
  postman_collection_json =  "{\n\t\"info\": {\n\t\t\"_postman_id\": \"d8f90314-8464-45b5-a24a-1d6ca618305b\",\n\t\t\"name\": \"New Collection\",\n\t\t\"schema\": \"https://schema.getpostman.com/json/collection/v2.1.0/collection.json\",\n\t\t\"_exporter_id\": \"40252431\"\n\t},\n\t\"item\": [\n\t\t{\n\t\t\t\"name\": \"http://localhost:20000/ApiAccounts\",\n\t\t\t\"protocolProfileBehavior\": {\n\t\t\t\t\"disableBodyPruning\": true\n\t\t\t},\n\t\t\t\"request\": {\n\t\t\t\t\"method\": \"GET\",\n\t\t\t\t\"header\": [],\n\t\t\t\t\"body\": {\n\t\t\t\t\t\"mode\": \"raw\",\n\t\t\t\t\t\"raw\": \"\",\n\t\t\t\t\t\"options\": {\n\t\t\t\t\t\t\"raw\": {\n\t\t\t\t\t\t\t\"language\": \"json\"\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t},\n\t\t\t\t\"url\": {\n\t\t\t\t\t\"raw\": \"http://localhost:20000/ApiAccounts\",\n\t\t\t\t\t\"protocol\": \"http\",\n\t\t\t\t\t\"host\": [\n\t\t\t\t\t\t\"localhost\"\n\t\t\t\t\t],\n\t\t\t\t\t\"port\": \"20000\",\n\t\t\t\t\t\"path\": [\n\t\t\t\t\t\t\"ApiAccounts\"\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"response\": []\n\t\t}\n\t]\n}"			
  predefined_variables = [{
			  "Key": "name",
			  "Value": "keyvalue"
			},]
  provider = itrs-uptrends.uptrendsauthenticated
}
```

## Monitor Type-Specific Attributes

### Http, WebserviceHttp, WebserviceHttps

> **Note:** Creation of `Http`, `WebserviceHttp`, and `WebserviceHttps` monitors is no longer supported.  
> These monitor types have been merged into the `Https` monitor type.
> You can still **import** existing monitors of these types and update their attributes if they were created previously.

### Https

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `url` - The URL to monitor
- `authentication_type` - Type of authentication (None, Basic, etc.)

**Optional:**
- `user_agent` - User agent string
- `ip_version` - IP version (IpV4, IpV6)
- `http_method` - HTTP method (Get, Post, etc.)
- `http_version` - HTTP version (Negotiate, HTTP/1.1, HTTP/2, HTTP/3)
- `tls_version` - TLS version (Tls12, etc.)
- `username` - Username for authentication
- `password_wo` - Password for authentication (write-only)
- `request_headers` - Custom request headers
- `request_body` - Request body content
- `check_certificate_errors` - Whether to check certificate errors
- `error_conditions` - List of error conditions and thresholds
- All common attributes

### Certificate

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `url` - The URL to check the certificate for
- `authentication_type` - Type of authentication

**Optional:**
- `check_certificate_errors` - Whether to check certificate errors
- `certificate_name` - Expected certificate name
- `certificate_organization` - Expected certificate organization
- `certificate_organizational_unit` - Expected certificate organizational unit
- `certificate_serial_number` - Expected certificate serial number
- `certificate_fingerprint` - Expected certificate fingerprint
- `certificate_issuer_name` - Expected certificate issuer name
- `certificate_issuer_company_name` - Expected certificate issuer company name
- `certificate_issuer_organizational_unit` - Expected certificate issuer organizational unit
- `certificate_expiration_warning_days` - Days before expiration to warn
- `username` - Username for authentication
- `password_wo` - Password for authentication (write-only)
- `ip_version` - IP version (IpV4, IpV6)
- `user_agent` - User agent string
- `error_conditions` - List of error conditions and thresholds
- All common attributes

### DNS

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `dns_query` - Type of DNS query (ARecord, CName, etc.)

**Optional:**
- `dns_server` - DNS server to use
- `dns_expected_result` - Expected DNS result
- `dns_test_value` - Test value for DNS query
- `port` - Port number for DNS query
- `ip_version` - IP version (IpV4, IpV6)
- `error_conditions` - List of error conditions and thresholds
- All common attributes

### FullPageCheck

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `url` - The URL to monitor
- `browser_type` - Browser type (Chrome, Firefox, etc.)
- `browser_window_dimensions` - Browser window configuration
- `authentication_type` - Type of authentication

**Optional:**
- `block_google_analytics` - Whether to block Google Analytics
- `block_uptrends_rum` - Whether to block Uptrends RUM
- `throttling_options` - Throttling configuration
- `domain_group_guid_specified` - Whether domain group is specified
- `ignore_external_elements` - Whether to ignore external elements
- `block_urls` - URLs to block
- `dns_bypasses` - DNS bypass configuration
- `username` - Username for authentication
- `password_wo` - Password for authentication (write-only)
- `request_headers` - Custom request headers
- `user_agent` - User agent string
- `error_conditions` - List of error conditions and thresholds
- `use_w3c_total_time` - Whether to use W3C total time
- All common attributes

### Transaction

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `self_service_transaction_script` - Transaction script in JSON format
- `browser_type` - Browser type (Chrome, Firefox, etc.)
- `browser_window_dimensions` - Browser window configuration
- `authentication_type` - Type of authentication

**Optional:**
- `user_agent` - User agent string
- `throttling_options` - Throttling configuration
- `block_google_analytics` - Whether to block Google Analytics
- `block_uptrends_rum` - Whether to block Uptrends RUM
- `block_urls` - URLs to block
- `dns_bypasses` - DNS bypass configuration
- `error_conditions` - List of error conditions and thresholds
- `username` - Username for authentication
- `password_wo` - Password for authentication (write-only)
- `request_headers` - Custom request headers
- `use_w3c_total_time` - Whether to use W3C total time
- All common attributes

### MultiStepApi

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `multi_step_api_transaction_script` - Multi-step API script in JSON format

**Optional:**
- `custom_metrics` - Custom metrics configuration
- All common attributes

### SFTP

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `network_address` - Network address to monitor
- `sftp_action` - SFTP action (ConnectOnly, etc.)
- `username` - Username for authentication

**Optional:**
- `ip_version` - IP version (IpV4, IpV6)
- `error_conditions` - List of error conditions and thresholds
- `sftp_action_path` - Path for SFTP action
- `password_wo` - Password for authentication (write-only)
- `port` - Port number (default: 22)
- All common attributes

### FTP

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `network_address` - Network address to monitor

**Optional:**
- `ip_version` - IP version (IpV4, IpV6)
- `error_conditions` - List of error conditions and thresholds
- `username` - Username for authentication
- `password_wo` - Password for authentication (write-only)
- `port` - Port number (default: 21)
- All common attributes

### SMTP

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `network_address` - Network address to monitor

**Optional:**
- `ip_version` - IP version (IpV4, IpV6)
- `error_conditions` - List of error conditions and thresholds
- `username` - Username for authentication
- `password_wo` - Password for authentication (write-only)
- `port` - Port number (default: 25)
- All common attributes

### POP3

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `network_address` - Network address to monitor

**Optional:**
- `ip_version` - IP version (IpV4, IpV6)
- `error_conditions` - List of error conditions and thresholds
- `username` - Username for authentication
- `password_wo` - Password for authentication (write-only)
- `port` - Port number (default: 110)
- All common attributes

### IMAP

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `network_address` - Network address to monitor
- `imap_secure_connection` - Whether to use secure connection

**Optional:**
- `ip_version` - IP version (IpV4, IpV6)
- `error_conditions` - List of error conditions and thresholds
- `username` - Username for authentication
- `password_wo` - Password for authentication (write-only)
- `port` - Port number
- All common attributes

### MSSQL

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `network_address` - Network address to monitor
- `database_name` - Database name to connect to

**Optional:**
- `ip_version` - IP version (IpV4, IpV6)
- `username` - Username for authentication
- `password_wo` - Password for authentication (write-only)
- `error_conditions` - List of error conditions and thresholds
- `port` - Port number (default: 1433)
- All common attributes

### MySQL

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `network_address` - Network address to monitor
- `database_name` - Database name to connect to

**Optional:**
- `ip_version` - IP version (IpV4, IpV6)
- `username` - Username for authentication
- `password_wo` - Password for authentication (write-only)
- `error_conditions` - List of error conditions and thresholds
- `port` - Port number
- All common attributes

### Ping

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `network_address` - Network address to monitor

**Optional:**
- `ip_version` - IP version (IpV4, IpV6)
- `error_conditions` - List of error conditions and thresholds
- All common attributes

### Connect

**Required:**
- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.
- `network_address` - Network address to monitor

**Optional:**
- `ip_version` - IP version (IpV4, IpV6)
- `error_conditions` - List of error conditions and thresholds
- `port` - Port number to connect to
- All common attributes

### PostmanApi

**Required:**
- `postman_collection_json` - The postman collection

**Optional:**
- `predefined_variables` - List of the predefined variables
- All common attributes

## Common attributes

All monitor types share these common attributes:

### Required

- `monitor_type` (String) The type of monitor. Must be one of the supported monitor types.

### Optional

- `name` (String) The name of the monitor.
- `check_interval` (Integer) The interval in minutes between checks.
- `is_active` (Boolean) Whether the monitor is active.
- `generate_alert` (Boolean) Whether to generate alerts for this monitor.
- `monitor_mode` (String) The monitor mode (Production, Staging, etc.).
- `notes` (String) Notes about the monitor.
- `custom_fields` (List) Custom fields for the monitor.
- `selected_checkpoints` (Map) Selected monitoring checkpoints.
- `use_primary_checkpoints_only` (Boolean) Whether to use only primary checkpoints.
- `use_concurrent_monitoring` (Boolean) Whether to use concurrent monitoring.
- `concurrent_unconfirmed_error_threshold` (Integer) Threshold for unconfirmed errors.
- `concurrent_confirmed_error_threshold` (Integer) Threshold for confirmed errors.
- `name_for_phone_alerts` (String) Name for phone alerts.

### Read-only

- `id` (String) The unique identifier of the monitor.
- `created_date` (String) The date when the monitor was created.

## Import

Import is supported using the following syntax:

```shell
# Monitor can be imported by specifying the unique identifier.
terraform import itrs-uptrends_monitor.example "046a727c-7a90-4776-9e41-ab050bdda5dc"
```

## Notes

- The `monitor_type` field is immutable and requires resource replacement when changed.
- Each monitor type has specific required and optional attributes.
- The resource automatically validates that all required attributes for the selected monitor type are provided.
- Write-only fields (marked with `_wo`) are sensitive and not stored in the Terraform state.
- Use `depends_on` to ensure proper resource creation order when referencing other resources.