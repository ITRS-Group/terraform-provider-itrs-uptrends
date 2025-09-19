resource "itrs-uptrends_operatorgroup" "operatorgroup123" {
  description = "Operator Group Description"
  provider    = itrs-uptrends.uptrendsauthenticated
}

# Import example:
# Import States available in the Uptrends APP for downloading as a tf file:
import {
  to = itrs-uptrends_operatorgroup.operatorgroup123
  id = "${itrs-uptrends_operatorgroup.operatorgroup123.id}" # Replace with the actual ID (e.g. "046a727c-7a90-4776-9e41-ab050bdda5dc")
  provider          = itrs-uptrends.uptrendsauthenticated
}