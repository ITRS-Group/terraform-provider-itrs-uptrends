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

var _ resource.Resource = &customIntegrationPermissionResource{}

type customIntegrationPermissionResource struct {
	client interfaces.ICustomIntegrationPermission
}

func NewCustomIntegrationPermissionResource(client interfaces.ICustomIntegrationPermission) resource.Resource {
	return &customIntegrationPermissionResource{
		client: client,
	}
}

type customIntegrationPermissionModel struct {
	ID                types.String `tfsdk:"id"`
	IntegrationID     types.String `tfsdk:"integration_id"`
	AuthorizationType types.String `tfsdk:"permission"`
	OperatorID        types.String `tfsdk:"operator_id"`
	OperatorGroupID   types.String `tfsdk:"operatorgroup_id"`
}

func (r *customIntegrationPermissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "itrs-uptrends_custom_integration_permission"
}

func (r *customIntegrationPermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = rschema.Schema{
		Attributes: map[string]rschema.Attribute{
			"id": rschema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier of the authorization (composite key in format `integration_id:authorization_id`).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"integration_id": rschema.StringAttribute{
				Required:    true,
				Description: "The GUID of the custom integration.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"permission": rschema.StringAttribute{
				Required:    true,
				Description: "The integration authorization type. Valid value: `UseIntegration`.",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"UseIntegration",
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

func (r *customIntegrationPermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state customIntegrationPermissionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	parts := strings.Split(state.ID.ValueString(), ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid ID", "Expected ID in the format `integration_id:authorization_id`.")
		return
	}
	integrationID := parts[0]
	authorizationID := parts[1]

	authorizations, err := r.client.GetIntegrationAuthorizations(integrationID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading integration authorization",
			fmt.Sprintf("Could not retrieve authorizations for integration %q: %s", integrationID, err.Error()),
		)
		return
	}

	for _, auth := range authorizations {
		if auth.AuthorizationId == authorizationID {
			state.IntegrationID = types.StringValue(integrationID)
			state.AuthorizationType = types.StringValue(auth.AuthorizationType)
			state.OperatorID = integrationAuthorizationString(auth.OperatorGuid)
			state.OperatorGroupID = integrationAuthorizationString(auth.OperatorGroupGuid)
			diags = resp.State.Set(ctx, state)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	resp.State.RemoveResource(ctx)
}

func (r *customIntegrationPermissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan customIntegrationPermissionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	auth := models.IntegrationAuthorization{
		AuthorizationType: plan.AuthorizationType.ValueString(),
	}
	if !plan.OperatorID.IsNull() && plan.OperatorID.ValueString() != "" {
		auth.OperatorGuid = plan.OperatorID.ValueString()
	}
	if !plan.OperatorGroupID.IsNull() && plan.OperatorGroupID.ValueString() != "" {
		auth.OperatorGroupGuid = plan.OperatorGroupID.ValueString()
	}

	created, err := r.client.CreateIntegrationAuthorization(plan.IntegrationID.ValueString(), auth)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating integration authorization",
			fmt.Sprintf("Could not create authorization for integration %q: %s", plan.IntegrationID.ValueString(), err.Error()),
		)
		return
	}

	plan.ID = types.StringValue(fmt.Sprintf("%s:%s", plan.IntegrationID.ValueString(), created.AuthorizationId))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *customIntegrationPermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *customIntegrationPermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state customIntegrationPermissionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	parts := strings.Split(state.ID.ValueString(), ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid ID", "Expected ID in the format `integration_id:authorization_id`.")
		return
	}
	integrationID := parts[0]
	authorizationID := parts[1]

	if err := r.client.DeleteIntegrationAuthorization(integrationID, authorizationID); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting integration authorization",
			fmt.Sprintf("Could not delete authorization %q from integration %q: %s", authorizationID, integrationID, err.Error()),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *customIntegrationPermissionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ":")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Error importing resource",
			"Expected ID in the format `integration_id:authorization_id`.",
		)
		return
	}

	integrationID := parts[0]
	authorizationID := parts[1]

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("integration_id"), integrationID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)

	authorizations, err := r.client.GetIntegrationAuthorizations(integrationID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error importing integration authorization",
			fmt.Sprintf("Could not retrieve authorizations for integration %q: %s", integrationID, err.Error()),
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
		fmt.Sprintf("No authorization %q found for integration %q.", authorizationID, integrationID),
	)
}

func integrationAuthorizationString(value string) types.String {
	if value == "" {
		return types.StringNull()
	}
	return types.StringValue(value)
}
