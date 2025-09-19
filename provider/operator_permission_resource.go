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
	"github.com/itrs-group/terraform-provider-itrs-uptrends/constants"
)

// Ensure operatorPermissionResource implements the resource.Resource interface.
var _ resource.Resource = &operatorPermissionResource{}

// operatorPermissionResource implements the Terraform resource for operator_permissions.
type operatorPermissionResource struct {
	client interfaces.IOperatorPermission
}

// NewResource returns a new instance of the operator_permission resource.
func NewOperatorPermissionResource(client interfaces.IOperatorPermission) resource.Resource {
	return &operatorPermissionResource{
		client: client,
	}
}

// operatorPermissionModel defines the schema model for the operator_permission resource.
type operatorPermissionModel struct {
	ID         types.String `tfsdk:"id"`
	OperatorID types.String `tfsdk:"operator_id"`
	Permission types.String `tfsdk:"permission"`
}

// Metadata returns the resource type name.
func (r *operatorPermissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "itrs-uptrends_operator_permission"
}

// Schema defines the schema for the operator_permission resource.
func (r *operatorPermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed:    true,
				Description: constants.OperatorDescription,
			},
			"permission": rschema.StringAttribute{
				Required:    true,
				Description: constants.OperatorDescription,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"AccountAccess",
						"AccountAdministrator",
						"AllowInfra",
						"FinancialOperator",
						"TechnicalContact",
						"ShareDashboards",
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
			"operator_id": rschema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: constants.OperatorDescription,
			},
		},
	}
}

// Read retrieves the current state of the resource from the API.
func (r *operatorPermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state operatorPermissionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve the current permission from the API.
	currentPermission, err := r.client.GetOperatorPermission(state.OperatorID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading permission",
			fmt.Sprintf("Could not retrieve permission for operator %q: %s", state.OperatorID.ValueString(), err.Error()),
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

// Create assigns the list of permission to the operator.
func (r *operatorPermissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan operatorPermissionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Assign the permission to the operator.
	if err := r.client.AssignOperatorPermission(plan.OperatorID.ValueString(), plan.Permission.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error creating permission",
			fmt.Sprintf("Could not assign permission %q to operator %q: %s", plan.Permission.ValueString(), plan.OperatorID.ValueString(), err.Error()),
		)
		return
	}

	// Set the composite ID.
	plan.ID = types.StringValue(fmt.Sprintf("%s:%s", plan.OperatorID.ValueString(), plan.Permission.ValueString()))

	// Save the state.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *operatorPermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// No implementation needed because RequiresReplace handles updates by replacing the resource.
}

// Delete removes permission from the operator.
func (r *operatorPermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state operatorPermissionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete the permission from the operator.
	if err := r.client.DeleteOperatorPermission(state.OperatorID.ValueString(), state.Permission.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting permission",
			fmt.Sprintf("Could not delete permission %q from operator %q: %s", state.Permission.ValueString(), state.OperatorID.ValueString(), err.Error()),
		)
		return
	}

	// Remove the resource from the state.
	resp.State.RemoveResource(ctx)
}

func (r *operatorPermissionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse the composite ID.
	parts := strings.Split(req.ID, ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Error importing resource",
			"Expected ID in the format `operator_id:permission`.",
		)
		return
	}

	operatorID := parts[0]
	permission := parts[1]

	// Set the attributes in the state.
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("operator_id"), operatorID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("permission"), permission)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
