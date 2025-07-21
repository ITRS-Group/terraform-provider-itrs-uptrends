terraform {
  required_providers {
    itrs-uptrends = {
      source  = "registry.terraform.io/ITRS-Group/itrs-uptrends"  # Must match the address in main.go
      version = "__NEW_BUILD_VERSION__"
    }
  }
}

provider "itrs-uptrends" {
  username = "username"  # Replace with actual username
  password = "password"  # Replace with a secure method in production
  debug    = false        # Set false in production if not debugging
  alias = "uptrendsauthenticated"
}
