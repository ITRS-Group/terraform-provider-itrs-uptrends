---
page_title: "itrs-uptrends_rum_website Data Source - itrs-uptrends"
subcategory: ""
description: |-
  Fetch a RUM website by GUID or description (name) when you need to reference it in other resources.
---

# itrs-uptrends_rum_website (Data Source)

Look up a Real User Monitoring (RUM) website to reuse its GUID and settings without hard-coding values.

## Example Usage

```terraform
data "itrs-uptrends_rum_website" "by_description" {
  description = "My Marketing Website"
}

data "itrs-uptrends_rum_website" "by_id" {
  id = "f5a3f2a3-1234-5678-9abc-def012345678"
}
```

## Schema

### Optional
- `id` (String) RUM website GUID. Provide this or `description`.
- `description` (String) RUM website description (name). Provide this or `id`. If the description is not unique it is going to give an error.

### Read-Only
- `url` (String) The URL of the monitored website.
- `is_spa` (Boolean) Whether the website is a Single Page Application (SPA).
- `include_url_fragment` (Boolean) Whether to include the URL fragment (hash) in monitoring.
- `rum_script` (String) The RUM JavaScript snippet to be embedded in your website. This script enables Uptrends to collect real user monitoring data from your web pages. You should copy and paste this script into the `<head>` section of your site as described under Notes.

## Notes

To collect RUM data for your website, the following script needs to be added to its web pages. This script, and the components it uses, are made available under a BSD license. The full text of this license can be found at https://hit.uptrendsdata.com/license.txt. Note that this script is specifically for tracking a single website.
1. Please make sure you have access to the code of your website, so the content of your pages can be changed.
2. Add the Javascript code to each page you want to monitor. Our advice is to add the script to the <HEAD> portion of your web pages.
3. Ensure that the updated version of your website is accessible through the domain you specified in the URL field.
4. RUM data will be tracked as soon as visitors are accessing your updated site. You should see data in the RUM overview dashboard right away.