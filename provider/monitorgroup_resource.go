package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/constants"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
	"github.com/samber/lo"
)

type monitorGroupResource struct {
	client interfaces.IMonitorGroupClient
}

func NewMonitorGroupResource(c interfaces.IMonitorGroupClient) resource.Resource {
	return &monitorGroupResource{
		client: c,
	}
}

// Resource model maps the schema attributes to Go types.
type monitorGroupResourceModel struct {
	ID                      types.String `tfsdk:"id"`
	Description             types.String `tfsdk:"description"`
	IsAll                   types.Bool   `tfsdk:"is_all"`
	IsQuotaUnlimited        types.Bool   `tfsdk:"is_quota_unlimited"`
	BasicMonitorQuota       types.Int64  `tfsdk:"basic_monitor_quota"`
	BrowserMonitorQuota     types.Int64  `tfsdk:"browser_monitor_quota"`
	TransactionMonitorQuota types.Int64  `tfsdk:"transaction_monitor_quota"`
	ApiMonitorQuota         types.Int64  `tfsdk:"api_monitor_quota"`
}

// Metadata returns the resource type name.
func (r *monitorGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "monitorgroup"
}

// Schema defines the resource schema with defaulting and plan modifiers.
func (r *monitorGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: constants.MonitorGroupDescription,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				Required:    true,
				Description: constants.MonitorGroupDescription,
			},
			"is_quota_unlimited": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: constants.MonitorGroupDescription,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"is_all": schema.BoolAttribute{
				Computed:    true,
				Description: constants.MonitorGroupDescription,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"basic_monitor_quota": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: constants.MonitorGroupDescription,
			},
			"browser_monitor_quota": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: constants.MonitorGroupDescription,
			},
			"transaction_monitor_quota": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: constants.MonitorGroupDescription,
			},
			"api_monitor_quota": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: constants.MonitorGroupDescription,
			},
		},
	}
}

func hasQuotaValues(config monitorGroupResourceModel) bool {
	return config.BasicMonitorQuota.ValueInt64() != 0 ||
		config.BrowserMonitorQuota.ValueInt64() != 0 ||
		config.TransactionMonitorQuota.ValueInt64() != 0 ||
		config.ApiMonitorQuota.ValueInt64() != 0
}

// ValidateConfig performs cross-field validation on the configuration.
func (r *monitorGroupResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config monitorGroupResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !config.IsQuotaUnlimited.IsNull() && config.IsQuotaUnlimited.ValueBool() && hasQuotaValues(config) {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"When 'is_quota_unlimited' is true, all quota values must be 0 (the default).",
		)
	}

	if !config.IsQuotaUnlimited.IsNull() && !config.IsQuotaUnlimited.ValueBool() && !hasQuotaValues(config) {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"When 'is_quota_unlimited' is false, at least one quota must be set to a non-zero value.",
		)
	}
}

// Create handles the creation of the resource.
func (r *monitorGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan monitorGroupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	payload := models.MonitorGroupRequest{
		Description: plan.Description.ValueString(),
	}
	// Determine IsQuotaUnlimited based on the plan and quota values.

	if !plan.IsQuotaUnlimited.IsNull() && !plan.IsQuotaUnlimited.IsUnknown() {
		isQuotaUnlimited := plan.IsQuotaUnlimited.ValueBool()
		payload.IsQuotaUnlimited = &isQuotaUnlimited
	} else {
		// Optionally, ensure that payload.IsQuotaUnlimited stays nil.
		payload.IsQuotaUnlimited = nil
	}
	if !plan.BasicMonitorQuota.IsNull() && !plan.BasicMonitorQuota.IsUnknown() {
		val := int(plan.BasicMonitorQuota.ValueInt64())
		payload.BasicMonitorQuota = &val
	}
	if !plan.BrowserMonitorQuota.IsNull() && !plan.BrowserMonitorQuota.IsUnknown() {
		val := int(plan.BrowserMonitorQuota.ValueInt64())
		payload.BrowserMonitorQuota = &val
	}
	if !plan.TransactionMonitorQuota.IsNull() && !plan.TransactionMonitorQuota.IsUnknown() {
		val := int(plan.TransactionMonitorQuota.ValueInt64())
		payload.TransactionMonitorQuota = &val
	}
	if !plan.ApiMonitorQuota.IsNull() && !plan.ApiMonitorQuota.IsUnknown() {
		val := int(plan.ApiMonitorQuota.ValueInt64())
		payload.ApiMonitorQuota = &val
	}

	result, statusCode, responseBody, err := r.client.CreateMonitorGroup(payload)

	if err != nil {
		resp.Diagnostics.AddError("Error creating monitor group", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to create monitor group",
			fmt.Sprintf("HTTP status code: %d and response %v", statusCode, responseBody),
		)
		return
	}

	var state monitorGroupResourceModel
	state.ID = types.StringValue(result.MonitorGroupGuid)
	state.Description = types.StringValue(result.Description)
	if result.IsQuotaUnlimited != nil {
		state.IsQuotaUnlimited = types.BoolValue(bool(*result.IsQuotaUnlimited))
	} else {
		state.IsQuotaUnlimited = types.BoolNull()
	}
	state.IsAll = types.BoolValue(result.IsAll)
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

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Read refreshes the Terraform state with the latest resource data.
func (r *monitorGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state monitorGroupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, _, err := r.client.GetMonitorGroup(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading monitor group", err.Error())
		return
	}

	state.ID = types.StringValue(result.MonitorGroupGuid)
	state.Description = types.StringValue(result.Description)
	if result.IsQuotaUnlimited != nil {
		state.IsQuotaUnlimited = types.BoolValue(bool(*result.IsQuotaUnlimited))
	} else {
		state.IsQuotaUnlimited = types.BoolNull()
	}
	state.IsAll = types.BoolValue(result.IsAll)
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

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update applies changes to the existing resource.
func (r *monitorGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan monitorGroupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Execute the appropriate update based on the group type.
	payload := &models.MonitorGroupRequest{
		MonitorGroupGuid: lo.ToPtr(plan.ID.ValueString()),
		Description:      plan.Description.ValueString(),
	}

	if !plan.IsQuotaUnlimited.IsNull() && !plan.IsQuotaUnlimited.IsUnknown() {
		isQuotaUnlimited := plan.IsQuotaUnlimited.ValueBool()
		payload.IsQuotaUnlimited = &isQuotaUnlimited
	} else {
		payload.IsQuotaUnlimited = nil
	}

	// Set quota fields only if provided.
	if !plan.BasicMonitorQuota.IsNull() && !plan.BasicMonitorQuota.IsUnknown() {
		val := int(plan.BasicMonitorQuota.ValueInt64())
		payload.BasicMonitorQuota = &val
	}
	if !plan.BrowserMonitorQuota.IsNull() && !plan.BrowserMonitorQuota.IsUnknown() {
		val := int(plan.BrowserMonitorQuota.ValueInt64())
		payload.BrowserMonitorQuota = &val
	}
	if !plan.TransactionMonitorQuota.IsNull() && !plan.TransactionMonitorQuota.IsUnknown() {
		val := int(plan.TransactionMonitorQuota.ValueInt64())
		payload.TransactionMonitorQuota = &val
	}
	if !plan.ApiMonitorQuota.IsNull() && !plan.ApiMonitorQuota.IsUnknown() {
		val := int(plan.ApiMonitorQuota.ValueInt64())
		payload.ApiMonitorQuota = &val
	}

	statusCode, msg, err := r.client.UpdateMonitorGroup(*payload, *payload.MonitorGroupGuid)
	if err != nil {
		resp.Diagnostics.AddError("Error updating monitor group", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to update monitor group",
			fmt.Sprintf("HTTP %d: %s", statusCode, msg),
		)
		return
	}

	// Refresh the resource state by retrieving the updated monitor group details.
	monitorGroupResp, _, _ := r.client.GetMonitorGroup(plan.ID.ValueString())

	plan.ID = types.StringValue(monitorGroupResp.MonitorGroupGuid)
	plan.Description = types.StringValue(monitorGroupResp.Description)
	if monitorGroupResp.IsQuotaUnlimited != nil {
		plan.IsQuotaUnlimited = types.BoolValue(bool(*monitorGroupResp.IsQuotaUnlimited))
	} else {
		plan.IsQuotaUnlimited = types.BoolNull()
	}
	plan.IsAll = types.BoolValue(monitorGroupResp.IsAll)
	if monitorGroupResp.BasicMonitorQuota != nil {
		plan.BasicMonitorQuota = types.Int64Value(int64(*monitorGroupResp.BasicMonitorQuota))
	} else {
		plan.BasicMonitorQuota = types.Int64Null()
	}
	if monitorGroupResp.BrowserMonitorQuota != nil {
		plan.BrowserMonitorQuota = types.Int64Value(int64(*monitorGroupResp.BrowserMonitorQuota))
	} else {
		plan.BrowserMonitorQuota = types.Int64Null()
	}
	if monitorGroupResp.TransactionMonitorQuota != nil {
		plan.TransactionMonitorQuota = types.Int64Value(int64(*monitorGroupResp.TransactionMonitorQuota))
	} else {
		plan.TransactionMonitorQuota = types.Int64Null()
	}
	if monitorGroupResp.ApiMonitorQuota != nil {
		plan.ApiMonitorQuota = types.Int64Value(int64(*monitorGroupResp.ApiMonitorQuota))
	} else {
		plan.ApiMonitorQuota = types.Int64Null()
	}
	// Persist the refreshed state.
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Delete removes the resource.
func (r *monitorGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state monitorGroupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	statusCode, msg, err := r.client.DeleteMonitorGroup(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting monitor group", err.Error())
		return
	}
	if statusCode >= 300 {
		resp.Diagnostics.AddError(
			"Failed to delete monitor group",
			fmt.Sprintf("HTTP %d: %s", statusCode, msg),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *monitorGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
