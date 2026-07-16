package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/constants"
)

var _ datasource.DataSource = &alertdefinitionDataSource{}

func NewAlertDefinitionDataSource(client interfaces.IAlertDefinition) datasource.DataSource {
	return &alertdefinitionDataSource{client: client}
}

type alertdefinitionDataSource struct {
	client interfaces.IAlertDefinition
}

func (d *alertdefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_alertdefinition"
}

func (d *alertdefinitionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Alert definition GUID. Provide this or name.",
				Optional:    true,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: constants.AlertDefinitionDescription,
				Optional:    true,
				Computed:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: constants.AlertDefinitionDescription,
				Computed:    true,
			},
			"escalation_levels": schema.ListNestedAttribute{
				Description: "List of escalation levels",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"escalation_mode": schema.StringAttribute{
							Computed: true,
						},
						"threshold_error_count": schema.Int64Attribute{
							Computed: true,
						},
						"threshold_minutes": schema.Int64Attribute{
							Computed: true,
						},
						"is_active": schema.BoolAttribute{
							Computed: true,
						},
						"message": schema.StringAttribute{
							Computed: true,
						},
						"number_of_reminders": schema.Int64Attribute{
							Computed: true,
						},
						"reminder_delay": schema.Int64Attribute{
							Computed: true,
						},
						"include_trace_route": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

type alertdefinitionDataSourceModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	IsActive       types.Bool   `tfsdk:"is_active"`
	EscalationList types.List   `tfsdk:"escalation_levels"`
}

func (d *alertdefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The alert definition client was not configured.")
		return
	}

	var data alertdefinitionDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idProvided := !data.ID.IsNull() && data.ID.ValueString() != ""
	nameProvided := !data.Name.IsNull() && data.Name.ValueString() != ""

	switch {
	case idProvided && nameProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide only one of id or name.")
		return
	case !idProvided && !nameProvided:
		resp.Diagnostics.AddError("Invalid configuration", "Provide either id or name.")
		return
	}

	var state alertDefinitionResourceModel

	if idProvided {
		def, err := d.client.GetAlertDefinition(data.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error reading alert definition", err.Error())
			return
		}
		state.AlertDefinitionGuid = types.StringValue(def.AlertDefinitionGuid)
		state.Name = types.StringValue(def.AlertName)
		state.IsActive = types.BoolValue(def.IsActive)
	} else {
		defs, statusCode, responseBody, err := d.client.GetAlertDefinitions()
		if err != nil {
			resp.Diagnostics.AddError("Error listing alert definitions", err.Error())
			return
		}
		if statusCode >= 300 {
			resp.Diagnostics.AddError("Failed to list alert definitions", fmt.Sprintf("HTTP %d: %s", statusCode, responseBody))
			return
		}
		name := data.Name.ValueString()
		found := false
		for _, def := range defs {
			if strings.EqualFold(def.AlertName, name) {
				if found {
					resp.Diagnostics.AddError("Alert definition not unique", fmt.Sprintf("More than one alert definition found with name %q", name))
					return
				}
				state.AlertDefinitionGuid = types.StringValue(def.AlertDefinitionGuid)
				state.Name = types.StringValue(def.AlertName)
				state.IsActive = types.BoolValue(def.IsActive)
				found = true
			}
		}
		if !found {
			resp.Diagnostics.AddError("Alert definition not found", fmt.Sprintf("No alert definition found with name %q", name))
			return
		}
	}

	levels, err := d.client.GetEscalationLevels(state.AlertDefinitionGuid.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error fetching escalation levels", err.Error())
		return
	}
	escalationLevelObjs := make([]escalationLevelResourceModel, len(levels))
	for i, level := range levels {
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
	if resp.Diagnostics.HasError() {
		return
	}
	state.EscalationLevels = listVal

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
