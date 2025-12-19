package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
)

var _ datasource.DataSource = &monitorGroupDataSource{}

func NewMonitorGroupDataSource(client interfaces.IMonitorGroupClient) datasource.DataSource {
	return &monitorGroupDataSource{client: client}
}

type monitorGroupDataSource struct {
	client interfaces.IMonitorGroupClient
}

func (d *monitorGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_monitorgroup"
}

func (d *monitorGroupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Monitor group GUID. Provide this or description.",
				Optional:    true,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the monitor group. Provide this or id.",
				Optional:    true,
				Computed:    true,
			},
			"is_all": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this group represents all monitors.",
			},
			"is_quota_unlimited": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether monitor quotas are unlimited.",
			},
			"basic_monitor_quota": schema.Int64Attribute{
				Computed:    true,
				Description: "Basic monitor quota.",
			},
			"browser_monitor_quota": schema.Int64Attribute{
				Computed:    true,
				Description: "Browser monitor quota.",
			},
			"transaction_monitor_quota": schema.Int64Attribute{
				Computed:    true,
				Description: "Transaction monitor quota.",
			},
			"api_monitor_quota": schema.Int64Attribute{
				Computed:    true,
				Description: "API monitor quota.",
			},
		},
	}
}

type monitorGroupDataSourceModel struct {
	ID                      types.String `tfsdk:"id"`
	Description             types.String `tfsdk:"description"`
	IsAll                   types.Bool   `tfsdk:"is_all"`
	IsQuotaUnlimited        types.Bool   `tfsdk:"is_quota_unlimited"`
	BasicMonitorQuota       types.Int64  `tfsdk:"basic_monitor_quota"`
	BrowserMonitorQuota     types.Int64  `tfsdk:"browser_monitor_quota"`
	TransactionMonitorQuota types.Int64  `tfsdk:"transaction_monitor_quota"`
	ApiMonitorQuota         types.Int64  `tfsdk:"api_monitor_quota"`
}

func (d *monitorGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The monitor group client was not configured.")
		return
	}

	var data monitorGroupDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idProvided := !data.ID.IsNull() && data.ID.ValueString() != ""
	descProvided := !data.Description.IsNull() && data.Description.ValueString() != ""

	switch {
	case idProvided && descProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide only one of id or description.")
		return
	case !idProvided && !descProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide either id or description.")
		return
	}

	var state monitorGroupResourceModel

	if idProvided {
		result, respBody, err := d.client.GetMonitorGroup(data.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading monitor group", fmt.Sprintf("%v - %s", err, respBody))
			return
		}
		state.ID = types.StringValue(result.MonitorGroupGuid)
		state.Description = types.StringValue(result.Description)
		state.IsAll = types.BoolValue(result.IsAll)
		if result.IsQuotaUnlimited != nil {
			state.IsQuotaUnlimited = types.BoolValue(*result.IsQuotaUnlimited)
		} else {
			state.IsQuotaUnlimited = types.BoolNull()
		}
		if result.BasicMonitorQuota != nil {
			state.BasicMonitorQuota = types.Int64Value(int64(*result.BasicMonitorQuota))
		} else {
			state.BasicMonitorQuota = types.Int64Null()
		}
		if result.BrowserMonitorQuota != nil {
			state.BrowserMonitorQuota = types.Int64Value(int64(*result.BrowserMonitorQuota))
		} else {
			state.BrowserMonitorQuota = types.Int64Null()
		}
		if result.TransactionMonitorQuota != nil {
			state.TransactionMonitorQuota = types.Int64Value(int64(*result.TransactionMonitorQuota))
		} else {
			state.TransactionMonitorQuota = types.Int64Null()
		}
		if result.ApiMonitorQuota != nil {
			state.ApiMonitorQuota = types.Int64Value(int64(*result.ApiMonitorQuota))
		} else {
			state.ApiMonitorQuota = types.Int64Null()
		}
	} else {
		groups, statusCode, responseBody, err := d.client.GetMonitorGroups()
		if err != nil {
			resp.Diagnostics.AddError("Error listing monitor groups", err.Error())
			return
		}
		if statusCode >= 300 {
			resp.Diagnostics.AddError("Failed to list monitor groups", fmt.Sprintf("HTTP %d: %s", statusCode, responseBody))
			return
		}
		desc := data.Description.ValueString()
		found := false
		for _, g := range groups {
			if strings.EqualFold(g.Description, desc) {
				if found {
					resp.Diagnostics.AddError("Monitor group not unique", fmt.Sprintf("More than one monitor group found with description %q", desc))
					return
				}
				state.ID = types.StringValue(g.MonitorGroupGuid)
				state.Description = types.StringValue(g.Description)
				state.IsAll = types.BoolValue(g.IsAll)
				if g.IsQuotaUnlimited != nil {
					state.IsQuotaUnlimited = types.BoolValue(*g.IsQuotaUnlimited)
				} else {
					state.IsQuotaUnlimited = types.BoolNull()
				}
				if g.BasicMonitorQuota != nil {
					state.BasicMonitorQuota = types.Int64Value(int64(*g.BasicMonitorQuota))
				} else {
					state.BasicMonitorQuota = types.Int64Null()
				}
				if g.BrowserMonitorQuota != nil {
					state.BrowserMonitorQuota = types.Int64Value(int64(*g.BrowserMonitorQuota))
				} else {
					state.BrowserMonitorQuota = types.Int64Null()
				}
				if g.TransactionMonitorQuota != nil {
					state.TransactionMonitorQuota = types.Int64Value(int64(*g.TransactionMonitorQuota))
				} else {
					state.TransactionMonitorQuota = types.Int64Null()
				}
				if g.ApiMonitorQuota != nil {
					state.ApiMonitorQuota = types.Int64Value(int64(*g.ApiMonitorQuota))
				} else {
					state.ApiMonitorQuota = types.Int64Null()
				}
				found = true
			}
		}
		if !found {
			resp.Diagnostics.AddError("Monitor group not found", fmt.Sprintf("No monitor group found with description %q", desc))
			return
		}
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
