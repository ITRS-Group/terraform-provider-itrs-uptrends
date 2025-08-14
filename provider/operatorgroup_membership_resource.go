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

// Ensure membershipResource implements the resource.Resource interface.
var _ resource.Resource = &membershipResource{}

// membershipResource implements the Terraform resource for memberships.
type membershipResource struct {
	client interfaces.IMembership
}

// NewMembershipResource returns a new instance of the membership resource.
func NewMembershipResource(client interfaces.IMembership) resource.Resource {
	return &membershipResource{
		client: client,
	}
}

// membershipModel defines the schema model for the membership resource.
type membershipModel struct {
	ID         types.String `tfsdk:"id"`
	OperatorID types.String `tfsdk:"operator_id"`
	GroupID    types.String `tfsdk:"operatorgroup_id"`
}

// Metadata returns the resource type name.
func (r *membershipResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_operatorgroup_membership"
}

// Schema defines the schema for the membership resource.
func (r *membershipResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed:    true,
				Description: constants.OperatorGroupDescription,
			},
			"operator_id": rschema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: constants.OperatorGroupDescription,
			},
			"operatorgroup_id": rschema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: constants.OperatorGroupDescription,
			},
		},
	}
}

// Create creates the membership resource by assigning an operator to an operator group.
func (r *membershipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan membershipModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the client's AssignOperator method to create the membership.
	err := r.client.AssignOperator(plan.GroupID.ValueString(), plan.OperatorID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating operatorgroup_membership",
			fmt.Sprintf("Could not assign operator to group: %s", err.Error()),
		)
		return
	}

	// Set a computed ID, e.g. a combination of operator and group IDs.
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s", plan.OperatorID.ValueString(), plan.GroupID.ValueString()))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read refreshes the state of the membership resource.
func (r *membershipResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state membershipModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve operatorgroup_memberships for the given group.
	memberships, err := r.client.GetMemberships(state.GroupID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading operatorgroup_membership",
			fmt.Sprintf("Could not retrieve operator group memberships: %s", err.Error()),
		)
		return
	}

	// Check if the operator exists in the returned memberships.
	found := lo.ContainsBy(memberships, func(m models.MembershipResponse) bool {
		return m.OperatorGuid == state.OperatorID.ValueString()
	})

	if !found {
		// The membership no longer exists, so remove it from state.
		resp.State.RemoveResource(ctx)
		return
	}

	// If the membership exists, the state remains unchanged.
}

func (r *membershipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not implemented for this resource as the API does not support updates.
}

// Delete removes the membership by calling the DeleteMembership method.
func (r *membershipResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state membershipModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteMembership(state.GroupID.ValueString(), state.OperatorID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting operatorgroup_membership",
			fmt.Sprintf("Could not delete operator group membership: %s", err.Error()),
		)
		return
	}
}

func (r *membershipResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Split the import ID into operator_id and group_id
	idParts := strings.Split(req.ID, ":")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: operator_id:operatorgroup_id. Got: %q", req.ID),
		)
		return
	}

	// Set the operator_id and group_id attributes in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("operator_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("operatorgroup_id"), idParts[1])...)

	// Set the ID attribute in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
