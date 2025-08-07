resource "monitor" "https_monitor" {
  provider = itrs-uptrends.uptrendsauthenticated
  name     = "HTTPS monitor"
  monitor_type      = "Https"
  generate_alert    = false
  is_active         = true
  check_interval    = 10
  monitor_mode      = "Staging"
  notes             = "This monitor pings https://example.com"
  url               = "https://example.com"
  user_agent        = "Chrome"
  username          = "1234"
  password_wo          = "abcdef"
  authentication_type = "Basic"
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