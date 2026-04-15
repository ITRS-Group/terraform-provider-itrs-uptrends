# Uptrends Terraform Provider

Welcome to the **Uptrends Terraform Provider**, enabling you to manage your Uptrends account resources—such as monitors, monitor groups, operators, and much more—using simple, declarative Terraform configuration. With this provider, you can seamlessly integrate Uptrends monitoring and alerts into your Infrastructure as Code (IaC) workflows.

## Features

- **Automated Setup**  
  Provision new monitors and monitor groups on Uptrends directly from Terraform.
- **Centralized Configuration**  
  Configure operators, operator groups, and alert escalations alongside the rest of your infrastructure, all in code.
- **Simplified Authentication**  
  Use a single Uptrends API account for automated deployments and updates.
- **Reliable Monitoring & Alerting**  
  Keep tabs on your infrastructure uptime and quickly respond to incidents via Uptrends’ robust alert mechanisms.

## Getting Started

### Authentication

To use the Uptrends Terraform Provider, you must first create a dedicated Uptrends API account. This differs from your regular login credentials but is still tied to your main Uptrends account. You can create it via the Uptrends API Swagger interface:

1. Go to [Uptrends Swagger page](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/Register).
2. Locate and expand the **`POST /Register`** method.
3. Click **Try it out**, then **Execute**.
4. Your browser will prompt for your usual Uptrends login credentials; provide them and click **OK**.
5. Once verified, the `Response body` shows a `UserName` and `Password` pair—your new API account credentials. These can be used for all Terraform calls.

### Provider Configuration

Use your newly created API username and password in your Terraform configuration:

```hcl
provider "itrs-uptrends" {
  # Replace with your API account credentials
  username = "username"
  password = "password"
  alias    = "uptrendsauthenticated"
}
```

- **`username`** and **`password`** should be stored securely according to your own company standards for this. Tools to help you with that are described [here](https://www.hashicorp.com/en/blog/terraform-1-10-improves-handling-secrets-in-state-with-ephemeral-values) and [here](https://spacelift.io/blog/terraform-secrets). 
- **`debug`** (optional) toggles the tool for debugging, which can be helpful in troubleshooting and validation.

In addition to the `provider` configuration, you need at least one `resource` configurations. Each `resource` to choose from can be found in the [resources document](resources.md). In case you want to get started with an example running in a Docker container, you can read the [Docker Instructions](docs/DockerExample/Instructions.md).

## Having Issues or Need Assistance?

If you encounter any difficulties or have questions about this ITRS Uptrends Terraform provider, please do not hesitate to reach out. The [Uptrends contact page](https://www.uptrends.com/contact) offers direct support and further assistance.