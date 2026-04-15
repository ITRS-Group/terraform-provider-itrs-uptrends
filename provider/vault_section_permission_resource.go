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
	models "github.com/itrs-group/terraform-provider-itrs-uptrends/client/models"
)

var _ resource.Resource = &vaultSectionPermissionResource{}

type vaultSectionPermissionResource struct {
	client interfaces.IVaultSectionPermission
}

func NewVaultSectionPermissionResource(client interfaces.IVaultSectionPermission) resource.Resource {
	return &vaultSectionPermissionResource{
		client: client,
	}
}

type vaultSectionPermissionModel struct {
	ID                types.String `tfsdk:"id"`
	VaultSectionID    types.String `tfsdk:"vault_section_id"`
	AuthorizationType types.String `tfsdk:"permission"`
	OperatorID        types.String `tfsdk:"operator_id"`
	OperatorGroupID   types.String `tfsdk:"operatorgroup_id"`
}

func (r *vaultSectionPermissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "itrs-uptrends_vault_section_permission"
}

func (r *vaultSectionPermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier of the authorization (composite key in format `vault_section_id:authorization_id`).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vault_section_id": rschema.StringAttribute{
				Required:    true,
				Description: "The GUID of the vault section.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"permission": rschema.StringAttribute{
				Required:    true,
				Description: "The authorization type. Valid values: `ViewVaultSection`, `ChangeVaultSection`.",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ViewVaultSection",
						"ChangeVaultSection",
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"operator_id": rschema.StringAttribute{
				Optional:    true,
				Description: "The GUID of the operator. Provide this or `operatorgroup_id`, not both.",
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(
						path.MatchRoot("operator_id"),
						path.MatchRoot("operatorgroup_id"),
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"operatorgroup_id": rschema.StringAttribute{
				Optional:    true,
				Description: "The GUID of the operator group. Provide this or `operator_id`, not both.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *vaultSectionPermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state vaultSectionPermissionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	parts := strings.Split(state.ID.ValueString(), ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid ID", "Expected ID in the format `vault_section_id:authorization_id`.")
		return
	}
	vaultSectionID := parts[0]
	authorizationID := parts[1]

	authorizations, err := r.client.GetVaultSectionAuthorizations(vaultSectionID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading vault section authorization",
			fmt.Sprintf("Could not retrieve authorizations for vault section %q: %s", vaultSectionID, err.Error()),
		)
		return
	}

	found := false
	for _, auth := range authorizations {
		if auth.AuthorizationId == authorizationID {
			state.VaultSectionID = types.StringValue(vaultSectionID)
			state.AuthorizationType = types.StringValue(auth.AuthorizationType)
			if auth.OperatorGuid != "" {
				state.OperatorID = types.StringValue(auth.OperatorGuid)
			} else {
				state.OperatorID = types.StringNull()
			}
			if auth.OperatorGroupGuid != "" {
				state.OperatorGroupID = types.StringValue(auth.OperatorGroupGuid)
			} else {
				state.OperatorGroupID = types.StringNull()
			}
			found = true
			break
		}
	}

	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *vaultSectionPermissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan vaultSectionPermissionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	auth := models.VaultSectionAuthorization{
		AuthorizationType: plan.AuthorizationType.ValueString(),
	}
	if !plan.OperatorID.IsNull() && plan.OperatorID.ValueString() != "" {
		auth.OperatorGuid = plan.OperatorID.ValueString()
	}
	if !plan.OperatorGroupID.IsNull() && plan.OperatorGroupID.ValueString() != "" {
		auth.OperatorGroupGuid = plan.OperatorGroupID.ValueString()
	}

	created, err := r.client.CreateVaultSectionAuthorization(plan.VaultSectionID.ValueString(), auth)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating vault section authorization",
			fmt.Sprintf("Could not create authorization for vault section %q: %s", plan.VaultSectionID.ValueString(), err.Error()),
		)
		return
	}

	plan.ID = types.StringValue(fmt.Sprintf("%s:%s", plan.VaultSectionID.ValueString(), created.AuthorizationId))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *vaultSectionPermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *vaultSectionPermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state vaultSectionPermissionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	parts := strings.Split(state.ID.ValueString(), ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid ID", "Expected ID in the format `vault_section_id:authorization_id`.")
		return
	}
	vaultSectionID := parts[0]
	authorizationID := parts[1]

	if err := r.client.DeleteVaultSectionAuthorization(vaultSectionID, authorizationID); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting vault section authorization",
			fmt.Sprintf("Could not delete authorization %q from vault section %q: %s", authorizationID, vaultSectionID, err.Error()),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *vaultSectionPermissionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Error importing resource",
			"Expected ID in the format `vault_section_id:authorization_id`.",
		)
		return
	}

	vaultSectionID := parts[0]
	authorizationID := parts[1]

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vault_section_id"), vaultSectionID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)

	authorizations, err := r.client.GetVaultSectionAuthorizations(vaultSectionID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error importing vault section authorization",
			fmt.Sprintf("Could not retrieve authorizations for vault section %q: %s", vaultSectionID, err.Error()),
		)
		return
	}

	for _, auth := range authorizations {
		if auth.AuthorizationId == authorizationID {
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("permission"), auth.AuthorizationType)...)
			if auth.OperatorGuid != "" {
				resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("operator_id"), auth.OperatorGuid)...)
			}
			if auth.OperatorGroupGuid != "" {
				resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("operatorgroup_id"), auth.OperatorGroupGuid)...)
			}
			return
		}
	}

	resp.Diagnostics.AddError(
		"Authorization not found",
		fmt.Sprintf("No authorization %q found for vault section %q.", authorizationID, vaultSectionID),
	)
}
