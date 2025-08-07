package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
)

// Ensure operatorGroupPermissionResource implements the resource.Resource interface.
var _ resource.Resource = &operatorGroupPermissionResource{}

// operatorGroupPermissionResource implements the Terraform resource for operatorgroup_permissions.
type operatorGroupPermissionResource struct {
	client interfaces.IOperatorGroupPermission
}

// NewResource returns a new instance of the operatorgroup_permission resource.
func NewOperatorGroupPermissionResource(client interfaces.IOperatorGroupPermission) resource.Resource {
	return &operatorGroupPermissionResource{
		client: client,
	}
}

// operatorGroupPermissionModel defines the schema model for the operatorgroup_permission resource.
type operatorGroupPermissionModel struct {
	ID         types.String `tfsdk:"id"`
	GroupID    types.String `tfsdk:"operatorgroup_id"`
	Permission types.String `tfsdk:"permission"`
}

// Metadata returns the resource type name.
func (r *operatorGroupPermissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "operatorgroup_permission"
}

// Schema defines the schema for the operatorgroup_permission resource.
func (r *operatorGroupPermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier of the operator group permission",
			},
			"permission": rschema.StringAttribute{
				Required:    true,
				Description: "The permission of the operator group",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ShareDashboards",
						"AllowInfra",
						"Administrator", // This is a special case, it should not be included in the list.
						"TechnicalContact",
						"FinancialOperator",
						"BasicOperator",
						"CreateAlertDefinition",
						"CreateIntegration",
						"CreatePrivateLocations",
						"ManageMonitorTemplates",
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"operatorgroup_id": rschema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The unique identifier of the operator group",
			},
		},
	}
}

// Read retrieves the current state of the resource from the API.
func (r *operatorGroupPermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state operatorGroupPermissionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve the current permission from the API.
	currentPermission, err := r.client.GetOperatorGroupPermission(state.GroupID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading permission",
			fmt.Sprintf("Could not retrieve permission for group %q: %s", state.GroupID.ValueString(), err.Error()),
		)
		return
	}

	// Check if the permission exists.
	found := false
	for _, auth := range currentPermission {
		if auth == state.Permission.ValueString() {
			found = true
			break
		}
	}

	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save the state.
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

// Create assigns the list of permission to the operator group.
func (r *operatorGroupPermissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan operatorGroupPermissionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Assign the permission to the operator group.
	if err := r.client.AssignOperatorGroupPermission(plan.GroupID.ValueString(), plan.Permission.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error creating permission",
			fmt.Sprintf("Could not assign permission %q to group %q: %s", plan.Permission.ValueString(), plan.GroupID.ValueString(), err.Error()),
		)
		return
	}

	// Set the composite ID.
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s", plan.GroupID.ValueString(), plan.Permission.ValueString()))

	// Save the state.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *operatorGroupPermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not implemented for this resource as the API does not support updates.
}

// Delete removes permission from the operator group.
func (r *operatorGroupPermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state operatorGroupPermissionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete the permission from the operator group.
	if err := r.client.DeleteOperatorGroupPermission(state.GroupID.ValueString(), state.Permission.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting permission",
			fmt.Sprintf("Could not delete permission %q from group %q: %s", state.Permission.ValueString(), state.GroupID.ValueString(), err.Error()),
		)
		return
	}

	// Remove the resource from the state.
	resp.State.RemoveResource(ctx)
}

func (r *operatorGroupPermissionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse the composite ID.
	parts := strings.Split(req.ID, ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Error importing resource",
			"Expected ID in the format `operatorgroup_id:permission`.",
		)
		return
	}

	groupID := parts[0]
	permission := parts[1]

	// Set the attributes in the state.
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("operatorgroup_id"), groupID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("permission"), permission)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
