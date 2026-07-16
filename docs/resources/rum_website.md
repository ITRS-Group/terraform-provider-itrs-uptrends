---
page_title: "rum_website Resource - itrs-uptrends"
subcategory: ""
description: |-
---

# itrs-uptrends_rum_website (Resource)
Manages a Real User Monitoring (RUM) website in the Uptrends monitoring platform.

A list of relevant fields and their meaning can be found in the [Uptrends API documentation (Swagger)](https://api.uptrends.com/v4/swagger/index.html?url=/v4/swagger/v1/swagger.json#/RumWebsite).

## Example usage

```terraform
# Manage a RUM website.
resource "itrs-uptrends_rum_website" "example" {
  description          = "My Marketing Website"
  url                  = "https://example.com"
  is_spa               = false
  include_url_fragment = false
  provider    = itrs-uptrends.uptrendsauthenticated
}
```

## Use cases

RUM websites define which public website(s) Uptrends should track for real-user performance and behavior metrics.

## Related resources

- [itrs-uptrends_monitor](monitor.md) - Manage synthetic monitors (separate from RUM)

## Schema

### Required

- `description` (String) The description (name) of the RUM website.
- `url` (String) The URL of the monitored website.

### Optional

- `is_spa` (Boolean) Indicates whether the website is a Single Page Application (SPA).
- `include_url_fragment` (Boolean) Specifies whether to include the URL fragment (hash) in monitoring. This can be true only if `is_spa = true`.

### Read-only

- `id` (String) The unique identifier (GUID) of the RUM website.
- `rum_script` (String) The RUM JavaScript snippet to be embedded in your website. This script enables Uptrends to collect real user monitoring data from your web pages. You should copy and paste this script into the `<head>` section of your site as described under Notes.

## Import

Import is supported using the following syntax:

```shell
# RUM website can be imported by specifying the unique identifier.
terraform import itrs-uptrends_rum_website.example "046a727c-7a90-4776-9e41-ab050bdda5dc"
```

## Notes

To collect RUM data for your website, the following script needs to be added to its web pages. This script, and the components it uses, are made available under a BSD license. The full text of this license can be found at https://hit.uptrendsdata.com/license.txt. Note that this script is specifically for tracking a single website.
1. Please make sure you have access to the code of your website, so the content of your pages can be changed.
2. Add the Javascript code to each page you want to monitor. Our advice is to add the script to the <HEAD> portion of your web pages.
3. Ensure that the updated version of your website is accessible through the domain you specified in the URL field.
4. RUM data will be tracked as soon as visitors are accessing your updated site. You should see data in the RUM overview dashboard right away.