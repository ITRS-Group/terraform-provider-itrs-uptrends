package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/general"
)

// Ensure alertdefinitionResource implements resource.Resource.
var _ resource.Resource = &alertdefinitionResource{}

type alertdefinitionResource struct {
	client interfaces.IAlertDefinition
}

// NewAlertdefinitionResource returns a new instance of the resource.
func NewAlertdefinitionResource(c interfaces.IAlertDefinition) resource.Resource {
	return &alertdefinitionResource{
		client: c,
	}
}

// Metadata sets the resource type name.
func (r *alertdefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "alertdefinition"
}

// Schema defines the schema for the alertdefinition resource.
func (r *alertdefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: general.AlertDefinitionDescription,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: general.AlertDefinitionDescription,
			},
			"is_active": schema.BoolAttribute{
				Required:    true,
				Description: general.AlertDefinitionDescription,
			},
		},
		Blocks: map[string]schema.Block{
			"escalation_level": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"threshold_error_count": schema.Int64Attribute{
							Required:    true,
							Description: general.AlertDefinitionDescription,
						},
						"is_active": schema.BoolAttribute{
							Required:    true,
							Description: general.AlertDefinitionDescription,
						},
						"number_of_reminders": schema.Int64Attribute{
							Required:    true,
							Description: general.AlertDefinitionDescription,
						},
						"reminder_delay": schema.Int64Attribute{
							Required:    true,
							Description: general.AlertDefinitionDescription,
						},
						"include_trace_route": schema.BoolAttribute{
							Required:    true,
							Description: general.AlertDefinitionDescription,
						},
					},
				},
				// Enforce exactly three escalation_level blocks.
				Validators: []validator.List{
					listvalidator.SizeBetween(3, 3),
				},
			},
		},
	}
}

// alertDefinitionResourceModel maps resource schema data.
type alertDefinitionResourceModel struct {
	ID               types.String                   `tfsdk:"id"`
	Name             types.String                   `tfsdk:"name"`
	IsActive         types.Bool                     `tfsdk:"is_active"`
	EscalationLevels []escalationLevelResourceModel `tfsdk:"escalation_level"`
}

// escalationLevelResourceModel maps the nested escalation_level block.
type escalationLevelResourceModel struct {
	ThresholdErrorCount types.Int64 `tfsdk:"threshold_error_count"`
	IsActive            types.Bool  `tfsdk:"is_active"`
	NumberOfReminders   types.Int64 `tfsdk:"number_of_reminders"`
	ReminderDelay       types.Int64 `tfsdk:"reminder_delay"`
	IncludeTraceRoute   types.Bool  `tfsdk:"include_trace_route"`
}

// Create handles the creation of a new alertdefinition resource.
func (r *alertdefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan alertDefinitionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the alert definition.
	createReq := models.CreateAlertDefinitionRequest{
		AlertName: plan.Name.ValueString(),
		IsActive:  plan.IsActive.ValueBool(),
	}
	alertDefItem, err := r.client.CreateAlertDefinition(createReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating alert definition",
			fmt.Sprintf("Could not create alert definition: %s", err),
		)
		return
	}

	// Patch each of the three escalation levels.
	for idx, level := range plan.EscalationLevels {
		patchReq := models.PatchEscalationLevelRequest{
			ThresholdErrorCount: int(level.ThresholdErrorCount.ValueInt64()),
			IsActive:            level.IsActive.ValueBool(),
			NumberOfReminders:   int(level.NumberOfReminders.ValueInt64()),
			ReminderDelay:       int(level.ReminderDelay.ValueInt64()),
			IncludeTraceRoute:   level.IncludeTraceRoute.ValueBool(),
		}
		// Assuming escalation level IDs start at 1.
		if err := r.client.PatchEscalationLevel(alertDefItem.AlertDefinitionGuid, idx+1, patchReq); err != nil {
			resp.Diagnostics.AddError(
				"Error patching escalation level",
				fmt.Sprintf("Could not patch escalation level %d: %s", idx+1, err),
			)
			return
		}
	}

	// Update plan with values returned from the API.
	plan.ID = types.StringValue(alertDefItem.AlertDefinitionGuid)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read refreshes the Terraform state with the latest data.
func (r *alertdefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state alertDefinitionResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	alertDefItem, err := r.client.GetAlertDefinition(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading alert definition",
			fmt.Sprintf("Could not read alert definition %s: %s", state.ID.ValueString(), err),
		)
		return
	}

	// Update state fields.
	state.Name = types.StringValue(alertDefItem.AlertName)
	state.IsActive = types.BoolValue(alertDefItem.IsActive)

	// Read escalation levels.
	escalationLevels, err := r.client.GetEscalationLevels(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading escalation levels",
			fmt.Sprintf("Could not read escalation levels for alert definition %s: %s", state.ID.ValueString(), err),
		)
		return
	}
	if len(escalationLevels) != 3 {
		resp.Diagnostics.AddError(
			"Unexpected number of escalation levels",
			fmt.Sprintf("Expected 3 escalation levels, got %d", len(escalationLevels)),
		)
		return
	}

	var levels []escalationLevelResourceModel
	for _, lvl := range escalationLevels {
		levelModel := escalationLevelResourceModel{
			ThresholdErrorCount: types.Int64Value(int64(lvl.ThresholdErrorCount)),
			IsActive:            types.BoolValue(lvl.IsActive),
			NumberOfReminders:   types.Int64Value(int64(lvl.NumberOfReminders)),
			ReminderDelay:       types.Int64Value(int64(lvl.ReminderDelay)),
			IncludeTraceRoute:   types.BoolValue(lvl.IncludeTraceRoute),
		}
		levels = append(levels, levelModel)
	}
	state.EscalationLevels = levels

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

// Update handles updating an existing alertdefinition resource.
func (r *alertdefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan alertDefinitionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update alert definition core fields.
	updateReq := models.UpdateAlertDefinitionRequest{
		AlertName: plan.Name.ValueString(),
		IsActive:  plan.IsActive.ValueBool(),
	}
	if err := r.client.UpdateAlertDefinition(plan.ID.ValueString(), updateReq); err != nil {
		resp.Diagnostics.AddError(
			"Error updating alert definition",
			fmt.Sprintf("Could not update alert definition %s: %s", plan.ID.ValueString(), err),
		)
		return
	}

	// Update each escalation level.
	for idx, level := range plan.EscalationLevels {
		patchReq := models.PatchEscalationLevelRequest{
			ThresholdErrorCount: int(level.ThresholdErrorCount.ValueInt64()),
			IsActive:            level.IsActive.ValueBool(),
			NumberOfReminders:   int(level.NumberOfReminders.ValueInt64()),
			ReminderDelay:       int(level.ReminderDelay.ValueInt64()),
			IncludeTraceRoute:   level.IncludeTraceRoute.ValueBool(),
		}
		if err := r.client.PatchEscalationLevel(plan.ID.ValueString(), idx+1, patchReq); err != nil {
			resp.Diagnostics.AddError(
				"Error updating escalation level",
				fmt.Sprintf("Could not update escalation level %d: %s", idx+1, err),
			)
			return
		}
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Delete removes the alertdefinition resource.
func (r *alertdefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state alertDefinitionResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteAlertDefinition(state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting alert definition",
			fmt.Sprintf("Could not delete alert definition %s: %s", state.ID.ValueString(), err),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *alertdefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
