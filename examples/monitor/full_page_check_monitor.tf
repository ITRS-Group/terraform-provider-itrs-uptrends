resource "monitor" "fullpagecheck_monitor" {
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
    	is_mobile= false
    	width= 1280
    	height= 800
    	pixel_ratio= 1
    	mobile_device= "iPhone SE"
    }
    provider = itrs-uptrends.uptrendsauthenticated
    }