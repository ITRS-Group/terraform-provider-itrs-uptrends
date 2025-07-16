package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/general"
)

// operatorGroupResource implements the Terraform resource interface.
type operatorGroupResource struct {
	client interfaces.IOperatorGroup
}

// NewOperatorGroupResource creates a new instance of the operator group resource.
func NewOperatorGroupResource(client interfaces.IOperatorGroup) resource.Resource {
	return &operatorGroupResource{
		client: client,
	}
}

// Metadata returns the resource type name.
func (r *operatorGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "operatorgroup"
}

// Schema defines the schema for the resource.
func (r *operatorGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed:    true,
				Description: general.OperatorGroupDescription,
			},
			"description": rschema.StringAttribute{
				Required:    true,
				Description: general.OperatorGroupDescription,
			},
		},
	}
}

// operatorGroupModel maps resource schema data.
type operatorGroupModel struct {
	Id          types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
}

// Create is called when the resource is created.
func (r *operatorGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan operatorGroupModel

	// Retrieve the plan.
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the underlying CreateOperatorGroup method.
	result, err, msg := r.client.CreateOperatorGroup(plan.Description.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error creating Operator Group", fmt.Sprintf("Error: %v. Message: %s", err, msg))
		return
	}

	// Update the plan with data from the response.
	plan.Id = types.StringValue(result.OperatorGroupGuid)
	plan.Description = types.StringValue(result.Description)

	// Set the state.
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Read is called to refresh the Terraform state.
func (r *operatorGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state operatorGroupModel

	// Retrieve the state.
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the underlying GetOperatorGroup method.
	result, err, msg := r.client.GetOperatorGroup(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading Operator Group", fmt.Sprintf("Error: %v. Message: %s", err, msg))
		return
	}

	// Update the state with the refreshed data.
	state.Description = types.StringValue(result.Description)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the state.
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update is called when the resource is modified.
func (r *operatorGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state operatorGroupModel
	// Retrieve the existing state to obtain the computed ID.
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config operatorGroupModel
	// Retrieve the plan.
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	operatorGroupID := state.Id.ValueString()

	// Use the merged plan.Id for the update API call.
	result, err := r.client.UpdateOperatorGroup(config.Description.ValueString(), operatorGroupID)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Operator Group", fmt.Sprintf("Error: %v", err))
		return
	}
	if result != "" {
		resp.Diagnostics.AddWarning("Operator Group Update Warning", fmt.Sprintf("Update response: %s", result))

	}
	state.Description = types.StringValue(config.Description.ValueString())
	state.Id = types.StringValue(operatorGroupID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the state in a single, consistent operation.
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Delete is called when the resource is destroyed.
func (r *operatorGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state operatorGroupModel

	// Retrieve the state.
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the underlying DeleteOperatorGroup method.
	err, msg := r.client.DeleteOperatorGroup(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting Operator Group", fmt.Sprintf("Error: %v. Message: %s", err, msg))
		return
	}

	// Remove the resource from state.
	resp.State.RemoveResource(ctx)
}

func (r *operatorGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
