package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/constants"
)

type alertDefinitionMonitorMembershipResource struct {
	client interfaces.IAlertDefinitionMonitorMember
}

// NewAlertDefinitionMonitorMembershipResource returns a new instance of the resource.
func NewAlertDefinitionMonitorMembershipResource(client interfaces.IAlertDefinitionMonitorMember) resource.Resource {
	return &alertDefinitionMonitorMembershipResource{
		client: client,
	}
}

// Resource model for itrs-uptrends_alertdefinition_monitor_membership.
type alertDefinitionMonitorMembershipModel struct {
	ID                types.String `tfsdk:"id"`
	AlertDefinitionID types.String `tfsdk:"alertdefinition_id"`
	MonitorID         types.String `tfsdk:"monitor_id"`
}

// Metadata returns the resource type name.
func (r *alertDefinitionMonitorMembershipResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "itrs-uptrends_alertdefinition_monitor_membership"
}

// Schema defines the schema for the resource.
func (r *alertDefinitionMonitorMembershipResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: constants.AlertDefinitionDescription,
				// Because we require replace on updates, we don't need to UseStateForUnknown()
			},
			"alertdefinition_id": schema.StringAttribute{
				Required:    true,
				Description: constants.AlertDefinitionDescription,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"monitor_id": schema.StringAttribute{
				Required:    true,
				Description: constants.AlertDefinitionDescription,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

// Create is called during resource creation.
func (r *alertDefinitionMonitorMembershipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan alertDefinitionMonitorMembershipModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the client's AssignMonitor function.
	assignResp, err := r.client.AssignMonitor(plan.AlertDefinitionID.ValueString(), plan.MonitorID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error assigning monitor",
			fmt.Sprintf("Could not assign monitor: %s", err.Error()),
		)
		return
	}

	// Use a composite ID based on the alert definition and monitor identifiers.
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s", assignResp.AlertDefinition, assignResp.Monitor))

	// Set state.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read refreshes the resource state.
func (r *alertDefinitionMonitorMembershipResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state alertDefinitionMonitorMembershipModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Verify that the assignment still exists.
	assignments, err := r.client.GetAssignments(state.AlertDefinitionID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading assignments",
			fmt.Sprintf("Could not retrieve assignments: %s", err.Error()),
		)
		return
	}

	found := false
	for _, assignment := range assignments {
		// Check if the monitor ID matches either the MonitorGuid or MonitorGroupGuid.
		if assignment.MonitorGuid != nil && *assignment.MonitorGuid == state.MonitorID.ValueString() {
			found = true
			break
		}
		if assignment.MonitorGroupGuid != nil && *assignment.MonitorGroupGuid == state.MonitorID.ValueString() {
			found = true
			break
		}
	}

	if !found {
		// If the assignment no longer exists, remove it from state.
		resp.State.RemoveResource(ctx)
		return
	}

	// If found, no changes to state are needed.
}

func (r *alertDefinitionMonitorMembershipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not implemented for this resource as the API does not support updates.
}

// Delete removes the monitor assignment.
func (r *alertDefinitionMonitorMembershipResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state alertDefinitionMonitorMembershipModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.RemoveAssignment(state.AlertDefinitionID.ValueString(), state.MonitorID.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting assignment",
			fmt.Sprintf("Could not remove assignment: %s", err.Error()),
		)
		return
	}
	// Upon successful deletion, Terraform automatically removes the state.
}

func (r *alertDefinitionMonitorMembershipResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Split the import ID into alertdefinition_id and monitor_id
	idParts := strings.Split(req.ID, ":")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: alertdefinition_id:monitor_id. Got: %q", req.ID),
		)
		return
	}

	// Set the alertdefinition_id and monitorgroup_id attributes in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alertdefinition_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("monitor_id"), idParts[1])...)

	// Set the ID attribute in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
