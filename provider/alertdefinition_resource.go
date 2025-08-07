package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/constants"
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
				Description: constants.AlertDefinitionDescription,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: constants.AlertDefinitionDescription,
			},
			"is_active": schema.BoolAttribute{
				Required:    true,
				Description: constants.AlertDefinitionDescription,
			},
			"escalation_levels": schema.ListNestedAttribute{
				Description: "List of escalation levels",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(3),
					listvalidator.SizeAtMost(4),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							Optional:    true,
							Description: "The ID of the escalation level.",
							Validators: []validator.Int64{
								int64validator.OneOf(1, 2, 3, 4),
							},
						},
						"escalation_mode": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: "Escalation mode: AlertOnErrorCount or AlertOnErrorDuration.",
							Validators: []validator.String{
								stringvalidator.OneOf("AlertOnErrorCount", "AlertOnErrorDuration"),
							},
						},
						"threshold_error_count": schema.Int64Attribute{
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							Optional:    true,
							Description: "Threshold for error count. This can be updated when escalation mode is AlertOnErrorCount",
						},
						"threshold_minutes": schema.Int64Attribute{
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							Optional:    true,
							Description: "Threshold for minutes. This can be updated when escalation mode is AlertOnErrorDuration",
						},
						"is_active": schema.BoolAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
							Optional:    true,
							Description: "Whether the escalation level is active.",
						},
						"message": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Optional:    true,
							Description: "Message for the escalation level.",
						},
						"number_of_reminders": schema.Int64Attribute{
							Computed:    true,
							Optional:    true,
							Description: "Number of reminders.",
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"reminder_delay": schema.Int64Attribute{
							Computed:    true,
							Optional:    true,
							Description: "Delay between reminders in minutes.",
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"include_trace_route": schema.BoolAttribute{
							Computed:    true,
							Optional:    true,
							Description: "Whether to include trace route.",
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

func (r *alertdefinitionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	// We need to validate in here because ListNestedAttribute has a plan drift if we don't make all the attributes are optional and computed in the schema.
	var config alertDefinitionResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate escalation_levels if provided
	if !config.EscalationLevels.IsNull() && !config.EscalationLevels.IsUnknown() {
		var levels []map[string]attr.Value
		elems := config.EscalationLevels.Elements()
		for idx, elem := range elems {
			if elem.IsNull() || elem.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("escalation_levels").AtListIndex(idx),
					"Escalation level block is null or unknown",
					"Each escalation level must be fully specified.",
				)
				continue
			}
			// Convert to map to check fields
			levelMap, ok := elem.(types.Object)
			if !ok {
				resp.Diagnostics.AddAttributeError(
					path.Root("escalation_levels").AtListIndex(idx),
					"Invalid escalation level type",
					"Each escalation level must be an object.",
				)
				continue
			}
			levels = append(levels, levelMap.Attributes())
		}

		// Validate required fields and uniqueness/range of "id" for each escalation level
		requiredFields := []string{
			"id",
			"escalation_mode",
			"threshold_error_count",
			"threshold_minutes",
			"is_active",
			"message",
			"number_of_reminders",
			"reminder_delay",
			"include_trace_route",
		}

		seenIDs := make(map[int64]int) // map id -> index
		numLevels := int64(len(levels))

		for idx, level := range levels {
			// Check all required fields are present and not null/unknown
			for _, field := range requiredFields {
				val, ok := level[field]
				if !ok || val.IsNull() || val.IsUnknown() {
					resp.Diagnostics.AddAttributeError(
						path.Root("escalation_levels").AtListIndex(idx).AtName(field),
						"Missing escalation level field",
						fmt.Sprintf("Field '%s' must be provided for each escalation level.", field),
					)
				}
			}

			// Validate id field: present, unique, and in range
			val, ok := level["id"]
			if ok && !val.IsNull() && !val.IsUnknown() {
				if idAttr, ok := val.(types.Int64); ok {
					idVal := idAttr.ValueInt64()
					if prevIdx, exists := seenIDs[idVal]; exists {
						resp.Diagnostics.AddAttributeError(
							path.Root("escalation_levels").AtListIndex(idx).AtName("id"),
							"Duplicate escalation level id",
							fmt.Sprintf("Duplicate id '%d' found in escalation_levels at indices %d and %d. Each escalation level must have a unique id.", idVal, prevIdx, idx),
						)
					} else {
						seenIDs[idVal] = idx
					}
					if idVal < 1 || idVal > numLevels {
						resp.Diagnostics.AddAttributeError(
							path.Root("escalation_levels").AtListIndex(idx).AtName("id"),
							"Escalation level id out of range",
							fmt.Sprintf("Escalation level id '%d' is out of range. It must be between 1 and %d.", idVal, numLevels),
						)
					}
				}
			}
		}

	}
}

// alertDefinitionResourceModel maps resource schema data.
type alertDefinitionResourceModel struct {
	AlertDefinitionGuid types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	IsActive            types.Bool   `tfsdk:"is_active"`
	EscalationLevels    types.List   `tfsdk:"escalation_levels"`
}

type escalationLevelResourceModel struct {
	Id                  types.Int64  `tfsdk:"id"`
	EscalationMode      types.String `tfsdk:"escalation_mode"`
	ThresholdErrorCount types.Int64  `tfsdk:"threshold_error_count"`
	ThresholdMinutes    types.Int64  `tfsdk:"threshold_minutes"`
	IsActive            types.Bool   `tfsdk:"is_active"`
	Message             types.String `tfsdk:"message"`
	NumberOfReminders   types.Int64  `tfsdk:"number_of_reminders"`
	ReminderDelay       types.Int64  `tfsdk:"reminder_delay"`
	IncludeTraceRoute   types.Bool   `tfsdk:"include_trace_route"`
}

// Create handles the creation of a new alertdefinition resource.
func (r *alertdefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan alertDefinitionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	payloadAlert := models.AlertDefinitionRequest{
		AlertName: plan.Name.ValueString(),
		IsActive:  plan.IsActive.ValueBool(),
	}

	responseAlert, err := r.client.CreateAlertDefinition(payloadAlert)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating alert definition",
			fmt.Sprintf("Could not create alert definition: %s", err),
		)
		return
	}
	// Update using patch with the clients escalation levels:
	if !plan.EscalationLevels.IsNull() && !plan.EscalationLevels.IsUnknown() {
		var userEscalationLevels []escalationLevelResourceModel
		diags := plan.EscalationLevels.ElementsAs(ctx, &userEscalationLevels, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		for _, level := range userEscalationLevels {
			payload := models.EscalationLevel{
				AlertDefinitionGuid: responseAlert.AlertDefinitionGuid,
				Id:                  int(level.Id.ValueInt64()),
				EscalationMode:      level.EscalationMode.ValueString(),
				ThresholdErrorCount: int(level.ThresholdErrorCount.ValueInt64()),
				ThresholdMinutes:    int(level.ThresholdMinutes.ValueInt64()),
				IsActive:            level.IsActive.ValueBool(),
				Message: func() string {
					if level.Message.IsNull() {
						return ""
					} else {
						return level.Message.ValueString()
					}
				}(),
				NumberOfReminders: int(level.NumberOfReminders.ValueInt64()),
				ReminderDelay:     int(level.ReminderDelay.ValueInt64()),
				IncludeTraceRoute: level.IncludeTraceRoute.ValueBool(),
			}
			// PATCH the escalation level (use the correct ID from the default levels or user input)
			err := r.client.UpdateEscalationLevel(payload)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error updating escalation level after creation",
					fmt.Sprintf("Could not update escalation level %d: %s", level.Id.ValueInt64(), err),
				)
				// We do not return to stop the update if an escalation level fails to update
			}
		}
	}

	getEscalationLevels, err := r.client.GetEscalationLevels(responseAlert.AlertDefinitionGuid)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching escalation levels",
			fmt.Sprintf("Could not fetch escalation levels for alert definition %s: %s", responseAlert.AlertDefinitionGuid, err),
		)
		return
	}

	// Update plan with values returned from the API.
	plan.AlertDefinitionGuid = types.StringValue(responseAlert.AlertDefinitionGuid)
	// Convert []escalationLevelResourceModel to types.List
	escalationLevelObjs := make([]escalationLevelResourceModel, len(getEscalationLevels))
	for i, level := range getEscalationLevels {
		escalationLevelObjs[i] = escalationLevelResourceModel{
			Id:                  types.Int64Value(int64(level.Id)),
			EscalationMode:      types.StringValue(level.EscalationMode),
			ThresholdErrorCount: types.Int64Value(int64(level.ThresholdErrorCount)),
			ThresholdMinutes:    types.Int64Value(int64(level.ThresholdMinutes)),
			IsActive:            types.BoolValue(level.IsActive),
			Message:             types.StringValue(level.Message),
			NumberOfReminders:   types.Int64Value(int64(level.NumberOfReminders)),
			ReminderDelay:       types.Int64Value(int64(level.ReminderDelay)),
			IncludeTraceRoute:   types.BoolValue(level.IncludeTraceRoute),
		}
	}
	listVal, diags := types.ListValueFrom(ctx, escalationLevelResourceModelType(), escalationLevelObjs)
	resp.Diagnostics.Append(diags...)
	plan.EscalationLevels = listVal

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

	alertDefItem, err := r.client.GetAlertDefinition(state.AlertDefinitionGuid.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading alert definition",
			fmt.Sprintf("Could not read alert definition %s: %s", state.AlertDefinitionGuid.ValueString(), err),
		)
		return
	}

	// Update state fields.
	state.Name = types.StringValue(alertDefItem.AlertName)
	state.IsActive = types.BoolValue(alertDefItem.IsActive)

	// Read escalation levels.
	escalationLevels, err := r.client.GetEscalationLevels(state.AlertDefinitionGuid.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading escalation levels",
			fmt.Sprintf("Could not read escalation levels for alert definition %s: %s", state.AlertDefinitionGuid.ValueString(), err),
		)
		return
	}

	var levels []escalationLevelResourceModel
	for _, lvl := range escalationLevels {
		levelModel := escalationLevelResourceModel{
			Id:                  types.Int64Value(int64(lvl.Id)),
			EscalationMode:      types.StringValue(lvl.EscalationMode),
			ThresholdErrorCount: types.Int64Value(int64(lvl.ThresholdErrorCount)),
			ThresholdMinutes:    types.Int64Value(int64(lvl.ThresholdMinutes)),
			IsActive:            types.BoolValue(lvl.IsActive),
			Message:             types.StringValue(lvl.Message),
			NumberOfReminders:   types.Int64Value(int64(lvl.NumberOfReminders)),
			ReminderDelay:       types.Int64Value(int64(lvl.ReminderDelay)),
			IncludeTraceRoute:   types.BoolValue(lvl.IncludeTraceRoute),
		}
		levels = append(levels, levelModel)
	}
	listVal, diags := types.ListValueFrom(ctx, escalationLevelResourceModelType(), levels)
	resp.Diagnostics.Append(diags...)
	state.EscalationLevels = listVal

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
	updateReq := models.AlertDefinitionRequest{
		AlertName: plan.Name.ValueString(),
		IsActive:  plan.IsActive.ValueBool(),
	}

	if err := r.client.UpdateAlertDefinition(plan.AlertDefinitionGuid.ValueString(), updateReq); err != nil {
		resp.Diagnostics.AddError(
			"Error updating alert definition",
			fmt.Sprintf("Could not update alert definition %s: %s", plan.AlertDefinitionGuid.ValueString(), err),
		)
		return
	}

	// Update escalation levels.
	var escalationLevelsSlice []escalationLevelResourceModel
	if !plan.EscalationLevels.IsNull() && !plan.EscalationLevels.IsUnknown() {
		diags := plan.EscalationLevels.ElementsAs(ctx, &escalationLevelsSlice, false)
		resp.Diagnostics.Append(diags...)
	}
	for _, level := range escalationLevelsSlice {
		payload := models.EscalationLevel{
			EscalationMode:      level.EscalationMode.ValueString(),
			ThresholdErrorCount: int(level.ThresholdErrorCount.ValueInt64()),
			ThresholdMinutes:    int(level.ThresholdMinutes.ValueInt64()),
			IsActive:            level.IsActive.ValueBool(),
			Message: func() string {
				if level.Message.IsNull() {
					return ""
				} else {
					return level.Message.ValueString()
				}
			}(),
			NumberOfReminders:   int(level.NumberOfReminders.ValueInt64()),
			ReminderDelay:       int(level.ReminderDelay.ValueInt64()),
			IncludeTraceRoute:   level.IncludeTraceRoute.ValueBool(),
			AlertDefinitionGuid: plan.AlertDefinitionGuid.ValueString(),
			Id:                  int(level.Id.ValueInt64()),
		}
		if err := r.client.UpdateEscalationLevel(payload); err != nil {
			resp.Diagnostics.AddError(
				"Error updating escalation level",
				fmt.Sprintf("Could not update escalation level %d: %s", level.Id.ValueInt64(), err),
			)
		}
	}

	escalationLevels, err := r.client.GetEscalationLevels(plan.AlertDefinitionGuid.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading escalation levels",
			fmt.Sprintf("Could not read escalation levels for alert definition %s: %s", plan.AlertDefinitionGuid.ValueString(), err),
		)
		return
	}

	var levels []escalationLevelResourceModel
	for _, lvl := range escalationLevels {
		levelModel := escalationLevelResourceModel{
			Id:                  types.Int64Value(int64(lvl.Id)),
			EscalationMode:      types.StringValue(lvl.EscalationMode),
			ThresholdErrorCount: types.Int64Value(int64(lvl.ThresholdErrorCount)),
			ThresholdMinutes:    types.Int64Value(int64(lvl.ThresholdMinutes)),
			IsActive:            types.BoolValue(lvl.IsActive),
			Message:             types.StringValue(lvl.Message),
			NumberOfReminders:   types.Int64Value(int64(lvl.NumberOfReminders)),
			ReminderDelay:       types.Int64Value(int64(lvl.ReminderDelay)),
			IncludeTraceRoute:   types.BoolValue(lvl.IncludeTraceRoute),
		}
		levels = append(levels, levelModel)
	}
	listVal, diags := types.ListValueFrom(ctx, escalationLevelResourceModelType(), levels)
	resp.Diagnostics.Append(diags...)
	plan.EscalationLevels = listVal

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

	if err := r.client.DeleteAlertDefinition(state.AlertDefinitionGuid.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting alert definition",
			fmt.Sprintf("Could not delete alert definition %s: %s", state.AlertDefinitionGuid.ValueString(), err),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *alertdefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import AlertDefinitionGuid and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Helper function for types.ObjectType for escalationLevelResourceModel
func escalationLevelResourceModelType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":                    types.Int64Type,
			"escalation_mode":       types.StringType,
			"threshold_error_count": types.Int64Type,
			"threshold_minutes":     types.Int64Type,
			"is_active":             types.BoolType,
			"message":               types.StringType,
			"number_of_reminders":   types.Int64Type,
			"reminder_delay":        types.Int64Type,
			"include_trace_route":   types.BoolType,
		},
	}
}
