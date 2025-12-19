package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	converters "github.com/itrs-group/terraform-provider-itrs-uptrends/converters/monitor"
	tfsdkmodels "github.com/itrs-group/terraform-provider-itrs-uptrends/provider/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &monitorDataSource{}
)

// NewMonitorDataSource constructs the monitor data source.
func NewMonitorDataSource(monitor interfaces.IMonitor) datasource.DataSource {
	return &monitorDataSource{client: monitor}
}

type monitorDataSource struct {
	client interfaces.IMonitor
}

// Metadata returns the data source type name.
func (d *monitorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_monitor"
}

// Schema defines the schema for the data source.
func (d *monitorDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Monitor GUID. Provide this or name.",
				Optional:    true,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Monitor name. Provide this or id.",
				Optional:    true,
				Computed:    true,
			},
			"monitor_type": schema.StringAttribute{
				Description: "Monitor type.",
				Computed:    true,
			},
			"generate_alert": schema.BoolAttribute{
				Description: "Flag to generate alerts.",
				Computed:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Monitor active status.",
				Computed:    true,
			},
			"check_interval": schema.Int64Attribute{
				Description: "Time interval in minutes between checks.",
				Computed:    true,
			},
			"check_interval_seconds": schema.Int64Attribute{
				Description: "Time interval in seconds between checks.",
				Computed:    true,
			},
			"monitor_mode": schema.StringAttribute{
				Description: "Monitor mode.",
				Computed:    true,
			},
			"notes": schema.StringAttribute{
				Description: "Monitor notes.",
				Computed:    true,
			},
			"custom_metrics": schema.ListNestedAttribute{
				Description: "List of custom metrics.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Metric name.",
							Computed:    true,
						},
						"variable_name": schema.StringAttribute{
							Description: "Variable name.",
							Computed:    true,
						},
					},
				},
			},
			"custom_fields": schema.ListNestedAttribute{
				Description: "List of custom fields.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Field name.",
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "Field value.",
							Computed:    true,
						},
					},
				},
			},
			"selected_checkpoints": schema.SingleNestedAttribute{
				Description: "Selected checkpoints configuration.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"checkpoints": schema.ListAttribute{
						Description: "List of checkpoint IDs.",
						ElementType: types.Int64Type,
						Computed:    true,
					},
					"regions": schema.ListAttribute{
						Description: "List of region IDs.",
						ElementType: types.Int64Type,
						Computed:    true,
					},
					"exclude_locations": schema.ListAttribute{
						Description: "List of location IDs to exclude.",
						ElementType: types.Int64Type,
						Computed:    true,
					},
				},
			},
			"use_primary_checkpoints_only": schema.BoolAttribute{
				Description: "Flag to use only primary checkpoints.",
				Computed:    true,
			},
			"use_w3c_total_time": schema.BoolAttribute{
				Description: "Whether to use W3C total time.",
				Computed:    true,
			},
			"self_service_transaction_script": schema.StringAttribute{
				Description: "Script for self-service transactions.",
				Computed:    true,
			},
			"multi_step_api_transaction_script": schema.StringAttribute{
				Description: "Multi-step API transaction script.",
				Computed:    true,
			},
			"block_google_analytics": schema.BoolAttribute{
				Description: "Block Google Analytics.",
				Computed:    true,
			},
			"block_uptrends_rum": schema.BoolAttribute{
				Description: "Block Uptrends RUM.",
				Computed:    true,
			},
			"block_urls": schema.ListAttribute{
				Description: "List of URLs to block.",
				ElementType: types.StringType,
				Computed:    true,
			},
			"request_headers": schema.ListNestedAttribute{
				Description: "List of request headers.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Header name.",
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "Header value.",
							Computed:    true,
						},
					},
				},
			},
			"user_agent": schema.StringAttribute{
				Description: "User agent string.",
				Computed:    true,
			},
			"username": schema.StringAttribute{
				Description: "Username for authentication.",
				Computed:    true,
			},
			"password_wo": schema.StringAttribute{
				Description: "Password for authentication.",
				Sensitive:   true,
				Computed:    true,
			},
			"name_for_phone_alerts": schema.StringAttribute{
				Description: "Name for phone alerts.",
				Computed:    true,
			},
			"authentication_type": schema.StringAttribute{
				Description: "Authentication type.",
				Computed:    true,
			},
			"throttling_options": schema.SingleNestedAttribute{
				Description: "Throttling options.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"throttling_type": schema.StringAttribute{
						Description: "Type of throttling.",
						Computed:    true,
					},
					"throttling_value": schema.StringAttribute{
						Description: "Throttling value.",
						Computed:    true,
					},
					"throttling_speed_up": schema.Int64Attribute{
						Description: "Throttling speed up.",
						Computed:    true,
					},
					"throttling_speed_down": schema.Int64Attribute{
						Description: "Throttling speed down.",
						Computed:    true,
					},
					"throttling_latency": schema.Int64Attribute{
						Description: "Throttling latency.",
						Computed:    true,
					},
				},
			},
			"dns_bypasses": schema.ListNestedAttribute{
				Description: "DNS bypass rules.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source": schema.StringAttribute{
							Description: "DNS bypass source.",
							Computed:    true,
						},
						"target": schema.StringAttribute{
							Description: "DNS bypass target.",
							Computed:    true,
						},
					},
				},
			},
			"certificate_name": schema.StringAttribute{
				Description: "Certificate name.",
				Computed:    true,
			},
			"certificate_organization": schema.StringAttribute{
				Description: "Certificate organization.",
				Computed:    true,
			},
			"certificate_organizational_unit": schema.StringAttribute{
				Description: "Certificate organizational unit.",
				Computed:    true,
			},
			"certificate_serial_number": schema.StringAttribute{
				Description: "Certificate serial number.",
				Computed:    true,
			},
			"certificate_fingerprint": schema.StringAttribute{
				Description: "Certificate fingerprint.",
				Computed:    true,
			},
			"certificate_issuer_name": schema.StringAttribute{
				Description: "Certificate issuer name.",
				Computed:    true,
			},
			"certificate_issuer_company_name": schema.StringAttribute{
				Description: "Certificate issuer company name.",
				Computed:    true,
			},
			"certificate_issuer_organizational_unit": schema.StringAttribute{
				Description: "Certificate issuer organizational unit.",
				Computed:    true,
			},
			"certificate_expiration_warning_days": schema.Int64Attribute{
				Description: "Certificate expiration warning days.",
				Computed:    true,
			},
			"check_certificate_errors": schema.BoolAttribute{
				Description: "Check certificate errors.",
				Computed:    true,
			},
			"ignore_external_elements": schema.BoolAttribute{
				Description: "Ignore external elements.",
				Computed:    true,
			},
			"domain_group_guid": schema.StringAttribute{
				Description: "Domain group GUID.",
				Computed:    true,
			},
			"domain_group_guid_specified": schema.BoolAttribute{
				Description: "Flag indicating if domain group GUID is specified.",
				Computed:    true,
			},
			"dns_server": schema.StringAttribute{
				Description: "DNS server.",
				Computed:    true,
			},
			"dns_query": schema.StringAttribute{
				Description: "DNS query.",
				Computed:    true,
			},
			"dns_expected_result": schema.StringAttribute{
				Description: "DNS expected result.",
				Computed:    true,
			},
			"dns_test_value": schema.StringAttribute{
				Description: "DNS test value.",
				Computed:    true,
			},
			"port": schema.Int64Attribute{
				Description: "Port number.",
				Computed:    true,
			},
			"ip_version": schema.StringAttribute{
				Description: "IP version (IpV4 or IpV6).",
				Computed:    true,
			},
			"database_name": schema.StringAttribute{
				Description: "Database name.",
				Computed:    true,
			},
			"network_address": schema.StringAttribute{
				Description: "Network address.",
				Computed:    true,
			},
			"imap_secure_connection": schema.BoolAttribute{
				Description: "IMAP secure connection.",
				Computed:    true,
			},
			"sftp_action": schema.StringAttribute{
				Description: "SFTP action.",
				Computed:    true,
			},
			"sftp_action_path": schema.StringAttribute{
				Description: "SFTP action path.",
				Computed:    true,
			},
			"http_method": schema.StringAttribute{
				Description: "HTTP method.",
				Computed:    true,
			},
			"http_version": schema.StringAttribute{
				Description: "HTTP version.",
				Computed:    true,
			},
			"tls_version": schema.StringAttribute{
				Description: "TLS version.",
				Computed:    true,
			},
			"request_body": schema.StringAttribute{
				Description: "Request body.",
				Computed:    true,
			},
			"url": schema.StringAttribute{
				Description: "URL to monitor.",
				Computed:    true,
			},
			"browser_type": schema.StringAttribute{
				Description: "Browser type.",
				Computed:    true,
			},
			"browser_window_dimensions": schema.SingleNestedAttribute{
				Description: "Browser window dimensions.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"is_mobile": schema.BoolAttribute{
						Description: "Is mobile device.",
						Computed:    true,
					},
					"width": schema.Int64Attribute{
						Description: "Window width.",
						Computed:    true,
					},
					"height": schema.Int64Attribute{
						Description: "Window height.",
						Computed:    true,
					},
					"pixel_ratio": schema.Int64Attribute{
						Description: "Pixel ratio.",
						Computed:    true,
					},
					"mobile_device": schema.StringAttribute{
						Description: "Mobile device name.",
						Computed:    true,
					},
				},
			},
			"use_concurrent_monitoring": schema.BoolAttribute{
				Description: "Use concurrent monitoring.",
				Computed:    true,
			},
			"concurrent_unconfirmed_error_threshold": schema.Int64Attribute{
				Description: "Concurrent unconfirmed error threshold.",
				Computed:    true,
			},
			"concurrent_confirmed_error_threshold": schema.Int64Attribute{
				Description: "Concurrent confirmed error threshold.",
				Computed:    true,
			},
			"error_conditions": schema.ListNestedAttribute{
				Description: "List of error conditions.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"error_condition_type": schema.StringAttribute{
							Description: "Error condition type.",
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "Error condition value.",
							Computed:    true,
						},
						"percentage": schema.StringAttribute{
							Description: "Error condition percentage.",
							Computed:    true,
						},
						"level": schema.StringAttribute{
							Description: "Error condition level.",
							Computed:    true,
						},
						"match_type": schema.StringAttribute{
							Description: "Error condition match type.",
							Computed:    true,
						},
						"effect": schema.StringAttribute{
							Description: "Error condition effect.",
							Computed:    true,
						},
					},
				},
			},
			"created_date": schema.StringAttribute{
				Description: "Created date.",
				Computed:    true,
			},
			"postman_collection_json": schema.StringAttribute{
				Description: "Postman collection JSON.",
				Computed:    true,
			},
			"predefined_variables": schema.ListNestedAttribute{
				Description: "List of predefined variables.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description: "Variable key.",
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "Variable value.",
							Computed:    true,
						},
					},
				},
			},
			"initial_monitor_group_id_wo": schema.StringAttribute{
				Description: "Initial monitor group Guid (write-only in resource).",
				Computed:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *monitorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The monitor client was not configured. This is an internal error in the provider.")
		return
	}

	var data tfsdkmodels.MonitorModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idProvided := !data.MonitorGuid.IsNull() && data.MonitorGuid.ValueString() != ""
	nameProvided := !data.Name.IsNull() && data.Name.ValueString() != ""

	switch {
	case idProvided && nameProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide only one of id or name to look up a monitor.")
		return
	case !idProvided && !nameProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide either id or name to look up a monitor.")
		return
	}

	var monitorResponse *tfsdkmodels.MonitorModel

	if idProvided {
		monitor, err := d.client.GetMonitor(data.MonitorGuid.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading monitor", err.Error())
			return
		}
		state := converters.UpdateStateConversion(monitor)
		monitorResponse = &state
	} else {
		monitors, statusCode, responseBody, err := d.client.GetMonitors()
		if err != nil {
			resp.Diagnostics.AddError("Error listing monitors", err.Error())
			return
		}
		if statusCode >= 300 {
			resp.Diagnostics.AddError(
				"Failed to list monitors",
				fmt.Sprintf("HTTP status code: %d with response body %v", statusCode, responseBody),
			)
			return
		}
		found := false
		name := data.Name.ValueString()
		for idx := range monitors {
			if strings.EqualFold(monitors[idx].Name, name) {
				if found {
					resp.Diagnostics.AddError("Monitor not unique", fmt.Sprintf("More than one monitor found with name %q", name))
					return
				}
				state := converters.UpdateStateConversion(&monitors[idx])
				monitorResponse = &state
				found = true
			}
		}
		if !found {
			resp.Diagnostics.AddError("Monitor not found", fmt.Sprintf("No monitor found with name %q", name))
			return
		}
	}

	diags = resp.State.Set(ctx, monitorResponse)
	resp.Diagnostics.Append(diags...)
}
