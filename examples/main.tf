terraform {
  required_providers {
    itrs-uptrends = {
      source  = "registry.terraform.io/ITRS-Group/itrs-uptrends"
      version = "__NEW_BUILD_VERSION__"
    }
  }
}

provider "itrs-uptrends" {
  username = "username"  # Replace with API username
  password = "password"  # Replace with API password
  alias = "uptrendsauthenticated"
}
