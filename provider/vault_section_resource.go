package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	interfaces "github.com/itrs-group/terraform-provider-itrs-uptrends/client/interfaces"
)

// Ensure vaultSectionResource implements resource.Resource.
var _ resource.Resource = &vaultSectionResource{}

type vaultSectionResource struct {
	client interfaces.IVaultSection
}

type vaultSectionResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// NewVaultSectionResource returns a new instance of the resource.
func NewVaultSectionResource(c interfaces.IVaultSection) resource.Resource {
	return &vaultSectionResource{
		client: c,
	}
}

// Metadata sets the resource type name.
func (r *vaultSectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vault_section"
}

// Schema defines the schema for the vault_section resource.
func (r *vaultSectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "The unique identifier of the vault section.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the vault section.",
			},
		},
	}
}

// Create handles the creation of a new vault_section resource.
func (r *vaultSectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan vaultSectionResourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the underlying CreateOperatorGroup method.
	result, err, msg := r.client.CreateVaultSection(plan.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error creating Operator Group", fmt.Sprintf("Error: %v. Message: %s", err, msg))
		return
	}

	// Update the plan with data from the response.
	plan.Id = types.StringValue(result.VaultSectionGuid)
	plan.Name = types.StringValue(result.Name)

	// Set the state.
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Read refreshes the Terraform state with the latest data.
func (r *vaultSectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state vaultSectionResourceModel

	// Retrieve the state.
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the underlying GetVaultSection method.
	result, err, msg := r.client.GetVaultSection(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading Vault Section", fmt.Sprintf("Error: %v. Message: %s", err, msg))
		return
	}

	// Update the state with the refreshed data.
	state.Name = types.StringValue(result.Name)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the state.
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update handles updating an existing vault_section resource.
func (r *vaultSectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state vaultSectionResourceModel
	// Retrieve the existing state to obtain the computed ID.
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config vaultSectionResourceModel
	// Retrieve the plan.
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	vaultSectionID := state.Id.ValueString()

	// Use the merged plan.Id for the update API call.
	result, updateErr := r.client.UpdateVaultSection(vaultSectionID, config.Name.ValueString())
	if updateErr != nil {
		resp.Diagnostics.AddError("Error updating Operator Group", fmt.Sprintf("Error: %v", updateErr))
		return
	}
	if result != "" {
		resp.Diagnostics.AddWarning("Operator Group Update Warning", fmt.Sprintf("Update response: %s", result))
	}

	state.Name = types.StringValue(config.Name.ValueString())
	state.Id = types.StringValue(vaultSectionID)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the state in a single, consistent operation.
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Delete removes the vault_section resource.
func (r *vaultSectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state vaultSectionResourceModel
	// Retrieve the state.
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the underlying DeleteVaultSection method.
	err, msg := r.client.DeleteVaultSection(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting Operator Group", fmt.Sprintf("Error: %v. Message: %s", err, msg))
		return
	}

	// Remove the resource from state.
	resp.State.RemoveResource(ctx)
}

func (r *vaultSectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
