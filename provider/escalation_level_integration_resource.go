package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ resource.Resource = &escalationLevelIntegrationResource{}

type escalationLevelIntegrationResource struct {
	client interfaces.IEscalationLevelIntegration
}

func NewEscalationLevelIntegrationResource(client interfaces.IEscalationLevelIntegration) resource.Resource {
	return &escalationLevelIntegrationResource{client: client}
}

type escalationLevelIntegrationModel struct {
	ID                        types.String `tfsdk:"id"`
	AlertDefinitionID         types.String `tfsdk:"alertdefinition_id"`
	EscalationLevelID         types.Int64  `tfsdk:"escalation_level_id"`
	IntegrationGuid           types.String `tfsdk:"integration_guid"`
	IsActive                  types.Bool   `tfsdk:"is_active_wo"`
	IsActiveVersion           types.Int64  `tfsdk:"is_active_wo_version"`
	SendOkAlerts              types.Bool   `tfsdk:"send_ok_alerts_wo"`
	SendOkAlertsVersion       types.Int64  `tfsdk:"send_ok_alerts_wo_version"`
	SendReminderAlerts        types.Bool   `tfsdk:"send_reminder_alerts_wo"`
	SendReminderAlertsVersion types.Int64  `tfsdk:"send_reminder_alerts_wo_version"`
	VariableValues            types.Map    `tfsdk:"variable_values"`
	ExtraEmailAddresses       types.List   `tfsdk:"extra_email_addresses"`
	StatusHubServiceList      types.List   `tfsdk:"status_hub_service_list"`
	IntegrationServices       types.List   `tfsdk:"integration_services"`
}

type statusHubServiceEntryModel struct {
	MonitorGuid            types.String `tfsdk:"monitor_guid"`
	IntegrationServiceGuid types.String `tfsdk:"integration_service_guid"`
}

var statusHubServiceObjectType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"monitor_guid":             types.StringType,
		"integration_service_guid": types.StringType,
	},
}

func (r *escalationLevelIntegrationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "itrs-uptrends_escalation_level_integration"
}

func (r *escalationLevelIntegrationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Composite identifier in format alertdefinition_id:escalation_level_id:integration_guid.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"alertdefinition_id": schema.StringAttribute{
				Required:    true,
				Description: "The GUID of the alert definition.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"escalation_level_id": schema.Int64Attribute{
				Required:    true,
				Description: "The escalation level ID (1-4).",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"integration_guid": schema.StringAttribute{
				Required:    true,
				Description: "The GUID of the integration to attach to the escalation level.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"is_active_wo": schema.BoolAttribute{
				Optional:    true,
				WriteOnly:   true,
				Description: "Whether the integration is active.",
			},
			"is_active_wo_version": schema.Int64Attribute{
				Optional:    true,
				Description: "Version of the is_active_wo field. Increment to re-send the value without changing other attributes.",
			},
			"send_ok_alerts_wo": schema.BoolAttribute{
				Optional:    true,
				WriteOnly:   true,
				Description: "Whether to send OK recovery alerts. Must be false for Phone and GenericWebhook integrations.",
			},
			"send_ok_alerts_wo_version": schema.Int64Attribute{
				Optional:    true,
				Description: "Version of the send_ok_alerts_wo field. Increment to re-send the value without changing other attributes.",
			},
			"send_reminder_alerts_wo": schema.BoolAttribute{
				Optional:    true,
				WriteOnly:   true,
				Description: "Whether to send reminder alerts. Must be false for GenericWebhook integrations.",
			},
			"send_reminder_alerts_wo_version": schema.Int64Attribute{
				Optional:    true,
				Description: "Version of the send_reminder_alerts_wo field. Increment to re-send the value without changing other attributes.",
			},
			"variable_values": schema.MapAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Key-value variable values for the integration.",
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.UseStateForUnknown(),
				},
			},
			"extra_email_addresses": schema.ListAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Extra email addresses for notifications.",
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
			"status_hub_service_list": schema.ListNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Status hub service mappings.",
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"monitor_guid": schema.StringAttribute{
							Required:    true,
							Description: "The GUID of the monitor.",
						},
						"integration_service_guid": schema.StringAttribute{
							Required:    true,
							Description: "The GUID of the integration service.",
						},
					},
				},
			},
			"integration_services": schema.ListAttribute{
				Computed:    true,
				Description: "Integration service GUIDs (read-only, returned by the API).",
				ElementType: types.StringType,
			},
		},
	}
}

func (r *escalationLevelIntegrationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan escalationLevelIntegrationModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config escalationLevelIntegrationModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.IsActive = config.IsActive
	plan.SendOkAlerts = config.SendOkAlerts
	plan.SendReminderAlerts = config.SendReminderAlerts

	payload := r.buildPayload(ctx, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	alertDefGuid := plan.AlertDefinitionID.ValueString()
	escalationLevelId := int(plan.EscalationLevelID.ValueInt64())

	_, err := r.client.AddIntegration(alertDefGuid, escalationLevelId, payload)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error adding integration",
			fmt.Sprintf("Could not add integration to escalation level: %s", err),
		)
		return
	}

	integrationGuid := plan.IntegrationGuid.ValueString()
	plan.ID = types.StringValue(fmt.Sprintf("%s:%d:%s", alertDefGuid, escalationLevelId, integrationGuid))

	integration := r.getIntegration(alertDefGuid, escalationLevelId, integrationGuid, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if integration != nil {
		r.mapResponseToState(ctx, &plan, integration, &resp.Diagnostics)
	}

	if !plan.IsActiveVersion.IsNull() {
		plan.IsActiveVersion = types.Int64Value(plan.IsActiveVersion.ValueInt64())
	}
	if !plan.SendOkAlertsVersion.IsNull() {
		plan.SendOkAlertsVersion = types.Int64Value(plan.SendOkAlertsVersion.ValueInt64())
	}
	if !plan.SendReminderAlertsVersion.IsNull() {
		plan.SendReminderAlertsVersion = types.Int64Value(plan.SendReminderAlertsVersion.ValueInt64())
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *escalationLevelIntegrationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state escalationLevelIntegrationModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	integration := r.getIntegration(
		state.AlertDefinitionID.ValueString(),
		int(state.EscalationLevelID.ValueInt64()),
		state.IntegrationGuid.ValueString(),
		&resp.Diagnostics,
	)
	if resp.Diagnostics.HasError() {
		return
	}
	if integration == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	r.mapResponseToState(ctx, &state, integration, &resp.Diagnostics)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *escalationLevelIntegrationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan escalationLevelIntegrationModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config escalationLevelIntegrationModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.IsActive = config.IsActive
	plan.SendOkAlerts = config.SendOkAlerts
	plan.SendReminderAlerts = config.SendReminderAlerts

	payload := r.buildPayload(ctx, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	alertDefGuid := plan.AlertDefinitionID.ValueString()
	escalationLevelId := int(plan.EscalationLevelID.ValueInt64())
	integrationGuid := plan.IntegrationGuid.ValueString()

	err := r.client.UpdateIntegration(alertDefGuid, escalationLevelId, integrationGuid, payload)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating integration",
			fmt.Sprintf("Could not update integration: %s", err),
		)
		return
	}

	integration := r.getIntegration(alertDefGuid, escalationLevelId, integrationGuid, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if integration != nil {
		r.mapResponseToState(ctx, &plan, integration, &resp.Diagnostics)
	}

	if !plan.IsActiveVersion.IsNull() {
		plan.IsActiveVersion = types.Int64Value(plan.IsActiveVersion.ValueInt64())
	}
	if !plan.SendOkAlertsVersion.IsNull() {
		plan.SendOkAlertsVersion = types.Int64Value(plan.SendOkAlertsVersion.ValueInt64())
	}
	if !plan.SendReminderAlertsVersion.IsNull() {
		plan.SendReminderAlertsVersion = types.Int64Value(plan.SendReminderAlertsVersion.ValueInt64())
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *escalationLevelIntegrationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state escalationLevelIntegrationModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.RemoveIntegration(
		state.AlertDefinitionID.ValueString(),
		int(state.EscalationLevelID.ValueInt64()),
		state.IntegrationGuid.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error removing integration",
			fmt.Sprintf("Could not remove integration: %s", err),
		)
	}
}

func (r *escalationLevelIntegrationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ":")
	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected format: alertdefinition_id:escalation_level_id:integration_guid. Got: %q", req.ID),
		)
		return
	}

	escalationLevelId, err := strconv.ParseInt(idParts[1], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid escalation_level_id",
			fmt.Sprintf("Could not parse escalation_level_id %q as integer: %s", idParts[1], err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alertdefinition_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("escalation_level_id"), escalationLevelId)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("integration_guid"), idParts[2])...)
}

// buildPayload constructs the API request from the Terraform plan.
// *Specified flags are only set to true when the user provides the corresponding field,
// because some integration types reject these fields as NotAvailable.
func (r *escalationLevelIntegrationResource) buildPayload(ctx context.Context, plan *escalationLevelIntegrationModel, diags *diag.Diagnostics) models.EscalationLevelIntegrationRequest {
	payload := models.EscalationLevelIntegrationRequest{
		IntegrationGuid: plan.IntegrationGuid.ValueString(),
	}

	if !plan.IsActive.IsNull() && !plan.IsActive.IsUnknown() {
		v := plan.IsActive.ValueBool()
		payload.IsActive = &v
	}
	if !plan.SendOkAlerts.IsNull() && !plan.SendOkAlerts.IsUnknown() {
		v := plan.SendOkAlerts.ValueBool()
		payload.SendOkAlerts = &v
	}
	if !plan.SendReminderAlerts.IsNull() && !plan.SendReminderAlerts.IsUnknown() {
		v := plan.SendReminderAlerts.ValueBool()
		payload.SendReminderAlerts = &v
	}

	if !plan.VariableValues.IsNull() && !plan.VariableValues.IsUnknown() {
		var varValues map[string]string
		d := plan.VariableValues.ElementsAs(ctx, &varValues, false)
		diags.Append(d...)
		payload.VariableValues = varValues
	}

	if !plan.ExtraEmailAddresses.IsNull() && !plan.ExtraEmailAddresses.IsUnknown() {
		specified := true
		payload.ExtraEmailAddressesSpecified = &specified
		var emails []string
		d := plan.ExtraEmailAddresses.ElementsAs(ctx, &emails, false)
		diags.Append(d...)
		payload.ExtraEmailAddresses = emails
	}

	if !plan.StatusHubServiceList.IsNull() && !plan.StatusHubServiceList.IsUnknown() {
		specified := true
		payload.StatusHubServiceListSpecified = &specified
		var entries []statusHubServiceEntryModel
		d := plan.StatusHubServiceList.ElementsAs(ctx, &entries, false)
		diags.Append(d...)
		apiEntries := make([]models.StatusHubServiceEntry, len(entries))
		for i, e := range entries {
			apiEntries[i] = models.StatusHubServiceEntry{
				MonitorGuid:            e.MonitorGuid.ValueString(),
				IntegrationServiceGuid: e.IntegrationServiceGuid.ValueString(),
			}
		}
		payload.StatusHubServiceList = apiEntries
	}

	return payload
}

func (r *escalationLevelIntegrationResource) getIntegration(alertDefGuid string, escalationLevelId int, integrationGuid string, diags *diag.Diagnostics) *models.EscalationLevelIntegrationResponse {
	result, err := r.client.GetIntegration(alertDefGuid, escalationLevelId, integrationGuid)
	if err != nil {
		diags.AddError("Error reading integration", fmt.Sprintf("Could not retrieve integration: %s", err))
		return nil
	}
	return result
}

// mapResponseToState maps API response fields onto the Terraform state.
// Fields not present in the GET response (is_active_wo, send_ok_alerts_wo, send_reminder_alerts_wo)
// are left untouched so their plan/state values are preserved.
func (r *escalationLevelIntegrationResource) mapResponseToState(ctx context.Context, state *escalationLevelIntegrationModel, integration *models.EscalationLevelIntegrationResponse, diags *diag.Diagnostics) {
	state.IntegrationGuid = types.StringValue(integration.IntegrationGuid)

	if integration.VariableValues != nil {
		mapVal, d := types.MapValueFrom(ctx, types.StringType, integration.VariableValues)
		diags.Append(d...)
		state.VariableValues = mapVal
	} else {
		state.VariableValues = types.MapNull(types.StringType)
	}

	if integration.ExtraEmailAddresses != "" {
		parts := strings.Split(integration.ExtraEmailAddresses, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		listVal, d := types.ListValueFrom(ctx, types.StringType, parts)
		diags.Append(d...)
		state.ExtraEmailAddresses = listVal
	} else {
		state.ExtraEmailAddresses = types.ListNull(types.StringType)
	}

	if len(integration.StatusHubServiceList) > 0 {
		entries := make([]statusHubServiceEntryModel, len(integration.StatusHubServiceList))
		for i, e := range integration.StatusHubServiceList {
			entries[i] = statusHubServiceEntryModel{
				MonitorGuid:            types.StringValue(e.MonitorGuid),
				IntegrationServiceGuid: types.StringValue(e.IntegrationServiceGuid),
			}
		}
		listVal, d := types.ListValueFrom(ctx, statusHubServiceObjectType, entries)
		diags.Append(d...)
		state.StatusHubServiceList = listVal
	} else {
		state.StatusHubServiceList = types.ListNull(statusHubServiceObjectType)
	}

	if len(integration.IntegrationServices) > 0 {
		listVal, d := types.ListValueFrom(ctx, types.StringType, integration.IntegrationServices)
		diags.Append(d...)
		state.IntegrationServices = listVal
	} else {
		state.IntegrationServices = types.ListNull(types.StringType)
	}
}
