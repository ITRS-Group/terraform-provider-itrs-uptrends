package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/constants"
)

// Ensure the implementation satisfies the Resource and ResourceWithImportState interfaces.
var _ resource.Resource = &alertDefinitionOperatorMembershipResource{}

// alertDefinitionOperatorMembershipResource implements the Terraform resource.
type alertDefinitionOperatorMembershipResource struct {
	client interfaces.IAlertDefinitionOperatorMembership
}

// NewAlertDefinitionOperatorMembershipResource instantiates the resource with a client.
func NewAlertDefinitionOperatorMembershipResource(cli interfaces.IAlertDefinitionOperatorMembership) resource.Resource {
	return &alertDefinitionOperatorMembershipResource{
		client: cli,
	}
}

// alertDefinitionOperatorMembershipResourceModel maps the resource schema.
type alertDefinitionOperatorMembershipResourceModel struct {
	ID                types.String `tfsdk:"id"`
	AlertDefinitionID types.String `tfsdk:"alertdefinition_id"`
	OperatorID        types.String `tfsdk:"operator_id"`
	EscalationLevel   types.Int64  `tfsdk:"escalationlevel"`
}

// Metadata returns the resource type name.
func (r *alertDefinitionOperatorMembershipResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "itrs-uptrends_alertdefinition_operator_membership"
}

// Schema defines the schema for this resource.
func (r *alertDefinitionOperatorMembershipResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Description: "Manages an operator membership for an escalation level in an alert definition.",
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed:    true,
				Description: constants.AlertDefinitionDescription,
			},
			"alertdefinition_id": rschema.StringAttribute{
				Required:    true,
				Description: constants.AlertDefinitionDescription,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"operator_id": rschema.StringAttribute{
				Required:    true,
				Description: constants.OperatorDescription,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"escalationlevel": rschema.Int64Attribute{
				Required:    true,
				Description: constants.AlertDefinitionDescription,
				Validators: []validator.Int64{
					int64validator.Between(1, 4),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
		},
	}
}

// Create handles the creation of a new resource.
func (r *alertDefinitionOperatorMembershipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan alertDefinitionOperatorMembershipResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the membership using the client.
	apiResp, err := r.client.CreateMembership(
		plan.AlertDefinitionID.ValueString(),
		int(plan.EscalationLevel.ValueInt64()),
		plan.OperatorID.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError("Error adding an operator to the escalation level in alert definition",
			fmt.Sprintf("Error adding an operator to the escalation level %d in alert definition %s %s",
				plan.EscalationLevel.ValueInt64(),
				plan.AlertDefinitionID.ValueString(),
				err.Error()))
		return
	}

	// Compose a unique ID using the API response and plan values.
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s:%d", apiResp.AlertDefinition, apiResp.Operator, apiResp.Escalationlevel))
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read refreshes the resource state.
func (r *alertDefinitionOperatorMembershipResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state alertDefinitionOperatorMembershipResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve memberships using the client.
	memberships, err := r.client.GetMembership(
		state.AlertDefinitionID.ValueString(),
		int(state.EscalationLevel.ValueInt64()),
	)
	if err != nil {
		resp.Diagnostics.AddError("Error reading which operators are assigned to the escalation levels of alert definition",
			fmt.Sprintf("Error reading which operators are assigned to the escalation levels of alert definition %s, %s",
				state.AlertDefinitionID.ValueString(),
				err.Error()))
		return
	}

	// Check if the membership still exists for the given operator.
	found := false
	for _, membership := range memberships {
		if membership.OperatorGuid == state.OperatorID.ValueString() {
			found = true
			break
		}
	}

	if !found {
		// Membership not found; remove resource from state.
		resp.State.RemoveResource(ctx)
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *alertDefinitionOperatorMembershipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not implemented for this resource as the API does not support updates.
}

// Delete handles deletion of the resource.
func (r *alertDefinitionOperatorMembershipResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state alertDefinitionOperatorMembershipResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteMembership(
		state.AlertDefinitionID.ValueString(),
		int(state.EscalationLevel.ValueInt64()),
		state.OperatorID.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError("Error removing operator from escalation level of alert definition",
			fmt.Sprintf("Error removing operator %s from escalation level %d of alert definition %s because %s",
				state.OperatorID.ValueString(),
				state.EscalationLevel.ValueInt64(),
				state.AlertDefinitionID.ValueString(),
				err.Error()))
		return
	}
}

func (r *alertDefinitionOperatorMembershipResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Split the import ID into alertdefinition_id, operator_id, and escalationlevel
	idParts := strings.Split(req.ID, ":")

	// Validate the format of the import ID
	if len(idParts) != 3 {
		resp.Diagnostics.AddError(
			"Invalid Import Identifier",
			fmt.Sprintf("Expected import identifier with format: alertdefinition_id:operator_id:escalationlevel. Got: %q", req.ID),
		)
		return
	}

	// Validate that none of the parts are empty
	if idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Invalid Import Identifier",
			fmt.Sprintf("One or more parts of the import identifier are empty. Got: %q", req.ID),
		)
		return
	}

	// Convert escalationlevel to an integer
	escalationLevel, err := strconv.Atoi(idParts[2])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Escalation Level",
			fmt.Sprintf("Escalation level must be a valid integer. Got: %q", idParts[2]),
		)
		return
	}

	// Set the alertdefinition_id, operator_id, and escalationlevel attributes in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alertdefinition_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("operator_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("escalationlevel"), escalationLevel)...)

	// Set the ID attribute in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
