package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/constants"
	"github.com/samber/lo"
)

// monitorgroupMembershipResource implements the Terraform resource for monitor group memberships.
type monitorgroupMembershipResource struct {
	client interfaces.IMonitorGroupMember
}

// Ensure the implementation satisfies the expected interfaces.
var _ resource.Resource = &monitorgroupMembershipResource{}

// monitorgroupMembershipResourceModel maps the resource schema data.
type monitorgroupMembershipResourceModel struct {
	ID             types.String `tfsdk:"id"`
	MonitorID      types.String `tfsdk:"monitor_id"`
	MonitorGroupID types.String `tfsdk:"monitorgroup_id"`
}

// NewMonitorgroupMembershipResource instantiates the resource.
func NewMonitorgroupMembershipResource(client interfaces.IMonitorGroupMember) resource.Resource {
	return &monitorgroupMembershipResource{
		client: client,
	}
}

// Metadata returns the resource type name.
func (r *monitorgroupMembershipResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "itrs-uptrends_monitorgroup_membership"
}

// Schema defines the schema for the resource.
func (r *monitorgroupMembershipResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed:    true,
				Description: constants.MonitorGroupDescription,
				// Because we require replace on updates, we don't need to UseStateForUnknown()
			},
			"monitor_id": rschema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: constants.MonitorGroupDescription,
			},
			"monitorgroup_id": rschema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: constants.MonitorGroupDescription,
			},
		},
	}
}

// Configure sets the provider client for the resource.
func (r *monitorgroupMembershipResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	// Assert that the provider data satisfies IMonitorGroupMember.
	r.client = req.ProviderData.(interfaces.IMonitorGroupMember)
}

// Create handles the creation of the resource.
func (r *monitorgroupMembershipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan monitorgroupMembershipResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the membership via the client API.
	err := r.client.AssignMembership(plan.MonitorGroupID.ValueString(), plan.MonitorID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating monitor group membership",
			fmt.Sprintf("Could not create monitor group membership: %s", err.Error()),
		)
		return
	}

	// Set a composite ID (e.g., monitor_id:monitorgroup_id).
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s", plan.MonitorID.ValueString(), plan.MonitorGroupID.ValueString()))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read refreshes the Terraform state with data from the API.
func (r *monitorgroupMembershipResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state monitorgroupMembershipResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve memberships for the monitor group.
	memberships, err := r.client.GetGroupMemberships(state.MonitorGroupID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading monitor group membership",
			fmt.Sprintf("Could not read monitor group memberships: %s", err.Error()),
		)
		return
	}

	// Check if the membership exists.
	exists := lo.ContainsBy(memberships, func(m models.MonitorMembershipResponse) bool {
		return m.MonitorGuid == state.MonitorID.ValueString()
	})

	if !exists {
		// If the membership no longer exists, remove it from state.
		resp.State.RemoveResource(ctx)
		return
	}
}

func (r *monitorgroupMembershipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not implemented for this resource as the API does not support updates.
}

// Delete removes the resource from the API.
func (r *monitorgroupMembershipResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state monitorgroupMembershipResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteMembership(state.MonitorGroupID.ValueString(), state.MonitorID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting monitor group membership",
			fmt.Sprintf("Could not delete monitor group membership: %s", err.Error()),
		)
		return
	}
}

func (r *monitorgroupMembershipResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Split the import ID into monitor_id and monitorgroup_id
	idParts := strings.Split(req.ID, ":")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: monitor_id:monitorgroup_id. Got: %q", req.ID),
		)
		return
	}

	// Set the monitor_id and monitorgroup_id attributes in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("monitor_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("monitorgroup_id"), idParts[1])...)

	// Set the ID attribute in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
