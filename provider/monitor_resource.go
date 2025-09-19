package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/constants"
	converters "github.com/itrs-group/terraform-provider-itrs-uptrends/converters/monitor"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/helpers"
	tfsdkmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/provider/models"
)

// Ensure the implementation satisfies the expected interfaces.
var _ resource.Resource = &monitorResource{}
var _ resource.ResourceWithConfigure = &monitorResource{}
var _ resource.ResourceWithValidateConfig = &monitorResource{}

// monitorResource implements the Terraform resource.
type monitorResource struct {
	client interfaces.IMonitor
}

// NewMonitorResource creates a new instance of monitorResource using the provided client.
func NewMonitorResource(client interfaces.IMonitor) resource.Resource {
	return &monitorResource{
		client: client,
	}
}

// Metadata implements resource.Resource.
func (r *monitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "itrs-uptrends_monitor"
}

func (r *monitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Description: "Unique identifier for the monitor returned by the API",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Monitor name",
				Required:    true,
			},
			"monitor_type": schema.StringAttribute{
				Description: "Monitor type",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Http",
						"Https",
						"Connect",
						"Ping",
						"POP3",
						"SMTP",
						"FTP",
						"MySQL",
						"MSSQL",
						"WebserviceHttp",
						"WebserviceHttps",
						"Transaction",
						"DNS",
						"FullPageCheck",
						"Certificate",
						"SFTP",
						"MultiStepApi",
						"IMAP",
						"PostmanApi",
					),
				},
			},
			"generate_alert": schema.BoolAttribute{
				Description: "Flag to generate alerts",
				Required:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Monitor active status",
				Required:    true,
			},
			"check_interval": schema.Int64Attribute{
				Description: "Time interval between checks",
				Required:    true,
			},
			"monitor_mode": schema.StringAttribute{
				Description: "Monitor mode",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Development",
						"Staging",
						"Production",
					),
				},
			},
			"notes": schema.StringAttribute{
				Description: "Monitor notes",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"custom_metrics": schema.ListNestedAttribute{
				Description: "List of custom metrics",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Metric name",
							Required:    true,
						},
						"variable_name": schema.StringAttribute{
							Description: "Variable name",
							Required:    true,
						},
					},
				},
			},
			"custom_fields": schema.ListNestedAttribute{
				Description: "List of custom fields",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Field name",
							Required:    true,
						},
						"value": schema.StringAttribute{
							Description: "Field value",
							Required:    true,
						},
					},
				},
			},
			"selected_checkpoints": schema.SingleNestedAttribute{
				Description: "Selected checkpoints configuration",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{
					"checkpoints": schema.ListAttribute{
						Description: "List of checkpoint IDs",
						ElementType: types.Int64Type,
						Optional:    true,
						Default:     nil,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"regions": schema.ListAttribute{
						Description: "List of region IDs",
						ElementType: types.Int64Type,
						Optional:    true,
						Default:     nil,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"exclude_locations": schema.ListAttribute{
						Description: "List of location IDs to exclude",
						ElementType: types.Int64Type,
						Optional:    true,
						Default:     nil,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"use_primary_checkpoints_only": schema.BoolAttribute{
				Description: "Flag to use only primary checkpoints",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"use_w3c_total_time": schema.BoolAttribute{
				Description: "Whether to use W3C total time",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"self_service_transaction_script": schema.StringAttribute{
				Description: "Script for self-service transactions",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"multi_step_api_transaction_script": schema.StringAttribute{
				Description: "Multi-step API transaction script",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"block_google_analytics": schema.BoolAttribute{
				Description: "Block Google Analytics",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"block_uptrends_rum": schema.BoolAttribute{
				Description: "Block Uptrends RUM",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"block_urls": schema.ListAttribute{
				Description: "List of URLs to block",
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
			"request_headers": schema.ListNestedAttribute{
				Description: "List of request headers",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Header name",
							Required:    true,
						},
						"value": schema.StringAttribute{
							Description: "Header value",
							Required:    true,
						},
					},
				},
			},
			"user_agent": schema.StringAttribute{
				Description: "User agent string",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"username": schema.StringAttribute{
				Description: "Username for authentication",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"password_wo": schema.StringAttribute{
				Description: "Password for authentication",
				Sensitive:   true,
				WriteOnly:   true,
				Optional:    true,
			},
			"name_for_phone_alerts": schema.StringAttribute{
				Description: "Name for phone alerts",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"authentication_type": schema.StringAttribute{
				Description: "Authentication type",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"None",
						"Basic",
						"NTLM",
						"Digest",
					),
				},
			},
			"throttling_options": schema.SingleNestedAttribute{
				Description: "Throttling options",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{
					"throttling_type": schema.StringAttribute{
						Description: "Type of throttling",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Inactive",
								"Browser",
								"Simulated",
							),
						},
					},
					"throttling_value": schema.StringAttribute{
						Description: "Throttling value",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Adsl",
								"Fiber",
								"Cable",
								"2G",
								"3G",
								"4G",
							),
						},
					},
					"throttling_speed_up": schema.Int64Attribute{
						Description: "Throttling speed up",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"throttling_speed_down": schema.Int64Attribute{
						Description: "Throttling speed down",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"throttling_latency": schema.Int64Attribute{
						Description: "Throttling latency",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"dns_bypasses": schema.ListNestedAttribute{
				Description: "DNS bypass rules",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source": schema.StringAttribute{
							Description: "DNS bypass source",
							Required:    true,
						},
						"target": schema.StringAttribute{
							Description: "DNS bypass target",
							Required:    true,
						},
					},
				},
			},
			"certificate_name": schema.StringAttribute{
				Description: "Certificate name",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate_organization": schema.StringAttribute{
				Description: "Certificate organization",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate_organizational_unit": schema.StringAttribute{
				Description: "Certificate organizational unit",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate_serial_number": schema.StringAttribute{
				Description: "Certificate serial number",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate_fingerprint": schema.StringAttribute{
				Description: "Certificate fingerprint",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate_issuer_name": schema.StringAttribute{
				Description: "Certificate issuer name",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate_issuer_company_name": schema.StringAttribute{
				Description: "Certificate issuer company name",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate_issuer_organizational_unit": schema.StringAttribute{
				Description: "Certificate issuer organizational unit",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"certificate_expiration_warning_days": schema.Int64Attribute{
				Description: "Certificate expiration warning days",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"check_certificate_errors": schema.BoolAttribute{
				Description: "Check certificate errors",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"ignore_external_elements": schema.BoolAttribute{
				Description: "Ignore external elements",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"domain_group_guid": schema.StringAttribute{
				Description: "Domain group GUID",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain_group_guid_specified": schema.BoolAttribute{
				Description: "Flag indicating if domain group GUID is specified",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"dns_server": schema.StringAttribute{
				Description: "DNS server",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"dns_query": schema.StringAttribute{
				Description: "DNS query",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ARecord",
						"CnameRecord",
						"MxRecord",
						"NsRecord",
						"TxtRecord",
						"SoaRecord",
						"RootServer",
						"AaaaRecord",
						"SrvRecord",
					),
				},
			},
			"dns_expected_result": schema.StringAttribute{
				Description: "DNS expected result",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"dns_test_value": schema.StringAttribute{
				Description: "DNS test value",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"port": schema.Int64Attribute{
				Description: "Port number",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"ip_version": schema.StringAttribute{
				Description: "IP version (IpV4 or IpV6)",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"IpV4",
						"IpV6",
					),
				},
			},
			"database_name": schema.StringAttribute{
				Description: "Database name",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_address": schema.StringAttribute{
				Description: "Network address",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"imap_secure_connection": schema.BoolAttribute{
				Description: "IMAP secure connection",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"sftp_action": schema.StringAttribute{
				Description: "SFTP action",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ConnectOnly",
						"TestFileExists",
						"TestFileDoesNotExist",
						"DownLoadFile",
					),
				},
			},
			"sftp_action_path": schema.StringAttribute{
				Description: "SFTP action path",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"http_method": schema.StringAttribute{
				Description: "HTTP method",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Get",
						"Post",
					),
				},
			},
			"http_version": schema.StringAttribute{
				Description: "HTTP version",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Negotiate",
						"Http1_1",
						"Http2",
						"Http3",
					),
				},
			},
			"tls_version": schema.StringAttribute{
				Description: "TLS version",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Tls12",
						"Tls11",
						"Tls10",
						"Tls12_Tls11",
						"Tls11_Tls10",
						"Tls12_Tls11_Tls10",
						"Tls12_Tls11_Tls10_WithFallback",
						"Tls13",
					),
				},
			},
			"request_body": schema.StringAttribute{
				Description: "Request body",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"url": schema.StringAttribute{
				Description: "URL to monitor",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"browser_type": schema.StringAttribute{
				Description: "Browser type",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Chrome",
						"Edge",
					),
				},
			},
			"browser_window_dimensions": schema.SingleNestedAttribute{
				Description: "Browser window dimensions",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{
					"is_mobile": schema.BoolAttribute{
						Description: "Is mobile device",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"width": schema.Int64Attribute{
						Description: "Window width",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"height": schema.Int64Attribute{
						Description: "Window height",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"pixel_ratio": schema.Int64Attribute{
						Description: "Pixel ratio",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"mobile_device": schema.StringAttribute{
						Description: "Mobile device name",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"use_concurrent_monitoring": schema.BoolAttribute{
				Description: "Use concurrent monitoring",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"concurrent_unconfirmed_error_threshold": schema.Int64Attribute{
				Description: "Concurrent unconfirmed error threshold",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"concurrent_confirmed_error_threshold": schema.Int64Attribute{
				Description: "Concurrent confirmed error threshold",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"error_conditions": schema.ListNestedAttribute{
				Description: "List of error conditions",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"error_condition_type": schema.StringAttribute{
							Description: "Error condition type",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"LoadTimeLimit1",
									"LoadTimeLimit2",
									"TotalMinBytes",
									"TotalMaxBytes",
									"ContentMatch",
									"ConsoleContentMatch",
									"ConsoleLevel",
									"HttpStatus",
									"PageElementMaxSizeWithPercentage",
									"PageElementFailedWithPercentage",
									"PageElementUrlMatch",
									"TimeToFirstByteMaximum",
									"RequestStartMaximum",
									"DomCompleteMaximum",
									"DomInteractiveMaximum",
									"UseRecommendedCoreWebVitals",
									"FirstContentfulPaintMaximum",
									"LargestContentfulPaintMaximum",
									"TimeToInteractiveMaximum",
									"CumulativeLayoutShiftMaximum",
									"TotalBlockingTimeMaximum",
								),
							},
						},
						"value": schema.StringAttribute{
							Description: "Error condition value",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"percentage": schema.StringAttribute{
							Description: "Error condition percentage",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"level": schema.StringAttribute{
							Description: "Error condition level",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"Info",
									"Warning",
									"Error",
								),
							},
						},
						"match_type": schema.StringAttribute{
							Description: "Error condition match type",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"ErrorWhenContains",
									"ErrorWhenNotContains",
									"ErrorWhenContainsRegex",
									"ErrorWhenNotContainsRegex",
								),
							},
						},
						"effect": schema.StringAttribute{
							Description: "Error condition effect",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"Indicate",
									"Error",
								),
							},
						},
					},
				},
			},
			"created_date": schema.StringAttribute{
				Description: "Created date",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"postman_collection_json": schema.StringAttribute{
				Description: "Postman collection JSON",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"predefined_variables": schema.ListNestedAttribute{
				Description: "List of predefined variables",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description: "Variable key",
							Required:    true,
						},
						"value": schema.StringAttribute{
							Description: "Variable value",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func (r *monitorResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config tfsdkmodels.MonitorModel
	diags := req.Config.Get(ctx, &config)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	monitorType := config.MonitorType.ValueString()
	var err error

	if monitorType == "Http" || monitorType == "WebserviceHttp" || monitorType == "WebserviceHttps" {
		// Skip required/allowed attribute validation for these monitor types
		return
	}

	// Check if all required attributes are provided
	required := helpers.GetRequiredAttributes(monitorType, constants.MonitorResourceAttributes)
	err = helpers.ValidateRequiredAttributes("itrs-uptrends_monitor", monitorType, config, required)
	if err != nil {
		resp.Diagnostics.AddError("Invalid configuration", err.Error())
	}

	// Check if all attributes are allowed
	allowed := helpers.GetAllowedAttributes(monitorType, constants.MonitorResourceAttributes)
	err = helpers.ValidateAllowedAttributes("itrs-uptrends_monitor", monitorType, config, allowed)
	if err != nil {
		resp.Diagnostics.AddError("Invalid configuration", err.Error())
	}
}

func (r *monitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state tfsdkmodels.MonitorModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Fetch items from the API
	getMonitor, err := r.client.GetMonitor(state.MonitorGuid.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading monitor", err.Error())
		return
	}

	state = converters.UpdateStateConversion(getMonitor)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *monitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config tfsdkmodels.MonitorModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := converters.PayloadConversion(config)

	result, statusCode, responseBody, err := r.client.CreateMonitor(payload)

	if err != nil {
		resp.Diagnostics.AddError("Error creating monitor", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to create monitor",
			fmt.Sprintf("HTTP status code: %d and response %v", statusCode, responseBody),
		)
		return
	}

	var state = converters.UpdateStateConversion(result)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *monitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config tfsdkmodels.MonitorModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state tfsdkmodels.MonitorModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	monitorGuid := state.MonitorGuid.ValueString()

	payload := converters.PayloadConversion(config)

	statusCode, msg, err := r.client.UpdateMonitor(monitorGuid, payload)
	if err != nil {
		resp.Diagnostics.AddError("Error updating monitor", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to update monitor",
			fmt.Sprintf("HTTP %d: %s", statusCode, msg),
		)
		return
	}

	// Re-read from the server to get the latest data
	getMonitor, _ := r.client.GetMonitor(monitorGuid)

	state = converters.UpdateStateConversion(getMonitor)

	// Update state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *monitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state tfsdkmodels.MonitorModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	monitorGuid := state.MonitorGuid.ValueString()

	statusCode, msg, err := r.client.DeleteMonitor(monitorGuid)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting monitor", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to delete monitor",
			fmt.Sprintf("HTTP %d: %s", statusCode, msg),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *monitorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return // Provider not configured (e.g. during validation).
	}
	clientReturned, ok := req.ProviderData.(interfaces.IMonitor)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Provider Data",
			fmt.Sprintf("Expected client.IMonitor, got: %T", req.ProviderData),
		)
		return
	}
	r.client = clientReturned
}

func (r *monitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
