resource "monitorgroup" "limited_quota_group" {
  description = "Monitor group with quota"
  is_quota_unlimited = false
  basic_monitor_quota = 2 
  browser_monitor_quota = 2 
  transaction_monitor_quota = 2 
  api_monitor_quota = 2 
  provider = itrs-uptrends.uptrendsauthenticated
}

resource "monitorgroup" "unlimited_quota_group" {
  description = "Monitor group with unlimited quota"
  is_quota_unlimited = true
  provider = itrs-uptrends.uptrendsauthenticated
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = monitorgroup.monitorgroup_imported
  id = "${monitorgroup.unlimited_quota_group.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrs-uptrends.uptrendsauthenticated
}